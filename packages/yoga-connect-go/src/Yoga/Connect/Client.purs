-- | A gRPC client on Go, derived from the service's row: a record with a function per method.
-- |
-- | ```purescript
-- | connection <- connect "http://127.0.0.1:50051"
-- | let greeterClient = client greeter connection
-- | reply <- greeterClient.sayHello { name: "Ada" }
-- | greeterClient.listGreetings { name: "Ada", count: 3 } # Strom.forEach (log <<< _.text)
-- | ```
-- |
-- | | Method | Client function |
-- | | --- | --- |
-- | | `Unary req res` | `req -> OmIO ctx err res` |
-- | | `ServerStream req res` | `req -> Strom ctx err res` |
-- | | `ClientStream req res` | `Strom ctx err req -> OmIO ctx err res` |
-- | | `BidiStream req res` | `Strom ctx err req -> Strom ctx err res` |
-- |
-- | where `err` has `connectError :: ConnectError`, which is how a call fails.
module Yoga.Connect.Client
  ( Connection
  , connect
  , client
  , class ClientMethods
  , clientMethods
  , class ClientMethod
  , clientMethod
  ) where

import Prelude

import Control.Monad.Error.Class (throwError)
import Control.Monad.Reader (ask)
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.Symbol (class IsSymbol)
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Exception (Error)
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, EffectFn5, runEffectFn1, runEffectFn2, runEffectFn3, runEffectFn5)
import Go.IO (forkIO)
import Go.IO.Internal (IO(..))
import Go.IO.IORef (newIORef, readIORef, writeIORef)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Record as Record
import Record.Builder (Builder)
import Record.Builder as Builder
import Type.Proxy (Proxy(..))
import Yoga.Connect.Service (BidiStream(..), ClientStream(..), Service, ServerStream(..), Unary(..), methodsOf)
import Yoga.Connect.Status (Code(..), ConnectError(..), codeFromNumber)
import Yoga.Connect.Strom (Strom(..), runStrom)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.Protobuf.Bytes (Bytes)
import Yoga.Protobuf.Codec (Codec)
import Yoga.Protobuf.Codec as Codec

-- | Where a server is, and the HTTP/2 connections to it.
foreign import data Connection :: Type

-- | One streaming call in flight.
foreign import data ClientCall :: Type

-- | Talks gRPC over HTTP/2 without TLS to the server at the address: `http://127.0.0.1:50051`.
connect :: String -> Effect Connection
connect address = runEffectFn1 connectImpl address

foreign import connectImpl :: EffectFn1 String Connection

-- | A function for each method of the service.
client
  :: forall methods list ctx err clients
   . RowToList methods list
  => ClientMethods list methods ctx err clients
  => Service methods
  -> Connection
  -> Record clients
client service connection = Builder.buildFromScratch (clientMethods (Proxy :: Proxy list) (Proxy :: Proxy ctx) (Proxy :: Proxy err) (methodsOf service) connection)

type Failing err = (connectError :: ConnectError | err)

class ClientMethods (list :: RowList Type) (methods :: Row Type) ctx (err :: Row Type) (clients :: Row Type) | list methods ctx err -> clients where
  clientMethods :: Proxy list -> Proxy ctx -> Proxy err -> Record methods -> Connection -> Builder {} (Record clients)

instance ClientMethods RowList.Nil methods ctx err () where
  clientMethods _ _ _ _ _ = identity

instance
  ( IsSymbol label
  , Row.Cons label method otherMethods methods
  , ClientMethod method ctx err function
  , ClientMethods tail methods ctx err otherClients
  , Row.Cons label function otherClients clients
  , Row.Lacks label otherClients
  ) =>
  ClientMethods (RowList.Cons label method tail) methods ctx err clients where
  clientMethods _ context errors methods connection =
    Builder.insert (Proxy :: Proxy label) (clientMethod context errors (Record.get (Proxy :: Proxy label) methods) connection)
      <<< clientMethods (Proxy :: Proxy tail) context errors methods connection

-- | The function that calls one method, by its kind.
class ClientMethod method ctx (err :: Row Type) function | method ctx err -> function where
  clientMethod :: Proxy ctx -> Proxy err -> method -> Connection -> function

instance ClientMethod (Unary request response) ctx err (request -> OmIO ctx (Failing err) response) where
  clientMethod _ _ (Unary method) connection request = do
    reply <- IO (runEffectFn5 callUnaryImpl failedWith Right connection method.procedure (Codec.encode method.request request)) # liftIO
    case reply of
      Left failure -> throwError (connectError failure)
      Right bytes -> decodeResponse method.response bytes

instance ClientMethod (ServerStream request response) ctx err (request -> Strom ctx (Failing err) response) where
  clientMethod _ _ (ServerStream method) connection request = Strom \sink -> do
    call <- IO (runEffectFn3 openServerStreamImpl connection method.procedure (Codec.encode method.request request)) # liftIO
    receiveAll call method.response sink

instance ClientMethod (ClientStream request response) ctx err (Strom ctx (Failing err) request -> OmIO ctx (Failing err) response) where
  clientMethod _ _ (ClientStream method) connection outgoing = do
    call <- IO (runEffectFn2 openClientStreamImpl connection method.procedure) # liftIO
    _ <- runStrom (sendTo call method.request) outgoing
    reply <- IO (runEffectFn3 closeAndReceiveImpl failedWith Right call) # liftIO
    case reply of
      Left failure -> throwError (connectError failure)
      Right bytes -> decodeResponse method.response bytes

-- | The requests go out on a thread of their own while the responses come in. If sending fails,
-- | the call is closed and the failure is the stream's.
instance ClientMethod (BidiStream request response) ctx err (Strom ctx (Failing err) request -> Strom ctx (Failing err) response) where
  clientMethod _ _ (BidiStream method) connection outgoing = Strom \sink -> do
    context <- ask
    call <- IO (runEffectFn2 openBidiStreamImpl connection method.procedure) # liftIO
    sendFailure <- newIORef Nothing # liftIO
    _ <- forkIO (sendAll context call sendFailure) # liftIO
    finished <- receiveAll call method.response sink
    failedSending <- readIORef sendFailure # liftIO
    case failedSending of
      Just failure -> throwError failure
      Nothing -> pure finished
    where
    sendAll context call sendFailure = do
      sent <- runOmIO context (runStrom (sendTo call method.request) outgoing)
      case sent of
        Left failure -> writeIORef sendFailure (Just failure) *> IO (runEffectFn1 closeImpl call)
        Right _ -> IO (runEffectFn1 closeSendingImpl call)

-- | Hands the responses to the sink as they arrive. If the sink stops early, the call is
-- | closed.
receiveAll :: forall ctx err response. Call -> Codec response -> (response -> OmIO ctx (Failing err) Boolean) -> OmIO ctx (Failing err) Boolean
receiveAll call codec sink = receiveNext
  where
  receiveNext = do
    received <- IO (runEffectFn1 receiveImpl call) # liftIO
    if received then do
      bytes <- IO (runEffectFn1 receivedImpl call) # liftIO
      response <- decodeResponse codec bytes
      more <- sink response
      if more then receiveNext else (IO (runEffectFn1 closeImpl call) # liftIO) $> false
    else do
      outcome <- IO (runEffectFn3 outcomeImpl failedWith (Right unit) call) # liftIO
      case outcome of
        Left failure -> throwError (connectError failure)
        Right _ -> pure true

sendTo :: forall ctx err request. Call -> Codec request -> request -> OmIO ctx err Boolean
sendTo call codec request = IO (Codec.encodeEffect codec request >>= runEffectFn2 sendImpl call) # liftIO

decodeResponse :: forall ctx err response. Codec response -> Bytes -> OmIO ctx (Failing err) response
decodeResponse codec bytes = case Codec.decode codec bytes of
  Left problem -> throwError (connectError (ConnectError { code: Internal, message: "the response doesn't decode: " <> show problem }))
  Right response -> pure response

connectError :: forall err. ConnectError -> Variant (exception :: Error | Failing err)
connectError failure = Variant.inj (Proxy :: Proxy "connectError") failure

failedWith :: forall a. Int -> String -> Either ConnectError a
failedWith code message = Left (ConnectError { code: codeFromNumber code, message })

type Call = ClientCall

foreign import callUnaryImpl :: EffectFn5 (Int -> String -> Either ConnectError Bytes) (Bytes -> Either ConnectError Bytes) Connection String Bytes (Either ConnectError Bytes)
foreign import openServerStreamImpl :: EffectFn3 Connection String Bytes ClientCall
foreign import openClientStreamImpl :: EffectFn2 Connection String ClientCall
foreign import openBidiStreamImpl :: EffectFn2 Connection String ClientCall

-- | `false` once the server is gone.
foreign import sendImpl :: EffectFn2 ClientCall Bytes Boolean

-- | Tells the server that all requests are sent.
foreign import closeSendingImpl :: EffectFn1 ClientCall Unit

foreign import closeAndReceiveImpl :: EffectFn3 (Int -> String -> Either ConnectError Bytes) (Bytes -> Either ConnectError Bytes) ClientCall (Either ConnectError Bytes)

-- | `true` with a response, `false` once there are no more.
foreign import receiveImpl :: EffectFn1 ClientCall Boolean
foreign import receivedImpl :: EffectFn1 ClientCall Bytes

-- | How the call ended, once `receiveImpl` says there are no more responses.
foreign import outcomeImpl :: EffectFn3 (Int -> String -> Either ConnectError Unit) (Either ConnectError Unit) ClientCall (Either ConnectError Unit)

-- | Ends the call early.
foreign import closeImpl :: EffectFn1 ClientCall Unit

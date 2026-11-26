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
-- | where `err` has `connectError :: ConnectError`, which is how a call fails. A streaming call
-- | lives as long as the stream: once its reader is done, the call is cancelled.
module Yoga.Connect.Client
  ( Connection
  , connect
  , client
  , clientWithin
  , class ClientMethods
  , clientMethods
  , class ClientMethod
  , clientMethod
  ) where

import Prelude

import Control.Monad.Error.Class (throwError, try)
import Data.Either (Either(..))
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Symbol (class IsSymbol)
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Exception (Error)
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, EffectFn4, EffectFn6, runEffectFn1, runEffectFn2, runEffectFn3, runEffectFn4, runEffectFn6)
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
import Yoga.Go.HTTP.OmIO (OmIO, liftIO)
import Yoga.OmIO.Strom (Exit(..), Scope, Step(..), Strom, addFinalizer, contextOf)
import Yoga.OmIO.Strom as Strom
import Yoga.OmIO.Strom.Context (Context, background)
import Yoga.OmIO.Strom.Scope (isInterruption, throwIfInterrupted)
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
client service connection = Builder.buildFromScratch (clientMethods (Proxy :: Proxy list) (Proxy :: Proxy ctx) (Proxy :: Proxy err) background (methodsOf service) connection)

-- | `client`, whose unary and client-streaming calls are made within the context: cancelling it
-- | cancels them, and its deadline is theirs, so that they fail with `DeadlineExceeded` once it
-- | has passed.
-- |
-- | ```purescript
-- | deadline /\ _ <- withTimeout (Milliseconds 500.0) background # liftIO
-- | reply <- (clientWithin deadline greeter connection).sayHello { name: "Ada" }
-- | ```
-- |
-- | A streamed response is made within the scope it's read in; `Strom.timeout` is its deadline.
clientWithin
  :: forall methods list ctx err clients
   . RowToList methods list
  => ClientMethods list methods ctx err clients
  => Context
  -> Service methods
  -> Connection
  -> Record clients
clientWithin context service connection = Builder.buildFromScratch (clientMethods (Proxy :: Proxy list) (Proxy :: Proxy ctx) (Proxy :: Proxy err) context (methodsOf service) connection)

type Failing err = (connectError :: ConnectError | err)

class ClientMethods (list :: RowList Type) (methods :: Row Type) ctx (err :: Row Type) (clients :: Row Type) | list methods ctx err -> clients where
  clientMethods :: Proxy list -> Proxy ctx -> Proxy err -> Context -> Record methods -> Connection -> Builder {} (Record clients)

instance ClientMethods RowList.Nil methods ctx err () where
  clientMethods _ _ _ _ _ _ = identity

instance
  ( IsSymbol label
  , Row.Cons label method otherMethods methods
  , ClientMethod method ctx err function
  , ClientMethods tail methods ctx err otherClients
  , Row.Cons label function otherClients clients
  , Row.Lacks label otherClients
  ) =>
  ClientMethods (RowList.Cons label method tail) methods ctx err clients where
  clientMethods _ ctx errors context methods connection =
    Builder.insert (Proxy :: Proxy label) (clientMethod ctx errors context (Record.get (Proxy :: Proxy label) methods) connection)
      <<< clientMethods (Proxy :: Proxy tail) ctx errors context methods connection

-- | The function that calls one method, by its kind, within the context.
class ClientMethod method ctx (err :: Row Type) function | method ctx err -> function where
  clientMethod :: Proxy ctx -> Proxy err -> Context -> method -> Connection -> function

instance ClientMethod (Unary request response) ctx err (request -> OmIO ctx (Failing err) response) where
  clientMethod _ _ context (Unary method) connection request = do
    reply <- IO (runEffectFn6 callUnaryImpl failedWith Right context connection method.procedure (Codec.encode method.request request)) # liftIO
    case reply of
      Left failure -> throwError (connectError failure)
      Right bytes -> decodeResponse method.response bytes

-- | The call is made within the context of the scope the stream is opened in: once the reader is
-- | done with the responses (`take` has enough, say), the call is cancelled, and the server's
-- | stream is interrupted. A `Strom.timeout` around it is the call's deadline.
instance ClientMethod (ServerStream request response) ctx err (request -> Strom ctx (Failing err) response) where
  clientMethod _ _ _ (ServerStream method) connection request = Strom.fromPull \scope -> do
    call <- IO (runEffectFn4 openServerStreamImpl (contextOf scope) connection method.procedure (Codec.encode method.request request)) # liftIO
    _ <- addFinalizer scope \_ -> closeCall call
    pure (receiveNext scope call method.response)

-- | The requests are read within the context too: once it's done, reading them stops and the call
-- | fails as the server saw it end, with `DeadlineExceeded` or `Canceled`.
instance ClientMethod (ClientStream request response) ctx err (Strom ctx (Failing err) request -> OmIO ctx (Failing err) response) where
  clientMethod _ _ context (ClientStream method) connection outgoing = do
    call <- IO (runEffectFn3 openClientStreamImpl context connection method.procedure) # liftIO
    sent <- try (outgoing # Strom.mapM (sendTo call method.request) # Strom.takeWhile identity # Strom.forEachWithin context (\_ -> pure unit))
    case sent of
      Left failure | not (isInterruption failure) -> throwError failure
      _ -> pure unit
    reply <- IO (runEffectFn3 closeAndReceiveImpl failedWith Right call) # liftIO
    case reply of
      Left failure -> throwError (connectError failure)
      Right bytes -> decodeResponse method.response bytes

-- | The requests go out on a goroutine of their own while the responses come in, both within
-- | the context of the scope the stream is opened in. If sending fails, the call is closed and the
-- | failure is the stream's.
instance ClientMethod (BidiStream request response) ctx err (Strom ctx (Failing err) request -> Strom ctx (Failing err) response) where
  clientMethod _ _ _ (BidiStream method) connection outgoing = Strom.fromPull \scope -> do
    call <- IO (runEffectFn3 openBidiStreamImpl (contextOf scope) connection method.procedure) # liftIO
    _ <- addFinalizer scope \_ -> closeCall call
    sendFailure <- newIORef Nothing # liftIO
    Strom.forkDrain scope (sending call sendFailure)
    pure do
      received <- try (receiveNext scope call method.response)
      failedSending <- readIORef sendFailure # liftIO
      case failedSending, received of
        Just failure, _ -> throwError failure
        Nothing, Left failure -> throwError failure
        Nothing, Right step -> pure step
    where
    sending call sendFailure = outgoing
      # Strom.mapM (sendTo call method.request)
      # Strom.takeWhile identity
      # Strom.onExit (afterSending call sendFailure)

    afterSending call sendFailure = case _ of
      Ended -> IO (runEffectFn1 closeSendingImpl call) # liftIO
      Failed failure -> do
        writeIORef sendFailure (Just failure) # liftIO
        closeCall call
      Interrupted -> pure unit

-- | The next response, a chunk each. Once the scope is interrupted, so is the stream.
receiveNext :: forall ctx err response. Scope ctx (Failing err) -> Call -> Codec response -> OmIO ctx (Failing err) (Step response)
receiveNext scope call codec = do
  received <- IO (runEffectFn1 receiveImpl call) # liftIO
  if received then do
    bytes <- IO (runEffectFn1 receivedImpl call) # liftIO
    response <- decodeResponse codec bytes
    pure (Emit [ response ])
  else do
    throwIfInterrupted scope
    outcome <- IO (runEffectFn3 outcomeImpl failedWith (Right unit) call) # liftIO
    case outcome of
      Left failure -> throwError (connectError failure)
      Right _ -> pure End

closeCall :: forall ctx err. Call -> OmIO ctx err Unit
closeCall call = IO (runEffectFn1 closeImpl call) # liftIO

-- | Sends the request, and says whether the server is still there to take more.
sendTo :: forall ctx err request. Call -> Codec request -> request -> OmIO ctx err Boolean
sendTo call codec request = IO (Codec.encodeEffect codec request >>= runEffectFn2 sendImpl call) # liftIO

decodeResponse :: forall ctx err response. Codec response -> Bytes -> OmIO ctx (Failing err) response
decodeResponse codec bytes = case Codec.decode codec bytes of
  Left problem -> throwError (connectError (ConnectError { code: Internal, message: "the response doesn't decode: " <> show problem, details: [] }))
  Right response -> pure response

connectError :: forall err. ConnectError -> Variant (exception :: Error | Failing err)
connectError failure = Variant.inj (Proxy :: Proxy "connectError") failure

-- | A failure, with the type names and encodings of its details.
failedWith :: forall a. Int -> String -> Array String -> Array Bytes -> Either ConnectError a
failedWith code message typeNames encodings = Left (ConnectError { code: codeFromNumber code, message, details: Array.zipWith { typeName: _, bytes: _ } typeNames encodings })

type Call = ClientCall

foreign import callUnaryImpl :: EffectFn6 (Int -> String -> Array String -> Array Bytes -> Either ConnectError Bytes) (Bytes -> Either ConnectError Bytes) Context Connection String Bytes (Either ConnectError Bytes)

-- | Opens a call within the context: cancelling it cancels the call, and its deadline is the
-- | call's.
foreign import openServerStreamImpl :: EffectFn4 Context Connection String Bytes ClientCall
foreign import openClientStreamImpl :: EffectFn3 Context Connection String ClientCall
foreign import openBidiStreamImpl :: EffectFn3 Context Connection String ClientCall

-- | `false` once the server is gone.
foreign import sendImpl :: EffectFn2 ClientCall Bytes Boolean

-- | Tells the server that all requests are sent.
foreign import closeSendingImpl :: EffectFn1 ClientCall Unit

foreign import closeAndReceiveImpl :: EffectFn3 (Int -> String -> Array String -> Array Bytes -> Either ConnectError Bytes) (Bytes -> Either ConnectError Bytes) ClientCall (Either ConnectError Bytes)

-- | `true` with a response, `false` once there are no more.
foreign import receiveImpl :: EffectFn1 ClientCall Boolean
foreign import receivedImpl :: EffectFn1 ClientCall Bytes

-- | How the call ended, once `receiveImpl` says there are no more responses.
foreign import outcomeImpl :: EffectFn3 (Int -> String -> Array String -> Array Bytes -> Either ConnectError Unit) (Either ConnectError Unit) ClientCall (Either ConnectError Unit)

-- | Ends the call early.
foreign import closeImpl :: EffectFn1 ClientCall Unit

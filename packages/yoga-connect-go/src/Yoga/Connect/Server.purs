-- | A gRPC server on Go's net/http, through connect-go, so it also answers gRPC-Web and Connect.
-- |
-- | ```purescript
-- | main = do
-- |   server <- createServer
-- |   serve greeter { sayHello, listGreetings, greetEveryone, chat } server
-- |   launchIO_ do
-- |     _ <- listen { host: Host "127.0.0.1", port: Port 50051 } server
-- |     log "listening"
-- | ```
-- |
-- | A handler runs in `OmIO`, on the call's own goroutine: decode, handler, encode. Its error row
-- | holds the status codes it rejects calls with (`Rejections`); any other exception is
-- | `internal`.
-- |
-- | | Method | Handler |
-- | | --- | --- |
-- | | `Unary req res` | `req -> OmIO ctx err res` |
-- | | `ServerStream req res` | `req -> Strom ctx err res` |
-- | | `ClientStream req res` | `Strom ctx err req -> OmIO ctx err res` |
-- | | `BidiStream req res` | `Strom ctx err req -> Strom ctx err res` |
module Yoga.Connect.Server
  ( Server
  , createServer
  , serve
  , serveWith
  , listen
  , reject
  , class ServeMethods
  , serveMethods
  , class ServeMethod
  , serveMethod
  , module Yoga.Go.HTTP.Om.Server
  ) where

import Prelude

import Control.Monad.Error.Class (class MonadThrow, throwError)
import Data.Either (Either(..), either)
import Data.Symbol (class IsSymbol)
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Exception (Error, error, message)
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, EffectFn5, runEffectFn1, runEffectFn2, runEffectFn3, runEffectFn5)
import Go.IO (IO, throwIO)
import Go.IO.Internal (IO(..))
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Record as Record
import Type.Equality (class TypeEquals, to)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Service (BidiStream(..), ClientStream(..), Service, ServerStream(..), Unary(..), methodsOf)
import Yoga.Connect.Status (class Rejection, Code(..), Rejections, codeNumber, rejectWith)
import Yoga.Connect.Strom (Strom(..), runStrom)
import Yoga.Go.HTTP.Om.Server (BoundAddress, Host(..), ListenOptions, Port(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.Protobuf.Bytes (Bytes)
import Yoga.Protobuf.Codec as Codec

-- | The routes, and once listening, the HTTP/2 server.
foreign import data Server :: Type

-- | One call a route answers: its request messages and its responses.
foreign import data Call :: Type

createServer :: Effect Server
createServer = createServerImpl

foreign import createServerImpl :: Effect Server

-- | Answers each method of the service with its handler. The handlers get no context.
serve
  :: forall methods list handlers
   . RowToList methods list
  => ServeMethods list methods handlers Unit
  => Service methods
  -> Record handlers
  -> Server
  -> Effect Unit
serve = serveWith unit

-- | Answers each method with its handler, which `ask`s for `context`.
serveWith
  :: forall methods list handlers ctx
   . RowToList methods list
  => ServeMethods list methods handlers ctx
  => ctx
  -> Service methods
  -> Record handlers
  -> Server
  -> Effect Unit
serveWith context service handlers server = serveMethods (Proxy :: Proxy list) context (methodsOf service) handlers server

-- | Binds the address and serves HTTP/1.1 and HTTP/2 without TLS in the background. While it
-- | listens, the server keeps the program running.
listen :: ListenOptions -> Server -> IO BoundAddress
listen { host: Host host, port: Port port } server = do
  bound <- IO (runEffectFn5 listenImpl Left Right host port server)
  either (throwIO <<< error) (\boundPort -> pure { address: host, port: Port boundPort }) bound

foreign import listenImpl :: EffectFn5 (String -> Either String Int) (Int -> Either String Int) String Int Server (Either String Int)

-- | Fails the call with a status code: `reject @"notFound" "no such user"`.
reject
  :: forall @code err others rest m a
   . IsSymbol code
  => Row.Cons code String others Rejections
  => Row.Cons code String rest (exception :: Error | err)
  => MonadThrow (Variant (exception :: Error | err)) m
  => String
  -> m a
reject explanation = throwError (Variant.inj (Proxy :: Proxy code) explanation)

-- | Registers a handler for each method of the row list.
class ServeMethods (list :: RowList Type) (methods :: Row Type) (handlers :: Row Type) ctx where
  serveMethods :: Proxy list -> ctx -> Record methods -> Record handlers -> Server -> Effect Unit

instance ServeMethods RowList.Nil methods handlers ctx where
  serveMethods _ _ _ _ _ = pure unit

instance
  ( IsSymbol label
  , Row.Cons label method otherMethods methods
  , Row.Cons label handler otherHandlers handlers
  , ServeMethod method handler ctx
  , ServeMethods tail methods handlers ctx
  ) =>
  ServeMethods (RowList.Cons label method tail) methods handlers ctx where
  serveMethods _ context methods handlers server = do
    serveMethod context (Record.get (Proxy :: Proxy label) methods) (Record.get (Proxy :: Proxy label) handlers) server
    serveMethods (Proxy :: Proxy tail) context methods handlers server

-- | Registers the handler of one method, by its kind.
class ServeMethod method handler ctx where
  serveMethod :: ctx -> method -> handler -> Server -> Effect Unit

-- | `TypeEquals` lets a handler that works in any context, `forall ctx. … OmIO ctx …`, take
-- | the server's.
instance
  ( RowToList err labels
  , Rejection labels err
  , TypeEquals ctx handlerContext
  ) =>
  ServeMethod (Unary request response) (request -> OmIO handlerContext err response) ctx where
  serveMethod context (Unary method) handler server = runEffectFn3 unaryImpl method.procedure answer server
    where
    answer call = case Codec.decode method.request (requestImpl call) of
      Left problem -> failCall call InvalidArgument (show problem)
      Right request -> runOmIO (to context) (handler request) >>= case _ of
        Left problem -> failWith call problem
        Right response -> respondWith call method.response response

instance
  ( RowToList err labels
  , Rejection labels err
  , TypeEquals ctx handlerContext
  ) =>
  ServeMethod (ServerStream request response) (request -> Strom handlerContext err response) ctx where
  serveMethod context (ServerStream method) handler server = runEffectFn3 serverStreamImpl method.procedure answer server
    where
    answer call = case Codec.decode method.request (requestImpl call) of
      Left problem -> failCall call InvalidArgument (show problem)
      Right request -> runOmIO (to context) (runStrom (sendTo call method.response) (handler request)) >>= case _ of
        Left problem -> failWith call problem
        Right _ -> pure unit

instance
  ( RowToList err labels
  , Rejection labels err
  , TypeEquals ctx handlerContext
  ) =>
  ServeMethod (ClientStream request response) (Strom handlerContext err request -> OmIO handlerContext err response) ctx where
  serveMethod context (ClientStream method) handler server = runEffectFn3 clientStreamImpl method.procedure answer server
    where
    answer call = runOmIO (to context) (handler (incoming call method.request)) >>= case _ of
      Left problem -> failWith call problem
      Right response -> respondWith call method.response response

instance
  ( RowToList err labels
  , Rejection labels err
  , TypeEquals ctx handlerContext
  ) =>
  ServeMethod (BidiStream request response) (Strom handlerContext err request -> Strom handlerContext err response) ctx where
  serveMethod context (BidiStream method) handler server = runEffectFn3 bidiStreamImpl method.procedure answer server
    where
    answer call = runOmIO (to context) (runStrom (sendTo call method.response) (handler (incoming call method.request))) >>= case _ of
      Left problem -> failWith call problem
      Right _ -> pure unit

-- | Encodes the response when the call answers with it.
respondWith :: forall response. Call -> Codec.Codec response -> response -> IO Unit
respondWith call codec response = IO (Codec.encodeEffect codec response >>= runEffectFn2 respondImpl call)

-- | Sends each response as it comes, and says to stop once the client is gone.
sendTo :: forall ctx err response. Call -> Codec.Codec response -> response -> OmIO ctx err Boolean
sendTo call codec response = IO (Codec.encodeEffect codec response >>= runEffectFn2 sendImpl call) # liftIO

-- | The client's messages, decoded as they arrive. A message that doesn't decode ends the stream
-- | and fails the call with `invalid_argument`.
incoming :: forall ctx err request. Call -> Codec.Codec request -> Strom ctx err request
incoming call codec = Strom receiveInto
  where
  receiveInto sink = do
    received <- IO (runEffectFn1 receiveImpl call) # liftIO
    if not received then pure true
    else do
      received <- IO (runEffectFn1 receivedImpl call) # liftIO
      case Codec.decode codec received of
        Left problem -> (failCall call InvalidArgument (show problem) # liftIO) $> false
        Right request -> do
          more <- sink request
          if more then receiveInto sink else pure false

failWith :: forall err labels. RowToList err labels => Rejection labels err => Call -> Variant (exception :: Error | err) -> IO Unit
failWith call = Variant.on (Proxy :: Proxy "exception") (failCall call Internal <<< message) (rejectWith (failCall call))

failCall :: Call -> Code -> String -> IO Unit
failCall call code explanation = IO (runEffectFn3 failImpl call (codeNumber code) explanation)

foreign import unaryImpl :: EffectFn3 String (Call -> IO Unit) Server Unit
foreign import serverStreamImpl :: EffectFn3 String (Call -> IO Unit) Server Unit
foreign import clientStreamImpl :: EffectFn3 String (Call -> IO Unit) Server Unit
foreign import bidiStreamImpl :: EffectFn3 String (Call -> IO Unit) Server Unit

-- | The request of a unary or server-streaming call.
foreign import requestImpl :: Call -> Bytes

foreign import respondImpl :: EffectFn2 Call Bytes Unit

-- | `false` once the client is gone.
foreign import sendImpl :: EffectFn2 Call Bytes Boolean

-- | `true` with a message from the client, `false` once it has sent all of them.
foreign import receiveImpl :: EffectFn1 Call Boolean

-- | The message `receiveImpl` got.
foreign import receivedImpl :: EffectFn1 Call Bytes

-- | The first failure wins; a call that failed answers with it, whatever the handler returns.
foreign import failImpl :: EffectFn3 Call Int String Unit

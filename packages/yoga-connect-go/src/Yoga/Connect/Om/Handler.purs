-- | Handlers of a service type's methods, which `ask` for the request and whatever else they
-- | need.
-- |
-- | ```purescript
-- | watch :: Handler (ServerStream Subscribe Quote (notFound :: String)) (market :: Market)
-- | watch = handle do
-- |   { request, market } <- ask
-- |   …
-- |   pure (Strom.mergeAll (listings <#> quotesOf))
-- | ```
-- |
-- | | Method | `handle`s |
-- | | --- | --- |
-- | | `Unary req res err` | `OmIO { request :: req \| ctx } err res` |
-- | | `ServerStream req res err` | `OmIO { request :: req \| ctx } err (Strom (Record ctx) err res)` |
-- | | `ClientStream req res err` | `OmIO { requests :: Strom (Record ctx) err req \| ctx } err res` |
-- | | `BidiStream req res err` | `OmIO { requests :: Strom (Record ctx) err req \| ctx } err (Strom (Record ctx) err res)` |
-- |
-- | where `req` and `res` are the plain records of the message types and `err` is the method's
-- | error row: `reject @"notFound" "no such symbol"` fails the call with `NOT_FOUND`. A response
-- | stream runs once the handler has answered with it, as long as the call lasts.
module Yoga.Connect.Om.Handler
  ( Handler(..)
  , class Handle
  , handle
  , withoutRequests
  , reject
  , registerService
  , registerServiceWith
  , class RegisterHandlers
  , registerHandlers
  ) where

import Prelude

import Control.Monad.Error.Class (class MonadThrow, throwError)
import Control.Monad.Reader (withReaderT)
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect.Exception (Error)
import Data.Newtype (class Newtype, un)
import Data.Symbol (class IsSymbol, reflectSymbol)
import Effect (Effect)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Record as Record
import Record.Builder as Builder
import Type.Proxy (Proxy(..))
import Yoga.Connect.Protobuf (class MethodMessage, methodCodec)
import Yoga.Connect.Schema (class ServiceType, BidiStream, ClientStream, ServerStream, Unary, procedureOf)
import Yoga.Connect.Rejection (class Rejections, rejections)
import Yoga.Connect.Server (Server, serveRejecting)
import Yoga.Connect.Service as Connect
import Yoga.Go.HTTP.Om.API (class PickFields, pickFields)
import Yoga.Go.HTTP.OmIO (OmIO(..))
import Yoga.OmIO.Strom (Strom)
import Yoga.OmIO.Strom as Strom

-- | The handler of a method, once it has the context it asks for: `(market :: Market)`.
newtype Handler :: Type -> Row Type -> Type
newtype Handler method ctx = Handler (Record ctx -> String -> Server -> Effect Unit)

derive instance Newtype (Handler method ctx) _

-- | Makes the handler of a method out of what it does, which depends on the method's kind.
class Handle :: Type -> Row Type -> Type -> Constraint
class Handle method ctx body | method ctx -> body where
  handle :: body -> Handler method ctx

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , RowToList errors labels
  , Rejections labels plainErrors
  , Row.Lacks "request" ctx
  , Row.Cons "request" plainRequest ctx withRequest
  ) =>
  Handle (Unary request response errors) ctx (OmIO (Record withRequest) plainErrors plainResponse) where
  handle body = Handler \dependencies procedure ->
    serveRejecting (rejections (Proxy :: Proxy labels)).encode dependencies (Connect.unary procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) answer
    where
    answer request = body # providing (Record.insert (Proxy :: Proxy "request") request)

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , RowToList errors labels
  , Rejections labels plainErrors
  , Row.Lacks "request" ctx
  , Row.Cons "request" plainRequest ctx withRequest
  ) =>
  Handle (ServerStream request response errors) ctx (OmIO (Record withRequest) plainErrors (Strom (Record ctx) plainErrors plainResponse)) where
  handle body = Handler \dependencies procedure ->
    serveRejecting (rejections (Proxy :: Proxy labels)).encode dependencies (Connect.serverStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) answer
    where
    answer request = body # providing (Record.insert (Proxy :: Proxy "request") request) # streamed

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , RowToList errors labels
  , Rejections labels plainErrors
  , Row.Lacks "requests" ctx
  , Row.Cons "requests" (Strom (Record ctx) plainErrors plainRequest) ctx withRequests
  ) =>
  Handle (ClientStream request response errors) ctx (OmIO (Record withRequests) plainErrors plainResponse) where
  handle body = Handler \dependencies procedure ->
    serveRejecting (rejections (Proxy :: Proxy labels)).encode dependencies (Connect.clientStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) answer
    where
    answer requests = body # providing (Record.insert (Proxy :: Proxy "requests") requests)

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , RowToList errors labels
  , Rejections labels plainErrors
  , Row.Lacks "requests" ctx
  , Row.Cons "requests" (Strom (Record ctx) plainErrors plainRequest) ctx withRequests
  ) =>
  Handle (BidiStream request response errors) ctx (OmIO (Record withRequests) plainErrors (Strom (Record ctx) plainErrors plainResponse)) where
  handle body = Handler \dependencies procedure ->
    serveRejecting (rejections (Proxy :: Proxy labels)).encode dependencies (Connect.bidiStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) answer
    where
    answer requests = body # providing (Record.insert (Proxy :: Proxy "requests") requests) # streamed

-- | Fails the call with one of the method's errors: `reject @"notFound" "no such symbol"`,
-- | `reject @"columnFull" 3`. The error row says how it goes over the wire.
reject :: forall @label payload others err m a. IsSymbol label => Row.Cons label payload others (exception :: Error | err) => MonadThrow (Variant (exception :: Error | err)) m => payload -> m a
reject payload = throwError (Variant.inj (Proxy :: Proxy label) payload)

-- | Runs what needs only the handler's context, such as reading the requests of a client
-- | stream to the end: `requests # Strom.runCollect # withoutRequests`.
withoutRequests :: forall ctx withRequests requests err a. Row.Lacks "requests" ctx => Row.Cons "requests" requests ctx withRequests => OmIO (Record ctx) err a -> OmIO (Record withRequests) err a
withoutRequests = providing (Record.delete (Proxy :: Proxy "requests"))

-- | The handler's answer, the stream it makes, as the stream of the call.
streamed :: forall ctx err a. OmIO ctx err (Strom ctx err a) -> Strom ctx err a
streamed answer = Strom.fromOm answer # Strom.flatMap identity

providing :: forall outer inner err a. (outer -> inner) -> OmIO inner err a -> OmIO outer err a
providing narrow (OmIO om) = OmIO (withReaderT narrow om)

-- | Answers each method of the service with its handler. The handlers ask for nothing.
registerService
  :: forall @service name methods list handlers
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => RegisterHandlers list handlers ()
  => Record handlers
  -> Server
  -> Effect Unit
registerService = registerServiceWith @service {}

-- | Answers each method of the service with its handler, which gets the fields of the context it
-- | asks for.
registerServiceWith
  :: forall @service name methods list handlers context
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => RegisterHandlers list handlers context
  => Record context
  -> Record handlers
  -> Server
  -> Effect Unit
registerServiceWith context handlers server = registerHandlers (Proxy :: Proxy list) (reflectSymbol (Proxy :: Proxy name)) context handlers server

class RegisterHandlers :: RowList Type -> Row Type -> Row Type -> Constraint
class RegisterHandlers list handlers context where
  registerHandlers :: Proxy list -> String -> Record context -> Record handlers -> Server -> Effect Unit

instance RegisterHandlers RowList.Nil handlers context where
  registerHandlers _ _ _ _ _ = pure unit

instance
  ( IsSymbol label
  , Row.Cons label (Handler method handlerContext) otherHandlers handlers
  , RowToList handlerContext handlerContextList
  , PickFields handlerContextList context handlerContext
  , RegisterHandlers tail handlers context
  ) =>
  RegisterHandlers (RowList.Cons label method tail) handlers context where
  registerHandlers _ serviceName context handlers server = do
    serving (Builder.buildFromScratch (pickFields (Proxy :: Proxy handlerContextList) context)) (procedureOf serviceName (reflectSymbol (Proxy :: Proxy label))) server
    registerHandlers (Proxy :: Proxy tail) serviceName context handlers server
    where
    serving = un Handler (Record.get (Proxy :: Proxy label) handlers)

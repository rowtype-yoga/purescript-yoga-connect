-- | A gRPC service as a type: its full name, and a record of its methods.
-- |
-- | ```purescript
-- | type Ticker = Service "ticker.v1.Ticker"
-- |   { watch :: ServerStream Subscribe Quote (notFound :: String)
-- |   , trade :: BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)
-- |   }
-- | ```
-- |
-- | Each method says what it takes, what it answers with, and how it fails: the last row has the
-- | gRPC status codes, by name, that its handler rejects calls with, each with its message. The
-- | method `watch` is `/ticker.v1.Ticker/Watch` on the wire.
-- |
-- | The server, the client and the `.proto` all come from this type. It says nothing about how
-- | the messages are encoded or carried: `connectService` works out the protocol buffer codecs,
-- | and yoga-connect-go serves and calls it over HTTP/2.
module Yoga.Connect.Schema
  ( Service
  , Unary
  , ServerStream
  , ClientStream
  , BidiStream
  , connectService
  , class ServiceType
  , class ServiceMethods
  , serviceMethods
  , methodShapes
  , class Method
  , connectMethod
  , methodSchema
  , MethodShape
  , procedureOf
  , methodNameOf
  ) where

import Prelude

import Data.Array as Array
import Data.String as String
import Data.String.CodeUnits as CodeUnits
import Data.Symbol (class IsSymbol, reflectSymbol)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Record.Builder (Builder)
import Record.Builder as Builder
import Type.Proxy (Proxy(..))
import Data.Maybe (Maybe)
import Yoga.Connect.Protobuf (class MethodMessage, MessageShape, methodCodec, methodName, methodShape)
import Yoga.Connect.Service as Connect

-- | A service called `name`, such as `"ticker.v1.Ticker"`, with a record of its methods.
data Service :: Symbol -> Type -> Type
data Service name methods

-- | One request, one response.
data Unary :: Type -> Type -> Row Type -> Type
data Unary request response errors

-- | One request, any number of responses.
data ServerStream :: Type -> Type -> Row Type -> Type
data ServerStream request response errors

-- | Any number of requests, one response.
data ClientStream :: Type -> Type -> Row Type -> Type
data ClientStream request response errors

-- | Requests and responses, both streams, at the same time.
data BidiStream :: Type -> Type -> Row Type -> Type
data BidiStream request response errors

-- | The name and the methods of a service type.
class ServiceType :: Type -> Symbol -> Row Type -> Constraint
class ServiceType service name methods | service -> name methods

instance ServiceType (Service name (Record methods)) name methods

-- | The routes and codecs of a service type's methods, which yoga-connect-go serves and calls.
connectService
  :: forall @service name methods list connect
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => ServiceMethods list connect
  => Connect.Service connect
connectService = Connect.service serviceName (Builder.buildFromScratch (serviceMethods (Proxy :: Proxy list) serviceName))
  where
  serviceName = reflectSymbol (Proxy :: Proxy name)

class ServiceMethods :: RowList Type -> Row Type -> Constraint
class ServiceMethods list connect | list -> connect where
  serviceMethods :: Proxy list -> String -> Builder {} (Record connect)
  methodShapes :: Proxy list -> Array { name :: String, schema :: MethodShape }

instance ServiceMethods RowList.Nil () where
  serviceMethods _ _ = identity
  methodShapes _ = []

instance
  ( IsSymbol label
  , Method method connectMethod
  , ServiceMethods tail connectTail
  , Row.Cons label connectMethod connectTail connect
  , Row.Lacks label connectTail
  ) =>
  ServiceMethods (RowList.Cons label method tail) connect where
  serviceMethods _ serviceName =
    Builder.insert (Proxy :: Proxy label) (connectMethod (Proxy :: Proxy method) (procedureOf serviceName (reflectSymbol (Proxy :: Proxy label))))
      <<< serviceMethods (Proxy :: Proxy tail) serviceName
  methodShapes _ = Array.cons { name: methodNameOf (reflectSymbol (Proxy :: Proxy label)), schema: methodSchema (Proxy :: Proxy method) } (methodShapes (Proxy :: Proxy tail))

-- | What a `.proto` says about a method.
type MethodShape =
  { clientStreams :: Boolean
  , serverStreams :: Boolean
  , request :: MessageShape
  , response :: MessageShape
  , requestName :: Maybe String
  , responseName :: Maybe String
  }

-- | A method kind, as the route and codecs yoga-connect's runtime needs.
class Method :: Type -> Type -> Constraint
class Method method connectMethod | method -> connectMethod where
  connectMethod :: Proxy method -> String -> connectMethod
  methodSchema :: Proxy method -> MethodShape

instance (MethodMessage request plainRequest, MethodMessage response plainResponse) => Method (Unary request response errors) (Connect.Unary plainRequest plainResponse) where
  connectMethod _ procedure = Connect.unary procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))
  methodSchema _ = { clientStreams: false, serverStreams: false, request: methodShape (Proxy :: Proxy request), response: methodShape (Proxy :: Proxy response), requestName: methodName (Proxy :: Proxy request), responseName: methodName (Proxy :: Proxy response) }

instance (MethodMessage request plainRequest, MethodMessage response plainResponse) => Method (ServerStream request response errors) (Connect.ServerStream plainRequest plainResponse) where
  connectMethod _ procedure = Connect.serverStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))
  methodSchema _ = { clientStreams: false, serverStreams: true, request: methodShape (Proxy :: Proxy request), response: methodShape (Proxy :: Proxy response), requestName: methodName (Proxy :: Proxy request), responseName: methodName (Proxy :: Proxy response) }

instance (MethodMessage request plainRequest, MethodMessage response plainResponse) => Method (ClientStream request response errors) (Connect.ClientStream plainRequest plainResponse) where
  connectMethod _ procedure = Connect.clientStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))
  methodSchema _ = { clientStreams: true, serverStreams: false, request: methodShape (Proxy :: Proxy request), response: methodShape (Proxy :: Proxy response), requestName: methodName (Proxy :: Proxy request), responseName: methodName (Proxy :: Proxy response) }

instance (MethodMessage request plainRequest, MethodMessage response plainResponse) => Method (BidiStream request response errors) (Connect.BidiStream plainRequest plainResponse) where
  connectMethod _ procedure = Connect.bidiStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))
  methodSchema _ = { clientStreams: true, serverStreams: true, request: methodShape (Proxy :: Proxy request), response: methodShape (Proxy :: Proxy response), requestName: methodName (Proxy :: Proxy request), responseName: methodName (Proxy :: Proxy response) }

-- | `/ticker.v1.Ticker/Watch` for the method `watch`.
procedureOf :: String -> String -> String
procedureOf serviceName label = "/" <> serviceName <> "/" <> methodNameOf label

-- | `Watch` for `watch`.
methodNameOf :: String -> String
methodNameOf label = String.toUpper (CodeUnits.take 1 label) <> CodeUnits.drop 1 label

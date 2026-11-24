-- | A service is a row of its methods, one entry per RPC:
-- |
-- | ```purescript
-- | type Greeter =
-- |   ( sayHello :: Unary HelloRequest HelloReply
-- |   , listGreetings :: ServerStream ListRequest Greeting
-- |   , chat :: BidiStream ChatMessage ChatMessage
-- |   )
-- | ```
-- |
-- | protoc-gen-purescript writes that row and a `Service Greeter` value with the route and the
-- | message codecs of each method. The Go server and the browser client both work out what they
-- | need from the row.
module Yoga.Connect.Service
  ( Service(..)
  , service
  , serviceName
  , methodsOf
  , Method
  , Unary(..)
  , ServerStream(..)
  , ClientStream(..)
  , BidiStream(..)
  , unary
  , serverStream
  , clientStream
  , bidiStream
  ) where

import Yoga.Protobuf.Codec (Codec)

-- | A service's full protocol buffer name and its methods.
newtype Service (methods :: Row Type) = Service { name :: String, methods :: Record methods }

service :: forall methods. String -> Record methods -> Service methods
service name methods = Service { name, methods }

-- | `greeter.v1.Greeter`.
serviceName :: forall methods. Service methods -> String
serviceName (Service { name }) = name

methodsOf :: forall methods. Service methods -> Record methods
methodsOf (Service { methods }) = methods

-- | What both sides need to call a method: its route, `/greeter.v1.Greeter/SayHello`, and how
-- | its messages go over the wire.
type Method request response =
  { procedure :: String
  , request :: Codec request
  , response :: Codec response
  }

-- | One request, one response.
newtype Unary request response = Unary (Method request response)

-- | One request, any number of responses.
newtype ServerStream request response = ServerStream (Method request response)

-- | Any number of requests, one response.
newtype ClientStream request response = ClientStream (Method request response)

-- | Requests and responses, both streams, at the same time.
newtype BidiStream request response = BidiStream (Method request response)

unary :: forall request response. String -> Codec request -> Codec response -> Unary request response
unary procedure request response = Unary { procedure, request, response }

serverStream :: forall request response. String -> Codec request -> Codec response -> ServerStream request response
serverStream procedure request response = ServerStream { procedure, request, response }

clientStream :: forall request response. String -> Codec request -> Codec response -> ClientStream request response
clientStream procedure request response = ClientStream { procedure, request, response }

bidiStream :: forall request response. String -> Codec request -> Codec response -> BidiStream request response
bidiStream procedure request response = BidiStream { procedure, request, response }

module Test.Yoga.Connect.ServiceSpec where

import Prelude

import Data.Either (Either(..))
import Demo.Proto.Greeter.V1.Greeter (greeter)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual)
import Yoga.Connect.Service (BidiStream(..), ClientStream(..), ServerStream(..), Unary(..), methodsOf, serviceName)
import Yoga.Protobuf.Codec as Codec

spec :: Spec Unit
spec = describe "a generated service" do
  it "has the service's full name" do
    serviceName greeter `shouldEqual` "greeter.v1.Greeter"

  it "routes each method to /package.Service/Method, whatever its kind" do
    case methodsOf greeter of
      { sayHello: Unary sayHello, listGreetings: ServerStream listGreetings, greetEveryone: ClientStream greetEveryone, chat: BidiStream chat } -> do
        sayHello.procedure `shouldEqual` "/greeter.v1.Greeter/SayHello"
        listGreetings.procedure `shouldEqual` "/greeter.v1.Greeter/ListGreetings"
        greetEveryone.procedure `shouldEqual` "/greeter.v1.Greeter/GreetEveryone"
        chat.procedure `shouldEqual` "/greeter.v1.Greeter/Chat"

  it "carries the codecs of each method's messages" do
    case methodsOf greeter of
      { listGreetings: ServerStream listGreetings, chat: BidiStream chat } -> do
        Codec.decode listGreetings.request (Codec.encode listGreetings.request { name: "Ada", count: 3 })
          `shouldEqual` Right { name: "Ada", count: 3 }
        Codec.decode listGreetings.response (Codec.encode listGreetings.response { index: 2, text: "hi" })
          `shouldEqual` Right { index: 2, text: "hi" }
        Codec.decode chat.request (Codec.encode chat.response { from: "Ada", text: "both ways" })
          `shouldEqual` Right { from: "Ada", text: "both ways" }

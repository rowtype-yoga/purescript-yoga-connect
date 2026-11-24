module Test.Yoga.Connect.CallSpec where

import Prelude

import Control.Monad.Error.Class (try)
import Control.Monad.Reader (ask)
import Data.Either (Either(..), either)
import Data.Maybe (Maybe(..))
import Data.Variant as Variant
import Demo.Proto.Greeter.V1.Greeter (ChatMessage, Greeting, HelloReply, HelloRequest, ListRequest, greeter)
import Effect.Class (liftEffect)
import Go.IO (toAff)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Client (Connection, client, connect)
import Yoga.Connect.Server (Host(..), Port(..), createServer, listen, reject, serveWith)
import Yoga.Connect.Status (Code(..), ConnectError(..))
import Yoga.Connect.Strom (Strom)
import Yoga.Connect.Strom as Strom
import Yoga.Go.HTTP.OmIO (OmIO, runOmIO)

spec :: Spec Unit
spec = describe "a server and a client derived from the same service row" do
  it "answers each kind of method, and a rejection with its code" do
    outcome <- toAff do
      server <- createServer # liftEffect
      serveWith { greeting: "Hi" } greeter { sayHello, listGreetings, greetEveryone, chat } server # liftEffect
      bound <- listen { host: Host "127.0.0.1", port: Port 0 } server
      connection <- connect ("http://127.0.0.1:" <> show (portNumber bound.port)) # liftEffect
      runOmIO unit (callEach connection)
    either (const Nothing) Just outcome `shouldEqual` Just
      { hello: "Hi, Ada"
      , greetings: [ { index: 1, text: "Ada" }, { index: 2, text: "Ada" } ]
      , everyone: "Ada, Grace, "
      , echoes: [ "one", "two" ]
      , rejected: Just (ConnectError { code: NotFound, message: "no one by that name" })
      }
  where
  portNumber (Port number) = number

callEach
  :: Connection
  -> OmIO Unit (connectError :: ConnectError)
       { hello :: String
       , greetings :: Array Greeting
       , everyone :: String
       , echoes :: Array String
       , rejected :: Maybe ConnectError
       }
callEach connection = do
  let greeterClient = client greeter connection
  hello <- greeterClient.sayHello { name: "Ada" }
  greetings <- greeterClient.listGreetings { name: "Ada", count: 2 } # Strom.collect
  everyone <- greeterClient.greetEveryone (Strom.fromArray [ { name: "Ada" }, { name: "Grace" } ])
  echoes <- greeterClient.chat (Strom.fromArray [ "one", "two" ] # map { from: "spec", text: _ }) # map _.text # Strom.collect
  rejected <- try (greeterClient.sayHello { name: "nobody" })
  pure
    { hello: hello.message
    , greetings
    , everyone: everyone.message
    , echoes
    , rejected: case rejected of
        Left failure -> Variant.on (Proxy :: Proxy "connectError") Just (const Nothing) failure
        Right _ -> Nothing
    }

sayHello :: HelloRequest -> OmIO { greeting :: String } (notFound :: String) HelloReply
sayHello { name } = do
  { greeting } <- ask
  if name == "nobody" then reject @"notFound" "no one by that name"
  else pure { message: greeting <> ", " <> name }

listGreetings :: forall ctx. ListRequest -> Strom ctx () Greeting
listGreetings { name, count } = Strom.range 1 count # map { index: _, text: name }

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # Strom.foldl (\names person -> names <> person.name <> ", ") ""
  pure { message: names }

chat :: forall ctx. Strom ctx () ChatMessage -> Strom ctx () ChatMessage
chat incoming = incoming # Strom.mapM \{ text } -> pure { from: "echo", text }

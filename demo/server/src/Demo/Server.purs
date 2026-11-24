-- | The demo's gRPC server: one method of each kind.
module Demo.Server where

import Prelude

import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Effect (Effect)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Demo.Proto.Greeter.V1.Greeter (ChatMessage, Greeting, HelloReply, HelloRequest, ListRequest, greeter)
import Yoga.Connect.Server (Host(..), Port(..), createServer, listen, reject, serve)
import Yoga.Connect.Strom (Strom)
import Yoga.Connect.Strom as Strom
import Yoga.Go.HTTP.OmIO (OmIO)

main :: Effect Unit
main = do
  server <- createServer
  serve greeter { sayHello, listGreetings, greetEveryone, chat } server
  launchIO_ do
    bound <- listen { host: Host "127.0.0.1", port: Port 50051 } server
    log ("greeter listening on " <> bound.address <> ":" <> show bound.port)

sayHello :: forall ctx. HelloRequest -> OmIO ctx (invalidArgument :: String) HelloReply
sayHello { name }
  | name == "" = reject @"invalidArgument" "say who you are"
  | otherwise = pure { message: "Hello, " <> name <> "!" }

listGreetings :: forall ctx. ListRequest -> Strom ctx () Greeting
listGreetings { name, count } = Strom.range 1 count # map greeting
  where
  greeting index = { index, text: "Greeting number " <> show index <> " for " <> name }

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # map _.name # Strom.collect
  pure { message: "Hello, " <> String.joinWith " and " names <> "!" }

-- | Answers each message as it arrives, while the server's own heartbeat ticks in between.
chat :: forall ctx. Strom ctx () ChatMessage -> Strom ctx () ChatMessage
chat incoming = Strom.merge answers heartbeats
  where
  answers = incoming # map answer
  heartbeats = Strom.every (Milliseconds 150.0) # Strom.take 3 # map heartbeat
  answer { from, text } = { from: "server", text: from <> " said " <> show (String.length text) <> " characters: " <> String.toUpper text }
  heartbeat tick = { from: "server", text: "heartbeat " <> show (tick + 1) }

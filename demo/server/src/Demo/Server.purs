-- | The demo's gRPC server: one method of each kind. Its streams say how they ended, so that you
-- | can see a client's cancel reach them.
module Demo.Server where

import Prelude

import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Effect (Effect)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Demo.Proto.Greeter.V1.Greeter (ChatMessage, Greeting, HelloReply, HelloRequest, ListRequest, greeter)
import Yoga.Connect.Server (Host(..), Port(..), createServer, listen, reject, serve)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO)
import Yoga.OmIO.Strom (Exit(..), Strom)
import Yoga.OmIO.Strom as Strom

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

-- | `count` greetings, one every 20 ms.
listGreetings :: forall ctx. ListRequest -> Strom ctx () Greeting
listGreetings { name, count } = Strom.every (Milliseconds 20.0)
  # Strom.take count
  # map greeting
  # Strom.onExit (sayHowItEnded ("listGreetings for " <> name))
  where
  greeting tick = { index: tick + 1, text: "Greeting number " <> show (tick + 1) <> " for " <> name }

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # map _.name # Strom.runCollect
  pure { message: "Hello, " <> String.joinWith " and " names <> "!" }

-- | Answers each message as it arrives, with the server's heartbeat merged in, until the client
-- | is done talking.
chat :: forall ctx. Strom ctx () ChatMessage -> Strom ctx () ChatMessage
chat incoming = Strom.mergeHaltLeft answers heartbeats
  where
  answers = incoming # map answer
  heartbeats = Strom.every (Milliseconds 150.0)
    # map heartbeat
    # Strom.onExit (sayHowItEnded "the heartbeat")
  answer { from, text } = { from: "server", text: from <> " said " <> show (String.length text) <> " characters: " <> String.toUpper text }
  heartbeat tick = { from: "server", text: "heartbeat " <> show (tick + 1) }

sayHowItEnded :: forall ctx err. String -> Exit err -> OmIO ctx err Unit
sayHowItEnded name exit = log (name <> " " <> howItEnded) # liftIO
  where
  howItEnded = case exit of
    Ended -> "ended"
    Failed _ -> "failed"
    Interrupted -> "was interrupted, and its finalizer ran"

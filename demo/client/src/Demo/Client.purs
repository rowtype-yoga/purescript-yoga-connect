-- | The demo's gRPC client: calls each method of the server once, chats a little, and shows
-- | how a rejected call comes back.
module Demo.Client where

import Prelude

import Control.Monad.Error.Class (try)
import Data.Either (Either(..))
import Data.Time.Duration (Milliseconds(..))
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Exception (Error, message)
import Go.IO (launchIO_, threadDelay)
import Go.IO.Console (log)
import Demo.Proto.Greeter.V1.Greeter (greeter)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Client (Connection, client, connect)
import Yoga.Connect.Status (ConnectError)
import Yoga.Connect.Strom as Strom
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)

main :: Effect Unit
main = launchIO_ do
  connection <- connect "http://127.0.0.1:50051" # liftEffect
  outcome <- runOmIO unit (talkTo connection)
  case outcome of
    Left failure -> log ("the demo failed: " <> describe failure)
    Right _ -> log "done"

talkTo :: Connection -> OmIO Unit (connectError :: ConnectError) Unit
talkTo connection = do
  let greeterClient = client greeter connection

  say "unary"
  reply <- greeterClient.sayHello { name: "Ada" }
  say reply.message

  say "server stream"
  greeterClient.listGreetings { name: "Ada", count: 3 }
    # Strom.forEach \{ index, text } -> say (show index <> ". " <> text)

  say "client stream"
  everyone <- greeterClient.greetEveryone (Strom.fromArray [ { name: "Ada" }, { name: "Grace" }, { name: "Barbara" } ])
  say everyone.message

  say "bidi stream"
  greeterClient.chat chatter
    # Strom.forEach \{ from, text } -> say (from <> ": " <> text)

  say "rejected"
  rejected <- try (greeterClient.sayHello { name: "" })
  case rejected of
    Left failure -> say (describe failure)
    Right answered -> say ("unexpectedly answered " <> answered.message)
  where
  chatter = Strom.fromArray [ "hi there", "how are the streams?", "bye" ]
    # Strom.mapM \text -> do
        threadDelay (Milliseconds 200.0) # liftIO
        say ("client: " <> text)
        pure { from: "client", text }

say :: forall ctx err. String -> OmIO ctx err Unit
say line = log line # liftIO

describe :: Variant (exception :: Error, connectError :: ConnectError) -> String
describe = Variant.on (Proxy :: Proxy "connectError") show (Variant.on (Proxy :: Proxy "exception") message Variant.case_)

module Test.Yoga.Connect.CallSpec where

import Prelude

import Control.Monad.Error.Class (try)
import Control.Monad.Reader (ask)
import Data.Array as Array
import Data.Bifunctor (lmap)
import Data.Either (Either(..), either)
import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested ((/\))
import Data.Variant as Variant
import Demo.Proto.Greeter.V1.Greeter (ChatMessage, Greeting, HelloReply, HelloRequest, ListRequest, greeter)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Exception (message)
import Go.IO (threadDelay, toAff)
import Go.IO.IORef (IORef, atomicModifyIORef', newIORef, readIORef)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual, shouldSatisfy)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Client (Connection, client, connect)
import Yoga.Connect.Server (Host(..), Port(..), createServer, listen, reject, serveWith)
import Yoga.Connect.Status (Code(..), ConnectError(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom (Exit(..), Strom)
import Yoga.OmIO.Strom as Strom

spec :: Spec Unit
spec = describe "a server and a client derived from the same service row" do
  it "answers each kind of method, and a rejection with its code" do
    { connection } <- startServer
    outcome <- runOmIO unit (callEach connection) # toAff
    either (const Nothing) Just outcome `shouldEqual` Just
      { hello: "Hi, Ada"
      , greetings: [ { index: 1, text: "Ada" }, { index: 2, text: "Ada" } ]
      , everyone: "Ada, Grace, "
      , echoes: [ "one", "two" ]
      , rejected: Just (ConnectError { code: NotFound, message: "no one by that name" })
      }

  it "cancels a server stream once the client has taken enough, which interrupts the server's stream" do
    { connection, log } <- startServer
    taken <- callWith connection \greeterClient ->
      greeterClient.listGreetings { name: "endless", count: 1000000 }
        # Strom.take 3
        # map _.index
        # Strom.runCollect
    taken `shouldEqual` Right [ 1, 2, 3 ]
    entriesAfterAWhile log >>= shouldEqual [ "listGreetings interrupted" ]

  it "sends a timeout as the call's deadline, which interrupts the server's stream" do
    { connection, log } <- startServer
    received <- callWith connection \greeterClient ->
      greeterClient.listGreetings { name: "endless", count: 1000000 }
        # Strom.timeout (Milliseconds 30.0)
        # Strom.runCollect
    map Array.length received `shouldSatisfy` either (const false) (_ > 0)
    entriesAfterAWhile log >>= shouldEqual [ "listGreetings interrupted" ]

  it "ends a server stream the client reads to its end" do
    { connection, log } <- startServer
    _ <- callWith connection \greeterClient -> greeterClient.listGreetings { name: "Ada", count: 2 } # Strom.runDrain
    entriesAfterAWhile log >>= shouldEqual [ "listGreetings ended" ]

  it "chats with a heartbeat merged in, and stops the heartbeat once the client is done" do
    { connection, log } <- startServer
    chatted <- callWith connection \greeterClient ->
      greeterClient.chat slowChatter
        # map _.from
        # Strom.runCollect
    entriesAfterAWhile log >>= shouldEqual [ "heartbeats interrupted" ]
    map (Array.filter (_ == "echo")) chatted `shouldEqual` Right [ "echo", "echo" ]
    map (Array.length <<< Array.filter (_ == "heartbeat")) chatted `shouldSatisfy` either (const false) (_ >= 2)
  where
  slowChatter = Strom.every (Milliseconds 40.0) # Strom.take 2 # map \tick -> { from: "spec", text: show tick }

type GreeterClient =
  ( sayHello :: HelloRequest -> Calling HelloReply
  , listGreetings :: ListRequest -> Streaming Greeting
  , greetEveryone :: Streaming HelloRequest -> Calling HelloReply
  , chat :: Streaming ChatMessage -> Streaming ChatMessage
  )

type Calling = OmIO Unit (connectError :: ConnectError)

type Streaming = Strom Unit (connectError :: ConnectError)

-- | Runs the calls with the service's client, and how they failed, if they did.
callWith
  :: forall a
   . Connection
  -> ({ | GreeterClient } -> OmIO Unit (connectError :: ConnectError) a)
  -> Aff (Either String a)
callWith connection calls = do
  outcome <- runOmIO unit (calls (client greeter connection)) # toAff
  pure (lmap describe outcome)
  where
  describe = Variant.case_
    # Variant.on (Proxy :: Proxy "connectError") show
    # Variant.on (Proxy :: Proxy "exception") message

type Log = IORef (Array String)

-- | A server with the specs' handlers on a free port, the connection to it, and what its
-- | streams' finalizers have logged.
startServer :: Aff { connection :: Connection, log :: Log }
startServer = toAff do
  log <- newIORef []
  server <- createServer # liftEffect
  serveWith { greeting: "Hi", log } greeter { sayHello, listGreetings, greetEveryone, chat: heartbeatChat } server # liftEffect
  bound <- listen { host: Host "127.0.0.1", port: Port 0 } server
  connection <- connect ("http://127.0.0.1:" <> show (portNumber bound.port)) # liftEffect
  pure { connection, log }
  where
  portNumber (Port number) = number

-- | What the server's finalizers logged, once the server has had a moment to notice.
entriesAfterAWhile :: Log -> Aff (Array String)
entriesAfterAWhile log = toAff do
  threadDelay (Milliseconds 50.0)
  readIORef log

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
  greetings <- greeterClient.listGreetings { name: "Ada", count: 2 } # Strom.runCollect
  everyone <- greeterClient.greetEveryone (Strom.fromArray [ { name: "Ada" }, { name: "Grace" } ])
  echoes <- greeterClient.chat (Strom.fromArray [ "one", "two" ] # map { from: "spec", text: _ })
    # Strom.filter (\{ from } -> from == "echo")
    # map _.text
    # Strom.runCollect
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

type Context = { greeting :: String, log :: Log }

sayHello :: HelloRequest -> OmIO Context (notFound :: String) HelloReply
sayHello { name } = do
  { greeting } <- ask
  if name == "nobody" then reject @"notFound" "no one by that name"
  else pure { message: greeting <> ", " <> name }

-- | `count` greetings, one each millisecond.
listGreetings :: ListRequest -> Strom Context () Greeting
listGreetings { name, count } = Strom.every (Milliseconds 1.0)
  # Strom.take count
  # map (\tick -> { index: tick + 1, text: name })
  # Strom.onExit (logExit "listGreetings")

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # Strom.runFold "" (\names person -> names <> person.name <> ", ")
  pure { message: names }

-- | Echoes each message, with a heartbeat merged in, until the client is done.
heartbeatChat :: Strom Context () ChatMessage -> Strom Context () ChatMessage
heartbeatChat incoming = Strom.mergeHaltLeft echoes heartbeats
  where
  echoes = incoming # map \{ text } -> { from: "echo", text }
  heartbeats = Strom.every (Milliseconds 10.0)
    # map (\tick -> { from: "heartbeat", text: show tick })
    # Strom.onExit (logExit "heartbeats")

logExit :: forall err. String -> Exit err -> OmIO Context err Unit
logExit name exit = do
  { log } <- ask
  atomicModifyIORef' log (\entries -> Array.snoc entries (name <> " " <> describeExit exit) /\ unit) # liftIO
  where
  describeExit = case _ of
    Ended -> "ended"
    Failed _ -> "failed"
    Interrupted -> "interrupted"

-- | yoga-connect on JavaScript: `spago build -p yoga-connect-test-js`, then
-- | `bun --eval 'import("./output/Test.JS.Main/index.js").then(m => m.main())'`.
module Test.JS.Main where

import Prelude

import Data.Either (Either(..))
import Data.Generic.Rep (class Generic)
import Data.Show.Generic (genericShow)
import Data.String (Pattern(..), contains)
import Data.Variant (Variant, inj)
import Effect (Effect)
import Effect.Aff (launchAff_)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual, shouldSatisfy)
import Test.Spec.Reporter.Console (consoleReporter)
import Test.Spec.Runner (defaultConfig, runSpec')
import Type.Proxy (Proxy(..))
import Yoga.Connect.Message (type (:=), Message, Named)
import Yoga.Connect.Proto (printProto)
import Yoga.Connect.Protobuf (decode, encode)
import Yoga.Connect.Schema (BidiStream, ServerStream, Service)
import Yoga.Connect.Status (Code(..), ConnectError(..), rejection)
import Yoga.Protobuf.Bytes as Bytes

main :: Effect Unit
main = launchAff_ do
  runSpec' defaultConfig [ consoleReporter ] spec

spec :: Spec Unit
spec = describe "yoga-connect on JavaScript" do
  describe "an enum field" do
    it "goes over the wire as its constructor's number and comes back" do
      Bytes.toArray (encode @Bubble { text: "hi", voice: Fading }) `shouldEqual` [ 10, 2, 104, 105, 16, 2 ]
      decode @Bubble (encode @Bubble { text: "hi", voice: Said }) `shouldEqual` Right { text: "hi", voice: Said }

    it "is its first constructor when it isn't there" do
      decode @Bubble Bytes.empty `shouldEqual` Right { text: "", voice: Typing }

  describe "a handler's rejection" do
    it "is the code its label names, with the message" do
      rejection notFound `shouldEqual` ConnectError { code: NotFound, message: "no such symbol", details: [] }

  describe "the .proto" do
    it "names the messages of a method by their package, so the method's own name doesn't hide them" do
      tickerProto `shouldSatisfy` contains (Pattern "rpc Trade(stream ticker.v1.Trade) returns (stream ticker.v1.Fill);")
  where
  notFound :: Variant (notFound :: String, failedPrecondition :: String)
  notFound = inj (Proxy :: Proxy "notFound") "no such symbol"

type Bubble = Message { text :: 1 := String, voice :: 2 := Voice }

data Voice = Typing | Said | Fading

derive instance Generic Voice _
derive instance Eq Voice

instance Show Voice where
  show = genericShow

-- | protoc reads it as it is.
tickerProto :: String
tickerProto = printProto @Ticker

type Ticker = Service "ticker.v1.Ticker"
  { watch :: ServerStream (Message { symbols :: 1 := Array String }) (Message { symbol :: 1 := String, price :: 2 := Number }) (notFound :: String)
  , trade :: BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)
  }

type Trade = Named "Trade" (Message { symbol :: 1 := String, quantity :: 2 := Int })

type Fill = Named "Fill" (Message { symbol :: 1 := String, quantity :: 2 := Int, price :: 3 := Number })

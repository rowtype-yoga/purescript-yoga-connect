-- | Watches the ticker for two seconds, then trades a little, selling more than it holds at the end.
module Ticker.Client where

import Prelude

import Control.Monad.Error.Class (try)
import Data.Either (either)
import Data.Monoid (power)
import Data.Number.Format (fixed, toStringWith)
import Data.Ord (abs)
import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Exception (Error, message)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Ticker.Service (Ticker)
import Yoga.Connect.Om (Connection, ConnectError(..), OmIO, client, codeName, connect, liftIO, runOmIO)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = launchIO_ do
  connection <- connect "http://127.0.0.1:50051" # liftEffect
  outcome <- runOmIO unit (watchThenTrade connection)
  log (either describeFailure (const "done") outcome)

-- | How a call can fail: as the service says its methods do, or as any call can.
type Failures = (notFound :: String, failedPrecondition :: String, connectError :: ConnectError)

watchThenTrade :: Connection -> OmIO Unit Failures Unit
watchThenTrade connection = do
  let ticker = client @Ticker connection
  say "watching ACME, GLOBEX and INITECH for two seconds"
  ticker.watch { symbols: [ "ACME", "GLOBEX", "INITECH" ] }
    # Strom.timeout (Milliseconds 2000.0)
    # Strom.forEach (say <<< describeQuote)
  say "trading"
  traded <- try do
    ticker.trade (oneEvery (Milliseconds 300.0) trades)
      # Strom.forEach (say <<< describeFill)
  traded # either (say <<< ("refused: " <> _) <<< describeFailure) pure
  where
  trades =
    [ buy 10 "ACME" 121.40
    , buy 25 "INITECH" 9.80
    , sell 5 "ACME" 122.10
    , sell 40 "INITECH" 9.90
    ]

type Trade = { symbol :: String, quantity :: Int, limit :: Number }

-- | The trades, one at a time, each said out loud as it goes.
oneEvery :: forall ctx err. Milliseconds -> Array Trade -> Strom.Strom ctx err Trade
oneEvery interval trades = Strom.zipWith const (Strom.fromArray trades) (Strom.every interval)
  # Strom.tap (say <<< describeTrade)

buy :: Int -> String -> Number -> Trade
buy quantity symbol limit = { symbol, quantity, limit }

sell :: Int -> String -> Number -> Trade
sell quantity symbol limit = { symbol, quantity: negate quantity, limit }

describeQuote :: { symbol :: String, price :: Number, average :: Number } -> String
describeQuote { symbol, price, average } = padded 8 symbol <> column price <> trend <> "  avg" <> column average
  where
  column amount = leftPadded 8 (money amount)
  trend = if price >= average then "  ▲" else "  ▼"

describeTrade :: Trade -> String
describeTrade { symbol, quantity, limit } =
  "→ " <> (if quantity > 0 then "buy " else "sell ") <> shares quantity symbol <> " at " <> money limit

describeFill :: { symbol :: String, quantity :: Int, price :: Number, position :: Int } -> String
describeFill { symbol, quantity, price, position } =
  "← " <> (if quantity > 0 then "bought " else "sold ") <> shares quantity symbol <> " at " <> money price <> ", holding " <> show position

shares :: Int -> String -> String
shares quantity symbol = show (abs quantity) <> " " <> symbol

describeFailure :: Variant (exception :: Error | Failures) -> String
describeFailure = Variant.match
  { notFound: \why -> "not_found: " <> why
  , failedPrecondition: \why -> "failed_precondition: " <> why
  , connectError: \(ConnectError { code, message: why }) -> codeName code <> ": " <> why
  , exception: message
  }

money :: Number -> String
money amount = toStringWith (fixed 2) amount

padded :: Int -> String -> String
padded width text = text <> power " " (width - String.length text)

leftPadded :: Int -> String -> String
leftPadded width text = power " " (width - String.length text) <> text

say :: forall ctx err. String -> OmIO ctx err Unit
say line = log line # liftIO

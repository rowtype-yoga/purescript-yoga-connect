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
import Ticker.Proto.Ticker.V1.Ticker (Fill, Quote, Trade, ticker)
import Yoga.Connect.Client (Connection, client, connect)
import Yoga.Connect.Status (ConnectError(..), codeName)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = launchIO_ do
  connection <- connect "http://127.0.0.1:50051" # liftEffect
  outcome <- runOmIO unit (watchThenTrade connection)
  log (either describeFailure (const "done") outcome)

watchThenTrade :: Connection -> OmIO Unit (connectError :: ConnectError) Unit
watchThenTrade connection = do
  let tickerClient = client ticker connection
  say "watching ACME, GLOBEX and INITECH for two seconds"
  tickerClient.watch { symbols: [ "ACME", "GLOBEX", "INITECH" ] }
    # Strom.timeout (Milliseconds 2000.0)
    # Strom.forEach (say <<< describeQuote)
  say "trading"
  traded <- try do
    tickerClient.trade (oneEvery (Milliseconds 300.0) trades)
      # Strom.forEach (say <<< describeFill)
  traded # either (say <<< ("refused: " <> _) <<< describeFailure) pure
  where
  trades =
    [ buy 10 "ACME" 121.40
    , buy 25 "INITECH" 9.80
    , sell 5 "ACME" 122.10
    , sell 40 "INITECH" 9.90
    ]

-- | The trades, one at a time, each said out loud as it goes.
oneEvery :: forall ctx err. Milliseconds -> Array Trade -> Strom.Strom ctx err Trade
oneEvery interval trades = Strom.zipWith const (Strom.fromArray trades) (Strom.every interval)
  # Strom.tap (say <<< describeTrade)

buy :: Int -> String -> Number -> Trade
buy quantity symbol limit = { symbol, quantity, limit }

sell :: Int -> String -> Number -> Trade
sell quantity symbol limit = { symbol, quantity: negate quantity, limit }

describeQuote :: Quote -> String
describeQuote { symbol, price, average } = padded 8 symbol <> column price <> trend <> "  avg" <> column average
  where
  column amount = leftPadded 8 (money amount)
  trend = if price >= average then "  ▲" else "  ▼"

describeTrade :: Trade -> String
describeTrade { symbol, quantity, limit } =
  "→ " <> (if quantity > 0 then "buy " else "sell ") <> shares quantity symbol <> " at " <> money limit

describeFill :: Fill -> String
describeFill { symbol, quantity, price, position } =
  "← " <> (if quantity > 0 then "bought " else "sold ") <> shares quantity symbol <> " at " <> money price <> ", holding " <> show position

shares :: Int -> String -> String
shares quantity symbol = show (abs quantity) <> " " <> symbol

describeFailure :: Variant (exception :: Error, connectError :: ConnectError) -> String
describeFailure = Variant.match
  { connectError: \(ConnectError { code, message: why }) -> codeName code <> ": " <> why
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

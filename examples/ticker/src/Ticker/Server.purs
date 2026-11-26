-- | A live stock ticker over gRPC. `watch` merges a random walk per symbol into one stream of
-- | quotes; `trade` fills trades as they arrive, until one can't go through.
module Ticker.Server where

import Prelude

import Data.Array as Array
import Data.Int (toNumber)
import Data.Int.Bits ((.&.))
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe, fromMaybe, isNothing, maybe)
import Data.Number as Number
import Data.String (joinWith)
import Data.Time.Duration (Milliseconds(..))
import Data.Traversable (traverse)
import Data.Tuple (snd)
import Data.Tuple.Nested ((/\))
import Effect (Effect)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Go.IO.IORef (IORef, modifyIORef, newIORef, readIORef)
import Ticker.Service (Ticker, TradeRejections, Trading, Watching)
import Yoga.Connect.Om (Handler, Host(..), OmIO, Port(..), ask, createServer, handle, liftIO, listen, registerServiceWith, reject)
import Yoga.OmIO.Strom (Exit(..), Strom)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = do
  server <- createServer
  registerServiceWith @Ticker { market } { watch, trade } server
  launchIO_ do
    bound <- listen { host: Host "127.0.0.1", port: Port 50051 } server
    log ("ticker listening on " <> bound.address <> ":" <> show bound.port)

-- | The quotes of the symbols asked for (all of them if none are), merged as they come.
watch :: Handler Watching (market :: Market)
watch = handle do
  { request: { symbols }, market } <- ask
  let wanted = if Array.null symbols then market <#> _.symbol else symbols
  listings <- traverse (listingIn market) wanted # maybe (reject @"notFound" ("we list " <> symbolsIn market)) pure
  pure (Strom.mergeAll (listings <#> quotesOf) # Strom.onExit (sayHowItEnded ("watching " <> joinWith ", " wanted)))

-- | Fills each trade at its limit. Selling more than you hold ends the session.
trade :: Handler Trading (market :: Market)
trade = handle do
  { requests: trades, market } <- ask
  pure (Strom.bracketExit openBook closeBook \book -> trades # Strom.mapM (fill market book))
  where
  openBook = newIORef Map.empty # liftIO
  closeBook book exit = do
    positions <- readIORef book # liftIO
    sayHowItEnded ("trading, holding " <> describePositions positions) exit

type Book = IORef (Map String Int)

-- | What `trade` gets and answers with: the plain records of `Trade` and `Fill`.
type Order = { symbol :: String, quantity :: Int, limit :: Number }

type Filled = { symbol :: String, quantity :: Int, price :: Number, position :: Int }

fill :: forall ctx. Market -> Book -> Order -> OmIO ctx TradeRejections Filled
fill market book { symbol, quantity, limit } = do
  when (isNothing (listingIn market symbol)) do
    reject @"notFound" (symbol <> " isn't listed")
  held <- holding symbol book
  let position = held + quantity
  when (position < 0) do
    reject @"failedPrecondition" ("you hold " <> show held <> " " <> symbol <> ", so you can't sell " <> show (negate quantity))
  modifyIORef book (Map.insert symbol position) # liftIO
  pure { symbol, quantity, price: limit, position }

holding :: forall ctx err. String -> Book -> OmIO ctx err Int
holding symbol book = readIORef book <#> (Map.lookup symbol >>> fromMaybe 0) # liftIO

-- | A symbol's random walk, a step every `pace`, with its exponential moving average.
quotesOf :: forall ctx err. Listing -> Strom ctx err { symbol :: String, price :: Number, average :: Number }
quotesOf { symbol, opening, volatility, pace, seed } = Strom.every pace
  # Strom.scan (\walk _ -> step walk) (Seed seed /\ opening)
  # map snd
  # Strom.scan (\(_ /\ average) price -> price /\ smooth average price) (opening /\ opening)
  # map \(price /\ average) -> { symbol, price: toCents price, average: toCents average }
  where
  step (Seed current /\ price) = Seed (nextSeed current) /\ price * (1.0 + volatility * wobble current)
  smooth average price = average + 0.2 * (price - average)

type Market = Array Listing

type Listing = { symbol :: String, opening :: Number, volatility :: Number, pace :: Milliseconds, seed :: Int }

market :: Market
market =
  [ { symbol: "ACME", opening: 121.50, volatility: 0.004, pace: Milliseconds 250.0, seed: 7 }
  , { symbol: "GLOBEX", opening: 48.20, volatility: 0.007, pace: Milliseconds 400.0, seed: 42 }
  , { symbol: "INITECH", opening: 9.75, volatility: 0.012, pace: Milliseconds 650.0, seed: 1999 }
  , { symbol: "UMBRELLA", opening: 302.10, volatility: 0.003, pace: Milliseconds 900.0, seed: 2026 }
  ]

listingIn :: Market -> String -> Maybe Listing
listingIn listings symbol = Array.find (_.symbol >>> eq symbol) listings

symbolsIn :: Market -> String
symbolsIn listings = joinWith ", " (listings <#> _.symbol)

-- | A linear congruential generator, which is all the randomness a ticker needs.
newtype Seed = Seed Int

nextSeed :: Int -> Int
nextSeed seed = (seed * 1103515245 + 12345) .&. 2147483647

-- | Between -1 and 1.
wobble :: Int -> Number
wobble seed = toNumber (seed `mod` 2001 - 1000) / 1000.0

toCents :: Number -> Number
toCents amount = Number.round (amount * 100.0) / 100.0

describePositions :: Map String Int -> String
describePositions positions
  | Map.isEmpty positions = "nothing"
  | otherwise = joinWith ", " (Map.toUnfoldable positions <#> \(symbol /\ shares) -> show shares <> " " <> symbol)

sayHowItEnded :: forall ctx err. String -> Exit err -> OmIO ctx err Unit
sayHowItEnded what exit = log (what <> ": " <> howItEnded) # liftIO
  where
  howItEnded = case exit of
    Ended -> "ended"
    Failed _ -> "failed, and its finalizer ran"
    Interrupted -> "cancelled, and its finalizer ran"

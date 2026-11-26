-- | The ticker's service, which the server, the client and the `.proto` all come from.
module Ticker.Service where

import Prelude

import Effect (Effect)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Yoga.Connect.Om (type (:=), BidiStream, Message, ServerStream, Service, printProto)

type Ticker = Service "ticker.v1.Ticker"
  { watch :: Watching
  , trade :: Trading
  }

-- | The quotes of every symbol, merged, as they come.
type Watching = ServerStream Subscribe Quote (notFound :: String)

-- | Fills each trade as it arrives; a trade that can't go through ends the session.
type Trading = BidiStream Trade Fill TradeRejections

type TradeRejections = (notFound :: String, failedPrecondition :: String)

-- | The symbols to watch, all of them if none.
type Subscribe = Message { symbols :: 1 := Array String }

-- | A symbol's latest price, and its moving average.
type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }

-- | Buys (a positive quantity) or sells (a negative one) at the limit price.
type Trade = Message { symbol :: 1 := String, quantity :: 2 := Int, limit :: 3 := Number }

-- | A trade that went through, and the position it left.
type Fill = Message { symbol :: 1 := String, quantity :: 2 := Int, price :: 3 := Number, position :: 4 := Int }

-- | Prints the `.proto`, for grpcurl and other languages.
main :: Effect Unit
main = launchIO_ (log (printProto @Ticker))

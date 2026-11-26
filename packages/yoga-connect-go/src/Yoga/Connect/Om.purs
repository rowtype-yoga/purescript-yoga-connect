-- | gRPC services as PureScript types, served and called on Go.
-- |
-- | ```purescript
-- | type Ticker = Service "ticker.v1.Ticker"
-- |   { watch :: ServerStream Subscribe Quote (notFound :: String)
-- |   , trade :: BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)
-- |   }
-- |
-- | type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }
-- |
-- | main = do
-- |   server <- createServer
-- |   registerServiceWith @Ticker { market } { watch, trade } server
-- |   launchIO_ (listen { host: Host "127.0.0.1", port: Port 50051 } server)
-- | ```
-- |
-- | Handlers are `handle do … ask …`, see `Yoga.Connect.Om.Handler`; `client @Ticker connection`
-- | calls them, see `Yoga.Connect.Om.Client`; `printProto @Ticker` is the `.proto`.
module Yoga.Connect.Om
  ( module Yoga.Connect.Message
  , module Yoga.Connect.Schema
  , module Yoga.Connect.Proto
  , module Yoga.Connect.Protobuf
  , module Yoga.Connect.Om.Handler
  , module Yoga.Connect.Om.Client
  , module Yoga.Connect.Server
  , module Yoga.Connect.Client
  , module Yoga.Connect.Status
  , module Yoga.Go.HTTP.OmIO
  , module Control.Monad.Reader
  ) where

import Control.Monad.Reader (ask)
import Yoga.Connect.Client (Connection, connect)
import Yoga.Connect.Message (type (:=), Message, Named, OneOf)
import Yoga.Connect.Om.Client (client, clientWithin)
import Yoga.Connect.Om.Handler (Handler, handle, registerService, registerServiceWith, withoutRequests)
import Yoga.Connect.Proto (printProto)
import Yoga.Connect.Protobuf (codecOf, decode, encode)
import Yoga.Connect.Schema (BidiStream, ClientStream, ServerStream, Service, Unary)
import Yoga.Connect.Server (BoundAddress, Host(..), Port(..), Server, createServer, listen, reject)
import Yoga.Connect.Status (Code(..), ConnectError(..), codeName)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)

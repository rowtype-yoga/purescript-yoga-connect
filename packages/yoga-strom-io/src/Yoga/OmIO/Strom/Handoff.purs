-- | How goroutines of one stream pass chunks to each other, and how many of them may work at once,
-- | waiting no longer than their context lets them: go-io's channels, in a select with the
-- | context's `done`.
module Yoga.OmIO.Strom.Handoff
  ( Handoff
  , newHandoff
  , offer
  , receive
  , Permits
  , newPermits
  , acquire
  , release
  , acquireMany
  , releaseMany
  , awaitAllReleased
  ) where

import Prelude

import Control.Monad.Rec.Class (Step(..), tailRecM)
import Data.Maybe (Maybe(..))
import Go.IO (IO)
import Go.IO.Chan (Chan, capacity, newBufferedChan, send)
import Go.IO.Chan as Chan
import Go.IO.Context (Context, isCancelled)
import Go.IO.Select (onDone, onReceive, onSend, select2)

-- | A Go channel. With room for none, each value waits until it's taken.
type Handoff a = Chan a

-- | Room for `room` values that nobody has taken yet.
newHandoff :: forall a. Int -> IO (Handoff a)
newHandoff room = newBufferedChan room

-- | Hands the value over, waiting while there's no room. `false` once the context is done, even
-- | if there is room.
offer :: forall a. Context -> Handoff a -> a -> IO Boolean
offer context handoff value = do
  cancelled <- isCancelled context
  if cancelled then pure false
  else select2
    (onSend handoff value (pure true))
    (onDone context (pure false))

-- | The next value, waiting until there is one. `Nothing` once the context is done.
receive :: forall a. Context -> Handoff a -> IO (Maybe a)
receive context handoff = select2
  (onReceive handoff pure)
  (onDone context (pure Nothing))

-- | A token for each permit taken, in a channel with room for them all.
type Permits = Chan Unit

newPermits :: Int -> IO Permits
newPermits count = newBufferedChan count

-- | Takes a permit, waiting while all are taken. `false` once the context is done.
acquire :: Context -> Permits -> IO Boolean
acquire context permits = offer context permits unit

release :: Permits -> IO Unit
release permits = void (Chan.receive permits)

-- | Takes `count` permits, one after the other.
acquireMany :: Context -> Permits -> Int -> IO Boolean
acquireMany context permits count = tailRecM acquireRemaining count
  where
  acquireRemaining remaining
    | remaining <= 0 = pure (Done true)
    | otherwise = acquire context permits <#> if _ then Loop (remaining - 1) else Done false

releaseMany :: Permits -> Int -> IO Unit
releaseMany permits count = times count (release permits)

-- | Waits until every permit is back: until every worker is done.
awaitAllReleased :: Permits -> IO Unit
awaitAllReleased permits = do
  times (capacity permits) (send permits unit)
  times (capacity permits) (release permits)

times :: Int -> IO Unit -> IO Unit
times count action = tailRecM once count
  where
  once remaining
    | remaining <= 0 = pure (Done unit)
    | otherwise = action $> Loop (remaining - 1)

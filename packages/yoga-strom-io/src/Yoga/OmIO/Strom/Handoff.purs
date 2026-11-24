-- | How goroutines of one stream pass chunks to each other, and how many of them may work at once,
-- | waiting no longer than their context lets them.
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

import Data.Maybe (Maybe(..))
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, EffectFn4, runEffectFn1, runEffectFn2, runEffectFn3, runEffectFn4)
import Go.IO.Internal (IO(..))
import Yoga.OmIO.Strom.Context (Context)

-- | A Go channel. With room for none, each value waits until it's taken.
foreign import data Handoff :: Type -> Type

-- | Room for `capacity` values that nobody has taken yet.
newHandoff :: forall a. Int -> IO (Handoff a)
newHandoff capacity = IO (runEffectFn1 newHandoffImpl capacity)

foreign import newHandoffImpl :: forall a. EffectFn1 Int (Handoff a)

-- | Hands the value over, waiting while there's no room. `false` once the context is done.
offer :: forall a. Context -> Handoff a -> a -> IO Boolean
offer context handoff value = IO (runEffectFn3 offerImpl context handoff value)

foreign import offerImpl :: forall a. EffectFn3 Context (Handoff a) a Boolean

-- | The next value, waiting until there is one. `Nothing` once the context is done.
receive :: forall a. Context -> Handoff a -> IO (Maybe a)
receive context handoff = IO (runEffectFn4 receiveImpl Nothing Just context handoff)

foreign import receiveImpl :: forall a r. EffectFn4 r (a -> r) Context (Handoff a) r

-- | A number of permits, for as many workers at once.
foreign import data Permits :: Type

newPermits :: Int -> IO Permits
newPermits count = IO (runEffectFn1 newPermitsImpl count)

foreign import newPermitsImpl :: EffectFn1 Int Permits

-- | Takes a permit, waiting while all are taken. `false` once the context is done.
acquire :: Context -> Permits -> IO Boolean
acquire context permits = IO (runEffectFn2 acquireImpl context permits)

foreign import acquireImpl :: EffectFn2 Context Permits Boolean

release :: Permits -> IO Unit
release permits = IO (runEffectFn1 releaseImpl permits)

foreign import releaseImpl :: EffectFn1 Permits Unit

-- | Takes `count` permits, one after the other.
acquireMany :: Context -> Permits -> Int -> IO Boolean
acquireMany context permits count = IO (runEffectFn3 acquireManyImpl context permits count)

foreign import acquireManyImpl :: EffectFn3 Context Permits Int Boolean

releaseMany :: Permits -> Int -> IO Unit
releaseMany permits count = IO (runEffectFn2 releaseManyImpl permits count)

foreign import releaseManyImpl :: EffectFn2 Permits Int Unit

-- | Waits until every permit is back: until every worker is done.
awaitAllReleased :: Permits -> IO Unit
awaitAllReleased permits = IO (runEffectFn1 awaitAllReleasedImpl permits)

foreign import awaitAllReleasedImpl :: EffectFn1 Permits Unit

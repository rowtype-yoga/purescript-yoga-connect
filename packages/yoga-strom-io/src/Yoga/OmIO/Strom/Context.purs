-- | Go's `context.Context`: how a stream learns that it should stop, because whoever reads it
-- | is done, a deadline passed or a gRPC client hung up.
module Yoga.OmIO.Strom.Context
  ( Context
  , CancelContext
  , background
  , withCancel
  , withTimeout
  , cancel
  , isDone
  , deadlinePassed
  , sleepUnlessDone
  ) where

import Prelude

import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested (type (/\), (/\))
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, runEffectFn1, runEffectFn2, runEffectFn3)
import Go.IO.Internal (IO(..))

-- | Go's `context.Context`.
foreign import data Context :: Type

-- | Go's `context.CancelFunc`.
foreign import data CancelContext :: Type

-- | The context that is never done.
foreign import background :: Context

-- | A context that is done once `cancel` is called on it, or once its parent is.
withCancel :: Context -> IO (Context /\ CancelContext)
withCancel parent = IO (runEffectFn2 withCancelImpl (/\) parent)

foreign import withCancelImpl :: EffectFn2 (Context -> CancelContext -> Context /\ CancelContext) Context (Context /\ CancelContext)

-- | `withCancel`, and done by itself after the duration.
withTimeout :: Milliseconds -> Context -> IO (Context /\ CancelContext)
withTimeout (Milliseconds duration) parent = IO (runEffectFn3 withTimeoutImpl (/\) duration parent)

foreign import withTimeoutImpl :: EffectFn3 (Context -> CancelContext -> Context /\ CancelContext) Number Context (Context /\ CancelContext)

cancel :: CancelContext -> IO Unit
cancel cancelling = IO (runEffectFn1 cancelImpl cancelling)

foreign import cancelImpl :: EffectFn1 CancelContext Unit

isDone :: Context -> IO Boolean
isDone context = IO (runEffectFn1 isDoneImpl context)

foreign import isDoneImpl :: EffectFn1 Context Boolean

-- | Done because its deadline, or one of its parents', passed.
deadlinePassed :: Context -> IO Boolean
deadlinePassed context = IO (runEffectFn1 deadlinePassedImpl context)

foreign import deadlinePassedImpl :: EffectFn1 Context Boolean

-- | Waits the duration, or less if the context is done first. `false` when it is.
sleepUnlessDone :: Context -> Milliseconds -> IO Boolean
sleepUnlessDone context (Milliseconds duration) = IO (runEffectFn2 sleepUnlessDoneImpl context duration)

foreign import sleepUnlessDoneImpl :: EffectFn2 Context Number Boolean

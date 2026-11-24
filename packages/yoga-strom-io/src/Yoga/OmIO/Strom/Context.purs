-- | Go's `context.Context`, from go-io: how a stream learns that it should stop, because whoever
-- | reads it is done, a deadline passed or a gRPC client hung up.
module Yoga.OmIO.Strom.Context
  ( module Exports
  , isDone
  , deadlinePassed
  , sleepUnlessDone
  ) where

import Prelude

import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds)
import Go.IO (IO, threadDelayWith)
import Go.IO.Context (Cancellation(..), Context, err, isCancelled)
import Go.IO.Context (Context, CancelContext, background, withCancel, withTimeout, cancel) as Exports

isDone :: Context -> IO Boolean
isDone context = isCancelled context

-- | Done because its deadline, or one of its parents', passed.
deadlinePassed :: Context -> IO Boolean
deadlinePassed context = err context <#> eq (Just DeadlineExceeded)

-- | Waits the duration, or less if the context is done first. `false` when it is.
sleepUnlessDone :: Context -> Milliseconds -> IO Boolean
sleepUnlessDone context duration = threadDelayWith context duration

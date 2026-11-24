-- | When to do something again, and when to give up: for `repeatOmWith` and `retry`.
-- |
-- | ```purescript
-- | Schedule.exponential (Milliseconds 10.0) <> Schedule.recurs 5
-- | ```
-- |
-- | waits 10, 20, 40, 80 and 160 ms, and gives up after the fifth time. `<>` goes on while both
-- | go on, and waits the longer of the two.
module Yoga.OmIO.Strom.Schedule
  ( Schedule
  , delayBefore
  , forever
  , recurs
  , spaced
  , exponential
  ) where

import Prelude

import Control.Apply (lift2)
import Data.Int as Int
import Data.Maybe (Maybe(..))
import Data.Number (pow)
import Data.Time.Duration (Milliseconds(..))

-- | For the first, second, … repetition, how long to wait before it, or `Nothing` to stop.
newtype Schedule = Schedule (Int -> Maybe Milliseconds)

instance Semigroup Schedule where
  append (Schedule first) (Schedule second) = Schedule \repetition -> lift2 longer (first repetition) (second repetition)
    where
    longer (Milliseconds one) (Milliseconds other) = Milliseconds (max one other)

instance Monoid Schedule where
  mempty = forever

-- | How long to wait before the repetition, counting from 1, or `Nothing` to stop.
delayBefore :: Schedule -> Int -> Maybe Milliseconds
delayBefore (Schedule delayFor) repetition = delayFor repetition

-- | Again and again, right away.
forever :: Schedule
forever = Schedule \_ -> Just (Milliseconds 0.0)

-- | Right away, `times` times.
recurs :: Int -> Schedule
recurs times = Schedule \repetition -> if repetition <= times then Just (Milliseconds 0.0) else Nothing

-- | The same wait each time.
spaced :: Milliseconds -> Schedule
spaced interval = Schedule \_ -> Just interval

-- | `first`, then twice as long each time.
exponential :: Milliseconds -> Schedule
exponential (Milliseconds first) = Schedule \repetition -> Just (Milliseconds (first * pow 2.0 (Int.toNumber (repetition - 1))))

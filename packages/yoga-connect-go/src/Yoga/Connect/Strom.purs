-- | Streams of messages for gRPC handlers and clients on Go: yoga-om-strom's `Strom`, running in
-- | `OmIO` on the call's goroutine instead of in `Aff`.
-- |
-- | ```purescript
-- | listGreetings { name, count } = Strom.range 1 count # map (greeting name)
-- | chat incoming = incoming # Strom.mapM reply
-- | ```
-- |
-- | A `Strom` pushes each element into whoever runs it and waits until that one has taken it, so
-- | a slow consumer slows the producer down: sending a message blocks while HTTP/2 flow control
-- | says the peer can't take more. Nothing is buffered.
module Yoga.Connect.Strom
  ( Strom(..)
  , Sink
  , empty
  , succeed
  , fromArray
  , fromOm
  , range
  , iterate
  , unfold
  , repeatOm
  , every
  , mapM
  , filter
  , tapM
  , scan
  , take
  , takeWhile
  , drop
  , merge
  , concatMap
  , forEach
  , foldl
  , collect
  , drain
  , runStrom
  ) where

import Prelude

import Control.Monad.Error.Class (throwError)
import Control.Monad.Reader (ask)
import Data.Array as Array
import Data.Either (Either(..))
import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds)
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant (Variant)
import Effect.Exception (Error)
import Go.IO (forkIO, threadDelay)
import Go.IO.IORef (newIORef, readIORef, writeIORef)
import Go.IO.MVar (newEmptyMVar, putMVar, takeMVar)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)

-- | Hands each element to the sink for as long as the sink says to go on. `true` when the
-- | stream ran out by itself, `false` when the sink stopped it.
newtype Strom ctx err a = Strom (Sink ctx err a -> OmIO ctx err Boolean)

-- | Takes an element and says whether it wants the next one.
type Sink ctx err a = a -> OmIO ctx err Boolean

instance Functor (Strom ctx err) where
  map f (Strom source) = Strom \sink -> source \element -> sink (f element)

-- | One stream, then the other.
instance Semigroup (Strom ctx err a) where
  append (Strom first) (Strom second) = Strom \sink -> do
    finished <- first sink
    if finished then second sink else pure false

instance Monoid (Strom ctx err a) where
  mempty = empty

-- | Runs the stream into the sink.
runStrom :: forall ctx err a. Sink ctx err a -> Strom ctx err a -> OmIO ctx err Boolean
runStrom sink (Strom source) = source sink

-- Making streams

empty :: forall ctx err a. Strom ctx err a
empty = Strom \_ -> pure true

succeed :: forall ctx err a. a -> Strom ctx err a
succeed element = Strom \sink -> sink element

fromArray :: forall ctx err a. Array a -> Strom ctx err a
fromArray elements = Strom \sink -> pushFrom sink 0
  where
  pushFrom sink index = case Array.index elements index of
    Nothing -> pure true
    Just element -> do
      more <- sink element
      if more then pushFrom sink (index + 1) else pure false

-- | The one element the computation gives.
fromOm :: forall ctx err a. OmIO ctx err a -> Strom ctx err a
fromOm computation = Strom \sink -> computation >>= sink

-- | `from` to `to`, both included.
range :: forall ctx err. Int -> Int -> Strom ctx err Int
range from to = Strom \sink -> countFrom sink from
  where
  countFrom sink number
    | number > to = pure true
    | otherwise = do
        more <- sink number
        if more then countFrom sink (number + 1) else pure false

-- | `start`, `next start`, `next (next start)`, … without end.
iterate :: forall ctx err a. (a -> a) -> a -> Strom ctx err a
iterate next start = Strom \sink -> pushFrom sink start
  where
  pushFrom sink element = do
    more <- sink element
    if more then pushFrom sink (next element) else pure false

-- | Elements from a state, until `step` says there are no more.
unfold :: forall ctx err a state. (state -> Maybe (a /\ state)) -> state -> Strom ctx err a
unfold step start = Strom \sink -> pushFrom sink start
  where
  pushFrom sink state = case step state of
    Nothing -> pure true
    Just (element /\ next) -> do
      more <- sink element
      if more then pushFrom sink next else pure false

-- | The computation's result, again and again.
repeatOm :: forall ctx err a. OmIO ctx err a -> Strom ctx err a
repeatOm computation = Strom \sink -> pushNext sink
  where
  pushNext sink = do
    element <- computation
    more <- sink element
    if more then pushNext sink else pure false

-- | 0, 1, 2, … one each interval, the first after one interval.
every :: forall ctx err. Milliseconds -> Strom ctx err Int
every interval = Strom \sink -> tickFrom sink 0
  where
  tickFrom sink tick = do
    threadDelay interval # liftIO
    more <- sink tick
    if more then tickFrom sink (tick + 1) else pure false

-- Changing streams

mapM :: forall ctx err a b. (a -> OmIO ctx err b) -> Strom ctx err a -> Strom ctx err b
mapM f (Strom source) = Strom \sink -> source \element -> f element >>= sink

filter :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
filter keep (Strom source) = Strom \sink -> source \element ->
  if keep element then sink element else pure true

-- | Runs the computation on each element before passing it on.
tapM :: forall ctx err a. (a -> OmIO ctx err Unit) -> Strom ctx err a -> Strom ctx err a
tapM f (Strom source) = Strom \sink -> source \element -> f element *> sink element

-- | Each intermediate result of a left fold, the start excluded.
scan :: forall ctx err a b. (b -> a -> b) -> b -> Strom ctx err a -> Strom ctx err b
scan step start (Strom source) = Strom \sink -> do
  state <- newIORef start # liftIO
  source \element -> do
    previous <- readIORef state # liftIO
    let next = step previous element
    writeIORef state next # liftIO
    sink next

-- | The first `count` elements. The stream stops right after the last one, without waiting for
-- | another.
take :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
take count (Strom source)
  | count <= 0 = empty
  | otherwise = Strom \sink -> do
      taken <- newIORef 0 # liftIO
      sinkStopped <- newIORef false # liftIO
      _ <- source \element -> do
        before <- readIORef taken # liftIO
        writeIORef taken (before + 1) # liftIO
        more <- sink element
        unless more (writeIORef sinkStopped true # liftIO)
        pure (more && before + 1 < count)
      stopped <- readIORef sinkStopped # liftIO
      pure (not stopped)

-- | Elements while they pass the test; the first that doesn't ends the stream.
takeWhile :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
takeWhile keep (Strom source) = Strom \sink -> do
  sinkStopped <- newIORef false # liftIO
  _ <- source \element ->
    if keep element then do
      more <- sink element
      unless more (writeIORef sinkStopped true # liftIO)
      pure more
    else pure false
  stopped <- readIORef sinkStopped # liftIO
  pure (not stopped)

drop :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
drop count (Strom source) = Strom \sink -> do
  dropped <- newIORef 0 # liftIO
  source \element -> do
    before <- readIORef dropped # liftIO
    if before < count then (writeIORef dropped (before + 1) # liftIO) $> true
    else sink element

-- | For each element, all elements of its stream.
concatMap :: forall ctx err a b. (a -> Strom ctx err b) -> Strom ctx err a -> Strom ctx err b
concatMap streamOf (Strom source) = Strom \sink -> source \element -> runStrom sink (streamOf element)

-- | Both streams at the same time, each on a thread of its own, in the order their elements
-- | come. A thread waits until its element is taken, and once the sink stops, both stop at their
-- | next element. A failure of either fails the whole.
merge :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
merge (Strom left) (Strom right) = Strom \sink -> do
  context <- ask
  handoff <- newEmptyMVar # liftIO
  stopped <- newIORef false # liftIO
  let produce source = runOmIO context (source (offer handoff stopped)) >>= (putMVar handoff <<< Finished)
  _ <- forkIO (produce left) # liftIO
  _ <- forkIO (produce right) # liftIO
  consume sink handoff stopped 2
  where
  offer handoff stopped element = liftIO do
    isStopped <- readIORef stopped
    if isStopped then pure false else putMVar handoff (Element element) $> true

  consume sink handoff stopped running = do
    item <- takeMVar handoff # liftIO
    case item of
      Element element -> do
        more <- sink element
        if more then consume sink handoff stopped running
        else stopBoth handoff stopped running $> false
      Finished (Left failure) -> do
        stopBoth handoff stopped (running - 1)
        throwError failure
      Finished (Right _)
        | running == 1 -> pure true
        | otherwise -> consume sink handoff stopped (running - 1)

  stopBoth handoff stopped running = liftIO do
    writeIORef stopped true
    void (forkIO (awaitFinished handoff running))

  awaitFinished handoff running
    | running <= 0 = pure unit
    | otherwise = takeMVar handoff >>= case _ of
        Element _ -> awaitFinished handoff running
        Finished _ -> awaitFinished handoff (running - 1)

-- | What a producer of `merge` hands over: an element, or how it ended.
data Handoff err a
  = Element a
  | Finished (Either (Variant (exception :: Error | err)) Boolean)

-- Running streams

forEach :: forall ctx err a. (a -> OmIO ctx err Unit) -> Strom ctx err a -> OmIO ctx err Unit
forEach f (Strom source) = void (source \element -> f element $> true)

foldl :: forall ctx err a b. (b -> a -> b) -> b -> Strom ctx err a -> OmIO ctx err b
foldl step start (Strom source) = do
  state <- newIORef start # liftIO
  _ <- source \element -> do
    previous <- readIORef state # liftIO
    writeIORef state (step previous element) # liftIO
    pure true
  readIORef state # liftIO

collect :: forall ctx err a. Strom ctx err a -> OmIO ctx err (Array a)
collect stream = foldl Array.snoc [] stream

drain :: forall ctx err a. Strom ctx err a -> OmIO ctx err Unit
drain stream = forEach (\_ -> pure unit) stream

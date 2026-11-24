-- | Streams in `OmIO` on Go, after ZIO's `ZStream`: pulled in chunks, failing with the errors of
-- | the Om row, cleaning up after themselves, and stopped from outside through Go's
-- | `context.Context`.
-- |
-- | ```purescript
-- | Strom.range 1 1001 # map (_ * 2) # Strom.take 10 # Strom.forEach send
-- | ```
-- |
-- | Running a stream opens it in a `Scope`, which gives a `Pull`. Each pull answers with the next
-- | chunk, or says the stream has ended, or fails. Nothing happens between pulls, so a slow reader
-- | slows the stream down and nothing piles up.
-- |
-- | Once the reader is done (it has what it wants, it failed, or its context was cancelled) the
-- | scope closes: goroutines the stream started stop at their next wait, and finalizers run, with
-- | how the stream ended. `take` closes what it takes from as soon as it has enough, and `merge`,
-- | `buffer` and `mapPar` run their sides on goroutines of their own.
-- |
-- | Names follow yoga-om-strom where it has them: `runFold`, `runCollect`, `runDrain`, `merge`,
-- | `mapPar`, `buffer`, `catchAll`, `ensuring`, `bracket`.
module Yoga.OmIO.Strom
  ( Strom(..)
  , Pull
  , Step(..)
  , Chunk
  , HaltStrategy(..)
  , open
  , fromPull
  , module Exports
  -- Making streams
  , empty
  , succeed
  , fail
  , fromArray
  , fromChunks
  , fromOm
  , fromChan
  , fromHandoff
  , range
  , iterate
  , unfold
  , unfoldM
  , repeat
  , repeatOm
  , repeatOmWith
  , every
  , bracket
  , bracketExit
  , acquireRelease
  , ensuring
  , onExit
  -- Changing streams
  , mapChunks
  , mapM
  , mapIO
  , filter
  , filterMap
  , tap
  , scan
  , take
  , takeWhile
  , takeUntil
  , drop
  , dropWhile
  , concatMap
  , flatMap
  , concat
  , zip
  , zipWith
  , grouped
  , rechunk
  , catchAll
  , mapError
  , orElse
  , retry
  , timeout
  -- At the same time
  , buffer
  , merge
  , mergeAll
  , mergeWith
  , mergeHaltEither
  , mergeHaltLeft
  , mergeHaltRight
  , race
  , raceAll
  , zipPar
  , zipLatest
  , zipLatestWith
  , mapPar
  , mapParUnordered
  , forkDrain
  -- Running streams
  , forEach
  , forEachWithin
  , runFold
  , runFoldM
  , runCollect
  , runDrain
  , runHead
  , runLast
  ) where

import Prelude

import Control.Alt ((<|>))
import Control.Monad.Error.Class (catchError, throwError, try)
import Control.Monad.Reader (ask)
import Data.Array as Array
import Data.Either (Either(..), either)
import Data.FoldableWithIndex (forWithIndex_)
import Data.Maybe (Maybe(..), maybe)
import Data.Time.Duration (Milliseconds(..))
import Data.Traversable (for, for_, traverse)
import Data.TraversableWithIndex (forWithIndex)
import Data.Tuple.Nested (type (/\), (/\))
import Go.IO (IO, finally, forkIO)
import Go.IO.Chan (Chan)
import Go.IO.Select (onDone, onReceive, select2)
import Go.IO.IORef (IORef, newIORef, readIORef, writeIORef)
import Go.IO.MVar (MVar, newEmptyMVar, putMVar, readMVar)
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom.Context (Context)
import Yoga.OmIO.Strom.Handoff (Handoff, Permits, acquire, acquireMany, awaitAllReleased, newHandoff, newPermits, offer, receive, release, releaseMany)
import Yoga.OmIO.Strom.Schedule (Schedule, delayBefore)
import Yoga.OmIO.Strom.Scope (Exit(..), Failure, Scope, addFinalizer, attempt, closeScope, contextOf, deadlinePassedIn, exitAfter, forkScope, forkScopeWithTimeout, interrupt, interrupted, isInterrupted, isInterruption, scoped, scopedWithin, sleep, throwIfInterrupted)
import Yoga.OmIO.Strom.Scope (Scope, Exit(..), Failure, addFinalizer, contextOf, sleep) as Exports

-- | Opened in a scope, a stream gives the pull that reads it.
newtype Strom ctx err a = Strom (Scope ctx err -> OmIO ctx err (Pull ctx err a))

-- | The next chunk, or the end. Once a pull has ended or failed, nobody pulls it again.
type Pull ctx err a = OmIO ctx err (Step a)

data Step a
  = Emit (Chunk a)
  | End

-- | The elements of one pull, together, so that most work is a loop over an array.
type Chunk a = Array a

-- | Which ends of a merge end it: all of them, the first, the left one or the right one.
data HaltStrategy
  = HaltBoth
  | HaltEither
  | HaltLeft
  | HaltRight

derive instance Eq HaltStrategy

open :: forall ctx err a. Scope ctx err -> Strom ctx err a -> OmIO ctx err (Pull ctx err a)
open scope (Strom openIn) = openIn scope

-- | A stream from what opens it: for streams over Go resources, such as a gRPC call.
fromPull :: forall ctx err a. (Scope ctx err -> OmIO ctx err (Pull ctx err a)) -> Strom ctx err a
fromPull = Strom

instance Functor (Strom ctx err) where
  map f stream = mapChunks (map f) stream

instance Apply (Strom ctx err) where
  apply = ap

instance Applicative (Strom ctx err) where
  pure = succeed

-- | Each element's stream, one after the other.
instance Bind (Strom ctx err) where
  bind stream streamOf = concatMap streamOf stream

instance Monad (Strom ctx err)

-- | One stream, then the other.
instance Semigroup (Strom ctx err a) where
  append first second = concat [ first, second ]

instance Monoid (Strom ctx err a) where
  mempty = empty

--------------------------------------------------------------------------------
-- Making streams
--------------------------------------------------------------------------------

-- | How many elements `range` and `repeat` put in a chunk.
defaultChunkSize :: Int
defaultChunkSize = 4096

empty :: forall ctx err a. Strom ctx err a
empty = Strom \_ -> pure (pure End)

succeed :: forall ctx err a. a -> Strom ctx err a
succeed element = fromArray [ element ]

-- | Fails when pulled.
fail :: forall ctx err a. Failure err -> Strom ctx err a
fail failure = Strom \_ -> pure (throwError failure)

-- | The array, as one chunk.
fromArray :: forall ctx err a. Array a -> Strom ctx err a
fromArray elements = fromChunks [ elements ]

fromChunks :: forall ctx err a. Array (Chunk a) -> Strom ctx err a
fromChunks chunks = Strom \_ -> do
  nextIndex <- newIORef 0 # liftIO
  pure do
    index <- readIORef nextIndex # liftIO
    writeIORef nextIndex (index + 1) # liftIO
    pure (maybe End Emit (Array.index chunks index))

-- | The computation's result.
fromOm :: forall ctx err a. OmIO ctx err a -> Strom ctx err a
fromOm computation = Strom \_ -> do
  done <- newIORef false # liftIO
  pure do
    alreadyDone <- readIORef done # liftIO
    if alreadyDone then pure End
    else do
      writeIORef done true # liftIO
      element <- computation
      pure (Emit [ element ])

-- | What is sent on go-io's channel, until it's closed. Waiting for the next value stops when the
-- | stream is interrupted.
fromChan :: forall ctx err a. Chan a -> Strom ctx err a
fromChan channel = Strom \scope -> pure do
  received <- receiveOrClosed (contextOf scope) # liftIO
  case received of
    Nothing -> throwError interrupted
    Just Nothing -> pure End
    Just (Just element) -> pure (Emit [ element ])
  where
  receiveOrClosed context = select2
    (onReceive channel (pure <<< Just))
    (onDone context (pure Nothing))

-- | What is offered to the handoff, without end. Waiting for the next value stops when the stream
-- | is interrupted.
fromHandoff :: forall ctx err a. Handoff a -> Strom ctx err a
fromHandoff handoff = Strom \scope -> pure do
  received <- receive (contextOf scope) handoff # liftIO
  maybe (throwError interrupted) (\element -> pure (Emit [ element ])) received

-- | From `from` up to `to`, without `to`, as yoga-om-strom and ZStream count.
range :: forall ctx err. Int -> Int -> Strom ctx err Int
range from to = Strom \_ -> do
  nextStart <- newIORef from # liftIO
  pure do
    start <- readIORef nextStart # liftIO
    if start >= to then pure End
    else do
      let stop = min to (start + defaultChunkSize)
      writeIORef nextStart stop # liftIO
      pure (Emit (Array.range start (stop - 1)))

-- | `start`, `next start`, `next (next start)`, … without end.
iterate :: forall ctx err a. (a -> a) -> a -> Strom ctx err a
iterate next start = unfold (\element -> Just (element /\ next element)) start

-- | Elements from a state, until `step` says there are no more.
unfold :: forall ctx err a state. (state -> Maybe (a /\ state)) -> state -> Strom ctx err a
unfold step start = unfoldM (pure <<< step) start

unfoldM :: forall ctx err a state. (state -> OmIO ctx err (Maybe (a /\ state))) -> state -> Strom ctx err a
unfoldM step start = Strom \_ -> do
  current <- newIORef start # liftIO
  pure do
    state <- readIORef current # liftIO
    stepped <- step state
    case stepped of
      Nothing -> pure End
      Just (element /\ next) -> do
        writeIORef current next # liftIO
        pure (Emit [ element ])

-- | The element, without end.
repeat :: forall ctx err a. a -> Strom ctx err a
repeat element = Strom \_ -> pure (pure (Emit chunk))
  where
  chunk = Array.replicate defaultChunkSize element

-- | The computation's result, again and again.
repeatOm :: forall ctx err a. OmIO ctx err a -> Strom ctx err a
repeatOm computation = Strom \_ -> pure do
  element <- computation
  pure (Emit [ element ])

-- | The computation's result right away, and again as the schedule says.
repeatOmWith :: forall ctx err a. Schedule -> OmIO ctx err a -> Strom ctx err a
repeatOmWith schedule computation = Strom \scope -> do
  repetitions <- newIORef 0 # liftIO
  pure do
    repetition <- readIORef repetitions # liftIO
    case waitBefore repetition of
      Nothing -> pure End
      Just delay -> do
        sleep scope delay
        element <- computation
        writeIORef repetitions (repetition + 1) # liftIO
        pure (Emit [ element ])
  where
  waitBefore repetition
    | repetition == 0 = Just (Milliseconds 0.0)
    | otherwise = delayBefore schedule repetition

-- | 0, 1, 2, … one each interval, the first after one interval.
every :: forall ctx err. Milliseconds -> Strom ctx err Int
every interval = Strom \scope -> do
  nextTick <- newIORef 0 # liftIO
  pure do
    sleep scope interval
    tick <- readIORef nextTick # liftIO
    writeIORef nextTick (tick + 1) # liftIO
    pure (Emit [ tick ])

-- | Acquires the resource when the stream opens and releases it once the stream ends, fails or
-- | is interrupted.
bracket
  :: forall ctx err resource a
   . OmIO ctx err resource
  -> (resource -> OmIO ctx err Unit)
  -> (resource -> Strom ctx err a)
  -> Strom ctx err a
bracket acquireResource releaseResource use = bracketExit acquireResource (\resource _ -> releaseResource resource) use

-- | `bracket`, telling the release how the stream ended.
bracketExit
  :: forall ctx err resource a
   . OmIO ctx err resource
  -> (resource -> Exit err -> OmIO ctx err Unit)
  -> (resource -> Strom ctx err a)
  -> Strom ctx err a
bracketExit acquireResource releaseResource use = Strom \scope -> do
  resource <- acquireResource
  owned <- forkScope scope
  _ <- addFinalizer owned (releaseResource resource)
  openClosingAtEnd owned (use resource)

acquireRelease
  :: forall ctx err resource a
   . OmIO ctx err resource
  -> (resource -> OmIO ctx err Unit)
  -> (resource -> Strom ctx err a)
  -> Strom ctx err a
acquireRelease = bracket

-- | Runs the finalizer once the stream ends, fails or is interrupted.
ensuring :: forall ctx err a. OmIO ctx err Unit -> Strom ctx err a -> Strom ctx err a
ensuring finalizer stream = onExit (const finalizer) stream

-- | Runs the finalizer with how the stream ended, once it ends, fails or is interrupted.
onExit :: forall ctx err a. (Exit err -> OmIO ctx err Unit) -> Strom ctx err a -> Strom ctx err a
onExit finalize stream = Strom \scope -> do
  owned <- forkScope scope
  _ <- addFinalizer owned finalize
  openClosingAtEnd owned stream

-- | Opens the stream in a scope of its own and closes that scope as soon as the stream ends or
-- | fails, rather than when whoever reads it is done.
openClosingAtEnd :: forall ctx err a. Scope ctx err -> Strom ctx err a -> OmIO ctx err (Pull ctx err a)
openClosingAtEnd owned stream = do
  opened <- try (open owned stream)
  case opened of
    Left failure -> closeAfter failure
    Right pull -> pure (pullClosingAtEnd pull)
  where
  pullClosingAtEnd pull = do
    pulled <- try pull
    case pulled of
      Left failure -> closeAfter failure
      Right End -> closeScope Ended owned $> End
      Right step -> pure step

  closeAfter :: forall b. Failure err -> OmIO ctx err b
  closeAfter failure = do
    exit <- exitAfter owned failure
    closeScope exit owned
    throwError failure

--------------------------------------------------------------------------------
-- Changing streams
--------------------------------------------------------------------------------

mapChunks :: forall ctx err a b. (Chunk a -> Chunk b) -> Strom ctx err a -> Strom ctx err b
mapChunks change stream = Strom \scope -> do
  pullUpstream <- open scope stream
  pure (pullUpstream <#> changeStep)
  where
  changeStep = case _ of
    Emit chunk -> Emit (change chunk)
    End -> End

mapChunksM :: forall ctx err a b. (Chunk a -> OmIO ctx err (Chunk b)) -> Strom ctx err a -> Strom ctx err b
mapChunksM change stream = Strom \scope -> do
  pullUpstream <- open scope stream
  pure do
    step <- pullUpstream
    case step of
      Emit chunk -> Emit <$> change chunk
      End -> pure End

mapM :: forall ctx err a b. (a -> OmIO ctx err b) -> Strom ctx err a -> Strom ctx err b
mapM f stream = mapChunksM (traverse f) stream

mapIO :: forall ctx err a b. (a -> IO b) -> Strom ctx err a -> Strom ctx err b
mapIO f stream = mapM (\element -> f element # liftIO) stream

filter :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
filter keep stream = mapChunks (Array.filter keep) stream

filterMap :: forall ctx err a b. (a -> Maybe b) -> Strom ctx err a -> Strom ctx err b
filterMap f stream = mapChunks (Array.mapMaybe f) stream

-- | Runs the computation on each element before passing it on.
tap :: forall ctx err a. (a -> OmIO ctx err Unit) -> Strom ctx err a -> Strom ctx err a
tap f stream = mapChunksM (\chunk -> forEachIn chunk f $> chunk) stream

-- | Each intermediate result of a left fold, the start excluded.
scan :: forall ctx err a b. (b -> a -> b) -> b -> Strom ctx err a -> Strom ctx err b
scan step start stream = Strom \scope -> do
  pullUpstream <- open scope stream
  accumulated <- newIORef start # liftIO
  pure do
    pulled <- pullUpstream
    case pulled of
      End -> pure End
      Emit chunk -> do
        before <- readIORef accumulated # liftIO
        let results = Array.scanl step before chunk
        for_ (Array.last results) \last -> writeIORef accumulated last # liftIO
        pure (Emit results)

-- | The first `count` elements. Once it has them, the stream it takes from is closed: its
-- | goroutines stop and its finalizers run, with `Interrupted`.
take :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
take count stream
  | count <= 0 = empty
  | otherwise = Strom \scope -> do
      upstream <- forkScope scope
      pullUpstream <- openClosingAtEnd upstream stream
      remaining <- newIORef count # liftIO
      pure do
        wanted <- readIORef remaining # liftIO
        if wanted <= 0 then pure End
        else do
          pulled <- pullUpstream
          case pulled of
            End -> pure End
            Emit chunk -> do
              let taken = Array.take wanted chunk
              let stillWanted = wanted - Array.length taken
              writeIORef remaining stillWanted # liftIO
              when (stillWanted <= 0) (closeScope Interrupted upstream)
              pure (Emit taken)

-- | Elements while they pass the test; the first that doesn't ends the stream and closes the one
-- | it takes from.
takeWhile :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
takeWhile keep stream = cutOff whileTheyPass stream
  where
  whileTheyPass chunk = Array.span keep chunk # \{ init, rest } -> init /\ not (Array.null rest)

-- | Elements until one passes the test, that one included.
takeUntil :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
takeUntil stop stream = cutOff upToTheFirstThatStops stream
  where
  upToTheFirstThatStops chunk = case Array.findIndex stop chunk of
    Nothing -> chunk /\ false
    Just index -> Array.take (index + 1) chunk /\ true

-- | Passes on what `cut` keeps of each chunk, until it says the stream is over.
cutOff :: forall ctx err a. (Chunk a -> Chunk a /\ Boolean) -> Strom ctx err a -> Strom ctx err a
cutOff cut stream = Strom \scope -> do
  upstream <- forkScope scope
  pullUpstream <- openClosingAtEnd upstream stream
  over <- newIORef false # liftIO
  pure do
    alreadyOver <- readIORef over # liftIO
    if alreadyOver then pure End
    else do
      pulled <- pullUpstream
      case pulled of
        End -> pure End
        Emit chunk -> do
          let kept /\ nowOver = cut chunk
          when nowOver do
            writeIORef over true # liftIO
            closeScope Interrupted upstream
          pure (Emit kept)

drop :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
drop count stream = Strom \scope -> do
  pullUpstream <- open scope stream
  remaining <- newIORef count # liftIO
  pure do
    pulled <- pullUpstream
    case pulled of
      End -> pure End
      Emit chunk -> do
        toDrop <- readIORef remaining # liftIO
        if toDrop <= 0 then pure (Emit chunk)
        else do
          writeIORef remaining (toDrop - Array.length chunk) # liftIO
          pure (Emit (Array.drop toDrop chunk))

dropWhile :: forall ctx err a. (a -> Boolean) -> Strom ctx err a -> Strom ctx err a
dropWhile skip stream = Strom \scope -> do
  pullUpstream <- open scope stream
  dropping <- newIORef true # liftIO
  pure do
    pulled <- pullUpstream
    case pulled of
      End -> pure End
      Emit chunk -> do
        stillDropping <- readIORef dropping # liftIO
        if not stillDropping then pure (Emit chunk)
        else do
          let rest = Array.dropWhile skip chunk
          unless (Array.null rest) (writeIORef dropping false # liftIO)
          pure (Emit rest)

-- | For each element, all elements of its stream. Each of those streams is closed as soon as
-- | it ends.
concatMap :: forall ctx err a b. (a -> Strom ctx err b) -> Strom ctx err a -> Strom ctx err b
concatMap streamOf stream = Strom \scope -> do
  pullOuter <- open scope stream
  queued <- newIORef ([] /\ 0) # liftIO
  current <- newIORef Nothing # liftIO
  pure (pullInnerOrNext scope pullOuter queued current)
  where
  pullInnerOrNext scope pullOuter queued current = do
    inner <- readIORef current # liftIO
    case inner of
      Just pullInner -> do
        pulled <- pullInner
        case pulled of
          End -> do
            writeIORef current Nothing # liftIO
            pullInnerOrNext scope pullOuter queued current
          step -> pure step
      Nothing -> do
        chunk /\ index <- readIORef queued # liftIO
        case Array.index chunk index of
          Just element -> do
            writeIORef queued (chunk /\ (index + 1)) # liftIO
            innerScope <- forkScope scope
            pullInner <- openClosingAtEnd innerScope (streamOf element)
            writeIORef current (Just pullInner) # liftIO
            pullInnerOrNext scope pullOuter queued current
          Nothing -> do
            pulled <- pullOuter
            case pulled of
              End -> pure End
              Emit next -> do
                writeIORef queued (next /\ 0) # liftIO
                pullInnerOrNext scope pullOuter queued current

flatMap :: forall ctx err a b. (a -> Strom ctx err b) -> Strom ctx err a -> Strom ctx err b
flatMap = concatMap

-- | The streams one after the other.
concat :: forall ctx err a. Array (Strom ctx err a) -> Strom ctx err a
concat streams = fromArray streams # concatMap identity

zip :: forall ctx err a b. Strom ctx err a -> Strom ctx err b -> Strom ctx err (a /\ b)
zip left right = zipWith (/\) left right

-- | Pairs up the elements of both streams in order, as long as both have one.
zipWith :: forall ctx err a b c. (a -> b -> c) -> Strom ctx err a -> Strom ctx err b -> Strom ctx err c
zipWith combine left right = Strom \scope -> do
  pullLeft <- open scope left
  pullRight <- open scope right
  leftovers <- newIORef ([] /\ []) # liftIO
  pure do
    fromLeft /\ fromRight <- readIORef leftovers # liftIO
    leftChunk <- refill pullLeft fromLeft
    rightChunk <- refill pullRight fromRight
    case leftChunk, rightChunk of
      Just lefts, Just rights -> do
        let paired = min (Array.length lefts) (Array.length rights)
        writeIORef leftovers (Array.drop paired lefts /\ Array.drop paired rights) # liftIO
        pure (Emit (Array.zipWith combine lefts rights))
      _, _ -> pure End

-- | The leftover if there is one, or else the next chunk that isn't empty. `Nothing` at the end.
refill :: forall ctx err a. Pull ctx err a -> Chunk a -> OmIO ctx err (Maybe (Chunk a))
refill pull leftover
  | not (Array.null leftover) = pure (Just leftover)
  | otherwise = do
      pulled <- pull
      case pulled of
        End -> pure Nothing
        Emit chunk -> refill pull chunk

-- | The elements in arrays of `size`; the last array may be shorter.
grouped :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err (Array a)
grouped size stream
  | size <= 0 = empty
  | otherwise = Strom \scope -> do
      pullUpstream <- open scope stream
      leftover <- newIORef (Just []) # liftIO
      pure (gather pullUpstream leftover)
  where
  gather pullUpstream leftover = do
    gathered <- readIORef leftover # liftIO
    case gathered of
      Nothing -> pure End
      Just elements
        | Array.length elements >= size -> do
            let { groups, rest } = groupsOf size elements
            writeIORef leftover (Just rest) # liftIO
            pure (Emit groups)
        | otherwise -> do
            pulled <- pullUpstream
            case pulled of
              Emit chunk -> do
                writeIORef leftover (Just (joined elements chunk)) # liftIO
                gather pullUpstream leftover
              End -> do
                writeIORef leftover Nothing # liftIO
                pure if Array.null elements then End else Emit [ elements ]

  joined elements chunk
    | Array.null elements = chunk
    | otherwise = elements <> chunk

-- | Full groups of `size`, and what is left over.
groupsOf :: forall a. Int -> Array a -> { groups :: Array (Array a), rest :: Array a }
groupsOf size elements = { groups: map group (Array.range 0 (fullGroups - 1)), rest: Array.drop (fullGroups * size) elements }
  where
  fullGroups = Array.length elements / size
  group index = Array.slice (index * size) ((index + 1) * size) elements

-- | The same elements, in chunks of `size`; the last may be smaller.
rechunk :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
rechunk size stream = grouped size stream # flattenChunks

-- | Chunks no larger than `size`, without joining smaller ones.
inChunksOfAtMost :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
inChunksOfAtMost size stream = stream # mapChunks splitUp # flattenChunks
  where
  splitUp chunk
    | Array.length chunk <= size = [ chunk ]
    | otherwise = groupsOf size chunk # \{ groups, rest } -> if Array.null rest then groups else Array.snoc groups rest

-- | Each element, an array, becomes a chunk of its own.
flattenChunks :: forall ctx err a. Strom ctx err (Array a) -> Strom ctx err a
flattenChunks stream = Strom \scope -> do
  pullUpstream <- open scope stream
  queued <- newIORef ([] /\ 0) # liftIO
  pure (nextChunk pullUpstream queued)
  where
  nextChunk pullUpstream queued = do
    chunks /\ index <- readIORef queued # liftIO
    case Array.index chunks index of
      Just chunk -> do
        writeIORef queued (chunks /\ (index + 1)) # liftIO
        pure (Emit chunk)
      Nothing -> do
        pulled <- pullUpstream
        case pulled of
          End -> pure End
          Emit next -> do
            writeIORef queued (next /\ 0) # liftIO
            nextChunk pullUpstream queued

-- | If the stream fails, goes on with the recovery's stream instead. Interruption isn't a
-- | failure to recover from.
catchAll :: forall ctx err a. (Failure err -> Strom ctx err a) -> Strom ctx err a -> Strom ctx err a
catchAll recover stream = switchOnFailure recoverOnce stream
  where
  recoverOnce _ failuresBefore failure
    | failuresBefore == 0 = pure (Just (recover failure))
    | otherwise = pure Nothing

-- | Fails with what `change` makes of the stream's failure, as it's opened or pulled. Being
-- | interrupted stays being interrupted.
mapError :: forall ctx err a. (Failure err -> Failure err) -> Strom ctx err a -> Strom ctx err a
mapError change stream = Strom \scope -> do
  pull <- open scope stream `catchError` failChanged
  pure (pull `catchError` failChanged)
  where
  failChanged :: forall b. Failure err -> OmIO ctx err b
  failChanged failure = throwError (if isInterruption failure then failure else change failure)

-- | If the stream fails, goes on with the other.
orElse :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
orElse other stream = catchAll (const other) stream

-- | If the stream fails, starts it over after the schedule's wait, until the schedule gives up.
retry :: forall ctx err a. Schedule -> Strom ctx err a -> Strom ctx err a
retry schedule stream = switchOnFailure startOver stream
  where
  startOver scope failuresBefore _ = case delayBefore schedule (failuresBefore + 1) of
    Nothing -> pure Nothing
    Just delay -> sleep scope delay $> Just stream

-- | Each time the stream in use fails, `decide` gets how many failed before and the failure, and
-- | says which stream to go on with, or `Nothing` to fail.
switchOnFailure
  :: forall ctx err a
   . (Scope ctx err -> Int -> Failure err -> OmIO ctx err (Maybe (Strom ctx err a)))
  -> Strom ctx err a
  -> Strom ctx err a
switchOnFailure decide stream = Strom \scope -> do
  failures <- newIORef 0 # liftIO
  firstPull <- openOwnedOrFailing scope stream
  current <- newIORef firstPull # liftIO
  pure (pullOrSwitch scope failures current)
  where
  pullOrSwitch scope failures current = do
    pull <- readIORef current # liftIO
    pulled <- try pull
    case pulled of
      Right step -> pure step
      Left failure -> do
        stopped <- isInterrupted scope
        failuresBefore <- readIORef failures # liftIO
        next <- if stopped || isInterruption failure then pure Nothing else decide scope failuresBefore failure
        case next of
          Nothing -> throwError failure
          Just nextStream -> do
            writeIORef failures (failuresBefore + 1) # liftIO
            nextPull <- openOwnedOrFailing scope nextStream
            writeIORef current nextPull # liftIO
            pullOrSwitch scope failures current

  openOwnedOrFailing scope opening = do
    owned <- forkScope scope
    opened <- try (openClosingAtEnd owned opening)
    pure (either throwError identity opened)

-- | Ends the stream once the duration has passed. A pull waiting at that moment stops waiting,
-- | and so does a gRPC call made in the stream, whose deadline it is.
timeout :: forall ctx err a. Milliseconds -> Strom ctx err a -> Strom ctx err a
timeout duration stream = Strom \scope -> do
  timed <- forkScopeWithTimeout duration scope
  pull <- openClosingAtEnd timed stream
  pure do
    timedOut <- deadlinePassedIn timed
    if timedOut then closeScope Interrupted timed $> End
    else do
      pulled <- try pull
      case pulled of
        Right step -> pure step
        Left failure -> do
          timedOutWhilePulling <- deadlinePassedIn timed
          outsideStopped <- isInterrupted scope
          if timedOutWhilePulling && not outsideStopped then pure End else throwError failure

--------------------------------------------------------------------------------
-- At the same time
--------------------------------------------------------------------------------

-- | What a goroutine that pulls a stream hands over, and from which of the streams.
data Signal err a
  = Pulled Int (Chunk a)
  | Finished Int
  | FailedWith Int (Failure err)

-- | A goroutine pulling a stream, in a scope of its own.
type Producer ctx err = { scope :: Scope ctx err, stopped :: MVar Unit }

-- | Pulls the stream on a goroutine of its own and hands over each chunk, then how it ended. It
-- | stops at its next handover once its scope is interrupted. The stream's finalizers run on that
-- | goroutine, and the scope it was started in waits for it when it closes.
startProducer :: forall ctx err a. Scope ctx err -> Handoff (Signal err a) -> Int -> Strom ctx err a -> OmIO ctx err (Producer ctx err)
startProducer scope handoff source stream = do
  producerScope <- forkScope scope
  stopped <- newEmptyMVar # liftIO
  _ <- addFinalizer scope \_ -> readMVar stopped # liftIO
  context <- ask
  _ <- forkIO (void (runOmIO context (produce producerScope)) `finally` putMVar stopped unit) # liftIO
  pure { scope: producerScope, stopped }
  where
  produce producerScope = do
    outcome <- attempt (open producerScope stream >>= handOverAll producerScope)
    case outcome of
      Right true -> do
        handOver producerScope (Finished source)
        closeScope Ended producerScope
      Right false -> closeScope Interrupted producerScope
      Left failure -> do
        exit <- exitAfter producerScope failure
        case exit of
          Interrupted -> pure unit
          _ -> handOver producerScope (FailedWith source failure)
        closeScope exit producerScope

  handOverAll producerScope pull = do
    pulled <- pull
    case pulled of
      End -> pure true
      Emit chunk -> do
        handedOver <- offer (contextOf producerScope) handoff (Pulled source chunk) # liftIO
        if handedOver then handOverAll producerScope pull else pure false

  handOver producerScope signal = void (offer (contextOf producerScope) handoff signal # liftIO)

-- | Interrupts the producer and waits until it has stopped and closed its scope.
stopProducer :: forall ctx err. Producer ctx err -> OmIO ctx err Unit
stopProducer { scope, stopped } = do
  interrupt scope # liftIO
  readMVar stopped # liftIO

-- | Reads the streams at the same time, each on a goroutine of its own, through a handoff with
-- | room for `capacity` chunks. It ends as the strategy says, and fails as soon as one fails;
-- | either way it stops the others first.
mergeStreams :: forall ctx err a. Int -> HaltStrategy -> Array (Strom ctx err a) -> Strom ctx err a
mergeStreams capacity strategy streams = Strom \scope -> do
  handoff <- newHandoff capacity # liftIO
  producers <- forWithIndex streams \source stream -> startProducer scope handoff source stream
  running <- newIORef (Array.length streams) # liftIO
  pure (receiveNext scope handoff producers running)
  where
  receiveNext scope handoff producers running = do
    stillRunning <- readIORef running # liftIO
    if stillRunning <= 0 then pure End
    else do
      received <- receive (contextOf scope) handoff # liftIO
      case received of
        Nothing -> throwError interrupted
        Just (Pulled _ chunk) -> pure (Emit chunk)
        Just (FailedWith _ failure) -> do
          writeIORef running 0 # liftIO
          for_ producers stopProducer
          throwError failure
        Just (Finished source)
          | halts source || stillRunning == 1 -> do
              writeIORef running 0 # liftIO
              for_ producers stopProducer
              pure End
          | otherwise -> do
              writeIORef running (stillRunning - 1) # liftIO
              receiveNext scope handoff producers running

  halts source = case strategy of
    HaltBoth -> false
    HaltEither -> true
    HaltLeft -> source == 0
    HaltRight -> source == 1

-- | Up to `capacity` elements read ahead, on a goroutine of their own, while whoever reads this
-- | stream is busy with earlier ones. Each element read ahead holds a permit until it's pulled.
buffer :: forall ctx err a. Int -> Strom ctx err a -> Strom ctx err a
buffer capacity stream
  | capacity <= 0 = stream
  | otherwise = Strom \scope -> do
      permits <- newPermits capacity # liftIO
      inChunksOfAtMost capacity stream
        # holdingPermits permits
        # readAheadBy capacity
        # releasingPermits permits
        # open scope
  where
  readAheadBy chunks readAhead = mergeStreams chunks HaltBoth [ readAhead ]

  holdingPermits permits upstream = Strom \producerScope -> do
    pullUpstream <- open producerScope upstream
    pure do
      pulled <- pullUpstream
      case pulled of
        End -> pure End
        Emit chunk -> do
          permitted <- acquireMany (contextOf producerScope) permits (Array.length chunk) # liftIO
          if permitted then pure (Emit chunk) else throwError interrupted

  releasingPermits permits readAhead = readAhead # tapChunks \chunk -> releaseMany permits (Array.length chunk) # liftIO

  tapChunks f readAhead = mapChunksM (\chunk -> f chunk $> chunk) readAhead

-- | Both streams at the same time, each on a goroutine of its own, in the order their chunks
-- | come. It ends once both have.
merge :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
merge left right = mergeAll [ left, right ]

mergeAll :: forall ctx err a. Array (Strom ctx err a) -> Strom ctx err a
mergeAll streams = mergeStreams 0 HaltBoth streams

mergeWith :: forall ctx err a b c. HaltStrategy -> (a -> c) -> (b -> c) -> Strom ctx err a -> Strom ctx err b -> Strom ctx err c
mergeWith strategy fromLeft fromRight left right = mergeStreams 0 strategy [ map fromLeft left, map fromRight right ]

-- | `merge`, ending as soon as either ends.
mergeHaltEither :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
mergeHaltEither left right = mergeStreams 0 HaltEither [ left, right ]

mergeHaltLeft :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
mergeHaltLeft left right = mergeStreams 0 HaltLeft [ left, right ]

mergeHaltRight :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
mergeHaltRight left right = mergeStreams 0 HaltRight [ left, right ]

-- | Whichever stream gives a chunk or ends first; the other is stopped.
race :: forall ctx err a. Strom ctx err a -> Strom ctx err a -> Strom ctx err a
race left right = raceAll [ left, right ]

raceAll :: forall ctx err a. Array (Strom ctx err a) -> Strom ctx err a
raceAll streams = Strom \scope -> do
  handoff <- newHandoff 0 # liftIO
  producers <- forWithIndex streams \source stream -> startProducer scope handoff source stream
  winner <- newIORef Nothing # liftIO
  finished <- newIORef (Array.null streams) # liftIO
  pure (receiveNext scope handoff producers winner finished)
  where
  receiveNext scope handoff producers winner finished = do
    alreadyFinished <- readIORef finished # liftIO
    if alreadyFinished then pure End
    else do
      received <- receive (contextOf scope) handoff # liftIO
      case received of
        Nothing -> throwError interrupted
        Just signal -> do
          decided <- readIORef winner # liftIO
          case decided of
            Nothing -> do
              writeIORef winner (Just (sourceOf signal)) # liftIO
              forWithIndex_ producers \source producer -> unless (source == sourceOf signal) (stopProducer producer)
              answer finished producers signal
            Just _ -> answer finished producers signal

  answer finished producers = case _ of
    Pulled _ chunk -> pure (Emit chunk)
    Finished _ -> do
      writeIORef finished true # liftIO
      pure End
    FailedWith _ failure -> do
      writeIORef finished true # liftIO
      for_ producers stopProducer
      throwError failure

  sourceOf = case _ of
    Pulled source _ -> source
    Finished source -> source
    FailedWith source _ -> source

-- | `zip`, with each side read ahead on a goroutine of its own.
zipPar :: forall ctx err a b. Strom ctx err a -> Strom ctx err b -> Strom ctx err (a /\ b)
zipPar left right = zip (buffer 1 left) (buffer 1 right)

-- | Once both have given an element, each element of either with the latest of the other. Ends
-- | once either ends.
zipLatest :: forall ctx err a b. Strom ctx err a -> Strom ctx err b -> Strom ctx err (a /\ b)
zipLatest left right = zipLatestWith (/\) left right

zipLatestWith :: forall ctx err a b c. (a -> b -> c) -> Strom ctx err a -> Strom ctx err b -> Strom ctx err c
zipLatestWith combine left right =
  mergeWith HaltEither Left Right left right
    # scan remember (Nothing /\ Nothing)
    # filterMap \(latestLeft /\ latestRight) -> combine <$> latestLeft <*> latestRight
  where
  remember (latestLeft /\ latestRight) = case _ of
    Left element -> Just element /\ latestRight
    Right element -> latestLeft /\ Just element

-- | `mapM`, with up to `parallelism` elements worked on at once, each on a goroutine of its own.
-- | The results keep the elements' order.
mapPar :: forall ctx err a b. Int -> (a -> OmIO ctx err b) -> Strom ctx err a -> Strom ctx err b
mapPar parallelism f stream
  | parallelism <= 1 = mapM f stream
  | otherwise =
      inChunksOfAtMost parallelism stream
        # startingWorkers
        # onAGoroutineOfItsOwn
        # mapChunksM (traverse awaitResult)
  where
  onAGoroutineOfItsOwn promises = mergeStreams 1 HaltBoth [ promises ]

  startingWorkers upstream = Strom \scope -> do
    permits <- newPermits parallelism # liftIO
    _ <- addFinalizer scope \_ -> awaitAllReleased permits # liftIO
    pullUpstream <- open scope upstream
    pure do
      pulled <- pullUpstream
      case pulled of
        End -> pure End
        Emit chunk -> Emit <$> for chunk (startWorker scope permits)

  startWorker scope permits element = do
    promise <- newEmptyMVar # liftIO
    runWorker scope permits (f element) (putMVar promise)
    pure promise

  awaitResult promise = do
    result <- readMVar promise # liftIO
    either throwError pure result

-- | `mapPar`, with the results in the order they're done.
mapParUnordered :: forall ctx err a b. Int -> (a -> OmIO ctx err b) -> Strom ctx err a -> Strom ctx err b
mapParUnordered parallelism f stream = Strom \scope -> do
  results <- newHandoff parallelism # liftIO
  producer <- startProducer scope results 0 (inChunksOfAtMost parallelism stream # feedingWorkers results)
  running <- newIORef true # liftIO
  pure (receiveNext scope results producer running)
  where
  feedingWorkers results upstream = Strom \workersScope -> do
    permits <- newPermits parallelism # liftIO
    _ <- addFinalizer workersScope \_ -> awaitAllReleased permits # liftIO
    pullUpstream <- open workersScope upstream
    pure (startAll workersScope permits results pullUpstream)

  startAll workersScope permits results pullUpstream = do
    pulled <- pullUpstream
    case pulled of
      End -> do
        awaitAllReleased permits # liftIO
        pure End
      Emit chunk -> do
        for_ chunk \element -> runWorker workersScope permits (f element) (handOverResult workersScope results)
        startAll workersScope permits results pullUpstream

  handOverResult workersScope results outcome = void (offer (contextOf workersScope) results (either (FailedWith 0) (Pulled 0 <<< Array.singleton) outcome))

  receiveNext scope results producer running = do
    stillRunning <- readIORef running # liftIO
    if not stillRunning then pure End
    else do
      received <- receive (contextOf scope) results # liftIO
      case received of
        Nothing -> throwError interrupted
        Just (Pulled _ chunk) -> pure (Emit chunk)
        Just (Finished _) -> do
          writeIORef running false # liftIO
          pure End
        Just (FailedWith _ failure) -> do
          writeIORef running false # liftIO
          stopProducer producer
          throwError failure

-- | Takes a permit, waiting while all are taken, and runs the work on a goroutine of its own,
-- | handing its outcome to `done`. Fails with `interrupted` if the scope is interrupted while
-- | waiting.
runWorker
  :: forall ctx err b
   . Scope ctx err
  -> Permits
  -> OmIO ctx err b
  -> (Either (Failure err) b -> IO Unit)
  -> OmIO ctx err Unit
runWorker scope permits work done = do
  permitted <- acquire (contextOf scope) permits # liftIO
  unless permitted (throwError interrupted)
  context <- ask
  _ <- forkIO (workAndRelease context) # liftIO
  pure unit
  where
  workAndRelease context = (runOmIO context (attempt work) >>= either (done <<< Left) done) `finally` release permits

-- | Drains the stream on a goroutine of its own, in a child of the scope, which waits for it when
-- | it closes. Closing the scope interrupts it.
forkDrain :: forall ctx err a. Scope ctx err -> Strom ctx err a -> OmIO ctx err Unit
forkDrain scope stream = do
  handoff <- newHandoff 1 # liftIO
  void (startProducer scope handoff 0 (drained stream))

-- | Reads the whole stream when first pulled, and gives nothing.
drained :: forall ctx err a b. Strom ctx err a -> Strom ctx err b
drained stream = Strom \scope -> do
  pullUpstream <- open scope stream
  pure (pullUntilEnd pullUpstream)
  where
  pullUntilEnd pullUpstream = do
    pulled <- pullUpstream
    case pulled of
      End -> pure End
      Emit _ -> pullUntilEnd pullUpstream

--------------------------------------------------------------------------------
-- Running streams
--------------------------------------------------------------------------------

-- | Runs the computation on each element, in order.
forEach :: forall ctx err a. (a -> OmIO ctx err Unit) -> Strom ctx err a -> OmIO ctx err Unit
forEach f stream = scoped \scope -> forEachChunkIn scope (\chunk -> forEachIn chunk f) stream

-- | `forEach`, interrupted once the context is done: a gRPC call's.
forEachWithin :: forall ctx err a. Context -> (a -> OmIO ctx err Unit) -> Strom ctx err a -> OmIO ctx err Unit
forEachWithin context f stream = scopedWithin context \scope -> forEachChunkIn scope (\chunk -> forEachIn chunk f) stream

runFold :: forall ctx err a b. b -> (b -> a -> b) -> Strom ctx err a -> OmIO ctx err b
runFold start step stream = foldChunks start (\accumulated chunk -> pure (Array.foldl step accumulated chunk)) stream

runFoldM :: forall ctx err a b. b -> (b -> a -> OmIO ctx err b) -> Strom ctx err a -> OmIO ctx err b
runFoldM start step stream = foldChunks start (\accumulated chunk -> foldIn chunk accumulated) stream
  where
  foldIn chunk accumulated = foldFrom chunk 0 accumulated
  foldFrom chunk index accumulated = case Array.index chunk index of
    Nothing -> pure accumulated
    Just element -> step accumulated element >>= foldFrom chunk (index + 1)

runCollect :: forall ctx err a. Strom ctx err a -> OmIO ctx err (Array a)
runCollect stream = Array.concat <$> foldChunks [] (\chunks chunk -> pure (Array.snoc chunks chunk)) stream

runDrain :: forall ctx err a. Strom ctx err a -> OmIO ctx err Unit
runDrain stream = foldChunks unit (\_ _ -> pure unit) stream

-- | The first element, if there is one. The stream is closed right after it.
runHead :: forall ctx err a. Strom ctx err a -> OmIO ctx err (Maybe a)
runHead stream = scoped \scope -> do
  pull <- open scope stream
  firstChunk <- refill pull []
  pure (firstChunk >>= Array.head)

runLast :: forall ctx err a. Strom ctx err a -> OmIO ctx err (Maybe a)
runLast stream = foldChunks Nothing (\latest chunk -> pure (Array.last chunk <|> latest)) stream

-- | Folds the chunks in a scope of its own, which closes at the end.
foldChunks :: forall ctx err a b. b -> (b -> Chunk a -> OmIO ctx err b) -> Strom ctx err a -> OmIO ctx err b
foldChunks start step stream = scoped \scope -> do
  pull <- open scope stream
  foldPulled step scope pull start

-- | Folds what the pull gives, chunk by chunk, until it ends.
foldPulled :: forall ctx err a b. (b -> Chunk a -> OmIO ctx err b) -> Scope ctx err -> Pull ctx err a -> b -> OmIO ctx err b
foldPulled step scope pull accumulated = do
  throwIfInterrupted scope
  pulled <- pull
  case pulled of
    End -> pure accumulated
    Emit chunk -> step accumulated chunk >>= foldPulled step scope pull

forEachChunkIn :: forall ctx err a. Scope ctx err -> (Chunk a -> OmIO ctx err Unit) -> Strom ctx err a -> OmIO ctx err Unit
forEachChunkIn scope eachChunk stream = do
  pull <- open scope stream
  foldPulled (\_ chunk -> eachChunk chunk) scope pull unit

-- | Runs the computation on each element of the chunk, in order.
forEachIn :: forall ctx err a. Chunk a -> (a -> OmIO ctx err Unit) -> OmIO ctx err Unit
forEachIn chunk f = forEachFrom 0
  where
  forEachFrom index = case Array.index chunk index of
    Nothing -> pure unit
    Just element -> do
      f element
      forEachFrom (index + 1)

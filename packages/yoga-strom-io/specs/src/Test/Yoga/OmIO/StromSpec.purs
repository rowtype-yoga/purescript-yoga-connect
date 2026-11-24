module Test.Yoga.OmIO.StromSpec where

import Prelude

import Control.Monad.Error.Class (throwError)
import Data.Array as Array
import Data.Bifunctor (lmap)
import Data.Either (Either(..))
import Data.Int as Int
import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested ((/\))
import Data.Variant as Variant
import Effect.Aff (Aff)
import Effect.Exception (message)
import Go.IO (forkIO, threadDelay, toAff)
import Go.IO.Chan (close, newBufferedChan, sendAll)
import Go.IO.IORef (IORef, atomicModifyIORef', newIORef, readIORef)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (fail, shouldEqual, shouldSatisfy)
import Type.Proxy (Proxy(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom (Exit(..), Failure, Strom)
import Yoga.OmIO.Strom as Strom
import Yoga.OmIO.Strom.Context (background, cancel, withCancel)
import Yoga.OmIO.Strom.Schedule as Schedule

type Stream a = Strom Unit (boom :: String) a

type Om a = OmIO Unit (boom :: String) a

spec :: Spec Unit
spec = describe "Yoga.OmIO.Strom" do
  chunkSpec
  combinatorSpec
  failureSpec
  finalizerSpec
  interruptionSpec
  concurrencySpec
  backpressureSpec

chunkSpec :: Spec Unit
chunkSpec = describe "chunks" do
  it "range comes in chunks of 4096" do
    chunkSizesOf (Strom.range 0 10000) >>= shouldEqual [ 4096, 4096, 1808 ]

  it "map and filter keep each chunk a chunk" do
    chunksOf (Strom.fromChunks [ [ 1, 2, 3 ], [ 4, 5 ] ] # map (_ * 10) # Strom.filter (_ /= 20))
      >>= shouldEqual [ [ 10, 30 ], [ 40, 50 ] ]

  it "take cuts the chunk it stops in" do
    chunkSizesOf (Strom.range 0 10000 # Strom.take 5000) >>= shouldEqual [ 4096, 904 ]

  it "rechunk changes where chunks end, and grouped gives the groups as elements" do
    chunksOf (Strom.fromChunks [ [ 1 ], [ 2, 3, 4, 5 ], [ 6, 7 ] ] # Strom.rechunk 3)
      >>= shouldEqual [ [ 1, 2, 3 ], [ 4, 5, 6 ], [ 7 ] ]
    collected (Strom.range 1 8 # Strom.grouped 3) >>= shouldEqual [ [ 1, 2, 3 ], [ 4, 5, 6 ], [ 7 ] ]

  it "maps and folds numbers chunk by chunk, as an array would" do
    let readings = Array.range 0 9999 <#> \index -> Int.toNumber (index `mod` 100) * 0.5
    total <- run (Strom.fromArray readings # Strom.rechunk 4096 # map (_ * 2.0) # Strom.runFold 0.0 (+))
    total `shouldEqual` Right (readings # map (_ * 2.0) # Array.foldl (+) 0.0)

  it "zips chunks of different sizes" do
    collected (Strom.zip (Strom.fromChunks [ [ 1, 2, 3 ], [ 4 ] ]) (Strom.fromChunks [ [ "a" ], [ "b", "c", "d", "e" ] ]))
      >>= shouldEqual [ 1 /\ "a", 2 /\ "b", 3 /\ "c", 4 /\ "d" ]

combinatorSpec :: Spec Unit
combinatorSpec = describe "making and changing streams" do
  it "reads a channel until it's closed" do
    channel <- toAff do
      channel <- newBufferedChan 3
      _ <- forkIO (sendAll channel [ 1, 2, 3 ] *> close channel)
      pure channel
    collected (Strom.fromChan channel) >>= shouldEqual [ 1, 2, 3 ]

  it "maps, filters and scans" do
    collected (Strom.range 1 11 # map (_ * 10) # Strom.filter (_ > 50) # Strom.scan (+) 0)
      >>= shouldEqual [ 60, 130, 210, 300, 400 ]

  it "takes and drops, while and until" do
    collected (Strom.range 1 11 # Strom.takeWhile (_ < 4)) >>= shouldEqual [ 1, 2, 3 ]
    collected (Strom.range 1 11 # Strom.takeUntil (_ == 4)) >>= shouldEqual [ 1, 2, 3, 4 ]
    collected (Strom.range 1 6 # Strom.drop 3) >>= shouldEqual [ 4, 5 ]
    collected (Strom.range 1 6 # Strom.dropWhile (_ < 4)) >>= shouldEqual [ 4, 5 ]

  it "iterates and unfolds" do
    collected (Strom.iterate (_ * 2) 1 # Strom.take 5) >>= shouldEqual [ 1, 2, 4, 8, 16 ]
    collected (Strom.unfold (\n -> if n > 3 then Nothing else Just (show n /\ (n + 1))) 1) >>= shouldEqual [ "1", "2", "3" ]

  it "goes on after a stream with the next, and for each element with its stream" do
    collected ((Strom.iterate (_ + 1) 1 # Strom.take 2) <> Strom.fromArray [ 100 ]) >>= shouldEqual [ 1, 2, 100 ]
    collected (Strom.range 1 4 # Strom.concatMap (\n -> Strom.fromArray [ n, n ])) >>= shouldEqual [ 1, 1, 2, 2, 3, 3 ]
    pairs <- collected do
      n <- Strom.range 1 3
      letter <- Strom.fromArray [ "a", "b" ]
      pure (show n <> letter)
    pairs `shouldEqual` [ "1a", "1b", "2a", "2b" ]

  it "runs effects for each element, in order" do
    log <- newLog
    _ <- run (Strom.range 1 4 # Strom.tap (\n -> record log (show n)) # Strom.runDrain)
    entriesOf log >>= shouldEqual [ "1", "2", "3" ]

  it "folds, and reads the first and the last element" do
    run (Strom.range 1 101 # Strom.runFold 0 (+)) >>= shouldEqual (Right 5050)
    run (Strom.range 1 101 # Strom.runFoldM 0 (\total n -> pure (total + n))) >>= shouldEqual (Right 5050)
    run (Strom.range 7 100 # Strom.runHead) >>= shouldEqual (Right (Just 7))
    run (Strom.range 7 100 # Strom.runLast) >>= shouldEqual (Right (Just 99))
    run (Strom.empty # Strom.runHead) >>= shouldEqual (Right (Nothing :: Maybe Int))

failureSpec :: Spec Unit
failureSpec = describe "typed failures" do
  it "fails with the error of its row" do
    run (Strom.fromArray [ 1, 2 ] <> Strom.fail (boom "no more") # Strom.runCollect)
      >>= shouldEqual (Left "boom: no more")

  it "fails where a computation throws" do
    run (Strom.range 1 10 # Strom.mapM (\n -> if n == 3 then throwError (boom "three") else pure n) # Strom.runCollect)
      >>= shouldEqual (Left "boom: three")

  it "recovers with catchAll, keeping what came before the failure" do
    collected (Strom.fromArray [ 1, 2 ] <> Strom.fail (boom "no more") # Strom.catchAll (\_ -> Strom.fromArray [ 100 ]))
      >>= shouldEqual [ 1, 2, 100 ]

  it "hands the failure to catchAll" do
    collected (Strom.fail (boom "why") # Strom.catchAll (\failure -> Strom.succeed (describeFailure failure)))
      >>= shouldEqual [ "boom: why" ]

  it "changes the failure with mapError, keeping what came before it" do
    run (Strom.fromArray [ 1, 2 ] <> Strom.fail (boom "late") # Strom.mapError (\_ -> boom "changed") # Strom.runCollect)
      >>= shouldEqual (Left "boom: changed")
    collected (Strom.fromArray [ 1, 2 ] # Strom.mapError (\_ -> boom "never")) >>= shouldEqual [ 1, 2 ]

  it "goes on with another stream with orElse" do
    collected (Strom.fail (boom "no") # Strom.orElse (Strom.fromArray [ 7 ])) >>= shouldEqual [ 7 ]

  it "starts over with retry, until the schedule gives up" do
    attempts <- newCounter
    outcome <- run (Strom.fromOm (increment attempts) <> Strom.fail (boom "again") # Strom.retry (Schedule.recurs 2) # Strom.runCollect)
    outcome `shouldEqual` Left "boom: again"
    countOf attempts >>= shouldEqual 3

finalizerSpec :: Spec Unit
finalizerSpec = describe "finalizers" do
  it "run when the stream ends" do
    log <- newLog
    run (Strom.range 1 4 # Strom.onExit (recordExit log "range") # Strom.runCollect) >>= shouldEqual (Right [ 1, 2, 3 ])
    entriesOf log >>= shouldEqual [ "range ended" ]

  it "leave a stream that waits alone until it's done" do
    log <- newLog
    collected (Strom.every (Milliseconds 2.0) # Strom.onExit (recordExit log "ticks") # Strom.take 3) >>= shouldEqual [ 0, 1, 2 ]
    collected (Strom.merge (Strom.every (Milliseconds 2.0) # Strom.onExit (recordExit log "merged ticks") # Strom.take 2) Strom.empty) >>= shouldEqual [ 0, 1 ]
    entriesOf log >>= shouldEqual [ "ticks interrupted", "merged ticks interrupted" ]

  it "run when the stream fails" do
    log <- newLog
    _ <- run (Strom.fail (boom "no") # Strom.onExit (recordExit log "failing") # Strom.runDrain :: Om Unit)
    entriesOf log >>= shouldEqual [ "failing failed" ]

  it "run when the stream is interrupted, as soon as take has enough" do
    log <- newLog
    taken <- run do
      Strom.iterate (_ + 1) 1
        # Strom.onExit (recordExit log "endless")
        # Strom.take 2
        # Strom.tap (\n -> record log ("took " <> show n))
        # Strom.runCollect
    taken `shouldEqual` Right [ 1, 2 ]
    entriesOf log >>= shouldEqual [ "took 1", "endless interrupted", "took 2" ]

  it "release a bracket's resource after the stream that uses it, last acquired first released" do
    log <- newLog
    let resource name = Strom.bracket (record log ("acquire " <> name) $> name) (\acquired -> record log ("release " <> acquired))
    _ <- run (resource "outer" (\_ -> resource "inner" \_ -> Strom.fromArray [ 1, 2 ]) # Strom.runDrain)
    entriesOf log >>= shouldEqual [ "acquire outer", "acquire inner", "release inner", "release outer" ]

  it "close each stream of concatMap as soon as it ends" do
    log <- newLog
    _ <- run do
      Strom.range 1 3
        # Strom.concatMap (\n -> Strom.succeed n # Strom.ensuring (record log ("closed " <> show n)))
        # Strom.tap (\n -> record log ("got " <> show n))
        # Strom.runDrain
    entriesOf log >>= shouldEqual [ "got 1", "closed 1", "got 2", "closed 2" ]

interruptionSpec :: Spec Unit
interruptionSpec = describe "interruption" do
  it "take stops a producer without end, and nothing pulls it afterwards" do
    pulls <- newCounter
    run (Strom.repeatOm (increment pulls) # Strom.take 5 # Strom.runDrain) >>= shouldEqual (Right unit)
    countOf pulls >>= shouldEqual 5

  it "take stops a producer on a goroutine of its own, and waits for its finalizers" do
    log <- newLog
    pulls <- newCounter
    _ <- run do
      Strom.repeatOm (increment pulls)
        # Strom.onExit (recordExit log "producer")
        # Strom.buffer 4
        # Strom.take 3
        # Strom.runDrain
    entriesOf log >>= shouldEqual [ "producer interrupted" ]
    pullsAfterRun <- countOf pulls
    pause 20.0
    countOf pulls >>= shouldEqual pullsAfterRun

  it "stops a sleeping stream as soon as its context is cancelled from outside" do
    log <- newLog
    outcome <- toAff do
      context /\ cancelContext <- withCancel background
      _ <- forkIO (threadDelay (Milliseconds 30.0) *> cancel cancelContext)
      runOmIO unit do
        Strom.every (Milliseconds 5.0)
          # Strom.onExit (recordExit log "ticks")
          # Strom.forEachWithin context (\_ -> pure unit)
    lmap describeFailure outcome `shouldEqual` Left "the stream was interrupted"
    entriesOf log >>= shouldEqual [ "ticks interrupted" ]

  it "ends a stream once its timeout passes" do
    ticks <- collected (Strom.every (Milliseconds 5.0) # Strom.timeout (Milliseconds 32.0))
    Array.length ticks `shouldSatisfy` (_ >= 3)
    Array.length ticks `shouldSatisfy` (_ <= 6)

  it "repeats on a schedule" do
    collected (Strom.repeatOmWith (Schedule.spaced (Milliseconds 1.0) <> Schedule.recurs 3) (pure "again")) >>= shouldEqual [ "again", "again", "again", "again" ]

concurrencySpec :: Spec Unit
concurrencySpec = describe "at the same time" do
  it "merges two streams running at the same time" do
    merged <- collected (Strom.merge (Strom.every (Milliseconds 5.0) # Strom.take 3 # map (_ + 100)) (Strom.fromArray [ 1, 2 ]))
    Array.sort merged `shouldEqual` [ 1, 2, 100, 101, 102 ]

  it "stops both sides of a merge once its reader has enough, and runs their finalizers" do
    log <- newLog
    taken <- run do
      Strom.merge (Strom.repeat 1 # Strom.onExit (recordExit log "left")) (Strom.repeat 2 # Strom.onExit (recordExit log "right"))
        # Strom.take 10
        # Strom.runCollect
    map Array.length taken `shouldEqual` Right 10
    map Array.sort (entriesOf log) >>= shouldEqual [ "left interrupted", "right interrupted" ]

  it "stops the other side of a merge when one fails" do
    log <- newLog
    outcome <- run do
      Strom.merge (Strom.repeat 1 # Strom.onExit (recordExit log "endless")) (Strom.fail (boom "broken"))
        # Strom.runDrain
    outcome `shouldEqual` Left "boom: broken"
    entriesOf log >>= shouldEqual [ "endless interrupted" ]

  it "ends mergeHaltEither with the first side that ends" do
    log <- newLog
    merged <- collected (Strom.mergeHaltEither (Strom.every (Milliseconds 2.0) # Strom.onExit (recordExit log "ticks") # map (const 0)) (Strom.fromArray [ 1 ]))
    merged `shouldSatisfy` Array.elem 1
    entriesOf log >>= shouldEqual [ "ticks interrupted" ]

  it "races two streams and stops the loser" do
    log <- newLog
    winners <- collected do
      Strom.race
        (Strom.every (Milliseconds 200.0) # map (const "slow") # Strom.onExit (recordExit log "slow"))
        (Strom.every (Milliseconds 2.0) # map (const "fast"))
        # Strom.take 3
    winners `shouldEqual` [ "fast", "fast", "fast" ]
    entriesOf log >>= shouldEqual [ "slow interrupted" ]

  it "zips each side read ahead on its own goroutine, and with the latest of the other" do
    collected (Strom.zipPar (Strom.range 1 4) (Strom.fromArray [ "a", "b", "c", "d" ])) >>= shouldEqual [ 1 /\ "a", 2 /\ "b", 3 /\ "c" ]
    latest <- collected do
      Strom.zipLatest
        (Strom.succeed "first" <> (Strom.every (Milliseconds 500.0) # map (const "later")))
        (Strom.every (Milliseconds 5.0) # Strom.take 3)
    latest `shouldEqual` [ "first" /\ 0, "first" /\ 1, "first" /\ 2 ]

  it "works on at most mapPar's bound at once, in order" do
    workload <- newWorkload
    results <- collected (Strom.range 1 21 # Strom.mapPar 3 (work workload))
    results `shouldEqual` map (_ * 10) (Array.range 1 20)
    mostAtOnce workload >>= shouldEqual 3

  it "works on at most mapParUnordered's bound at once, as they're done" do
    workload <- newWorkload
    results <- collected (Strom.range 1 21 # Strom.mapParUnordered 4 (work workload))
    Array.sort results `shouldEqual` map (_ * 10) (Array.range 1 20)
    mostAtOnce workload >>= shouldEqual 4

  it "fails mapPar with the first failure" do
    run (Strom.range 1 10 # Strom.mapPar 3 (\n -> if n == 5 then throwError (boom "five") else pure n) # Strom.runDrain)
      >>= shouldEqual (Left "boom: five")

backpressureSpec :: Spec Unit
backpressureSpec = describe "backpressure" do
  it "pulls no element before its reader wants it" do
    leads <- leadsOfProducerOverSlowReader identity
    leads `shouldSatisfy` Array.all (_ == 0)

  it "lets a buffer read ahead by no more than its capacity, and the one it waits to hand over" do
    leads <- leadsOfProducerOverSlowReader (Strom.buffer 4)
    Array.foldl max 0 leads `shouldSatisfy` (_ >= 4)
    Array.foldl max 0 leads `shouldSatisfy` (_ <= 4 + 1)

-- | For each element a slow reader takes, how many more the producer has made by then.
leadsOfProducerOverSlowReader :: (Stream Unit -> Stream Unit) -> Aff (Array Int)
leadsOfProducerOverSlowReader between = do
  produced <- newCounter
  read <- newCounter
  leads <- newIORef [] # toAff
  _ <- run do
    Strom.repeatOm (void (increment produced))
      # between
      # Strom.take 30
      # Strom.forEach \_ -> do
          readSoFar <- increment read
          threadDelay (Milliseconds 1.0) # liftIO
          producedSoFar <- current produced
          atomicModifyIORef' leads (\soFar -> Array.snoc soFar (producedSoFar - readSoFar) /\ unit) # liftIO
  readIORef leads # toAff

-- Helpers

collected :: forall a. Stream a -> Aff (Array a)
collected stream = run (Strom.runCollect stream) >>= case _ of
  Left failure -> fail failure $> []
  Right elements -> pure elements

chunksOf :: forall a. Stream a -> Aff (Array (Array a))
chunksOf stream = collected (stream # Strom.mapChunks Array.singleton)

chunkSizesOf :: forall a. Stream a -> Aff (Array Int)
chunkSizesOf stream = map Array.length <$> chunksOf stream

run :: forall a. Om a -> Aff (Either String a)
run computation = lmap describeFailure <$> toAff (runOmIO unit computation)

boom :: String -> Failure (boom :: String)
boom = Variant.inj (Proxy :: Proxy "boom")

describeFailure :: Failure (boom :: String) -> String
describeFailure = Variant.case_
  # Variant.on (Proxy :: Proxy "boom") ("boom: " <> _)
  # Variant.on (Proxy :: Proxy "exception") message

type Log = IORef (Array String)

newLog :: Aff Log
newLog = newIORef [] # toAff

record :: forall ctx err. Log -> String -> OmIO ctx err Unit
record log entry = atomicModifyIORef' log (\entries -> Array.snoc entries entry /\ unit) # liftIO

recordExit :: forall ctx err. Log -> String -> Exit err -> OmIO ctx err Unit
recordExit log name exit = record log (name <> " " <> describeExit exit)
  where
  describeExit = case _ of
    Ended -> "ended"
    Failed _ -> "failed"
    Interrupted -> "interrupted"

entriesOf :: Log -> Aff (Array String)
entriesOf log = readIORef log # toAff

type Counter = IORef Int

newCounter :: Aff Counter
newCounter = newIORef 0 # toAff

increment :: forall ctx err. Counter -> OmIO ctx err Int
increment counter = atomicModifyIORef' counter (\count -> (count + 1) /\ (count + 1)) # liftIO

current :: forall ctx err. Counter -> OmIO ctx err Int
current counter = readIORef counter # liftIO

countOf :: Counter -> Aff Int
countOf counter = readIORef counter # toAff

pause :: Number -> Aff Unit
pause milliseconds = threadDelay (Milliseconds milliseconds) # toAff

-- | Work that takes a while, counting how many are at it at once.
type Workload = { atWork :: Counter, most :: Counter }

newWorkload :: Aff Workload
newWorkload = do
  atWork <- newCounter
  most <- newCounter
  pure { atWork, most }

work :: forall ctx err. Workload -> Int -> OmIO ctx err Int
work { atWork, most } n = do
  working <- increment atWork
  atomicModifyIORef' most (\before -> max before working /\ unit) # liftIO
  threadDelay (Milliseconds 3.0) # liftIO
  atomicModifyIORef' atWork (\count -> (count - 1) /\ unit) # liftIO
  pure (n * 10)

mostAtOnce :: Workload -> Aff Int
mostAtOnce { most } = countOf most

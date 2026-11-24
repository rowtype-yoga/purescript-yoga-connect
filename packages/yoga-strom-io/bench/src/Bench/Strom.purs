-- | Scales a few million sensor readings and adds them up, through a stream, an array and a
-- | loop, and says how long each took.
module Bench.Strom where

import Prelude

import Data.Array as Array
import Data.Either (either)
import Data.Int (toNumber)
import Data.Maybe (Maybe(..))
import Data.Number.Format (fixed, toStringWith)
import Effect (Effect)
import Go.IO (IO, getNumCapabilities, launchIO_)
import Go.IO.Console (log)
import Go.IO.IORef (modifyIORef, newIORef, readIORef)
import Go.IO.Internal (IO(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = launchIO_ do
  cores <- getNumCapabilities
  let readings = sensorReadings (cores * 500000)
  log ("readings: " <> show (Array.length readings))
  timed "stream, chunks of 4096" \_ -> streamTotal readings
  timed "array map and foldl" \_ -> pure (arrayTotal readings)
  timed "scalar loop" \_ -> pure (loopTotal readings)
  timed "range, map, take, forEach" \_ -> rangeTotal (cores * 500000)
  timed "buffer 16, chunks of 16" \_ -> bufferedTotal readings
  timed "merge of two, chunks of 16" \_ -> mergedTotal readings
  timed "mapPar 8 over 20000" \_ -> mapParTotal 20000

-- | A handover between goroutines for every 16 readings.
bufferedTotal :: Array Number -> IO Number
bufferedTotal readings = Strom.fromArray readings
  # Strom.rechunk 16
  # Strom.buffer 16
  # map calibrated
  # sumOf

-- | Two producers handing chunks of 16 to one reader.
mergedTotal :: Array Number -> IO Number
mergedTotal readings = Strom.merge half half
  # map calibrated
  # sumOf
  where
  half = Strom.fromArray (Array.take (Array.length readings / 2) readings) # Strom.rechunk 16

-- | A worker goroutine and a permit for every element.
mapParTotal :: Int -> IO Number
mapParTotal count = Strom.range 0 count
  # Strom.mapPar 8 (\index -> pure (toNumber index * 1.19))
  # sumOf

sumOf :: Strom.Strom Unit () Number -> IO Number
sumOf stream = do
  total <- runOmIO unit (Strom.runFold 0.0 (+) stream)
  pure (either (const 0.0) identity total)

sensorReadings :: Int -> Array Number
sensorReadings count = Array.range 0 (count - 1) <#> \index -> toNumber (index `mod` 1000) * 0.5

calibrated :: Number -> Number
calibrated reading = reading * 1.19

streamTotal :: Array Number -> IO Number
streamTotal readings = do
  total <- runOmIO unit (calibratedTotal readings)
  pure (either (const 0.0) identity total)

calibratedTotal :: Array Number -> OmIO Unit () Number
calibratedTotal readings = Strom.fromArray readings
  # Strom.rechunk 4096
  # map calibrated
  # Strom.runFold 0.0 (+)

arrayTotal :: Array Number -> Number
arrayTotal readings = readings # map calibrated # Array.foldl (+) 0.0

loopTotal :: Array Number -> Number
loopTotal readings = addFrom 0 0.0
  where
  addFrom index total = case Array.index readings index of
    Nothing -> total
    Just reading -> addFrom (index + 1) (total + calibrated reading)

-- | The pipeline from the README: a range, mapped, taken from, and each element sent on.
rangeTotal :: Int -> IO Number
rangeTotal count = do
  sent <- newIORef 0
  _ <- runOmIO unit (sendTripled count \tripled -> modifyIORef sent (_ + tripled) # liftIO)
  toNumber <$> readIORef sent

sendTripled :: Int -> (Int -> OmIO Unit () Unit) -> OmIO Unit () Unit
sendTripled count send = Strom.range 0 count
  # map (_ * 3)
  # Strom.take (count / 2)
  # Strom.forEach send

timed :: String -> (Unit -> IO Number) -> IO Unit
timed name computation = do
  _ <- computation unit
  started <- nowInMilliseconds
  totals <- Array.replicate rounds unit # traverseInOrder computation
  finished <- nowInMilliseconds
  log (name <> ": " <> toStringWith (fixed 2) ((finished - started) / toNumber rounds) <> " ms a run, total " <> show (Array.last totals))
  where
  rounds = 20

traverseInOrder :: forall a b. (a -> IO b) -> Array a -> IO (Array b)
traverseInOrder f elements = Array.foldM (\done element -> Array.snoc done <$> f element) [] elements

nowInMilliseconds :: IO Number
nowInMilliseconds = IO nowInMillisecondsImpl

foreign import nowInMillisecondsImpl :: Effect Number

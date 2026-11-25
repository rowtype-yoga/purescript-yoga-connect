module Main where

import Prelude

import Control.Monad.Error.Class (throwError)
import Data.Array as Array
import Data.DateTime.Instant (unInstant)
import Data.Either (Either(..))
import Data.Foldable (maximumBy)
import Data.Int (round, toNumber)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (un)
import Data.Number as Number
import Data.Number.Format (fixed, toStringWith)
import Data.Ord (comparing)
import Data.String (Pattern(..))
import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant as Variant
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Exception (message)
import Effect.Now (now)
import Go.IO (IO, launchIO_, threadDelay)
import Go.IO.Console (log)
import Go.IO.IORef (newIORef, readIORef, writeIORef)
import Type.Proxy (Proxy(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom (Exit(..), Failure, Strom)
import Yoga.OmIO.Strom as Strom

data Service = Api | Auth | Billing

derive instance Eq Service

data Level = Info | Warn | Error

derive instance Eq Level

newtype Route = Route String

newtype OnCall = OnCall String

type LogLine = { service :: Service, level :: Level, latency :: Milliseconds, route :: Route }

type Window = { lines :: Int, errors :: Int, meanLatency :: Milliseconds, slowest :: Milliseconds, noisiest :: Service }

type Totals = { lines :: Int, errors :: Int, slowest :: Milliseconds }

type Errors = (rotated :: Service)

main :: Effect Unit
main = launchIO_ do
  started <- now # liftEffect
  outcome <- runOmIO unit (Strom.runLast summaryOfTheLogs)
  finished <- now # liftEffect
  report outcome (millisecondsBetween started finished)
  where
  millisecondsBetween from to = un Milliseconds (unInstant to) - un Milliseconds (unInstant from)

summaryOfTheLogs :: Strom Unit Errors Totals
summaryOfTheLogs = Strom.mergeAll (tailFollowingRotations <$> services)
  # Strom.filterMap parseLine
  # Strom.grouped 50_000
  # Strom.take 12
  # map summarize
  # Strom.mapPar 4 findWhoIsOnCall
  # Strom.tap (printWindow >>> liftIO)
  # Strom.scan addUp { lines: 0, errors: 0, slowest: Milliseconds 0.0 }
  # Strom.timeout (Milliseconds 2000.0)

tailFollowingRotations :: Service -> Strom Unit Errors String
tailFollowingRotations service = tailOf service # Strom.catchAll reopenIfRotated
  where
  reopenIfRotated failure = failure # Variant.on (Proxy :: _ "rotated") (\_ -> tailFollowingRotations service) (\_ -> Strom.fail failure)

tailOf :: Service -> Strom Unit Errors String
tailOf service = Strom.bracketExit opened closed \linesRead -> Strom.range 0 1_000_000_000
  # Strom.tap (countOrRotate linesRead)
  # map (lineNumbered service)
  where
  opened = log ("opened " <> fileOf service) *> newIORef 0 # liftIO
  closed linesRead exit = do
    count <- readIORef linesRead # liftIO
    log ("closed " <> fileOf service <> " after " <> show count <> " lines, " <> howItEnded exit) # liftIO
  countOrRotate linesRead index
    | service == Billing && index == 100_000 = throwError (Variant.inj (Proxy :: _ "rotated") service)
    | otherwise = writeIORef linesRead (index + 1) # liftIO

parseLine :: String -> Maybe LogLine
parseLine line = case String.split (Pattern " ") line of
  [ service, level, latency, route ] -> { service: _, level: _, latency: _, route: Route route } <$> serviceNamed service <*> levelNamed level <*> inMilliseconds latency
  _ -> Nothing
  where
  serviceNamed name = Array.find (nameOf >>> eq name) services
  levelNamed = case _ of
    "INFO" -> Just Info
    "WARN" -> Just Warn
    "ERROR" -> Just Error
    _ -> Nothing
  inMilliseconds text = String.stripSuffix (Pattern "ms") text >>= Number.fromString <#> Milliseconds

summarize :: Array LogLine -> Window
summarize lines = { lines: Array.length lines, errors: Array.length errors, meanLatency, slowest, noisiest }
  where
  latencies = lines <#> _.latency >>> un Milliseconds
  meanLatency = Milliseconds (Array.foldl (+) 0.0 latencies / toNumber (Array.length latencies))
  slowest = Milliseconds (Array.foldl max 0.0 latencies)
  errors = Array.filter (_.level >>> eq Error) lines
  noisiest = maximumBy (comparing errorsFrom) services # fromMaybe Api
  errorsFrom service = Array.length (Array.filter (_.service >>> eq service) errors)

findWhoIsOnCall :: Window -> OmIO Unit Errors (Window /\ OnCall)
findWhoIsOnCall window = threadDelay (Milliseconds 40.0) # liftIO <#> \_ -> window /\ onCallFor window.noisiest
  where
  onCallFor = case _ of
    Api -> OnCall "Ada"
    Auth -> OnCall "Grace"
    Billing -> OnCall "Edsger"

printWindow :: Window /\ OnCall -> IO Unit
printWindow ({ lines, errors, meanLatency, slowest, noisiest } /\ OnCall person) = log
  (show lines <> " lines, " <> show errors <> " errors, mostly " <> fileOf noisiest <> " (paging " <> person <> "), mean " <> inMs meanLatency <> ", slowest " <> inMs slowest)

addUp :: Totals -> Window /\ OnCall -> Totals
addUp totals (window /\ _) = { lines: totals.lines + window.lines, errors: totals.errors + window.errors, slowest: max totals.slowest window.slowest }

report :: Either (Failure Errors) (Maybe Totals) -> Number -> IO Unit
report outcome elapsed = log case outcome of
  Right (Just { lines, errors, slowest }) -> "In total " <> show lines <> " lines and " <> show errors <> " errors, the slowest took " <> inMs slowest <> ", " <> show (round (toNumber lines / elapsed * 1000.0)) <> " lines a second"
  Right Nothing -> "No logs."
  Left failure -> "Gave up: " <> described failure

services :: Array Service
services = [ Api, Auth, Billing ]

nameOf :: Service -> String
nameOf = case _ of
  Api -> "api"
  Auth -> "auth"
  Billing -> "billing"

fileOf :: Service -> String
fileOf service = nameOf service <> ".log"

-- | Line `index` of a made-up log. Every service has bad stretches, with more errors and slower
-- | answers, and errors are slowest of all.
lineNumbered :: Service -> Int -> String
lineNumbered service index = String.joinWith " " [ nameOf service, level, show latency <> "ms", route ]
  where
  noise = (index `mod` 10007 * 7919 + String.length (nameOf service)) `mod` 10007
  inABadStretch = (index / 30_000 + String.length (nameOf service)) `mod` 3 == 0
  isError = noise `mod` (if inABadStretch then 25 else 150) == 0
  level = if isError then "ERROR" else if noise `mod` 7 == 0 then "WARN" else "INFO"
  latency = noise `mod` (if inABadStretch then 600 else 200) + (if isError then noise `mod` 3000 else 0)
  route = fromMaybe "/" (Array.index [ "/checkout", "/login", "/cart", "/search" ] (noise `mod` 4))

howItEnded :: Exit Errors -> String
howItEnded = case _ of
  Ended -> "the end"
  Failed failure -> described failure
  Interrupted -> "nobody needs more"

described :: Failure Errors -> String
described = Variant.case_
  # Variant.on (Proxy :: _ "rotated") (\service -> fileOf service <> " was rotated")
  # Variant.on (Proxy :: _ "exception") message

inMs :: Milliseconds -> String
inMs (Milliseconds ms) = toStringWith (fixed 1) ms <> " ms"

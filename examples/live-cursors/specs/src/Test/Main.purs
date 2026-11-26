module Test.Main where

import Prelude

import Bubbletea (MouseAction(..), MouseButton(..), Running)
import Bubbletea as Tea
import Bubbletea.Hooks as Hooks
import Bubbletea.Key (KeyCode(..))
import Bubbletea.Key as Key
import Data.String.CodeUnits as CodeUnits
import Bubbletea.Theme (charmtonePantera)
import Cursors.Client (canvasApp)
import Cursors.Server (newRoom, serveCanvas)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String (Pattern(..), contains)
import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Effect (Effect)
import Effect.Aff (Aff, launchAff_)
import Go.IO (IO, threadDelay, toAff)
import Test.Spec (describe, it)
import Data.Foldable (for_)
import Test.Spec.Assertions (shouldSatisfy)
import Test.Spec.Reporter.Console (consoleReporter)
import Test.Spec.Runner (defaultConfig, runSpec')
import Yoga.Connect.Om (Port(..))

main :: Effect Unit
main = launchAff_ do
  runSpec' defaultConfig { timeout = Just (Milliseconds 8000.0) } [ consoleReporter ] do
    describe "the shared canvas" do
      it "shows one participant's cursor, with their name, to the other" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 3
          settle
          frameOf bob
        lineOf 3 frame `shouldSatisfy` startsWithAt 10 "▲ada"
        lineOf statusRow frame `shouldSatisfy` contains (Pattern "● bob (you)")
        lineOf statusRow frame `shouldSatisfy` contains (Pattern "● ada")

      it "draws what one participant drags over, for everyone" do
        address <- startCanvas
        frames <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          drag ada [ { x: 5, y: 5 }, { x: 6, y: 5 }, { x: 7, y: 5 } ]
          settle
          seenByAda <- frameOf ada
          seenByBob <- frameOf bob
          pure { seenByAda, seenByBob }
        lineOf 5 frames.seenByBob `shouldSatisfy` startsWithAt 5 "███"
        lineOf 5 frames.seenByAda `shouldSatisfy` startsWithAt 5 "███"

      it "shows a newcomer what was drawn before they came" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          settle
          drag ada [ { x: 2, y: 1 }, { x: 3, y: 1 } ]
          settle
          cleo <- participant "cleo" address
          settle
          frameOf cleo
        lineOf 1 frame `shouldSatisfy` startsWithAt 2 "██"

      it "takes a participant's cursor away when they leave" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 3
          settle
          Tea.stop ada
          settle
          frameOf bob
        lineOf 3 frame `shouldSatisfy` (not <<< contains (Pattern "ada"))
        lineOf statusRow frame `shouldSatisfy` (not <<< contains (Pattern "ada"))

      it "turns down a name that's taken" do
        address <- startCanvas
        frame <- toAff do
          _ <- participant "ada" address
          settle
          impostor <- participant "ada" address
          settle
          frameOf impostor
        lineOf statusRow frame `shouldSatisfy` contains (Pattern "ada is here already, pick another name")

      it "says it's joining again while the canvas is away" do
        frame <- toAff do
          lonely <- participant "ada" "http://127.0.0.1:1"
          settle
          frameOf lonely
        lineOf statusRow frame `shouldSatisfy` contains (Pattern "joining again, unavailable")

    describe "saying something" do
      it "shows the others every letter as it's typed, keys like q and 1 included, then the line sent" do
        address <- startCanvas
        frames <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 5
          press ada Enter
          typeOut ada "q1 hi"
          settle
          whileTyping <- frameOf bob
          typedByAda <- frameOf ada
          press ada Backspace
          typeOut ada "ey"
          press ada Enter
          settle
          afterSending <- frameOf bob
          pure { whileTyping, typedByAda, afterSending }
        lineOf 4 frames.whileTyping `shouldSatisfy` startsWithAt 10 " q1 hi▌"
        lineOf 5 frames.whileTyping `shouldSatisfy` startsWithAt 10 "▲ada"
        lineOf statusRow frames.typedByAda `shouldSatisfy` contains (Pattern "› q1 hi▌")
        lineOf 4 frames.afterSending `shouldSatisfy` startsWithAt 10 " q1 hey "

      it "takes the bubble away when the line's given up on with Esc" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 5
          press ada Enter
          typeOut ada "oops"
          settle
          press ada Esc
          settle
          frameOf bob
        lineOf 4 frame `shouldSatisfy` (not <<< contains (Pattern "oops"))
        lineOf 5 frame `shouldSatisfy` startsWithAt 10 "▲ada"

      it "takes the bubble away when the speaker leaves" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 5
          press ada Enter
          typeOut ada "brb"
          settle
          Tea.stop ada
          settle
          frameOf bob
        lineOf 4 frame `shouldSatisfy` (not <<< contains (Pattern "brb"))

      it "shows a newcomer what's being typed right now" do
        address <- startCanvas
        frame <- toAff do
          ada <- participant "ada" address
          settle
          moveTo ada 10 5
          press ada Enter
          typeOut ada "hey cleo"
          settle
          cleo <- participant "cleo" address
          settle
          frameOf cleo
        lineOf 4 frame `shouldSatisfy` startsWithAt 10 " hey cleo▌"

      it "lets a line that was said go after a few seconds" do
        address <- startCanvas
        frames <- toAff do
          ada <- participant "ada" address
          bob <- participant "bob" address
          settle
          moveTo ada 10 5
          press ada Enter
          typeOut ada "bye"
          press ada Enter
          settle
          justSaid <- frameOf bob
          threadDelay (Milliseconds 4500.0)
          aWhileLater <- frameOf bob
          pure { justSaid, aWhileLater }
        lineOf 4 frames.justSaid `shouldSatisfy` startsWithAt 10 " bye "
        lineOf 4 frames.aWhileLater `shouldSatisfy` (not <<< contains (Pattern "bye"))

-- | A canvas server on a port of its own.
startCanvas :: Aff String
startCanvas = toAff do
  room <- newRoom
  bound <- serveCanvas room (Port 0)
  pure ("http://127.0.0.1:" <> show (portNumber bound.port))
  where
  portNumber (Port number) = number

-- | A client without a terminal, in a 40 by 12 window.
participant :: String -> String -> IO (Running Unit)
participant name address = Hooks.program charmtonePantera (canvasApp { name, address }) >>= Tea.startProgramHeadless window

window :: Tea.Size
window = { width: 40, height: 12 }

statusRow :: Int
statusRow = window.height - 1

moveTo :: Running Unit -> Int -> Int -> IO Unit
moveTo running x y = mouse running Motion NoButton x y

drag :: Running Unit -> Array { x :: Int, y :: Int } -> IO Unit
drag running points = do
  for_ (Array.head points) \{ x, y } -> mouse running Click LeftButton x y
  for_ (Array.drop 1 points) \{ x, y } -> mouse running Motion LeftButton x y
  for_ (Array.last points) \{ x, y } -> mouse running Release LeftButton x y

mouse :: Running Unit -> MouseAction -> MouseButton -> Int -> Int -> IO Unit
mouse running action button x y = Tea.sendEvent running (Tea.Mouse { x, y, button, action, ctrl: false, alt: false, shift: false } :: Tea.Event Unit)

press :: Running Unit -> KeyCode -> IO Unit
press running code = Tea.sendEvent running (Tea.KeyPress (Key.key code) :: Tea.Event Unit)

-- | One key press a letter, the way someone types.
typeOut :: Running Unit -> String -> IO Unit
typeOut running text = for_ (CodeUnits.toCharArray text) \character -> press running (if character == ' ' then Space else Char character)

frameOf :: Running Unit -> IO String
frameOf running = Tea.lastFrame running <#> Tea.stripAnsi

lineOf :: Int -> String -> String
lineOf row frame = String.split (Pattern "\n") frame # flip Array.index row # fromMaybe ""

-- | Whether the text is at that column of the line.
startsWithAt :: Int -> String -> String -> Boolean
startsWithAt column text line = String.take (String.length text) (String.drop column line) == text

-- | Long enough for a scene to get through the server and back, cursor moves going out every
-- | 25 milliseconds.
settle :: IO Unit
settle = threadDelay (Milliseconds 150.0)

module Test.Main where

import Prelude

import Bubbletea (MouseAction(..), MouseButton(..), Running)
import Bubbletea as Tea
import Bubbletea.Hooks as Hooks
import Bubbletea.Theme (charmtonePantera)
import Data.Array as Array
import Data.Foldable (for_)
import Data.Maybe (fromMaybe)
import Data.String (Pattern(..))
import Data.String as String
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested (type (/\), (/\))
import Draw.Client (canvasApp)
import Draw.Server (newRoom, serveCanvas)
import Effect (Effect)
import Effect.Aff (Aff, launchAff_)
import Go.IO (IO, threadDelay, toAff)
import Test.Spec (describe, it)
import Test.Spec.Assertions (shouldEqual)
import Test.Spec.Reporter.Console (consoleReporter)
import Test.Spec.Runner (defaultConfig, runSpec')
import Yoga.Connect.Om (Port(..))

main :: Effect Unit
main = launchAff_ do
  runSpec' defaultConfig [ consoleReporter ] do
    describe "the shared canvas" do
      it "shows bob what ada drew, and ada what bob drew" do
        address <- startCanvas
        seen <- toAff do
          ada <- clientOn address
          bob <- clientOn address
          drag ada [ 2 /\ 1, 3 /\ 1, 4 /\ 1 ]
          settle
          drag bob [ 0 /\ 3 ]
          settle
          byAda <- frameOf ada
          byBob <- frameOf bob
          pure { byAda, byBob }
        lineOf 1 seen.byBob `shouldEqual` "  ███"
        lineOf 3 seen.byBob `shouldEqual` "█    "
        lineOf 3 seen.byAda `shouldEqual` "█    "

-- | A canvas server on a free port.
startCanvas :: Aff String
startCanvas = toAff do
  room <- newRoom
  bound <- serveCanvas room (Port 0)
  pure ("http://127.0.0.1:" <> show (portNumber bound.port))
  where
  portNumber (Port number) = number

-- | A client without a terminal.
clientOn :: String -> IO (Running Unit)
clientOn address = Hooks.program charmtonePantera (canvasApp address) >>= Tea.startProgramHeadless { width: 40, height: 12 }

-- | Presses the left button on the first spot, moves over the rest and lets go on the last.
drag :: Running Unit -> Array (Int /\ Int) -> IO Unit
drag running spots = do
  for_ (Array.head spots) (mouse Click)
  for_ (Array.drop 1 spots) (mouse Motion)
  for_ (Array.last spots) (mouse Release)
  where
  mouse action (x /\ y) = Tea.sendEvent running (Tea.Mouse { x, y, button: LeftButton, action, ctrl: false, alt: false, shift: false } :: Tea.Event Unit)

frameOf :: Running Unit -> IO String
frameOf running = Tea.lastFrame running <#> Tea.stripAnsi

lineOf :: Int -> String -> String
lineOf row frame = String.split (Pattern "\n") frame # flip Array.index row # fromMaybe ""

-- | Long enough for a cell to get to the server and back.
settle :: IO Unit
settle = threadDelay (Milliseconds 150.0)

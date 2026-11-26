-- | The canvas in your terminal. Click or drag to paint, `q` quits.
module Draw.Client (main, canvasApp) where

import Prelude

import Bubbletea as Tea
import Bubbletea.Hooks (Action, Component, command, component, get, liftIO, modify, quit, render, runComponent, useCommand, useEvent, useKey, useState)
import Bubbletea.Key (KeyCode(..))
import Data.Array as Array
import Data.Either (Either)
import Data.Foldable (maximum)
import Data.Maybe (maybe)
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.Tuple (fst, snd)
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant (Variant)
import Draw.Service (Canvas)
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Exception (Error)
import Go.IO (IO)
import Go.IO.Chan (Chan, newBufferedChan, receive, send, trySend)
import Go.IO.Context (Context)
import Yoga.Connect.Om (ConnectError, client, connect, runOmIO)
import Yoga.Connect.Om as Om
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = runComponent (canvasApp "http://127.0.0.1:50053")

canvasApp :: String -> Component Unit Void
canvasApp address = component \_ -> do
  canvas <- useState Set.empty
  mine <- newBufferedChan 256 # liftIO
  theirs <- newBufferedChan 1024 # liftIO
  useCommand (Tea.performWith \context -> drawAlong address mine theirs context $> pure unit)
  useCommand (showNext theirs \spot -> modify canvas (Set.insert spot))
  useEvent case _ of
    Tea.Mouse { x, y, button: Tea.LeftButton, action } | action /= Tea.Release -> void (trySend mine (x /\ y) # liftIO)
    _ -> pure unit
  useKey [ Char 'q' ] quit
  render (get canvas <#> \spots -> Tea.view (picture spots) # Tea.mouse Tea.CellMotion # Tea.altScreen)

-- | Sends what you paint and hands over what anyone paints, until you quit or the server goes away.
drawAlong :: String -> Chan Spot -> Chan Spot -> Context -> IO (Either HungUp Unit)
drawAlong address mine theirs context = do
  connection <- connect address # liftEffect
  runOmIO unit ((client @Canvas connection).draw yourCells # Strom.forEachWithin context \{ x, y } -> send theirs (x /\ y) # Om.liftIO)
  where
  yourCells = Strom.fromChan mine <#> \(x /\ y) -> { x, y }

-- | Each spot from the call's goroutine, onto Bubble Tea's.
showNext :: Chan Spot -> (Spot -> Action Unit) -> Tea.Cmd (Action Unit)
showNext theirs show = Tea.perform (receive theirs <#> maybe (pure unit) \spot -> show spot *> command (showNext theirs show))

-- | A block for every painted spot, as far down and across as anyone painted.
picture :: Set Spot -> String
picture spots = String.joinWith "\n" (upTo snd <#> \y -> String.joinWith "" (upTo fst <#> \x -> glyphAt x y))
  where
  glyphAt x y = if Set.member (x /\ y) spots then "█" else " "
  upTo coordinate = maybe [] (Array.range 0) (maximum (Array.fromFoldable spots <#> coordinate))

type Spot = Int /\ Int

type HungUp = Variant (connectError :: ConnectError, exception :: Error)

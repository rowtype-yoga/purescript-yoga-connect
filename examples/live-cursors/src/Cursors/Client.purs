-- | The shared canvas in your terminal: everyone else's mouse as a colored marker with their name,
-- | what everyone drew, what everyone's saying, and who's here. Click or drag to draw, 1 to 6 picks
-- | a color, Enter starts saying something and sends it. It joins again when the server comes back.
module Cursors.Client
  ( main
  , canvasApp
  , Settings
  ) where

import Prelude

import Bubbletea as Tea
import Control.Alt ((<|>))
import Bubbletea.Hooks (Action, Component, Hooks, State, command, component, get, liftIO, modify, mountEach, quit, render, runComponent, set, useCommand, useEvent, useKeystroke, useSize, useState)
import Bubbletea.Key (Key, KeyCode(..))
import Bubbletea.Key as Key
import Cursors.Plain (Bubble, Gesture, Scene, Someone)
import Cursors.Service (Canvas, Color(..), Voice(..))
import Data.Array as Array
import Data.Either (either)
import Data.Foldable (foldl, for_)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, maybe)
import Data.String as String
import Data.String.CodeUnits as CodeUnits
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Class (liftEffect)
import Effect.Exception (Error, message)
import Go.IO (IO, threadDelayWith)
import Go.IO.Chan (Chan, newBufferedChan, receive, send, trySend)
import Go.IO.Context (Context)
import Go.OS (Arg(..), args)
import Lipgloss as L
import Lipgloss.Charmtone as C
import Lipgloss.Sx (Sx(..))
import Lipgloss.Sx as Style
import Type.Proxy (Proxy(..))
import Yoga.Connect.Om (ConnectError(..), client, codeName, connect, runOmIO)
import Yoga.Connect.Om as Om
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = do
  given <- args <#> map \(Arg arg) -> arg
  runComponent (canvasApp { name: Array.index given 0 # fromMaybe "guest", address: Array.index given 1 # fromMaybe "http://127.0.0.1:50052" })

type Settings = { name :: String, address :: String }

canvasApp :: Settings -> Component Unit Void
canvasApp settings = component \_ -> do
  canvas <- useState nobodyYet
  link <- useState Joining
  brush <- useState Nothing
  size <- useSize
  gestures <- newBufferedChan 256 # liftIO
  useCanvas settings gestures case _ of
    Arrived scene -> set link Live *> modify canvas (seeing scene)
    Lost why -> set link (Rejoining why) *> set canvas nobodyYet
  useEvent case _ of
    Tea.Mouse event -> do
      painting <- get brush
      seen <- get canvas
      window <- get size
      for_ (gestureOf event (painting <|> (seen.you <#> _.color)) window) \gesture -> void (trySend gestures gesture # liftIO)
    _ -> pure unit
  draft <- useState Nothing
  useEvent case _ of
    Tea.KeyPress pressed -> get draft >>= \typed -> keyPressed { draft, brush, gestures } typed pressed
    _ -> pure unit
  useKeystroke [ "ctrl+c" ] quit
  badges <- mountEach badge (get canvas <#> everyone) \_ -> absurd
  render ado
    seen <- get canvas
    status <- get link
    window <- get size
    painting <- get brush
    typed <- get draft
    shown <- badges
    in Tea.view (screen settings window seen status (painting <|> (seen.you <#> _.color)) typed (shown <#> Tea.content)) # Tea.mouse Tea.AllMotion # Tea.altScreen

-- | On the canvas, Enter starts saying something, 1 to 6 pick a color and q or Esc leave. While
-- | saying something every key types and everyone sees it at once; Enter sends it, Esc gives up.
keyPressed :: Keys -> Maybe String -> Key -> Action Unit
keyPressed keys typed pressed = typed # maybe onCanvas whileTyping
  where
  onCanvas
    | Key.is Enter pressed = set keys.draft (Just "")
    | Key.is Esc pressed || Key.is (Char 'q') pressed = quit
    | otherwise = for_ colorPicked \color -> set keys.brush (Just color)
  whileTyping text
    | Key.is Enter pressed = set keys.draft Nothing *> tell (if text == "" then typing "" else said text)
    | Key.is Esc pressed = set keys.draft Nothing *> tell (typing "")
    | Key.is Backspace pressed = retype (String.take (String.length text - 1) text)
    | pressed.text /= "" && not pressed.ctrl && not pressed.alt = retype (text <> pressed.text)
    | otherwise = pure unit
  retype text = set keys.draft (Just text) *> tell (typing text)
  tell gesture = void (trySend keys.gestures gesture # liftIO)
  typing text = Variant.inj (Proxy :: Proxy "typing") { text }
  said text = Variant.inj (Proxy :: Proxy "said") { text }
  colorPicked = Array.mapWithIndex (\index color -> digit (index + 1) /\ color) palette
    # Array.find (\(character /\ _) -> Key.is (Char character) pressed)
    <#> \(_ /\ color) -> color
  digit n = CodeUnits.charAt 0 (show n) # fromMaybe '0'

-- | What a key press can change: what you're saying, your color, and what goes to the canvas.
type Keys = { draft :: State (Maybe String), brush :: State (Maybe Color), gestures :: Chan Gesture }

data Link = Joining | Live | Rejoining String

data Arrival = Arrived Scene | Lost String

-- | What this participant knows of the canvas.
type CanvasView =
  { you :: Maybe Someone
  , others :: Map String Someone
  , cells :: Map (Int /\ Int) Color
  , bubbles :: Map String Bubble
  }

nobodyYet :: CanvasView
nobodyYet = { you: Nothing, others: Map.empty, cells: Map.empty, bubbles: Map.empty }

seeing :: Scene -> CanvasView -> CanvasView
seeing = Variant.match
  { welcome: \{ you, others, cells, bubbles } _ -> { you: Just you, others: Map.fromFoldable (others <#> \other -> other.name /\ other), cells: cellMap cells, bubbles: bubbleMap bubbles }
  , cursors: \{ moved } seen -> seen { others = foldl (\others { name, at } -> Map.update (\other -> Just other { at = Just at }) name others) seen.others moved }
  , drawn: \{ cells } seen -> seen { cells = Map.union (cellMap cells) seen.cells }
  , joined: \other seen -> seen { others = Map.insert other.name other seen.others }
  , left: \{ name } seen -> seen { others = Map.delete name seen.others, bubbles = Map.delete name seen.bubbles }
  , speech: \{ bubbles, hushed } seen -> seen { bubbles = Map.union (bubbleMap bubbles) (foldl (flip Map.delete) seen.bubbles hushed) }
  }
  where
  cellMap cells = Map.fromFoldable (cells <#> \{ x, y, color } -> (x /\ y) /\ color)
  bubbleMap bubbles = Map.fromFoldable (bubbles <#> \bubble -> bubble.name /\ bubble)

-- | Moving the mouse moves your cursor; clicking, or moving with the button down, draws.
gestureOf :: Tea.MouseEvent -> Maybe Color -> Tea.Size -> Maybe Gesture
gestureOf { x, y, button, action } painting window
  | y >= canvasHeight window = Nothing
  | button == Tea.LeftButton && (action == Tea.Click || action == Tea.Motion) = painting <#> \color -> Variant.inj (Proxy :: Proxy "drew") { x, y, color }
  | action == Tea.Motion = Just (Variant.inj (Proxy :: Proxy "moved") { x, y })
  | otherwise = Nothing

-- | Joins on a goroutine of its own and hands each scene to the action, on Bubble Tea's. When the
-- | canvas goes away it joins again every second; quitting cancels the call, which is leaving.
useCanvas :: Settings -> Chan Gesture -> (Arrival -> Action Unit) -> Hooks Unit
useCanvas settings gestures arrive = do
  arrivals <- newBufferedChan 1024 # liftIO
  useCommand (Tea.performWith (joinInto settings gestures arrivals))
  useCommand (handleNext arrivals)
  where
  handleNext arrivals = Tea.perform (receive arrivals <#> maybe (pure unit) \arrival -> arrive arrival *> command (handleNext arrivals))

joinInto :: Settings -> Chan Gesture -> Chan Arrival -> Context -> IO (Action Unit)
joinInto settings gestures arrivals context = do
  connection <- connect settings.address # liftEffect
  joined <- runOmIO unit do
    (client @Canvas connection).join (Strom.succeed hello <> Strom.fromChan gestures)
      # Strom.forEachWithin context \scene -> send arrivals (Arrived scene) # Om.liftIO
  send arrivals (Lost (either describeFailure (const "the canvas hung up") joined))
  stillHere <- threadDelayWith context (Milliseconds 1000.0)
  if stillHere then joinInto settings gestures arrivals context else pure (pure unit)
  where
  hello = Variant.inj (Proxy :: Proxy "hello") { name: settings.name }

describeFailure :: Variant (exception :: Error | Failures) -> String
describeFailure = Variant.match
  { nameTaken: identity
  , invalidArgument: identity
  , connectError: \(ConnectError { code }) -> codeName code
  , exception: message
  }

-- | One participant in the status line, in their color.
badge :: Component { person :: Someone, you :: Boolean } Void
badge = component \{ props } -> render ado
  { person, you } <- props
  in Tea.view (Style.render (Fg (colorOf person.color) <> Bold you) ("● " <> person.name <> (if you then " (you)" else "")))

everyone :: CanvasView -> Array { person :: Someone, you :: Boolean }
everyone seen = (seen.you # maybe [] \you -> [ { person: you, you: true } ]) <> (Map.values seen.others # Array.fromFoldable <#> { person: _, you: false })

-- | The canvas, and below it who's here and what you can do, or what you're saying.
screen :: Settings -> Tea.Size -> CanvasView -> Link -> Maybe Color -> Maybe String -> Array String -> String
screen settings window seen link painting typed badges = L.joinVertical L.left [ canvasOf window seen, statusLine ]
  where
  statusLine = String.joinWith "  " (badges <> [ status ])
  status = case link, typed of
    Joining, _ -> Style.render faint ("joining " <> settings.address <> "…")
    Rejoining why, _ -> Style.render alarm ("joining again, " <> why)
    Live, Just text -> Style.render yours ("› " <> text <> "▌") <> Style.render faint " enter sends · esc cancels"
    Live, Nothing -> painted <> Style.render faint " 1–6 color · click or drag draws · enter says · q quits"
  painted = painting # maybe "" \color -> Style.render (Fg (colorOf color)) "■"
  yours = painting # maybe None (Fg <<< colorOf)

-- | The cells everyone drew, everyone else's cursor with their name after it and what they're
-- | saying above it, as far as the terminal reaches.
canvasOf :: Tea.Size -> CanvasView -> String
canvasOf window seen
  | window.width <= 0 = ""
  | otherwise = String.joinWith "\n" (rowsOf (canvasHeight window) \y -> String.joinWith "" (columnsOf window.width \x -> glyphAt x y))
  where
  glyphs = Map.unions [ bubbles, markers, seen.cells <#> \color -> "█" /\ Fg (colorOf color) ]
  markers = Map.fromFoldable (othersOnCanvas >>= markerOf)
  markerOf other = other.at # maybe [] \at -> lettersAt at (Fg (colorOf other.color)) ("▲" <> other.name)
  bubbles = Map.fromFoldable (othersOnCanvas >>= bubbleOf)
  bubbleOf other = fromMaybe [] ado
    at <- other.at
    saying <- Map.lookup other.name seen.bubbles
    in lettersAt (aboveOrBelow at) (bubbleStyle other.color saying.voice) (bubbleText saying)
  othersOnCanvas = Map.values seen.others # Array.fromFoldable
  aboveOrBelow { x, y } = { x, y: if y == 0 then 1 else y - 1 }
  glyphAt x y = Map.lookup (x /\ y) glyphs # maybe " " \(letter /\ style) -> Style.render style letter
  rowsOf count f = if count <= 0 then [] else Array.range 0 (count - 1) <#> f
  columnsOf count f = Array.range 0 (count - 1) <#> f

-- | One letter a cell, from where the text starts.
lettersAt :: { x :: Int, y :: Int } -> Sx -> String -> Array ((Int /\ Int) /\ (String /\ Sx))
lettersAt { x, y } style text = String.toCodePointArray text # Array.mapWithIndex \offset letter -> ((x + offset) /\ y) /\ (String.singleton letter /\ style)

-- | The end of what's being typed, with a caret, or what was said.
bubbleText :: Bubble -> String
bubbleText { text, voice } = " " <> latest <> (if voice == Typing then "▌" else " ")
  where
  latest
    | String.length text <= widestBubble = text
    | otherwise = "…" <> String.drop (String.length text - widestBubble + 1) text

widestBubble :: Int
widestBubble = 30

-- | In the speaker's color, and faint once it's on its way out.
bubbleStyle :: Color -> Voice -> Sx
bubbleStyle color = case _ of
  Fading -> Fg (colorOf color) <> Faint true
  _ -> Bg (colorOf color) <> Fg C.pepper

-- | Everything but the status line.
canvasHeight :: Tea.Size -> Int
canvasHeight window = window.height - 1

palette :: Array Color
palette = [ Coral, Tang, Butter, Guac, Malibu, Dolly ]

colorOf :: Color -> L.Color
colorOf = case _ of
  Coral -> C.coral
  Tang -> C.tang
  Butter -> C.butter
  Guac -> C.guac
  Malibu -> C.malibu
  Dolly -> C.dolly

faint :: Sx
faint = Fg C.squid

alarm :: Sx
alarm = Fg C.coral

type Failures = (nameTaken :: String, invalidArgument :: String, connectError :: ConnectError)

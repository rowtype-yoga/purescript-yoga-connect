-- | The canvas everyone shares. Each participant's call is two streams: their gestures coming in,
-- | the scenes going out. Cursor moves pile up and go out together every 25 milliseconds, the
-- | latest position of each cursor winning, so a fast mouse doesn't flood anyone.
module Cursors.Server
  ( main
  , serveCanvas
  , Room
  , newRoom
  ) where

import Prelude

import Cursors.Plain (Cell, Gesture, Point, Scene, Someone)
import Cursors.Service (Canvas, Color(..), Joining)
import Data.Array as Array
import Data.Foldable (for_)
import Data.Map (Map)
import Data.Map as Map
import Data.Int as Int
import Data.Maybe (Maybe(..), fromMaybe, isJust, maybe)
import Data.Newtype (un)
import Data.Time.Duration (Milliseconds(..))
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect (Effect)
import Effect.Class (liftEffect)
import Go.IO (IO, forkIO, launchIO_)
import Go.IO.Chan (Chan, newBufferedChan, trySend)
import Go.IO.Console (log)
import Go.IO.IORef (IORef, atomicModifyIORef', newIORef, readIORef, writeIORef)
import Go.OS (Arg(..), args)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Om (BoundAddress, Handler, Host(..), OmIO, Port(..), ask, createServer, handle, liftIO, listen, registerServiceWith, reject, runOmIO)
import Yoga.OmIO.Strom (Exit, Strom)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = launchIO_ do
  port <- args <#> (Array.head >=> portOf) >>> fromMaybe 50052 # liftEffect
  room <- newRoom
  bound <- serveCanvas room (Port port)
  log ("canvas on 127.0.0.1:" <> show (un Port bound.port))
  where
  portOf (Arg given) = Int.fromString given

-- | Serves the room's canvas on the port, and sends everyone the moves that piled up, every 25
-- | milliseconds.
serveCanvas :: Room -> Port -> IO BoundAddress
serveCanvas room port = do
  server <- createServer # liftEffect
  registerServiceWith @Canvas { room } { join } server # liftEffect
  _ <- forkIO (everyTick (sendWhatPiledUp room))
  listen { host: Host "127.0.0.1", port } server

-- | Runs the action every 25 milliseconds, for as long as the server runs.
everyTick :: IO Unit -> IO Unit
everyTick action = void (runOmIO unit ticking)
  where
  ticking :: OmIO Unit () Unit
  ticking = Strom.every (Milliseconds 25.0) # Strom.forEach \_ -> action # liftIO

-- | Says hello, then moves and draws until the participant hangs up, which is when they leave.
join :: Handler Joining (room :: Room)
join = handle do
  { requests: gestures, room } <- ask
  pure (Strom.bracketExit (arrive # liftIO) (leave room) \visit -> scenesFor room visit gestures)
  where
  arrive = do
    outbox <- newBufferedChan 1024
    name <- newIORef Nothing
    pure { outbox, name }

-- | What the participant sees, while what they do changes the canvas.
scenesFor :: forall ctx. Room -> Visit -> Strom ctx Failures Gesture -> Strom ctx Failures Scene
scenesFor room visit gestures = Strom.mergeHaltLeft doing (Strom.fromChan visit.outbox)
  where
  doing = gestures # Strom.mapM (follow room visit) # Strom.filterMap (const Nothing)

follow :: forall ctx. Room -> Visit -> Gesture -> OmIO ctx Failures Unit
follow room visit = Variant.match
  { hello: \{ name } -> do
      already <- readIORef visit.name # liftIO
      when (isJust already) (reject @"invalidArgument" "you said hello already")
      when (name == "") (reject @"invalidArgument" "say who you are")
      welcomed <- welcome room visit name # liftIO
      unless welcomed (reject @"nameTaken" (name <> " is here already, pick another name"))
  , moved: \at -> asSomeone \name -> atomicModifyIORef' room (\canvas -> moveTo name at canvas /\ unit) # liftIO
  , drew: \cell -> asSomeone \_ -> atomicModifyIORef' room (\canvas -> canvas { cells = Map.insert (cell.x /\ cell.y) cell.color canvas.cells, drawn = Array.snoc canvas.drawn cell } /\ unit) # liftIO
  }
  where
  asSomeone action = readIORef visit.name # liftIO >>= maybe (reject @"invalidArgument" "say hello first") action
  moveTo name at canvas = canvas
    { people = Map.update (\person -> Just person { at = Just at }) name canvas.people
    , moves = Map.insert name at canvas.moves
    }

-- | Takes the name if nobody has it, picks a color, shows the newcomer the canvas and tells
-- | everyone else.
welcome :: Room -> Visit -> String -> IO Boolean
welcome room visit name = do
  arrived <- atomicModifyIORef' room \canvas ->
    if Map.member name canvas.people then canvas /\ Nothing
    else do
      let person = { color: freeColor canvas, at: Nothing, outbox: visit.outbox }
      canvas { people = Map.insert name person canvas.people } /\ Just (canvas /\ person)
  case arrived of
    Nothing -> pure false
    Just (before /\ person) -> do
      writeIORef visit.name (Just name)
      let you = someone name person
      _ <- trySend visit.outbox (Variant.inj (Proxy :: Proxy "welcome") { you, others: Map.toUnfoldable before.people <#> \(other /\ them) -> someone other them, cells: cellsOf before })
      tellEveryoneBut name before (Variant.inj (Proxy :: Proxy "joined") you)
      pure true
  where
  cellsOf canvas = Map.toUnfoldable canvas.cells <#> \((x /\ y) /\ color) -> { x, y, color }

-- | Takes the participant's cursor off the canvas, however they left.
leave :: forall ctx err. Room -> Visit -> Exit err -> OmIO ctx err Unit
leave room visit _ = do
  named <- readIORef visit.name # liftIO
  for_ named \name -> do
    canvas <- atomicModifyIORef' room (\canvas -> canvas { people = Map.delete name canvas.people, moves = Map.delete name canvas.moves } /\ canvas) # liftIO
    tellEveryoneBut name canvas (Variant.inj (Proxy :: Proxy "left") { name }) # liftIO
    log (name <> " left") # liftIO

-- | Everyone gets the cursors that moved and the cells drawn since last time, in one scene each.
sendWhatPiledUp :: Room -> IO Unit
sendWhatPiledUp room = do
  canvas <- atomicModifyIORef' room \canvas -> canvas { moves = Map.empty, drawn = [] } /\ canvas
  unless (Map.isEmpty canvas.moves) do
    tellEveryone canvas (Variant.inj (Proxy :: Proxy "cursors") { moved: Map.toUnfoldable canvas.moves <#> \(name /\ at) -> { name, at } })
  unless (Array.null canvas.drawn) do
    tellEveryone canvas (Variant.inj (Proxy :: Proxy "drawn") { cells: canvas.drawn })

-- | A participant who can't keep up misses scenes rather than holding everyone else up.
tellEveryone :: CanvasState -> Scene -> IO Unit
tellEveryone canvas scene = for_ canvas.people \person -> void (trySend person.outbox scene)

tellEveryoneBut :: String -> CanvasState -> Scene -> IO Unit
tellEveryoneBut name canvas scene = tellEveryone canvas { people = Map.delete name canvas.people } scene

-- | The first color nobody has, or the first one if everyone's taken.
freeColor :: CanvasState -> Color
freeColor canvas = Array.find (\color -> not (Array.elem color taken)) palette # maybe Coral identity
  where
  taken = Map.values canvas.people # Array.fromFoldable <#> _.color
  palette = [ Coral, Tang, Butter, Guac, Malibu, Dolly ]

someone :: String -> Person -> { name :: String, color :: Color, at :: Maybe Point }
someone name { color, at } = { name, color, at }

-- | The canvas, who's on it, and what piled up since everyone was last told.
type Room = IORef CanvasState

type CanvasState =
  { people :: Map String Person
  , cells :: Map (Int /\ Int) Color
  , moves :: Map String Point
  , drawn :: Array Cell
  }

type Person = { color :: Color, at :: Maybe Point, outbox :: Chan Scene }

-- | One call: where its scenes go, and the name it said hello with.
type Visit = { outbox :: Chan Scene, name :: IORef (Maybe String) }

newRoom :: IO Room
newRoom = newIORef { people: Map.empty, cells: Map.empty, moves: Map.empty, drawn: [] }

type Failures = (nameTaken :: String, invalidArgument :: String)

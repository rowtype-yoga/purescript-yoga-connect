-- | Keeps the canvas, shows it to whoever joins, and passes every painted cell on to everyone.
module Draw.Server (main, serveCanvas, newRoom) where

import Prelude

import Data.Foldable (for_)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple.Nested (type (/\), (/\))
import Draw.Service (Canvas, Cell)
import Effect (Effect)
import Effect.Class (liftEffect)
import Go.IO (IO, launchIO_)
import Go.IO.Chan (Chan, newBufferedChan, trySend)
import Go.IO.Console (log)
import Go.IO.IORef (IORef, atomicModifyIORef', newIORef)
import Yoga.Connect.Om (BidiStream, BoundAddress, Handler, Host(..), Port(..), ask, createServer, handle, liftIO, listen, registerServiceWith)
import Yoga.OmIO.Strom as Strom

main :: Effect Unit
main = launchIO_ do
  room <- newRoom
  _ <- serveCanvas room (Port 50053)
  log "canvas on 127.0.0.1:50053"

serveCanvas :: Room -> Port -> IO BoundAddress
serveCanvas room port = do
  server <- createServer # liftEffect
  registerServiceWith @Canvas { room } { draw } server # liftEffect
  listen { host: Host "127.0.0.1", port } server

-- | You get the canvas as it is, then every cell anyone paints, yours included, until you hang up.
draw :: Handler (BidiStream Cell Cell ()) (room :: Room)
draw = handle do
  { requests: yourCells, room } <- ask
  pure (Strom.bracketExit (join room # liftIO) (\you _ -> leave room you # liftIO) \you ->
    everyonesCells you # whilePainting room yourCells)
  where
  everyonesCells you = Strom.fromArray you.before <> Strom.fromChan you.outbox <#> \(x /\ y) -> { x, y }
  whilePainting room yourCells = Strom.mergeHaltLeft (yourCells # Strom.mapIO (paint room) # Strom.filterMap (const Nothing))

-- | Puts the cell on the canvas and in everyone's outbox. Someone who can't keep up misses it.
paint :: Room -> { x :: Int, y :: Int } -> IO Unit
paint room { x, y } = do
  canvas <- atomicModifyIORef' room \canvas -> do
    let painted = canvas { cells = Set.insert (x /\ y) canvas.cells }
    painted /\ painted
  for_ canvas.outboxes \outbox -> trySend outbox (x /\ y)

join :: Room -> IO You
join room = do
  outbox <- newBufferedChan 1024
  atomicModifyIORef' room \canvas ->
    canvas { outboxes = Map.insert canvas.nextId outbox canvas.outboxes, nextId = canvas.nextId + 1 }
      /\ { id: canvas.nextId, outbox, before: Set.toUnfoldable canvas.cells }

leave :: Room -> You -> IO Unit
leave room you = atomicModifyIORef' room \canvas -> canvas { outboxes = Map.delete you.id canvas.outboxes } /\ unit

-- | What's painted, and where to send what gets painted next.
type Room = IORef { cells :: Set Spot, outboxes :: Map Int (Chan Spot), nextId :: Int }

-- | Who you are to the room, where your cells go, and what was painted before you came.
type You = { id :: Int, outbox :: Chan Spot, before :: Array Spot }

type Spot = Int /\ Int

newRoom :: IO Room
newRoom = newIORef { cells: Set.empty, outboxes: Map.empty, nextId: 0 }

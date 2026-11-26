-- | A shared canvas: everyone's mouse, what everyone draws, and what everyone's saying, as one
-- | stream each way.
module Cursors.Service where

import Prelude

import Data.Generic.Rep (class Generic)
import Data.Maybe (Maybe)
import Effect (Effect)
import Go.IO (launchIO_)
import Go.IO.Console (log)
import Yoga.Connect.Om (type (:=), BidiStream, Message, Named, OneOf, Service, printProto)

type Canvas = Service "canvas.v1.Canvas"
  { join :: Joining
  }

-- | Say hello with a name, then move, draw and talk. The first scene back is the canvas as it is.
type Joining = BidiStream Gesture Scene (nameTaken :: String, invalidArgument :: String)

-- | `typing` is the whole line typed so far, on every keystroke, empty when it's given up on;
-- | `said` is the line sent.
type Gesture = Named "Gesture" (OneOf
  ( hello :: 1 := Message { name :: 1 := String }
  , moved :: 2 := Point
  , drew :: 3 := Cell
  , typing :: 4 := Message { text :: 1 := String }
  , said :: 5 := Message { text :: 1 := String }
  ))

-- | `speech` is the bubbles that changed since last time, and the names whose bubble is gone.
type Scene = Named "Scene" (OneOf
  ( welcome :: 1 := Welcome
  , cursors :: 2 := Message { moved :: 1 := Array Cursor }
  , drawn :: 3 := Message { cells :: 1 := Array Cell }
  , joined :: 4 := Someone
  , left :: 5 := Message { name :: 1 := String }
  , speech :: 6 := Message { bubbles :: 1 := Array Bubble, hushed :: 2 := Array String }
  ))

-- | Who you are, who else is here and where, what's drawn so far and what's being said right now.
type Welcome = Message
  { you :: 1 := Someone
  , others :: 2 := Array Someone
  , cells :: 3 := Array Cell
  , bubbles :: 4 := Array Bubble
  }

type Someone = Named "Someone" (Message { name :: 1 := String, color :: 2 := Color, at :: 3 := Maybe Point })

type Cursor = Named "Cursor" (Message { name :: 1 := String, at :: 2 := Point })

type Cell = Named "Cell" (Message { x :: 1 := Int, y :: 2 := Int, color :: 3 := Color })

type Point = Named "Point" (Message { x :: 1 := Int, y :: 2 := Int })

-- | What someone is saying, next to their cursor.
type Bubble = Named "Bubble" (Message { name :: 1 := String, text :: 2 := String, voice :: 3 := Voice })

data Color = Coral | Tang | Butter | Guac | Malibu | Dolly

derive instance Generic Color _
derive instance Eq Color
derive instance Ord Color

-- | Still typing, just said it, or said it a while ago and about to go.
data Voice = Typing | Said | Fading

derive instance Generic Voice _
derive instance Eq Voice
derive instance Ord Voice

-- | Prints the `.proto`.
main :: Effect Unit
main = launchIO_ (log (printProto @Canvas))

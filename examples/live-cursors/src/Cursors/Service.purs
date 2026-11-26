-- | A shared canvas: everyone's mouse, and what everyone draws, as one stream each way.
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

-- | Say hello with a name, then move and draw. The first scene back is the canvas as it is.
type Joining = BidiStream Gesture Scene (alreadyExists :: String, invalidArgument :: String)

type Gesture = Named "Gesture" (OneOf
  ( hello :: 1 := Message { name :: 1 := String }
  , moved :: 2 := Point
  , drew :: 3 := Cell
  ))

type Scene = Named "Scene" (OneOf
  ( welcome :: 1 := Welcome
  , cursors :: 2 := Message { moved :: 1 := Array Cursor }
  , drawn :: 3 := Message { cells :: 1 := Array Cell }
  , joined :: 4 := Someone
  , left :: 5 := Message { name :: 1 := String }
  ))

-- | Who you are, who else is here and where, and what's drawn so far.
type Welcome = Message { you :: 1 := Someone, others :: 2 := Array Someone, cells :: 3 := Array Cell }

type Someone = Named "Someone" (Message { name :: 1 := String, color :: 2 := Color, at :: 3 := Maybe Point })

type Cursor = Named "Cursor" (Message { name :: 1 := String, at :: 2 := Point })

type Cell = Named "Cell" (Message { x :: 1 := Int, y :: 2 := Int, color :: 3 := Color })

type Point = Named "Point" (Message { x :: 1 := Int, y :: 2 := Int })

data Color = Coral | Tang | Butter | Guac | Malibu | Dolly

derive instance Generic Color _
derive instance Eq Color
derive instance Ord Color

-- | Prints the `.proto`.
main :: Effect Unit
main = launchIO_ (log (printProto @Canvas))

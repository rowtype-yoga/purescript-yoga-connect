-- | One canvas, one call: send the cells you paint, get back every cell anyone paints.
module Draw.Service where

import Yoga.Connect.Om (type (:=), BidiStream, Message, Service)

type Canvas = Service "draw.v1.Canvas" { draw :: BidiStream Cell Cell () }

type Cell = Message { x :: 1 := Int, y :: 2 := Int }

-- | The service's messages as the server and the client see them: plain records and variants.
-- | The compiler checks them against `Cursors.Service`.
module Cursors.Plain where

import Data.Maybe (Maybe)
import Data.Variant (Variant)
import Cursors.Service (Color)

type Point = { x :: Int, y :: Int }

type Cell = { x :: Int, y :: Int, color :: Color }

type Someone = { name :: String, color :: Color, at :: Maybe Point }

type Gesture = Variant (hello :: { name :: String }, moved :: Point, drew :: Cell)

type Scene = Variant
  ( welcome :: { you :: Someone, others :: Array Someone, cells :: Array Cell }
  , cursors :: { moved :: Array { name :: String, at :: Point } }
  , drawn :: { cells :: Array Cell }
  , joined :: Someone
  , left :: { name :: String }
  )

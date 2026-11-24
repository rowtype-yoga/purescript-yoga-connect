module Test.Yoga.Connect.StromSpec where

import Prelude

import Data.Array as Array
import Data.Either (Either(..), either)
import Data.Time.Duration (Milliseconds(..))
import Effect.Aff (Aff)
import Go.IO (toAff)
import Go.IO.IORef (modifyIORef, newIORef, readIORef)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual)
import Yoga.Connect.Strom (Strom)
import Yoga.Connect.Strom as Strom
import Yoga.Go.HTTP.OmIO (liftIO, runOmIO)

spec :: Spec Unit
spec = describe "Yoga.Connect.Strom" do
  it "maps, filters and scans" do
    collected (Strom.range 1 10 # map (_ * 10) # Strom.filter (_ > 50) # Strom.scan (+) 0)
      >>= shouldEqual [ 60, 130, 210, 300, 400 ]

  it "takes from a stream without end, and stops it" do
    collected (Strom.iterate (_ * 2) 1 # Strom.take 5) >>= shouldEqual [ 1, 2, 4, 8, 16 ]

  it "goes on after a taken stream" do
    collected ((Strom.iterate (_ + 1) 1 # Strom.take 2) <> Strom.fromArray [ 100 ])
      >>= shouldEqual [ 1, 2, 100 ]

  it "takes while a test passes, and drops" do
    collected (Strom.range 1 10 # Strom.takeWhile (_ < 4)) >>= shouldEqual [ 1, 2, 3 ]
    collected (Strom.range 1 5 # Strom.drop 3) >>= shouldEqual [ 4, 5 ]

  it "runs effects for each element, in order" do
    seen <- toAff do
      ref <- newIORef []
      _ <- runOmIO unit (Strom.range 1 3 # Strom.tapM (\n -> modifyIORef ref (_ <> [ n ]) # liftIO) # Strom.drain)
      readIORef ref
    seen `shouldEqual` [ 1, 2, 3 ]

  it "concatenates the stream of each element" do
    collected (Strom.range 1 3 # Strom.concatMap (\n -> Strom.fromArray [ n, n ])) >>= shouldEqual [ 1, 1, 2, 2, 3, 3 ]

  it "merges two streams running at the same time" do
    merged <- collected (Strom.merge (Strom.every (Milliseconds 5.0) # Strom.take 3 # map (_ + 100)) (Strom.fromArray [ 1, 2 ]))
    Array.sort merged `shouldEqual` [ 1, 2, 100, 101, 102 ]

  it "stops both sides of a merge when its consumer stops" do
    collected (Strom.merge (Strom.iterate identity 1) (Strom.iterate identity 2) # Strom.take 4 # map (const 0))
      >>= shouldEqual [ 0, 0, 0, 0 ]

  it "folds" do
    total <- runOmIO unit (Strom.foldl (+) 0 (Strom.range 1 100)) # toAff
    either (const (-1)) identity total `shouldEqual` 5050

collected :: forall a. Strom Unit () a -> Aff (Array a)
collected stream = do
  outcome <- runOmIO unit (Strom.collect stream) # toAff
  case outcome of
    Left _ -> pure []
    Right elements -> pure elements

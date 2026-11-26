module Test.Yoga.Connect.ProtobufSpec where

import Prelude

import Data.Either (Either(..), isLeft)
import Data.Generic.Rep (class Generic)
import Data.Int64 as Int64
import Data.Maybe (Maybe(..))
import Data.Show.Generic (genericShow)
import Data.Variant (Variant, inj)
import Effect.Aff (Aff)
import Test.Fixtures.TickerProto as Generated
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual, shouldSatisfy)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Message (type (:=), Message, Named, OneOf)
import Yoga.Connect.Protobuf (class ProtobufMessage, decode, encode)
import Yoga.Protobuf.Bytes as Bytes
import Yoga.Protobuf.Error (DecodeError(..))
import Yoga.Protobuf.Unsigned (UInt32(..))

spec :: Spec Unit
spec = describe "protocol buffer codecs worked out from message types" do
  describe "what handlers see" do
    it "is a plain record of the fields' values, numbers gone" do
      compiles (plainOf @Quote :: Proxy { symbol :: String, price :: Number, average :: Number })

    it "maps arrays, maybes, nested messages, enums and oneofs to their PureScript types" do
      compiles (plainOf @Drawing :: Proxy PlainDrawing)

  describe "round trips" do
    it "gives back what it encoded, every kind of field set" do
      decode @Drawing (encode @Drawing busyDrawing) `shouldEqual` Right busyDrawing

    it "gives back what it encoded, every field at its default" do
      decode @Drawing (encode @Drawing emptyDrawing) `shouldEqual` Right emptyDrawing

    it "leaves out every field at its default, but a nested message and the oneof's choice" do
      Bytes.toArray (encode @Drawing emptyDrawing) `shouldEqual` [ 58, 0, 88, 0 ]

    it "decodes no bytes as every field at its default, the oneof at its lowest-numbered alternative" do
      decode @Drawing Bytes.empty `shouldEqual` Right emptyDrawing { focus = inj (Proxy :: Proxy "point") { x: 0, y: 0 } }

    it "keeps an optional field that's there at its default apart from one that isn't there" do
      (decode @Drawing (encode @Drawing emptyDrawing { note = Just "" }) <#> _.note) `shouldEqual` Right (Just "")

  describe "the wire, as protoc's code has it" do
    it "encodes the very bytes the protoc-generated encoder does" do
      encode @Quote quote `shouldEqual` Generated.encodeQuote quote
      encode @Subscribe subscribe `shouldEqual` Generated.encodeSubscribe subscribe
      encode @Fill fill `shouldEqual` Generated.encodeFill fill

    it "decodes what the protoc-generated encoder wrote" do
      decode @Quote (Generated.encodeQuote quote) `shouldEqual` Right quote
      decode @Fill (Generated.encodeFill fill) `shouldEqual` Right fill
      Generated.decodeFill (encode @Fill fill) `shouldEqual` Right fill

    it "writes repeated numbers packed, and reads them packed or one per tag" do
      Bytes.toArray (encode @Counts { counts: [ 1, 2 ] }) `shouldEqual` [ 42, 2, 1, 2 ]
      decode @Counts (Bytes.fromArray [ 40, 1, 40, 2 ]) `shouldEqual` Right { counts: [ 1, 2 ] }
      decode @Counts (Bytes.fromArray [ 42, 1, 1, 40, 2 ]) `shouldEqual` Right { counts: [ 1, 2 ] }

    it "writes an enum as its constructor's number" do
      Bytes.toArray (encode @Drawing emptyDrawing { color = Blue }) `shouldEqual` [ 16, 2, 58, 0, 88, 0 ]

    it "skips the fields it doesn't know, and takes the last of a field that comes twice" do
      decode @Counted (Bytes.fromArray [ 8, 1, 16, 7, 26, 1, 65, 16, 9 ]) `shouldEqual` Right { count: 9 }

    it "says where the bytes are broken" do
      decode @Quote (Bytes.fromArray [ 10, 5, 65 ]) `shouldEqual` Left (Truncated 1)
      decode @Quote (Bytes.fromArray [ 10, 1, 255 ]) `shouldEqual` Left (InvalidUtf8 2)
      decode @Empty (Bytes.fromArray [ 255 ]) `shouldSatisfy` isLeft

-- | The plain type of a message, which only compiles with the right one.
plainOf :: forall @message plain. ProtobufMessage message plain => Proxy plain
plainOf = Proxy

-- | Checked by the compiler alone.
compiles :: forall a. Proxy a -> Aff Unit
compiles _ = pure unit

type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }

type Subscribe = Message { symbols :: 1 := Array String }

type Fill = Message { symbol :: 1 := String, quantity :: 2 := Int, price :: 3 := Number, position :: 4 := Int }

type Counts = Message { counts :: 5 := Array Int }

type Counted = Message { count :: 2 := Int }

type Empty = Message {}

data Color = Red | Green | Blue

derive instance Generic Color _
derive instance Eq Color

instance Show Color where
  show = genericShow

type Point = Named "Point" (Message { x :: 1 := Int, y :: 2 := Int })

type Shape = OneOf (circle :: 1 := Message { radius :: 1 := Number }, square :: 2 := Message { side :: 1 := Number })

type Drawing = Message
  { title :: 1 := String
  , color :: 2 := Color
  , shapes :: 3 := Array Shape
  , tags :: 4 := Array String
  , counts :: 5 := Array Int
  , note :: 6 := Maybe String
  , origin :: 7 := Point
  , big :: 8 := Int64.Int64
  , visible :: 9 := Boolean
  , flags :: 12 := UInt32
  , focus :: OneOf (point :: 10 := Point, everything :: 11 := Boolean)
  }

type PlainDrawing =
  { title :: String
  , color :: Color
  , shapes :: Array (Variant (circle :: { radius :: Number }, square :: { side :: Number }))
  , tags :: Array String
  , counts :: Array Int
  , note :: Maybe String
  , origin :: { x :: Int, y :: Int }
  , big :: Int64.Int64
  , visible :: Boolean
  , flags :: UInt32
  , focus :: Variant (point :: { x :: Int, y :: Int }, everything :: Boolean)
  }

busyDrawing :: PlainDrawing
busyDrawing =
  { title: "Bücher 📚"
  , color: Green
  , shapes: [ inj (Proxy :: Proxy "circle") { radius: 2.5 }, inj (Proxy :: Proxy "square") { side: 0.0 } ]
  , tags: [ "a", "", "c" ]
  , counts: [ 0, -1, 300, 2147483647 ]
  , note: Just "hello"
  , origin: { x: -5, y: 7 }
  , big: Int64.fromLowHighBits 1 1
  , visible: true
  , flags: UInt32 (-1)
  , focus: inj (Proxy :: Proxy "everything") false
  }

emptyDrawing :: PlainDrawing
emptyDrawing =
  { title: ""
  , color: Red
  , shapes: []
  , tags: []
  , counts: []
  , note: Nothing
  , origin: { x: 0, y: 0 }
  , big: zero
  , visible: false
  , flags: UInt32 0
  , focus: inj (Proxy :: Proxy "everything") false
  }

quote :: { symbol :: String, price :: Number, average :: Number }
quote = { symbol: "ACME", price: 121.02, average: 121.4 }

subscribe :: { symbols :: Array String }
subscribe = { symbols: [ "ACME", "GLOBEX" ] }

fill :: { symbol :: String, quantity :: Int, price :: Number, position :: Int }
fill = { symbol: "INITECH", quantity: -40, price: 9.9, position: -15 }

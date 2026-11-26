-- | Messages as types: a record whose fields carry their protocol buffer numbers.
-- |
-- | ```purescript
-- | type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }
-- | ```
-- |
-- | The numbers are what goes over the wire, so a field keeps its number for good: rename it as
-- | you like, but give a new field a new number. Handlers and clients never see the numbers,
-- | they get `{ symbol :: String, price :: Number, average :: Number }`.
-- |
-- | A field is one of
-- |
-- | | PureScript | protocol buffers |
-- | | --- | --- |
-- | | `String`, `Boolean`, `Number`, `Bytes` | `string`, `bool`, `double`, `bytes` |
-- | | `Int` | `int32` |
-- | | `Int64`, `UInt32`, `UInt64` | `int64`, `uint32`, `uint64` |
-- | | a sum type of constructors without arguments, with `Generic` | an `enum`, the first constructor is 0 |
-- | | another message | a nested `message` |
-- | | `Array a` | `repeated a`, packed if it's a number |
-- | | `Maybe a` | `optional a` |
-- | | `OneOf (circle :: 4 := Circle, square :: 5 := Square)` | a `oneof`, a `Variant` to handlers |
-- |
-- | This module only says what a message is. How it goes over the wire is up to a codec, and
-- | `Yoga.Connect.Protobuf` is the one there is.
module Yoga.Connect.Message
  ( Message
  , Numbered
  , type (:=)
  , OneOf
  , Named
  ) where

-- | Marks a record of numbered fields as a message. It's only there to be read: any record of
-- | numbered fields is a message.
type Message :: Type -> Type
type Message fields = fields

-- | A field and its number: `1 := String`.
data Numbered :: Int -> Type -> Type
data Numbered number field

infixr 0 type Numbered as :=

-- | Fields of which a message has one at a time, each with its own number. Handlers get a
-- | `Variant` of them. As a message of its own, `type Shape = OneOf (…)` is a message with only
-- | the `oneof`.
data OneOf :: Row Type -> Type
data OneOf alternatives

-- | A message or enum with a name of its own in the `.proto`, rather than one after its field:
-- | `type Point = Named "Point" (Message { x :: 1 := Int, y :: 2 := Int })`. Names don't go over
-- | the wire, so naming one later changes nothing for clients.
data Named :: Symbol -> Type -> Type
data Named name value


-- | How a method's errors go over the wire, worked out from its error row.
-- |
-- | ```purescript
-- | type Playing = BidiStream Move Event
-- |   ( notYourTurn :: String
-- |   , columnFull :: Int
-- |   , nameTaken :: WithCode "alreadyExists" String
-- |   , notFound :: String
-- |   )
-- | ```
-- |
-- | - A label that's a gRPC code with a `String`, `notFound :: String`, is that code and the
-- |   message, as every gRPC client knows it.
-- | - Any other label is `FAILED_PRECONDITION`, or the code `WithCode` gives it, and carries its
-- |   value as an error detail: a message `NotYourTurn` (after the label) with the value as its
-- |   field `value = 1`, or the message itself when the value is one. A `String` is the status
-- |   message too, anything else has the label as its message.
-- |
-- | Clients made from the service type read the detail back, so `reject @"columnFull" 3` on the
-- | server is `columnFull` with `3` on the client. Other clients see the code and the message,
-- | and the detail's type in the `.proto`.
module Yoga.Connect.Rejection
  ( WithCode
  , Rejected
  , RejectionCodec
  , ErrorShape
  , class Rejections
  , rejections
  , rejectionsOf
  , class CodeOf
  , class Detailed
  , class PlainOrDetailed
  , plainOrDetailed
  , detailed
  , class ErrorText
  , errorText
  , detailName
  ) where

import Prelude

import Control.Alt ((<|>))
import Data.Array as Array
import Data.Either (hush)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String (Pattern(..))
import Data.String as String
import Data.String.CodeUnits as CodeUnits
import Data.Symbol (class IsSymbol, reflectSymbol)
import Data.Variant (Variant)
import Data.Variant as Variant
import Prim.Boolean (False, True)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Type.Proxy (Proxy(..))
import Yoga.Connect.Protobuf (class MethodMessage, MessageShape, methodCodec, methodShape)
import Yoga.Connect.Status (class RejectionCode, Code, ErrorDetail, rejectionCode)
import Yoga.Protobuf.Bytes (Bytes)
import Yoga.Protobuf.Codec as Codec

-- | An error that goes out with another code than `FAILED_PRECONDITION`:
-- | `nameTaken :: WithCode "alreadyExists" String`. Handlers and clients see the `String`.
data WithCode :: Symbol -> Type -> Type
data WithCode code payload

-- | What a call fails with on the wire. The detail's name is the message's, without the package.
type Rejected = { code :: Code, message :: String, detail :: Maybe { name :: String, bytes :: Bytes } }

-- | Turns a handler's error into what goes over the wire, and a call's failure back into the
-- | error, if it's one of the method's.
type RejectionCodec plain =
  { encode :: Variant plain -> Rejected
  , decode :: Code -> String -> Array ErrorDetail -> Maybe (Variant plain)
  , shapes :: Array ErrorShape
  }

-- | A detail's message, for the `.proto`.
type ErrorShape = { name :: String, shape :: MessageShape }

-- | The codec of a method's error row.
rejectionsOf :: forall @errors list plain. RowToList errors list => Rejections list plain => RejectionCodec plain
rejectionsOf = rejections (Proxy :: Proxy list)

class Rejections :: RowList Type -> Row Type -> Constraint
class Rejections list plain | list -> plain where
  rejections :: Proxy list -> RejectionCodec plain

instance Rejections RowList.Nil () where
  rejections _ = { encode: Variant.case_, decode: \_ _ _ -> Nothing, shapes: [] }

instance
  ( IsSymbol label
  , Detailed label payload plainPayload
  , Rejections tail plainTail
  , Row.Cons label plainPayload plainTail plain
  , Row.Cons label plainPayload () single
  , Row.Union plainTail single plain
  ) =>
  Rejections (RowList.Cons label payload tail) plain where
  rejections _ =
    { encode: Variant.on (Proxy :: Proxy label) this.encode rest.encode
    , decode: \code message details ->
        (this.decode code message details <#> Variant.inj (Proxy :: Proxy label)) <|> (rest.decode code message details <#> Variant.expand)
    , shapes: this.shapes <> rest.shapes
    }
    where
    this = detailed (Proxy :: Proxy label) (Proxy :: Proxy payload)
    rest = rejections (Proxy :: Proxy tail)

-- | The code of a label: its own if it's a gRPC code's name, `failedPrecondition` if not.
class CodeOf :: Symbol -> Symbol -> Boolean -> Constraint
class CodeOf label code isCode | label -> code isCode

instance CodeOf "canceled" "canceled" True
else instance CodeOf "unknown" "unknown" True
else instance CodeOf "invalidArgument" "invalidArgument" True
else instance CodeOf "deadlineExceeded" "deadlineExceeded" True
else instance CodeOf "notFound" "notFound" True
else instance CodeOf "alreadyExists" "alreadyExists" True
else instance CodeOf "permissionDenied" "permissionDenied" True
else instance CodeOf "resourceExhausted" "resourceExhausted" True
else instance CodeOf "failedPrecondition" "failedPrecondition" True
else instance CodeOf "aborted" "aborted" True
else instance CodeOf "outOfRange" "outOfRange" True
else instance CodeOf "unimplemented" "unimplemented" True
else instance CodeOf "internal" "internal" True
else instance CodeOf "unavailable" "unavailable" True
else instance CodeOf "dataLoss" "dataLoss" True
else instance CodeOf "unauthenticated" "unauthenticated" True
else instance CodeOf label "failedPrecondition" False

-- | How one error goes over the wire.
class Detailed :: Symbol -> Type -> Type -> Constraint
class Detailed label payload plain | label payload -> plain where
  detailed :: Proxy label -> Proxy payload -> { encode :: plain -> Rejected, decode :: Code -> String -> Array ErrorDetail -> Maybe plain, shapes :: Array ErrorShape }

instance (IsSymbol label, RejectionCode code, MethodMessage payload plain, ErrorText plain) => Detailed label (WithCode code payload) plain where
  detailed label _ = withDetail label (rejectionCode @code) (Proxy :: Proxy payload)
else instance (IsSymbol label, CodeOf label code isCode, RejectionCode code, PlainOrDetailed isCode label) => Detailed label String String where
  detailed label _ = plainOrDetailed (Proxy :: Proxy isCode) label (rejectionCode @code)
else instance (IsSymbol label, CodeOf label code isCode, RejectionCode code, MethodMessage payload plain, ErrorText plain) => Detailed label payload plain where
  detailed label _ = withDetail label (rejectionCode @code) (Proxy :: Proxy payload)

-- | A gRPC code's own label with a `String` is the code and the message, nothing more.
class PlainOrDetailed :: Boolean -> Symbol -> Constraint
class PlainOrDetailed isCode label where
  plainOrDetailed :: Proxy isCode -> Proxy label -> Code -> { encode :: String -> Rejected, decode :: Code -> String -> Array ErrorDetail -> Maybe String, shapes :: Array ErrorShape }

instance PlainOrDetailed True label where
  plainOrDetailed _ _ code =
    { encode: \message -> { code, message, detail: Nothing }
    , decode: \failedWith message _ -> if failedWith == code then Just message else Nothing
    , shapes: []
    }

instance IsSymbol label => PlainOrDetailed False label where
  plainOrDetailed _ label code = withDetail label code (Proxy :: Proxy String)

withDetail
  :: forall label payload plain
   . IsSymbol label
  => MethodMessage payload plain
  => ErrorText plain
  => Proxy label
  -> Code
  -> Proxy payload
  -> { encode :: plain -> Rejected, decode :: Code -> String -> Array ErrorDetail -> Maybe plain, shapes :: Array ErrorShape }
withDetail label code payload =
  { encode: \value -> { code, message: errorText value # fromMaybe labelName, detail: Just { name, bytes: Codec.encode payloadCodec value } }
  , decode: \failedWith _ details ->
      if failedWith /= code then Nothing
      else Array.find (\detail -> detailName detail.typeName == name) details >>= \detail -> hush (Codec.decode payloadCodec detail.bytes)
  , shapes: [ { name, shape: methodShape payload } ]
  }
  where
  labelName = reflectSymbol label
  name = String.toUpper (CodeUnits.take 1 labelName) <> CodeUnits.drop 1 labelName
  payloadCodec = methodCodec payload

-- | `NameTaken` of `canvas.v1.NameTaken`.
detailName :: String -> String
detailName typeName = case String.lastIndexOf (Pattern ".") typeName of
  Just dot -> CodeUnits.drop (dot + 1) typeName
  Nothing -> typeName

-- | The status message of an error that has one of its own.
class ErrorText a where
  errorText :: a -> Maybe String

instance ErrorText String where
  errorText = Just
else instance ErrorText a where
  errorText _ = Nothing

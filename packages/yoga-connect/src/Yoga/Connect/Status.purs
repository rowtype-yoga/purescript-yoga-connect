-- | How a call fails: one of the sixteen gRPC status codes and a message.
-- |
-- | A server handler rejects a call with a label of `Rejections`, `reject @"notFound" "no such
-- | user"`, and the client sees `ConnectError { code: NotFound, message: "no such user" }`.
module Yoga.Connect.Status
  ( Code(..)
  , codeNumber
  , codeFromNumber
  , codeName
  , codeFromName
  , ConnectError(..)
  , ErrorDetail
  , Rejections
  , fromRejection
  , rejection
  , rejectWith
  , class RejectionCode
  , rejectionCode
  , class Rejection
  , rejectionOf
  ) where

import Prelude

import Data.Maybe (Maybe(..))
import Data.Symbol (class IsSymbol)
import Data.Variant (Variant, case_, match, on)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Type.Proxy (Proxy(..))
import Yoga.Protobuf.Bytes (Bytes)

-- | The status codes, without `OK`.
data Code
  = Canceled
  | Unknown
  | InvalidArgument
  | DeadlineExceeded
  | NotFound
  | AlreadyExists
  | PermissionDenied
  | ResourceExhausted
  | FailedPrecondition
  | Aborted
  | OutOfRange
  | Unimplemented
  | Internal
  | Unavailable
  | DataLoss
  | Unauthenticated

derive instance Eq Code
derive instance Ord Code

instance Show Code where
  show code = codeName code

-- | The number gRPC sends in `grpc-status`.
codeNumber :: Code -> Int
codeNumber = case _ of
  Canceled -> 1
  Unknown -> 2
  InvalidArgument -> 3
  DeadlineExceeded -> 4
  NotFound -> 5
  AlreadyExists -> 6
  PermissionDenied -> 7
  ResourceExhausted -> 8
  FailedPrecondition -> 9
  Aborted -> 10
  OutOfRange -> 11
  Unimplemented -> 12
  Internal -> 13
  Unavailable -> 14
  DataLoss -> 15
  Unauthenticated -> 16

-- | Numbers outside 1 to 16 are `Unknown`, as gRPC clients treat them.
codeFromNumber :: Int -> Code
codeFromNumber = case _ of
  1 -> Canceled
  3 -> InvalidArgument
  4 -> DeadlineExceeded
  5 -> NotFound
  6 -> AlreadyExists
  7 -> PermissionDenied
  8 -> ResourceExhausted
  9 -> FailedPrecondition
  10 -> Aborted
  11 -> OutOfRange
  12 -> Unimplemented
  13 -> Internal
  14 -> Unavailable
  15 -> DataLoss
  16 -> Unauthenticated
  _ -> Unknown

-- | The name the Connect protocol sends in an error's `code`: `not_found`.
codeName :: Code -> String
codeName = case _ of
  Canceled -> "canceled"
  Unknown -> "unknown"
  InvalidArgument -> "invalid_argument"
  DeadlineExceeded -> "deadline_exceeded"
  NotFound -> "not_found"
  AlreadyExists -> "already_exists"
  PermissionDenied -> "permission_denied"
  ResourceExhausted -> "resource_exhausted"
  FailedPrecondition -> "failed_precondition"
  Aborted -> "aborted"
  OutOfRange -> "out_of_range"
  Unimplemented -> "unimplemented"
  Internal -> "internal"
  Unavailable -> "unavailable"
  DataLoss -> "data_loss"
  Unauthenticated -> "unauthenticated"

codeFromName :: String -> Maybe Code
codeFromName = case _ of
  "canceled" -> Just Canceled
  "unknown" -> Just Unknown
  "invalid_argument" -> Just InvalidArgument
  "deadline_exceeded" -> Just DeadlineExceeded
  "not_found" -> Just NotFound
  "already_exists" -> Just AlreadyExists
  "permission_denied" -> Just PermissionDenied
  "resource_exhausted" -> Just ResourceExhausted
  "failed_precondition" -> Just FailedPrecondition
  "aborted" -> Just Aborted
  "out_of_range" -> Just OutOfRange
  "unimplemented" -> Just Unimplemented
  "internal" -> Just Internal
  "unavailable" -> Just Unavailable
  "data_loss" -> Just DataLoss
  "unauthenticated" -> Just Unauthenticated
  _ -> Nothing

-- | A failed call, as the client sees it, with the details the server added.
newtype ConnectError = ConnectError { code :: Code, message :: String, details :: Array ErrorDetail }

-- | A message the server added to a failure: its full type name, `canvas.v1.NameTaken`, and its
-- | encoding.
type ErrorDetail = { typeName :: String, bytes :: Bytes }

derive instance Eq ConnectError

instance Show ConnectError where
  show (ConnectError { code, message }) = "(ConnectError " <> codeName code <> " " <> show message <> ")"

-- | The labels a handler rejects a call with, one per code, each with the message.
type Rejections =
  ( canceled :: String
  , unknown :: String
  , invalidArgument :: String
  , deadlineExceeded :: String
  , notFound :: String
  , alreadyExists :: String
  , permissionDenied :: String
  , resourceExhausted :: String
  , failedPrecondition :: String
  , aborted :: String
  , outOfRange :: String
  , unimplemented :: String
  , internal :: String
  , unavailable :: String
  , dataLoss :: String
  , unauthenticated :: String
  )

-- | The code a label of `Rejections` stands for.
class RejectionCode (label :: Symbol) where
  rejectionCode :: Code

instance RejectionCode "canceled" where
  rejectionCode = Canceled

instance RejectionCode "unknown" where
  rejectionCode = Unknown

instance RejectionCode "invalidArgument" where
  rejectionCode = InvalidArgument

instance RejectionCode "deadlineExceeded" where
  rejectionCode = DeadlineExceeded

instance RejectionCode "notFound" where
  rejectionCode = NotFound

instance RejectionCode "alreadyExists" where
  rejectionCode = AlreadyExists

instance RejectionCode "permissionDenied" where
  rejectionCode = PermissionDenied

instance RejectionCode "resourceExhausted" where
  rejectionCode = ResourceExhausted

instance RejectionCode "failedPrecondition" where
  rejectionCode = FailedPrecondition

instance RejectionCode "aborted" where
  rejectionCode = Aborted

instance RejectionCode "outOfRange" where
  rejectionCode = OutOfRange

instance RejectionCode "unimplemented" where
  rejectionCode = Unimplemented

instance RejectionCode "internal" where
  rejectionCode = Internal

instance RejectionCode "unavailable" where
  rejectionCode = Unavailable

instance RejectionCode "dataLoss" where
  rejectionCode = DataLoss

instance RejectionCode "unauthenticated" where
  rejectionCode = Unauthenticated

-- | The code and message of a rejection whose labels are known: one test per label of the row,
-- | rather than one per code.
class Rejection (labels :: RowList Type) (row :: Row Type) where
  rejectionOf :: forall result. Proxy labels -> (Code -> String -> result) -> Variant row -> result

instance Rejection RowList.Nil () where
  rejectionOf _ _ = case_

instance
  ( IsSymbol label
  , RejectionCode label
  , Row.Cons label String rest row
  , Rejection tail rest
  ) =>
  Rejection (RowList.Cons label String tail) row where
  rejectionOf _ rejected = on (Proxy :: Proxy label) (rejected (rejectionCode @label)) (rejectionOf (Proxy :: Proxy tail) rejected)

-- | Hands the code and message of a handler's rejection, whichever labels of `Rejections` its
-- | row has, to `rejected`.
rejectWith :: forall row labels result. RowToList row labels => Rejection labels row => (Code -> String -> result) -> Variant row -> result
rejectWith = rejectionOf (Proxy :: Proxy labels)

-- | The code and message of a handler's rejection.
rejection :: forall row labels. RowToList row labels => Rejection labels row => Variant row -> ConnectError
rejection = rejectWith \code message -> ConnectError { code, message, details: [] }

fromRejection :: Variant Rejections -> ConnectError
fromRejection = match
  { canceled: failed Canceled
  , unknown: failed Unknown
  , invalidArgument: failed InvalidArgument
  , deadlineExceeded: failed DeadlineExceeded
  , notFound: failed NotFound
  , alreadyExists: failed AlreadyExists
  , permissionDenied: failed PermissionDenied
  , resourceExhausted: failed ResourceExhausted
  , failedPrecondition: failed FailedPrecondition
  , aborted: failed Aborted
  , outOfRange: failed OutOfRange
  , unimplemented: failed Unimplemented
  , internal: failed Internal
  , unavailable: failed Unavailable
  , dataLoss: failed DataLoss
  , unauthenticated: failed Unauthenticated
  }
  where
  failed code message = ConnectError { code, message, details: [] }

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
  , Rejections
  , fromRejection
  ) where

import Prelude

import Data.Maybe (Maybe(..))
import Data.Variant (Variant, match)

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

-- | A failed call, as the client sees it.
newtype ConnectError = ConnectError { code :: Code, message :: String }

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
  failed code message = ConnectError { code, message }

module Test.Yoga.Connect.StatusSpec where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Variant (Variant, inj)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Status (Code(..), ConnectError(..), codeFromName, codeFromNumber, codeName, codeNumber, rejection)

spec :: Spec Unit
spec = describe "Yoga.Connect.Status" do
  it "numbers the codes as gRPC does, 1 to 16" do
    map codeNumber everyCode `shouldEqual` Array.range 1 16
    map (codeFromNumber <<< codeNumber) everyCode `shouldEqual` everyCode

  it "treats numbers outside 1 to 16 as unknown" do
    codeFromNumber 0 `shouldEqual` Unknown
    codeFromNumber 17 `shouldEqual` Unknown

  it "names the codes as the Connect protocol does" do
    codeName NotFound `shouldEqual` "not_found"
    map (codeFromName <<< codeName) everyCode `shouldEqual` map Just everyCode
    codeFromName "ok" `shouldEqual` Nothing

  it "turns a handler's rejection into the code of its label" do
    rejection (inj (Proxy :: Proxy "notFound") "no such user" :: HandlerRejections)
      `shouldEqual` ConnectError { code: NotFound, message: "no such user", details: [] }
    rejection (inj (Proxy :: Proxy "permissionDenied") "not yours" :: HandlerRejections)
      `shouldEqual` ConnectError { code: PermissionDenied, message: "not yours", details: [] }
  where
  everyCode =
    [ Canceled, Unknown, InvalidArgument, DeadlineExceeded, NotFound, AlreadyExists, PermissionDenied
    , ResourceExhausted, FailedPrecondition, Aborted, OutOfRange, Unimplemented, Internal, Unavailable
    , DataLoss, Unauthenticated
    ]

type HandlerRejections = Variant (notFound :: String, permissionDenied :: String)

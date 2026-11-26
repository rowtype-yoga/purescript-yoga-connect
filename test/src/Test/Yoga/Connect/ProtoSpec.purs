module Test.Yoga.Connect.ProtoSpec where

import Prelude

import Data.Maybe (Maybe)
import Test.Spec (Spec, describe, it)
import Test.Spec.Assertions (shouldEqual)
import Test.Yoga.Connect.ProtobufSpec (Drawing)
import Yoga.Connect.Message (type (:=), Message, Named, OneOf)
import Yoga.Connect.Proto (printProto)
import Yoga.Connect.Rejection (WithCode)
import Yoga.Connect.Schema (BidiStream, ClientStream, ServerStream, Service, Unary)

type Studio = Service "studio.v1.Studio"
  { save :: Unary Drawing (Named "Saved" (Message { id :: 1 := String })) (invalidArgument :: String)
  , list :: Unary {} (Array String) ()
  , follow :: ServerStream (Message { drawingId :: 1 := String }) Change ()
  , upload :: ClientStream Bytes' (Message { size :: 1 := Int }) (tooBig :: Int, gone :: WithCode "notFound" (Message { since :: 1 := String }))
  , sketch :: BidiStream Change Change (failedPrecondition :: String, notFound :: String)
  }

type Bytes' = Message { chunk :: 1 := Maybe String }

type Change = OneOf (moved :: 1 := Message { dx :: 1 := Int, dy :: 2 := Int }, renamed :: 2 := String)

spec :: Spec Unit
spec = describe "the .proto of a service type" do
  it "has a request and a response message per method, nested messages and enums named after their fields" do
    printProto @Studio `shouldEqual` studioProto

-- | protoc reads it as it is.
studioProto :: String
studioProto =
  """syntax = "proto3";

package studio.v1;

service Studio {
  rpc Follow(FollowRequest) returns (stream FollowResponse);
  rpc List(ListRequest) returns (ListResponse);
  rpc Save(SaveRequest) returns (Saved);
  rpc Sketch(stream SketchRequest) returns (stream SketchResponse);
  rpc Upload(stream UploadRequest) returns (UploadResponse);
}

message FollowRequest {
  string drawing_id = 1;
}

message FollowResponse {
  oneof kind {
    Moved moved = 1;
    string renamed = 2;
  }
}

message ListRequest {
}

message ListResponse {
  repeated string value = 1;
}

message SaveRequest {
  string title = 1;
  Color color = 2;
  repeated Shapes shapes = 3;
  repeated string tags = 4;
  repeated int32 counts = 5;
  optional string note = 6;
  Point origin = 7;
  int64 big = 8;
  bool visible = 9;
  oneof focus {
    Point point = 10;
    bool everything = 11;
  }
  uint32 flags = 12;
}

message Saved {
  string id = 1;
}

message SketchRequest {
  oneof kind {
    Moved moved = 1;
    string renamed = 2;
  }
}

message SketchResponse {
  oneof kind {
    Moved moved = 1;
    string renamed = 2;
  }
}

message UploadRequest {
  optional string chunk = 1;
}

message UploadResponse {
  int32 size = 1;
}

message Gone {
  string since = 1;
}

message TooBig {
  int32 value = 1;
}

message Circle {
  double radius = 1;
}

enum Color {
  COLOR_RED = 0;
  COLOR_GREEN = 1;
  COLOR_BLUE = 2;
}

message Moved {
  int32 dx = 1;
  int32 dy = 2;
}

message Point {
  int32 x = 1;
  int32 y = 2;
}

message Shapes {
  oneof kind {
    Circle circle = 1;
    Square square = 2;
  }
}

message Square {
  double side = 1;
}
"""

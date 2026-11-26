package ffi

import (
	"slices"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"
)

// rawBytes has connect-go pass messages through as bytes: the PureScript codecs decode and
// encode them. It stands in for the "proto" codec, so the content types stay
// application/grpc+proto and application/proto. connect-go also hands it the google.rpc.Status
// of a failure with details, which goes through protobuf as usual.
var rawBytes = connect.WithCodec(bytesCodec{})

type bytesCodec struct{}

func (bytesCodec) Name() string { return "proto" }

func (bytesCodec) Marshal(message any) ([]byte, error) {
	if raw, ok := message.(*[]byte); ok {
		return *raw, nil
	}
	return proto.Marshal(message.(proto.Message))
}

// Unmarshal copies: connect-go reuses the buffer it hands over.
func (bytesCodec) Unmarshal(data []byte, message any) error {
	if raw, ok := message.(*[]byte); ok {
		*raw = slices.Clone(data)
		return nil
	}
	return proto.Unmarshal(data, message.(proto.Message))
}

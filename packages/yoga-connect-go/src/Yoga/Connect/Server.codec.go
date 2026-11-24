package ffi

import (
	"slices"

	"connectrpc.com/connect"
)

// rawBytes has connect-go pass messages through as bytes: the generated PureScript codecs
// decode and encode them. It stands in for the "proto" codec, so the content types stay
// application/grpc+proto and application/proto.
var rawBytes = connect.WithCodec(bytesCodec{})

type bytesCodec struct{}

func (bytesCodec) Name() string { return "proto" }

func (bytesCodec) Marshal(message any) ([]byte, error) { return *message.(*[]byte), nil }

// Unmarshal copies: connect-go reuses the buffer it hands over.
func (bytesCodec) Unmarshal(data []byte, message any) error {
	*message.(*[]byte) = slices.Clone(data)
	return nil
}

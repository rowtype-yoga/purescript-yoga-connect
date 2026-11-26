package codec

import (
	bytes "iris.local/draw-test/Yoga_Protobuf_Bytes_ffi"
	writer "iris.local/draw-test/Yoga_Protobuf_Writer_ffi"
	"iris.local/draw-test/shared"
	"iris.local/draw-test/shared/Data.Either+Yoga.Protobuf.Error"
)

type Codec struct {
	Read  func(bytes.Bytes, int, int) (shared.XY, either_error.DecodeErrorFailure)
	Write func(writer.Writer, shared.XY)
}

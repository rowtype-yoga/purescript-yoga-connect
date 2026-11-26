package rejection

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Variant"
	"iris.local/draw-test/Type/Proxy"
	"iris.local/draw-test/Yoga/Connect/Protobuf"
	"iris.local/draw-test/Yoga/Connect/Status"
	"iris.local/draw-test/shared"
)

type Rejected struct {
	Detail  maybe.MaybeBytesName
	Message string
	Code    status.Code
}

type ErrorShape struct {
	Name  string
	Shape []*protobuf.Member
}

type RejectionCodec struct {
	Decode func(status.Code, string, []shared.ErrorDetail) variant.MaybeVariant
	Encode func(variant.Variant) Rejected
	Shapes []ErrorShape
}

func Rejections(ignored proxy.Proxy) RejectionCodec {
	return RejectionCodec{Encode: func(_ variant.Variant) Rejected {
		panic("Data.Variant: pattern match failure")
	}, Decode: func(_ status.Code, _ string, _ []shared.ErrorDetail) variant.MaybeVariant {
		return variant.NoVariant
	}, Shapes: []ErrorShape{}}
}

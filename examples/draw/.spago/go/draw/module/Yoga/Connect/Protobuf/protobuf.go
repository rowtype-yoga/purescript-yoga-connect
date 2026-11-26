package protobuf

import (
	"github.com/i-am-the-slime/iris-go/array"
	"iris.local/draw/Data/Maybe"
	"iris.local/draw/Data/Tuple"
	"iris.local/draw/Type/Proxy"
	"iris.local/draw/Yoga/Protobuf/Codec"
	"iris.local/draw/Yoga/Protobuf/Error"
	"iris.local/draw/Yoga/Protobuf/Writer"
	bytes "iris.local/draw/Yoga_Protobuf_Bytes_ffi"
	fields "iris.local/draw/Yoga_Protobuf_Fields_ffi"
	wire "iris.local/draw/Yoga_Protobuf_Wire_ffi"
	writerffi "iris.local/draw/Yoga_Protobuf_Writer_ffi"
	"iris.local/draw/shared"
	"iris.local/draw/shared/Data.Either+Yoga.Protobuf.Error"
	"slices"
)

type NumberWireType struct {
	Number   int
	WireType int
}

type NumberedWriter struct {
	Write  func(writerffi.Writer, shared.XY)
	Number int
}

type Cardinality uint8

const (
	CardinalityOnce Cardinality = iota
	CardinalityRepeated
	CardinalityOptional
)

type FieldShape struct {
	Kind        *ValueKind
	Name        string
	Number      int
	Cardinality Cardinality
}

type Member struct {
	Single  FieldShape
	Choice0 string
	Choice1 []FieldShape
	Tag     uint8
}

const (
	MemberSingle = iota
	MemberChoice
)

type ValueKind struct {
	Scalar       string
	Nested0      maybe.String
	Nested1      []*Member
	Enumeration0 maybe.String
	Enumeration1 []string
	Tag          uint8
}

const (
	ValueKindScalar = iota
	ValueKindNested
	ValueKindEnumeration
)

type FieldCodecResult struct {
	Kind        *ValueKind
	Noted       []NumberWireType
	Read        func(bytes.Bytes, fields.Fields, int, int) (int, either_error.DecodeErrorFailure)
	Write       func(writerffi.Writer, int)
	Cardinality Cardinality
}

type ValueCodec struct {
	IsDefault func(int) bool
	Kind      *ValueKind
	Read      func(bytes.Bytes, int, int) (tuple.TupleInt, either_error.DecodeErrorFailure)
	Write     func(writerffi.Writer, int)
	Default   int
	WireType  int
}

func Scalar(scalar string) *ValueKind {
	return &ValueKind{Tag: ValueKindScalar, Scalar: scalar}
}

//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:220:12
func MessageCodec(ignored proxy.Proxy) codec.Codec {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:220:12
	this := FieldCodec(proxy.ProxyProxy, 1)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:220:12
	this2 := FieldCodec(proxy.ProxyProxy, 2)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:120:25
	noted := array.Append(FieldCodec(proxy.ProxyProxy, 1).Noted, array.Append(FieldCodec(proxy.ProxyProxy, 2).Noted, []NumberWireType{}))
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:15
	accumulator := 0
	for numberWireTypeIndex := range noted {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:15
		number := noted[numberWireTypeIndex].Number
		if accumulator < number {
			accumulator = number
		}
	}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:15
	accumulator2 := array.Replicate(accumulator+1, -1)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:15
	accumulator2 = slices.Clip(accumulator2)
	copied := false
	for numberWireTypeIndex := range noted {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:15
		number := noted[numberWireTypeIndex].Number
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:59
		if number >= 0 && number < len(accumulator2) {
			updated := accumulator2
			if !copied {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:59
				updated = slices.Clone(updated)
				copied = true
			}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:556:59
			updated[number] = noted[numberWireTypeIndex].WireType
			accumulator2 = updated
		}
	}
	wireTypes := accumulator2
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:278:12
	this3 := FieldCodec(proxy.ProxyProxy, 1)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:278:12
	this4 := FieldCodec(proxy.ProxyProxy, 2)
//line ../../../../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
	sorted := slices.Clone(array.Append([]NumberedWriter{{Number: 1, Write: func(writer writerffi.Writer, message shared.XY) {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:279:28
		this3.Write(writer, message.X)
	}}}, array.Append([]NumberedWriter{{Number: 2, Write: func(writer writerffi.Writer, message shared.XY) {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:279:28
		this4.Write(writer, message.Y)
	}}}, []NumberedWriter{})))
//line ../../../../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
	slices.SortStableFunc(sorted, func(x NumberedWriter, y NumberedWriter) int {
		key := -x.Number
		key2 := -y.Number
		if key < key2 {
			return -1
		}
		if key > key2 {
			return 1
		}
		return 0
	})
	writers := sorted
	return codec.Codec{Read: func(input bytes.Bytes, start int, end int) (shared.XY, either_error.DecodeErrorFailure) {
//line ../../../../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
		b, err := MessageCodecLoop(wireTypes, input, end, start, fields.NewImpl(input, len(wireTypes)))
		if err.Failed {
			return shared.XY{}, err
		}
//line ../../../../../../p/either-6.1.0/src/Data/Either.purs:152:43
		i, err := this.Read(input, b, start, end)
		if err.Failed {
			return shared.XY{}, err
		}
//line ../../../../../../p/either-6.1.0/src/Data/Either.purs:152:43
		i2, err := this2.Read(input, b, start, end)
		if err.Failed {
			return shared.XY{}, err
		}
		return shared.XY{Y: i2, X: i}, either_error.DecodeErrorFailure{}
	}, Write: func(writer writerffi.Writer, message shared.XY) {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:122:28
		for _, numberedWriter := range writers {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:122:55
			numberedWriter.Write(writer, message)
		}
	}}
}

//line protobuf.go:184
func FieldCodec(ignored proxy.Proxy, number int) FieldCodecResult {
	value := ValueCodec{WireType: 0, Read: func(input bytes.Bytes, position int, end int) (tuple.TupleInt, either_error.DecodeErrorFailure) {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:668:18
		next := wire.VarintEndImpl(input, position, end)
		if next < 0 {
			var result protobuferror.DecodeError
			switch next {
			case -2:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:669:7
				result = protobuferror.MalformedVarint(position)
			case -3:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:669:7
				result = protobuferror.InvalidWireType(position)
			case -4:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:669:7
				result = protobuferror.InvalidUtf8(position)
			default:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:669:7
				result = protobuferror.Truncated(position)
			}
			return tuple.TupleInt{}, either_error.DecodeErrorFailure{Failed: true, Left: result}
		}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:669:58
		return tuple.TupleInt{First: wire.Varint32Impl(input, position), Second: next}, either_error.DecodeErrorFailure{}
	}, Write: writer.Int32_, IsDefault: func(i int) bool {
		return 0 == i
	}, Default: 0, Kind: Scalar("int32")}
	var result2 []NumberWireType
	if number < 1024 {
		result2 = []NumberWireType{{Number: number, WireType: value.WireType}}
	} else {
		result2 = []NumberWireType{}
	}
	return FieldCodecResult{Read: func(input bytes.Bytes, fields fields.Fields, ignored2 int, end int) (int, either_error.DecodeErrorFailure) {
		return Once(value, number, input, fields, ignored2, end)
	}, Noted: result2, Write: func(writer writerffi.Writer, field int) {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:619:39
		if !value.IsDefault(field) {
			value.Write(writer, field)
			writerffi.Uint32Impl(writer, number<<3|value.WireType)
		}
	}, Cardinality: CardinalityOnce, Kind: value.Kind}
}

//line protobuf.go:229
func Once(value ValueCodec, number int, input bytes.Bytes, fields2 fields.Fields, ignored int, end int) (int, either_error.DecodeErrorFailure) {
	if number < 1024 {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:505:22
		position := fields.AtImpl(fields2, number)
		if position < 0 {
			return value.Default, either_error.DecodeErrorFailure{}
		}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:506:53
		tuple, err := value.Read(input, position, end)
		if err.Failed {
			return 0, err
		}
		return tuple.First, either_error.DecodeErrorFailure{}
	}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:609:3
	return OnceLoop(value, number, input, end, ignored, value.Default)
}

//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:562:5
func MessageCodecLoop(wireTypes []int, input bytes.Bytes, end int, position int, fields2 fields.Fields) (fields.Fields, either_error.DecodeErrorFailure) {
	for {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:562:5
		if position >= end {
			return fields2, either_error.DecodeErrorFailure{}
		}
		afterTag := wire.VarintEndImpl(input, position, end)
		if afterTag < 0 {
			var result protobuferror.DecodeError
			switch afterTag {
			case -2:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:566:11
				result = protobuferror.MalformedVarint(position)
			case -3:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:566:11
				result = protobuferror.InvalidWireType(position)
			case -4:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:566:11
				result = protobuferror.InvalidUtf8(position)
			default:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:566:11
				result = protobuferror.Truncated(position)
			}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:566:11
			return *new(fields.Fields), either_error.DecodeErrorFailure{Failed: true, Left: result}
		}
		tag := wire.Varint32Impl(input, position)
		next := wire.SkipImpl(input, tag, afterTag, end)
		if next < 0 {
			var result protobuferror.DecodeError
			switch next {
			case -2:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:13
				result = protobuferror.MalformedVarint(afterTag)
			case -3:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:13
				result = protobuferror.InvalidWireType(afterTag)
			case -4:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:13
				result = protobuferror.InvalidUtf8(afterTag)
			default:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:13
				result = protobuferror.Truncated(afterTag)
			}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:13
			return *new(fields.Fields), either_error.DecodeErrorFailure{Failed: true, Left: result}
		}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:570:64
		position, fields2 = next, fields.NoteImpl(fields2, wireTypes, tag, afterTag)
	}
}

//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:596:5
func OnceLoop(value ValueCodec, number int, input bytes.Bytes, end int, position int, kept int) (int, either_error.DecodeErrorFailure) {
	for {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:596:5
		if position >= end {
			return kept, either_error.DecodeErrorFailure{}
		}
		position2 := wire.VarintEndImpl(input, position, end)
		if position2 < 0 {
			var result protobuferror.DecodeError
			switch position2 {
			case -2:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:600:11
				result = protobuferror.MalformedVarint(position)
			case -3:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:600:11
				result = protobuferror.InvalidWireType(position)
			case -4:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:600:11
				result = protobuferror.InvalidUtf8(position)
			default:
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:600:11
				result = protobuferror.Truncated(position)
			}
			return 0, either_error.DecodeErrorFailure{Failed: true, Left: result}
		}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:602:23
		tag := wire.Varint32Impl(input, position)
		decodeError := protobuferror.MalformedVarint(position2)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:603:13
		decodeError2 := protobuferror.InvalidWireType(position2)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:603:13
		decodeError3 := protobuferror.InvalidUtf8(position2)
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:603:13
		decodeError4 := protobuferror.Truncated(position2)
		if int(uint(tag)>>3) == number {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:603:49
			var tuple2 tuple.TupleInt
			var err either_error.DecodeErrorFailure
			if tag&7 == value.WireType {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:510:46
				tuple2, err = value.Read(input, position2, end)
			} else {
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:613:14
				next := wire.SkipImpl(input, tag, position2, end)
				if next < 0 {
					var result protobuferror.DecodeError
					switch next {
					case -2:
						result = decodeError
					case -3:
						result = decodeError2
					case -4:
						result = decodeError3
					default:
						result = decodeError4
					}
					err = either_error.DecodeErrorFailure{Failed: true, Left: result}
				} else {
					tuple2, err = tuple.TupleInt{First: kept, Second: next}, either_error.DecodeErrorFailure{}
				}
			}
			if err.Failed {
				return 0, err
			}
			position, kept = tuple2.Second, tuple2.First
			continue
		}
//line ../../../../../../p/yoga-connect/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect/src/Yoga/Connect/Protobuf.purs:607:26
		next2 := wire.SkipImpl(input, tag, position2, end)
		if next2 < 0 {
			var result protobuferror.DecodeError
			switch next2 {
			case -2:
				result = decodeError
			case -3:
				result = decodeError2
			case -4:
				result = decodeError3
			default:
				result = decodeError4
			}
			return 0, either_error.DecodeErrorFailure{Failed: true, Left: result}
		}
		position = next2
	}
}

package protobuferror

type DecodeError struct {
	Field0 int
	Tag    uint8
}

const (
	DecodeErrorTruncated = iota
	DecodeErrorMalformedVarint
	DecodeErrorInvalidWireType
	DecodeErrorInvalidUtf8
)

func MalformedVarint(field0 int) DecodeError {
	return DecodeError{Tag: DecodeErrorMalformedVarint, Field0: field0}
}

func InvalidWireType(field0 int) DecodeError {
	return DecodeError{Tag: DecodeErrorInvalidWireType, Field0: field0}
}

func InvalidUtf8(field0 int) DecodeError {
	return DecodeError{Tag: DecodeErrorInvalidUtf8, Field0: field0}
}

func Truncated(field0 int) DecodeError {
	return DecodeError{Tag: DecodeErrorTruncated, Field0: field0}
}

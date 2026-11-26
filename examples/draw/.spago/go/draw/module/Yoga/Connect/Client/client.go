package client

import (
	"iris.local/draw/Yoga/Connect/Status"
	bytes "iris.local/draw/Yoga_Protobuf_Bytes_ffi"
	"iris.local/draw/shared"
)

//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
func FailedWith(code int, message string, typeNames []string, encodings []bytes.Bytes) status.EitherConnectError {
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
	var result status.Code
	switch code {
	case 1:
		result = status.CodeCanceled
	case 3:
		result = status.CodeInvalidArgument
	case 4:
		result = status.CodeDeadlineExceeded
	case 5:
		result = status.CodeNotFound
	case 6:
		result = status.CodeAlreadyExists
	case 7:
		result = status.CodePermissionDenied
	case 8:
		result = status.CodeResourceExhausted
	case 9:
		result = status.CodeFailedPrecondition
	case 10:
		result = status.CodeAborted
	case 11:
		result = status.CodeOutOfRange
	case 12:
		result = status.CodeUnimplemented
	case 13:
		result = status.CodeInternal
	case 14:
		result = status.CodeUnavailable
	case 15:
		result = status.CodeDataLoss
	case 16:
		result = status.CodeUnauthenticated
	default:
		result = status.CodeUnknown
	}
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
	count := min(len(typeNames), len(encodings))
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
	built := make([]shared.ErrorDetail, count)
	for index := range count {
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
		value := typeNames[index]
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
		value2 := encodings[index]
		element := shared.ErrorDetail{TypeName: value, Bytes: value2}
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
		built[index] = element
	}
//line ../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:220:47
	return status.LeftEitherConnectError(status.ConnectError{Code: result, Message: message, Details: built})
}

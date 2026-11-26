package client

import (
	"github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/prim"
	dataeither "iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Go/IO"
	chanffi "iris.local/draw-test/Go_IO_Chan_ffi"
	context "iris.local/draw-test/Go_IO_Context_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
	internal "iris.local/draw-test/Go_IO_Internal_ffi"
	mvar "iris.local/draw-test/Go_IO_MVar_ffi"
	selectffi "iris.local/draw-test/Go_IO_Select_ffi"
	ioffi "iris.local/draw-test/Go_IO_ffi"
	"iris.local/draw-test/Type/Proxy"
	"iris.local/draw-test/Yoga/Connect/Client"
	"iris.local/draw-test/Yoga/Connect/Protobuf"
	"iris.local/draw-test/Yoga/Connect/Rejection"
	"iris.local/draw-test/Yoga/Connect/Status"
	"iris.local/draw-test/Yoga/OmIO/Strom"
	"iris.local/draw-test/Yoga/Protobuf/Codec"
	"iris.local/draw-test/Yoga/Protobuf/Error"
	clientffi "iris.local/draw-test/Yoga_Connect_Client_ffi"
	bytes "iris.local/draw-test/Yoga_Protobuf_Bytes_ffi"
	writer "iris.local/draw-test/Yoga_Protobuf_Writer_ffi"
	"iris.local/draw-test/shared"
	"iris.local/draw-test/shared/Yoga.OmIO.Strom.Scope+Yoga.Connect.Status"
	"slices"
	"strconv"
)

type EitherStepXY struct {
	Left  status.VariantConnectErrorException
	Right strom.StepXY
	Tag   uint8
}

const (
	EitherStepXYLeft = iota
	EitherStepXYRight
)

type Draw struct {
	Draw func(func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure)) func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure)
}

type Method struct {
	Procedure string
	Request   codec.Codec
	Response  codec.Codec
}

type Signal struct {
	Field1      []any
	FailedWith1 status.VariantConnectErrorException
	Field0      int
	FailedWith0 int
	Tag         uint8
}

const (
	SignalPulled = iota
	SignalFinished
	SignalFailedWith
)

type EitherStepBoolean struct {
	Left  status.VariantConnectErrorException
	Right strom.StepBoolean
	Tag   uint8
}

const (
	EitherStepBooleanLeft = iota
	EitherStepBooleanRight
)

type EitherStep struct {
	Left  status.VariantConnectErrorException
	Right *strom.Step
	Tag   uint8
}

const (
	EitherStepLeft = iota
	EitherStepRight
)

type TupleMaybeInt struct {
	First  scope_status.Finalizers
	Second maybe.Int
}

type TupleUnit struct {
	Second struct{}
	First  scope_status.Finalizers
}

type InitRest struct {
	Init []bool
	Rest []bool
}

func LeftEitherStepXY(left status.VariantConnectErrorException) EitherStepXY {
	return EitherStepXY{Tag: EitherStepXYLeft, Left: left}
}

func RightEitherStepXY(right strom.StepXY) EitherStepXY {
	return EitherStepXY{Tag: EitherStepXYRight, Right: right}
}

func RightEitherStepBoolean(right strom.StepBoolean) EitherStepBoolean {
	return EitherStepBoolean{Tag: EitherStepBooleanRight, Right: right}
}

func LeftEitherStepBoolean(left status.VariantConnectErrorException) EitherStepBoolean {
	return EitherStepBoolean{Tag: EitherStepBooleanLeft, Left: left}
}

//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Om/Client.purs:94:100
func Client(connection clientffi.Connection) Draw {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Om/Client.purs:94:100
	procedure := prim.Concat("/draw.v1.Canvas/", prim.StringToUpper(prim.CodeUnitTake(1, "draw")), prim.CodeUnitDrop(1, "draw"))
	exitRecordConnectError := scope_status.Exit{Tag: scope_status.ExitInterrupted}
	return Draw{Draw: func(requests func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure)) func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Om/Client.purs:156:42
		method := Method{Procedure: procedure, Request: protobuf.MessageCodec(proxy.ProxyProxy), Response: protobuf.MessageCodec(proxy.ProxyProxy)}
		return func(scope scope_status.Scope, r struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:164:17
			call := clientffi.OpenBidiStreamImpl(scope.Context, connection, method.Procedure)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			_, err := AddFinalizer(scope, func(_ scope_status.Exit, _ struct{}) status.EitherVariantConnectErrorExceptionUnit {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:204:22
				clientffi.CloseImpl(call)
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:204:22
				return status.RightEitherVariantConnectErrorExceptionUnit(struct{}{})
			}, r)
			if err.Failed {
				a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:40
				var result2 bool
				switch a.Tag {
				case status.VariantConnectErrorExceptionConnectError:
					result2 = false
				default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
					result2 = exception.Message(a.Exception) == "the stream was interrupted"
				}
				var result status.VariantConnectErrorException
				if result2 {
					result = a
				} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:81
					result = ClientChange(a)
				}
				return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: result}
			}
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:166:20
			sendFailure := ioref.New(status.NoVariantConnectErrorException)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:1050:14
			handoff := chanffi.NewBufferedChanImpl[*Signal](1)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			producerScope, err := ForkScope(scope, r)
			if err.Failed {
				a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:40
				var result bool
				switch a.Tag {
				case status.VariantConnectErrorExceptionConnectError:
					result = false
				default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
					result = exception.Message(a.Exception) == "the stream was interrupted"
				}
				var result2 status.VariantConnectErrorException
				if result {
					result2 = a
				} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:81
					result2 = ClientChange(a)
				}
				return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: result2}
			}
//line ../../../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/MVar.purs:35:19
			mVar := mvar.NewEmptyMVarImpl[struct{}]()
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			_, err = AddFinalizer(scope, func(_ scope_status.Exit, _ struct{}) status.EitherVariantConnectErrorExceptionUnit {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:769:33
				mvar.ReadMVarImpl(mVar)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:769:33
				return status.RightEitherVariantConnectErrorExceptionUnit(struct{}{})
			}, r)
			if err.Failed {
				a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:40
				var result bool
				switch a.Tag {
				case status.VariantConnectErrorExceptionConnectError:
					result = false
				default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
					result = exception.Message(a.Exception) == "the stream was interrupted"
				}
				var result2 status.VariantConnectErrorException
				if result {
					result2 = a
				} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:81
					result2 = ClientChange(a)
				}
				return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: result2}
			}
//line ../../../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:52:43
			ioffi.ForkIOImpl(func() {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:771:16
				io.Finally(func() {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
					either2, err3 := internal.CatchResultsImpl(func(x exception.Error) (status.EitherBoolean, dataeither.ErrorFailure) {
						return status.EitherBoolean{}, dataeither.ErrorFailure{Failed: true, Left: x}
					}, func() (status.EitherBoolean, dataeither.ErrorFailure) {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						owned, err2 := ForkScope(producerScope, r)
						if err2.Failed {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
							return status.LeftEitherBoolean(err2.Left), dataeither.ErrorFailure{}
						}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						_, err2 = AddFinalizer(owned, func(value scope_status.Exit, _ struct{}) status.EitherVariantConnectErrorExceptionUnit {
							switch value.Tag {
							case scope_status.ExitEnded:
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:182:20
								clientffi.CloseSendingImpl(call)
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:182:20
								return status.RightEitherVariantConnectErrorExceptionUnit(struct{}{})
							case scope_status.ExitFailed:
								sendFailure.Store(status.JustVariantConnectErrorException(value.Failed))
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:204:22
								clientffi.CloseImpl(call)
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:204:22
								return status.RightEitherVariantConnectErrorExceptionUnit(struct{}{})
							default: // scope_status.ExitInterrupted
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
								return status.RightEitherVariantConnectErrorExceptionUnit(struct{}{})
							}
						}, r)
						if err2.Failed {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
							return status.LeftEitherBoolean(err2.Left), dataeither.ErrorFailure{}
						}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						f, err2 := OpenClosingAtEndStream2(owned, requests, method, call, r)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						var either status.EitherBoolean
						if err2.Failed {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							either = status.LeftEitherBoolean(err2.Left)
						} else {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
							either = status.EitherBooleanFromResults(ClientHandOverAll(handoff, producerScope, func(r2 struct{}) *EitherStep {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
								return ClientPullUntilEnd(f, r2)
							}, r))
						}
						return either, dataeither.ErrorFailure{}
					})
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
					var result status.EitherBoolean
					if err3.Failed {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						result = status.LeftEitherBoolean(status.Exception(err3.Left))
					} else {
						result = either2
					}
					if result.Tag == status.EitherBooleanRight && result.Right {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						_, err2 := ClientHandOver(handoff, producerScope, &Signal{Tag: SignalFinished, Field0: 0}, r)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						if !err2.Failed {
							scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitEnded}, producerScope, r)
						}
					} else if result.Tag == status.EitherBooleanRight && !result.Right {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						scope_status.CloseScope(exitRecordConnectError, producerScope, r)
					} else {
						failure := result.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
						boolean := context.IsCancelledImpl(producerScope.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
						var result2 bool
						switch failure.Tag {
						case status.VariantConnectErrorExceptionConnectError:
							result2 = false
						default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
							result2 = exception.Message(failure.Exception) == "the stream was interrupted"
						}
						var field scope_status.Exit
						if boolean || result2 {
							field = exitRecordConnectError
						} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
							field = scope_status.Failed(failure)
						}
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						switch field.Tag {
						case scope_status.ExitInterrupted:
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							scope_status.CloseScope(field, producerScope, r)
						default:
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							_, err2 := ClientHandOver(handoff, producerScope, &Signal{Tag: SignalFailedWith, FailedWith0: 0, FailedWith1: failure}, r)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							if !err2.Failed {
								scope_status.CloseScope(field, producerScope, r)
							}
						}
					}
				}, func() {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:771:73
					mvar.PutMVarImpl(mVar, struct{}{})
				})
			})
			return func(_ struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:191:19
				boolean := clientffi.ReceiveImpl(call)
				var step strom.StepXY
				var err2 status.VariantConnectErrorExceptionFailure
				if boolean {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:193:18
					bytes2 := clientffi.ReceivedImpl(call)
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:193:18
					xy, err3 := method.Response.Read(bytes2, 0, bytes.LengthImpl(bytes2))
					if err3.Failed {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
						maybe2 := sendFailure.Load()
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						switch maybe2.Tag {
						case maybe.TagJust:
							err2 = status.VariantConnectErrorExceptionFailure{Failed: true, Left: maybe2.Just}
						default: // maybe.TagNothing
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
							var result string
							switch err3.Left.Tag {
							case protobuferror.DecodeErrorTruncated:
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:24:27
								result = "(Truncated " + strconv.Itoa(err3.Left.Field0) + ")"
							case protobuferror.DecodeErrorMalformedVarint:
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:25:33
								result = "(MalformedVarint " + strconv.Itoa(err3.Left.Field0) + ")"
							case protobuferror.DecodeErrorInvalidWireType:
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:26:33
								result = "(InvalidWireType " + strconv.Itoa(err3.Left.Field0) + ")"
							default: // protobuferror.DecodeErrorInvalidUtf8
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:27:29
								result = "(InvalidUtf8 " + strconv.Itoa(err3.Left.Field0) + ")"
							}
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
							err2 = status.VariantConnectErrorExceptionFailure{Failed: true, Left: status.ConnectErrorVariantConnectErrorException(status.ConnectError{Code: status.CodeInternal, Message: "the response doesn't decode: " + result, Details: []shared.ErrorDetail{}})}
						}
					} else {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
						maybe2 := sendFailure.Load()
						switch maybe2.Tag {
						case maybe.TagJust:
							err2 = status.VariantConnectErrorExceptionFailure{Failed: true, Left: maybe2.Just}
						default: // maybe.TagNothing
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
							step, err2 = strom.EmitStepXY([]shared.XY{xy}), status.VariantConnectErrorExceptionFailure{}
						}
					}
				} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
					boolean2 := context.IsCancelledImpl(scope.Context)
					if boolean2 {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
						maybe2 := sendFailure.Load()
						switch maybe2.Tag {
						case maybe.TagJust:
							err2 = status.VariantConnectErrorExceptionFailure{Failed: true, Left: maybe2.Just}
						default: // maybe.TagNothing
							err2 = status.VariantConnectErrorExceptionFailure{Failed: true, Left: scope_status.Interrupted}
						}
					} else {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:198:20
						either := clientffi.OutcomeImpl(func(code int) func(string) func([]string) func([]bytes.Bytes) status.EitherConnectError {
							return func(message string) func([]string) func([]bytes.Bytes) status.EitherConnectError {
								return func(typeNames []string) func([]bytes.Bytes) status.EitherConnectError {
									return func(encodings []bytes.Bytes) status.EitherConnectError {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:198:20
										return client.FailedWith(code, message, typeNames, encodings)
									}
								}
							}
						}, status.RightEitherConnectError(struct{}{}), call)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						switch either.Tag {
						case status.EitherConnectErrorLeft:
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							step, err2 = ClientJoin(sendFailure, LeftEitherStepXY(status.ConnectErrorVariantConnectErrorException(either.Left)))
						default: // status.EitherConnectErrorRight
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							step, err2 = ClientJoin(sendFailure, RightEitherStepXY(strom.StepXY{Tag: strom.StepXYEnd}))
						}
					}
				}
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				if err2.Failed {
					a := err2.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:40
					var result bool
					switch a.Tag {
					case status.VariantConnectErrorExceptionConnectError:
						result = false
					default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
						result = exception.Message(a.Exception) == "the stream was interrupted"
					}
					var result2 status.VariantConnectErrorException
					if result {
						result2 = a
					} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:682:81
						result2 = ClientChange(a)
					}
					return strom.StepXY{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: result2}
				}
				return step, status.VariantConnectErrorExceptionFailure{}
			}, status.VariantConnectErrorExceptionFailure{}
		}
	}}
}

//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
func AddFinalizer(scope scope_status.Scope, finalize func(scope_status.Exit, struct{}) status.EitherVariantConnectErrorExceptionUnit, r struct{}) (int, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
	maybe2 := ioref.AtomicModifyIORefImpl(FstMaybeFinalizerKey, SndMaybeFinalizerKey, func(value scope_status.Finalizers) TupleMaybeInt {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:145:3
		switch value.Tag {
		case scope_status.FinalizersOpen:
			key := value.Open0
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:146:28
			return TupleMaybeInt{First: scope_status.Open(key+1, array.Snoc(value.Open1, scope_status.Registered{First: key, Second: finalize})), Second: maybe.JustInt(key)}
		default: // scope_status.FinalizersClosed
			return TupleMaybeInt{First: scope_status.Finalizers{Tag: scope_status.FinalizersClosed}, Second: maybe.NoInt}
		}
	}, scope.Finalizers)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch maybe2.Tag {
	case maybe.TagJust:
		return maybe2.Just, status.VariantConnectErrorExceptionFailure{}
	default: // maybe.TagNothing
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		either := finalize(scope_status.Exit{Tag: scope_status.ExitInterrupted}, r)
		var i int
		var err status.VariantConnectErrorExceptionFailure
		switch either.Tag {
		case status.EitherVariantConnectErrorExceptionUnitLeft:
			err = status.VariantConnectErrorExceptionFailure{Failed: true, Left: either.Left}
		default: // status.EitherVariantConnectErrorExceptionUnitRight
			i, err = -1, status.VariantConnectErrorExceptionFailure{}
		}
		return i, err
	}
}

//line client.go:469
func ForkScope(parent scope_status.Scope, r struct{}) (scope_status.Scope, status.VariantConnectErrorExceptionFailure) {
	x := func() scope_status.Scope {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:121:12
		return scope_status.NewScope(parent.Context)
	}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	scope := x()
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	i, err := AddFinalizer(parent, func(_ scope_status.Exit, unit struct{}) status.EitherVariantConnectErrorExceptionUnit {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		return status.EitherVariantConnectErrorExceptionUnitFromResults(scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitInterrupted}, scope, unit))
	}, r)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	if err.Failed {
		return scope_status.Scope{}, err
	}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:134:3
	scope.DetachFromParent.Store(func() {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
		ioref.AtomicModifyIORefImpl(FstUnit, func(tuple TupleUnit) struct{} {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
			SndUnit(tuple)
			return struct{}{}
		}, func(state scope_status.Finalizers) TupleUnit {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:87
			var result scope_status.Finalizers
			switch state.Tag {
			case scope_status.FinalizersOpen:
				kept := []scope_status.Registered{}
				for _, x2 := range state.Open1 {
					if x2.First != i {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
						kept = append(kept, x2)
					}
				}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
				result = scope_status.Open(state.Open0, kept)
			default: // scope_status.FinalizersClosed
				result = scope_status.Finalizers{Tag: scope_status.FinalizersClosed}
			}
			return TupleUnit{First: result, Second: struct{}{}}
		}, parent.Finalizers)
	})
	return scope, status.VariantConnectErrorExceptionFailure{}
}

//line client.go:516
func FstMaybeFinalizerKey(tuple TupleMaybeInt) scope_status.Finalizers {
	return tuple.First
}

func SndMaybeFinalizerKey(tuple TupleMaybeInt) maybe.Int {
	return tuple.Second
}

func FstUnit(tuple TupleUnit) scope_status.Finalizers {
	return tuple.First
}

func SndUnit(tuple TupleUnit) {
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func OpenClosingAtEndStream(owned scope_status.Scope, f func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure), method Method, clientCall clientffi.ClientCall, r struct{}) (func(struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	f2, err := f(owned, r)
	exitRecordConnectError := scope_status.Exit{Tag: scope_status.ExitInterrupted}
	if err.Failed {
		a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		var result2 bool
		switch a.Tag {
		case status.VariantConnectErrorExceptionConnectError:
			result2 = false
		default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
			result2 = exception.Message(a.Exception) == "the stream was interrupted"
		}
		var result scope_status.Exit
		if boolean || result2 {
			result = exitRecordConnectError
		} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			result = scope_status.Failed(a)
		}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		_, err2 := scope_status.CloseScope(result, owned, r)
		if err2.Failed {
			return nil, err2
		}
		return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: a}
	}
	return func(r2 struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err2 := f2(r2)
		if err2.Failed {
			a := err2.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
			boolean := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			var result bool
			switch a.Tag {
			case status.VariantConnectErrorExceptionConnectError:
				result = false
			default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
				result = exception.Message(a.Exception) == "the stream was interrupted"
			}
			var result2 scope_status.Exit
			if boolean || result {
				result2 = exitRecordConnectError
			} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
				result2 = scope_status.Failed(a)
			}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			_, err3 := scope_status.CloseScope(result2, owned, r2)
			if err3.Failed {
				return strom.StepBoolean{}, err3
			}
			return strom.StepBoolean{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: a}
		}
		switch step.Tag {
		case strom.StepXYEmit:
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			array_, err3 := OpenClosingAtEndStreamFrom(method, clientCall, step.Emit, 0, []bool{})
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err3.Failed {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				return OpenClosingAtEndStreamJoin(owned, r2, strom.StepBoolean{}, err3)
			}
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			return OpenClosingAtEndStreamJoin(owned, r2, strom.EmitStepBoolean(array_), status.VariantConnectErrorExceptionFailure{})
		default: // strom.StepXYEnd
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			_, err3 := scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitEnded}, owned, r2)
			var step2 strom.StepBoolean
			var err4 status.VariantConnectErrorExceptionFailure
			if err3.Failed {
				err4 = err3
			} else {
				step2, err4 = strom.StepBoolean{Tag: strom.StepBooleanEnd}, status.VariantConnectErrorExceptionFailure{}
			}
			return step2, err4
		}
	}, status.VariantConnectErrorExceptionFailure{}
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func OpenClosingAtEndStream2(owned scope_status.Scope, f func(scope_status.Scope, struct{}) (func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure), method Method, clientCall clientffi.ClientCall, r struct{}) (func(struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure), status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	upstream, err := ForkScope(owned, r)
	exitRecordConnectError := scope_status.Exit{Tag: scope_status.ExitInterrupted}
	exitRecordConnectError2 := scope_status.Exit{Tag: scope_status.ExitEnded}
	stepBoolean := strom.StepBoolean{Tag: strom.StepBooleanEnd}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	eitherVariantConnectErrorExceptionStepBoolean := RightEitherStepBoolean(stepBoolean)
	if err.Failed {
		a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		var result2 bool
		switch a.Tag {
		case status.VariantConnectErrorExceptionConnectError:
			result2 = false
		default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
			result2 = exception.Message(a.Exception) == "the stream was interrupted"
		}
		var result scope_status.Exit
		if boolean || result2 {
			result = exitRecordConnectError
		} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			result = scope_status.Failed(a)
		}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		_, err2 := scope_status.CloseScope(result, owned, r)
		if err2.Failed {
			return nil, err2
		}
		return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: a}
	}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	f2, err := OpenClosingAtEndStream(upstream, f, method, clientCall, r)
	if err.Failed {
		a := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		var result bool
		switch a.Tag {
		case status.VariantConnectErrorExceptionConnectError:
			result = false
		default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
			result = exception.Message(a.Exception) == "the stream was interrupted"
		}
		var result2 scope_status.Exit
		if boolean || result {
			result2 = exitRecordConnectError
		} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			result2 = scope_status.Failed(a)
		}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		_, err2 := scope_status.CloseScope(result2, owned, r)
		if err2.Failed {
			return nil, err2
		}
		return nil, status.VariantConnectErrorExceptionFailure{Failed: true, Left: a}
	}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:474:11
	ioRef := ioref.New(false)
	return func(r2 struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure) {
		boolean := ioRef.Load()
		if boolean {
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			_, err2 := scope_status.CloseScope(exitRecordConnectError2, owned, r2)
			var result EitherStepBoolean
			if err2.Failed {
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
				result = LeftEitherStepBoolean(err2.Left)
			} else {
				result = eitherVariantConnectErrorExceptionStepBoolean
			}
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			return EitherStepBooleanResults(result)
		}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err3 := f2(r2)
		if err3.Failed {
			a := err3.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
			boolean2 := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			var result bool
			switch a.Tag {
			case status.VariantConnectErrorExceptionConnectError:
				result = false
			default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
				result = exception.Message(a.Exception) == "the stream was interrupted"
			}
			var result2 scope_status.Exit
			if boolean2 || result {
				result2 = exitRecordConnectError
			} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
				result2 = scope_status.Failed(a)
			}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			_, err2 := scope_status.CloseScope(result2, owned, r2)
			if err2.Failed {
				return strom.StepBoolean{}, err2
			}
			return strom.StepBoolean{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: a}
		}
		switch step.Tag {
		case strom.StepBooleanEnd:
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			_, err2 := scope_status.CloseScope(exitRecordConnectError2, owned, r2)
			var result EitherStepBoolean
			if err2.Failed {
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
				result = LeftEitherStepBoolean(err2.Left)
			} else {
				result = eitherVariantConnectErrorExceptionStepBoolean
			}
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			return EitherStepBooleanResults(result)
		default: // strom.StepBooleanEmit
			chunk := step.Emit
//line ../../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1022:3
			scrutinee := strom.OpenClosingAtEndStreamGo(chunk, 0)
			var result InitRest
			if scrutinee.Tag == maybe.TagJust && scrutinee.Just == 0 {
				result = InitRest{Init: []bool{}, Rest: chunk}
			} else if scrutinee.Tag == maybe.TagJust {
				i := scrutinee.Just
//line ../../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1026:7
				result = InitRest{Init: array.Slice(0, i, chunk), Rest: array.Slice(i, len(chunk), chunk)}
			} else {
				result = InitRest{Init: chunk, Rest: []bool{}}
			}
			init_ := result.Init
//line ../../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1022:3
			stepBoolean2 := strom.EmitStepBoolean(init_)
//line ../../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1022:3
			eitherVariantConnectErrorExceptionStepBoolean2 := RightEitherStepBoolean(stepBoolean2)
//line ../../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1022:3
			if len(result.Rest) != 0 {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:485:13
				ioRef.Store(true)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				_, err2 := scope_status.CloseScope(exitRecordConnectError, upstream, r2)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				if err2.Failed {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
					return OpenClosingAtEndStreamJoin(owned, r2, strom.StepBoolean{}, err2)
				}
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				return OpenClosingAtEndStreamJoin(owned, r2, eitherVariantConnectErrorExceptionStepBoolean2.Right, EitherStepBooleanFailure(eitherVariantConnectErrorExceptionStepBoolean2))
			}
			return stepBoolean2, status.VariantConnectErrorExceptionFailure{}
		}
	}, status.VariantConnectErrorExceptionFailure{}
}

//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Om/Client.purs:165:1
func ClientChange(failure status.VariantConnectErrorException) status.VariantConnectErrorException {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Om/Client.purs:165:1
	switch failure.Tag {
	case status.VariantConnectErrorExceptionConnectError:
		scrutinee := rejection.Rejections(proxy.ProxyProxy).Decode(failure.ConnectError.Code, failure.ConnectError.Message, failure.ConnectError.Details)
		switch scrutinee.Tag {
		case maybe.TagNothing:
			return failure
		default: // maybe.TagJust
//line ../../../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
			panic("Failed pattern match")
		}
	default: // status.VariantConnectErrorExceptionException
		return failure
	}
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func ClientHandOverAll(handoff chanffi.Chan[*Signal], producerScope scope_status.Scope, pull func(struct{}) *EitherStep, r struct{}) (bool, status.VariantConnectErrorExceptionFailure) {
	for {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		either := pull(r)
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		switch either.Tag {
		case EitherStepLeft:
			return false, status.VariantConnectErrorExceptionFailure{Failed: true, Left: either.Left}
		default: // EitherStepRight
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			switch either.Right.Tag {
			case strom.StepEnd:
				return true, status.VariantConnectErrorExceptionFailure{}
			default: // strom.StepEmit
				context2 := producerScope.Context
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:39:16
				cancelled := context.IsCancelledImpl(context2)
				var boolean bool
				if cancelled {
					boolean = false
				} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:41:8
					boolean = selectffi.Select2Impl(selectffi.OnSendImpl(handoff, &Signal{Tag: SignalPulled, Field0: 0, Field1: either.Right.Emit}, func() bool {
						return true
					}), selectffi.OnReceiveImpl(context.DoneImpl(context2), func(_ struct{}, _ bool) bool {
						return false
					}))
				}
				if boolean {
					continue
				}
				return false, status.VariantConnectErrorExceptionFailure{}
			}
		}
	}
}

//line client.go:838
func ClientHandOver(handoff chanffi.Chan[*Signal], producerScope scope_status.Scope, signal *Signal, unit struct{}) (struct{}, status.VariantConnectErrorExceptionFailure) {
	context2 := producerScope.Context
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:39:16
	cancelled := context.IsCancelledImpl(context2)
	if !cancelled {
		selectffi.Select2Impl(selectffi.OnSendImpl(handoff, signal, func() bool {
			return true
		}), selectffi.OnReceiveImpl(context.DoneImpl(context2), func(_ struct{}, _ bool) bool {
			return false
		}))
	}
	return struct{}{}, status.VariantConnectErrorExceptionFailure{}
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func ClientPullUntilEnd(pullUpstream func(struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure), r struct{}) *EitherStep {
	for {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err := pullUpstream(r)
		if err.Failed {
			return &EitherStep{Tag: EitherStepLeft, Left: err.Left}
		}
		switch step.Tag {
		case strom.StepBooleanEnd:
			return &EitherStep{Tag: EitherStepRight, Right: &strom.Step{Tag: strom.StepEnd}}
		default: // strom.StepBooleanEmit
			continue
		}
	}
}

//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
func ClientJoin(sendFailure ioref.IORef[status.MaybeVariantConnectErrorException], b EitherStepXY) (strom.StepXY, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/yoga-connect-go/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-connect-go/src/Yoga/Connect/Client.purs:170:24
	maybe2 := sendFailure.Load()
	if maybe2.Tag == maybe.TagJust {
		return strom.StepXY{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: maybe2.Just}
	}
	if maybe2.Tag == maybe.TagNothing && b.Tag == EitherStepXYLeft {
		return strom.StepXY{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: b.Left}
	}
	return b.Right, status.VariantConnectErrorExceptionFailure{}
}

//line client.go:883
func OpenClosingAtEndStreamFrom(method Method, clientCall clientffi.ClientCall, chunk []shared.XY, index int, field0 []bool) ([]bool, status.VariantConnectErrorExceptionFailure) {
	field0 = slices.Clip(field0)
	for {
		if index < len(chunk) {
			element := chunk[index]
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:68:32
			bytes := OpenClosingAtEndStreamFromAttempt(element, method, 0)
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:68:32
			boolean := clientffi.SendImpl(clientCall, bytes)
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:68:32
			index, field0 = index+1, append(field0, boolean)
			continue
		}
		return field0, status.VariantConnectErrorExceptionFailure{}
	}
}

//line client.go:901
const (
	OpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromRetryStep = iota
	OpenClosingAtEndStreamFromAttemptStep
)

func OpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromRetryOrOpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromAttempt(step int, element shared.XY, method Method, writer2 writer.Writer, capacity int) bytes.Bytes {
	for {
		switch step {
		case OpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromRetryStep:
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:76:17
			tooSmall := writer.CapacityImpl(writer2)
			step, capacity = OpenClosingAtEndStreamFromAttemptStep, 2*tooSmall
		default: // OpenClosingAtEndStreamFromAttemptStep
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:71:15
			writer3 := writer.AcquireImpl(capacity)
			method.Request.Write(writer3, element)
			overflowed := writer.OverflowedImpl(writer3)
			if overflowed {
				step, writer2 = OpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromRetryStep, writer3
				continue
			}
//line ../../../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:74:42
			return writer.FinishImpl(writer3)
		}
	}
}

//line client.go:929
func OpenClosingAtEndStreamFromAttempt(element shared.XY, method Method, capacity int) bytes.Bytes {
	var writer2 writer.Writer
	return OpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromRetryOrOpenClosingAtEndUnitRecordConnectErrorBooleanStreamFromAttempt(OpenClosingAtEndStreamFromAttemptStep, element, method, writer2, capacity)
}

//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
func OpenClosingAtEndStreamJoin(owned scope_status.Scope, r struct{}, step strom.StepBoolean, err status.VariantConnectErrorExceptionFailure) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
	if err.Failed {
		failure := err.Left
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(owned.Context)
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		var result2 bool
		switch failure.Tag {
		case status.VariantConnectErrorExceptionConnectError:
			result2 = false
		default: // status.VariantConnectErrorExceptionException
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
			result2 = exception.Message(failure.Exception) == "the stream was interrupted"
		}
		var result scope_status.Exit
		if boolean || result2 {
			result = scope_status.Exit{Tag: scope_status.ExitInterrupted}
		} else {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			result = scope_status.Failed(failure)
		}
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
		_, err2 := scope_status.CloseScope(result, owned, r)
		if err2.Failed {
			return strom.StepBoolean{}, err2
		}
		return strom.StepBoolean{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: failure}
	}
	if !err.Failed && step.Tag == strom.StepBooleanEnd {
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		_, err2 := scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitEnded}, owned, r)
		var step2 strom.StepBoolean
		var err3 status.VariantConnectErrorExceptionFailure
		if err2.Failed {
			err3 = err2
		} else {
			step2, err3 = strom.StepBoolean{Tag: strom.StepBooleanEnd}, status.VariantConnectErrorExceptionFailure{}
		}
		return step2, err3
	}
	return step, status.VariantConnectErrorExceptionFailure{}
}

//line client.go:980
func EitherStepBooleanFailure(either EitherStepBoolean) status.VariantConnectErrorExceptionFailure {
	if either.Tag == EitherStepBooleanLeft {
		return status.VariantConnectErrorExceptionFailure{Failed: true, Left: either.Left}
	}
	return status.VariantConnectErrorExceptionFailure{}
}

func EitherStepBooleanResults(either EitherStepBoolean) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure) {
	if either.Tag == EitherStepBooleanLeft {
		return strom.StepBoolean{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: either.Left}
	}
	return either.Right, status.VariantConnectErrorExceptionFailure{}
}

package server

import (
	"github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/List/Types"
	"iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Set"
	"iris.local/draw-test/Data/Tuple"
	"iris.local/draw-test/Data/Variant"
	"iris.local/draw-test/Go/IO"
	chanffi "iris.local/draw-test/Go_IO_Chan_ffi"
	context "iris.local/draw-test/Go_IO_Context_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
	internalffi "iris.local/draw-test/Go_IO_Internal_ffi"
	mvar "iris.local/draw-test/Go_IO_MVar_ffi"
	selectffi "iris.local/draw-test/Go_IO_Select_ffi"
	ioffi "iris.local/draw-test/Go_IO_ffi"
	"iris.local/draw-test/Type/Proxy"
	"iris.local/draw-test/Yoga/Connect/Protobuf"
	"iris.local/draw-test/Yoga/OmIO/Strom"
	stromscope "iris.local/draw-test/Yoga/OmIO/Strom/Scope"
	"iris.local/draw-test/Yoga/Protobuf/Codec"
	"iris.local/draw-test/Yoga/Protobuf/Error"
	server "iris.local/draw-test/Yoga_Connect_Server_ffi"
	bytes "iris.local/draw-test/Yoga_Protobuf_Bytes_ffi"
	writer "iris.local/draw-test/Yoga_Protobuf_Writer_ffi"
	"iris.local/draw-test/shared"
	"iris.local/draw-test/shared/Data.Maybe+Data.Tuple"
	"iris.local/draw-test/shared/Yoga.OmIO.Strom+Data.Map.Internal"
	"iris.local/draw-test/shared/Yoga.OmIO.Strom.Scope+Data.Map.Internal"
	"slices"
	"strconv"
)

type TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY struct {
	First  []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure)
	Second int
}

type RequestsRoom struct {
	Requests func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure)
	Room     ioref.IORef[internal.Canvas]
}

type TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt struct {
	First  []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure)
	Second int
}

type TupleUnit struct {
	Second struct{}
	First  scope_internal.Finalizers
}

type TupleMaybeInt struct {
	First  scope_internal.Finalizers
	Second maybe.Int
}

type TupleFinalizersRoomNoRequest struct {
	First  scope_internal.Finalizers
	Second scope_internal.Finalizers
}

func EmitStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt(emit []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure)) strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt {
	return strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt{Tag: strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEmit, Emit: emit}
}

func EmitStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY(emit []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure)) strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY {
	return strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY{Tag: strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEmit, Emit: emit}
}

func Open(open0 int, open1 []scope_internal.Registered) scope_internal.Finalizers {
	return scope_internal.Finalizers{Tag: scope_internal.FinalizersOpen, Open0: open0, Open1: open1}
}

//line ../../../../../../src/Draw/Server.purs:72:11
func NewRoom() ioref.IORef[internal.Canvas] {
//line ../../../../../../src/Draw/Server.purs:72:11
	return ioref.New(internal.Canvas{Cells: internal.SharedMapTupleIntIntLeaf, Outboxes: internal.SharedMapIntLeaf, NextId: 0})
}

//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Om/Handler.purs:123:108
func Draw(dependencies internal.Room, procedure string, server2 server.Server) {
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Om/Handler.purs:123:108
	methodRequest := protobuf.MessageCodec(proxy.ProxyProxy)
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Om/Handler.purs:123:147
	methodResponse := protobuf.MessageCodec(proxy.ProxyProxy)
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:187:3
	stepRecordXY := strom.StepXY{Tag: strom.StepXYEnd}
	stepTupleIntInt := strom.StepTupleIntInt{Tag: strom.StepTupleIntIntEnd}
	exitRecord := stromscope.Exit{Tag: stromscope.ExitInterrupted}
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:187:3
	server.BidiStreamImpl(procedure, func(call server.Call) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:91:12
		scope := NewScope(server.ContextImpl(call))
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		either3, err7 := internalffi.CatchResultsImpl(func(x exception.Error) (variant.EitherUnit, either.ErrorFailure) {
			return variant.EitherUnit{}, either.ErrorFailure{Failed: true, Left: x}
		}, func() (variant.EitherUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:212:11
			ioRef := ioref.New(false)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:525:13
			ioRef2 := ioref.New(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY{First: []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure){}, Second: 0})
			ioRef3 := ioref.New((func(internal.Room) (strom.StepXY, variant.ExceptionFailure))(nil))
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			either2 := variant.EitherUnitFromResults(FoldPulledStepFF(methodResponse, call, scope, func(r2 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				return DrawPullInnerOrNext(scope, func(r internal.Room) (strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:214:20
					boolean := ioRef.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
					if boolean {
						return strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY{Tag: strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEnd}, variant.ExceptionFailure{}
					}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:217:7
					ioRef.Store(true)
					scrutinee := RequestsRoom{Room: r.Room, Requests: func(_ scope_internal.Scope, _ internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure) {
						return func(_ internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:215:21
							boolean2 := server.ReceiveImpl(call)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							if boolean2 {
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:225:18
								bytes2 := server.ReceivedImpl(call)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
								xy, err := methodRequest.Read(bytes2, 0, bytes.LengthImpl(bytes2))
								if err.Failed {
									problem := err.Left
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:229:40
									var explanation string
									switch problem.Tag {
									case protobuferror.DecodeErrorTruncated:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:24:27
										explanation = "(Truncated " + strconv.Itoa(problem.Field0) + ")"
									case protobuferror.DecodeErrorMalformedVarint:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:25:33
										explanation = "(MalformedVarint " + strconv.Itoa(problem.Field0) + ")"
									case protobuferror.DecodeErrorInvalidWireType:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:26:33
										explanation = "(InvalidWireType " + strconv.Itoa(problem.Field0) + ")"
									default: // protobuferror.DecodeErrorInvalidUtf8
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:27:29
										explanation = "(InvalidUtf8 " + strconv.Itoa(problem.Field0) + ")"
									}
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:251:38
									server.FailImpl(call, 3, explanation)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
									var result string
									switch problem.Tag {
									case protobuferror.DecodeErrorTruncated:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:24:27
										result = "(Truncated " + strconv.Itoa(problem.Field0) + ")"
									case protobuferror.DecodeErrorMalformedVarint:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:25:33
										result = "(MalformedVarint " + strconv.Itoa(problem.Field0) + ")"
									case protobuferror.DecodeErrorInvalidWireType:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:26:33
										result = "(InvalidWireType " + strconv.Itoa(problem.Field0) + ")"
									default: // protobuferror.DecodeErrorInvalidUtf8
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Error.purs:27:29
										result = "(InvalidUtf8 " + strconv.Itoa(problem.Field0) + ")"
									}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
									return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: variant.Exception{Exception: exception.NewError("a request doesn't decode: " + result)}}
								}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
								return strom.EmitStepXY([]shared.XY{xy}), variant.ExceptionFailure{}
							} else {
//line ../../../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:218:20
								boolean3 := context.IsCancelledImpl(server.ContextImpl(call))
								boolean4 := server.FailedImpl(call)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
								if boolean3 {
									return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: stromscope.Interrupted}
								}
								if boolean4 {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
									return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: variant.Exception{Exception: exception.NewError("the client's stream broke off")}}
								}
								return stepRecordXY, variant.ExceptionFailure{}
							}
						}, variant.ExceptionFailure{}
					}}
					room := scrutinee.Room
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:217:7
					return EmitStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY([]func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure){func(scope2 scope_internal.Scope, r3 internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure) {
//line ../../../../../../src/Draw/Server.purs:55:13
						outbox := chanffi.NewBufferedChanImpl[tuple.TupleInt](1024)
						you := ioref.AtomicModifyIORefImpl(internal.FstYou, internal.SndYou, func(canvas internal.Canvas) internal.TupleYou {
							updated := canvas
							k := canvas.NextId
//line ../../../../../../src/Draw/Server.purs:57:5
							updated.Outboxes = DrawGo(outbox, k, canvas.Outboxes)
							updated.NextId = canvas.NextId + 1
//line ../../../../../../src/Draw/Server.purs:57:5
							return internal.TupleYou{First: updated, Second: tuple.You{Id: canvas.NextId, Outbox: outbox, Before: array.Unfoldr(func(value types.MaybeTupleTupleIntIntListTupleIntInt) bool {
								return value.Tag == maybe.TagNothing
							}, func(value types.MaybeTupleTupleIntIntListTupleIntInt) types.Tuple {
								return value.Just
							}, func(value types.Tuple) tuple.TupleInt {
								return value.First
							}, func(value types.Tuple) *types.ListTupleIntInt {
								return value.Second
							}, func(xs *types.ListTupleIntInt) types.MaybeTupleTupleIntIntListTupleIntInt {
//line ../../../../../p/lists-7.0.0/src/Data/List.purs:122:32
								if xs == types.SharedListTupleIntIntNil {
									return types.NoTupleTupleIntIntListTupleIntInt
								}
//line ../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:33:21
								return types.JustTupleTupleIntIntListTupleIntInt(types.Tuple{First: xs.Cons0, Second: xs.Cons1})
							}, set.DrawGo(canvas.Cells, types.SharedListTupleIntIntNil))}}
						}, room)
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						owned, err := ForkScope(scope2, r3)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						if err.Failed {
							return nil, err
						}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						_, err = AddFinalizer(owned, func(_ stromscope.Exit, _ internal.Room) variant.EitherUnit {
//line ../../../../../../src/Draw/Server.purs:39:59
							ioref.AtomicModifyIORefImpl(internal.FstUnit, func(tuple internal.TupleUnit) struct{} {
//line ../../../../../../src/Draw/Server.purs:39:59
								internal.SndUnit(tuple)
								return struct{}{}
							}, func(canvas internal.Canvas) internal.TupleUnit {
//line ../../../../../../src/Draw/Server.purs:61:53
								updated := canvas
								k := you.Id
//line ../../../../../../src/Draw/Server.purs:61:53
								updated.Outboxes = DrawGo2(k, canvas.Outboxes)
								return internal.TupleUnit{First: updated, Second: struct{}{}}
							}, room)
//line ../../../../../../src/Draw/Server.purs:39:59
							return variant.RightEitherUnit(struct{}{})
						}, r3)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						if err.Failed {
							return nil, err
						}
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
						chunks := [][]tuple.TupleInt{you.Before}
						chunks2 := [][]func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure){{func(_ scope_internal.Scope, _ internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:203:16
							ioRef4 := ioref.New(0)
							return func(_ internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure) {
								i := ioRef4.Load()
								ioRef4.Store(i + 1)
								var result strom.StepTupleIntInt
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:207:11
								if i >= 0 && i < len(chunks) {
//line ../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
									result = strom.EmitStepTupleIntInt(chunks[i])
								} else {
									result = stepTupleIntInt
								}
								return result, variant.ExceptionFailure{}
							}, variant.ExceptionFailure{}
						}, func(scope3 scope_internal.Scope, _ internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure) {
							x := func() maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:225:15
								return selectffi.Select2Impl(selectffi.OnReceiveImpl(chanffi.ReceiveEndImpl(you.Outbox), func(value tuple.TupleInt, open bool) maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
									var result maybe_tuple.MaybeTupleIntInt
									if open {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:62
										result = maybe_tuple.JustTupleIntInt(value)
									} else {
										result = maybe_tuple.NoTupleIntInt
									}
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
									return maybe_tuple.JustMaybeTupleIntInt(result)
								}), selectffi.OnReceiveImpl(context.DoneImpl(scope3.Context), func(_ struct{}, _ bool) maybe_tuple.MaybeMaybeTupleIntInt {
									return maybe_tuple.NoMaybeTupleIntInt
								}))
							}
							return func(_ internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
								maybe2 := x()
								if maybe2.Tag == maybe.TagNothing {
									return strom.StepTupleIntInt{}, variant.ExceptionFailure{Failed: true, Left: stromscope.Interrupted}
								}
								if maybe2.Tag == maybe.TagJust && maybe2.Just.Tag == maybe.TagNothing {
									return stepTupleIntInt, variant.ExceptionFailure{}
								}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
								return strom.EmitStepTupleIntInt([]tuple.TupleInt{maybe2.Just.Just}), variant.ExceptionFailure{}
							}, variant.ExceptionFailure{}
						}}}
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
						f2, err := MergeStreams(0, strom.HaltStrategyHaltLeft, []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure){func(scope3 scope_internal.Scope, r4 internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
							f, err2 := scrutinee.Requests(scope3, r4)
							if err2.Failed {
								return nil, err2
							}
							return func(room2 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
								step, err3 := f(room2)
								if err3.Failed {
									return strom.StepXY{}, err3
								}
								switch step.Tag {
								case strom.StepXYEmit:
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
									array_, err4 := DrawFrom(room, step.Emit, 0, []struct{}{})
									var step2 strom.StepUnit
									var err5 variant.ExceptionFailure
									if err4.Failed {
										err5 = err4
									} else {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
										step2, err5 = strom.EmitStepUnit(array_), variant.ExceptionFailure{}
									}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:9
									var step3 strom.StepXY
									var err6 variant.ExceptionFailure
									if err5.Failed {
										err6 = err5
									} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:387:3
										var result strom.StepXY
										switch step2.Tag {
										case strom.StepUnitEmit:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:388:19
											bound := make([]shared.XY, 0, 0*len(step2.Emit))
											for range step2.Emit {
											}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:388:19
											result = strom.EmitStepXY(bound)
										default: // strom.StepUnitEnd
											result = stepRecordXY
										}
										step3, err6 = result, variant.ExceptionFailure{}
									}
									return step3, err6
								default: // strom.StepXYEnd
									return stepRecordXY, variant.ExceptionFailure{}
								}
							}, variant.ExceptionFailure{}
						}, func(scope3 scope_internal.Scope, _ internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:203:16
							ioRef4 := ioref.New(0)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:525:13
							ioRef5 := ioref.New(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt{First: []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure){}, Second: 0})
							ioRef6 := ioref.New((func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure))(nil))
							field := func(r4 internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:526:14
								return DrawPullInnerOrNext2(scope3, func(_ internal.Room) (strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:205:14
									i := ioRef4.Load()
									ioRef4.Store(i + 1)
									var result strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:207:11
									if i >= 0 && i < len(chunks2) {
//line ../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
										result = EmitStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt(chunks2[i])
									} else {
										result = strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt{Tag: strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEnd}
									}
									return result, variant.ExceptionFailure{}
								}, ioRef5, ioRef6, r4)
							}
							return func(room2 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
								step, err2 := field(room2)
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:9
								var step2 strom.StepXY
								var err3 variant.ExceptionFailure
								if err2.Failed {
									err3 = err2
								} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:387:3
									var result strom.StepXY
									switch step.Tag {
									case strom.StepTupleIntIntEmit:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:388:19
										mapped := make([]shared.XY, len(step.Emit))
										for index, tuple := range step.Emit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:388:19
											mapped[index] = shared.XY{X: tuple.First, Y: tuple.Second}
										}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:388:19
										result = strom.EmitStepXY(mapped)
									default: // strom.StepTupleIntIntEnd
										result = stepRecordXY
									}
									step2, err3 = result, variant.ExceptionFailure{}
								}
								return step2, err3
							}, variant.ExceptionFailure{}
						}}, owned, r3)
						if err.Failed {
							failure := err.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
							boolean2 := context.IsCancelledImpl(owned.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
							var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
							if boolean2 || exception.Message(failure.Exception) == "the stream was interrupted" {
								result = exitRecord
							} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
								result = stromscope.Failed(failure)
							}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
							_, err2 := CloseScope(result, owned, r3)
							if err2.Failed {
								return nil, err2
							}
							return nil, variant.ExceptionFailure{Failed: true, Left: failure}
						}
						return func(r4 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
							step, err2 := f2(r4)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							return DrawPullInnerOrNextJoin(owned, r4, step, err2)
						}, variant.ExceptionFailure{}
					}}), variant.ExceptionFailure{}
				}, ioRef2, ioRef3, r2)
			}, struct{}{}, dependencies))
			return either2, either.ErrorFailure{}
		})
		var field2 variant.EitherUnit
		if err7.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			field2 = variant.LeftEitherUnit(variant.Exception{Exception: err7.Left})
		} else {
			field2 = either3
		}
		switch field2.Tag {
		case variant.EitherUnitRight:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (variant.EitherUnit, either.ErrorFailure) {
				return variant.EitherUnit{}, either.ErrorFailure{Failed: true, Left: x}
			}, func() (variant.EitherUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:22
				either2 := variant.EitherUnitFromResults(CloseScope(stromscope.Exit{Tag: stromscope.ExitEnded}, scope, dependencies))
				return either2, either.ErrorFailure{}
			})
			switch field2.Tag {
			case variant.EitherUnitLeft:
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				server.FailImpl(call, 13, exception.Message(field2.Left.Exception))
			default: // variant.EitherUnitRight
				var result variant.EitherUnit
				if err.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
					result = variant.LeftEitherUnit(variant.Exception{Exception: err.Left})
				} else {
					result = either4
				}
				switch result.Tag {
				case variant.EitherUnitLeft:
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
					server.FailImpl(call, 13, exception.Message(result.Left.Exception))
				default: // variant.EitherUnitRight
				}
			}
		default: // variant.EitherUnitLeft
			failure := field2.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
			boolean := context.IsCancelledImpl(scope.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (variant.EitherUnit, either.ErrorFailure) {
				return variant.EitherUnit{}, either.ErrorFailure{Failed: true, Left: x}
			}, func() (variant.EitherUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
					result = exitRecord
				} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
					result = stromscope.Failed(failure)
				}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				either2 := variant.EitherUnitFromResults(CloseScope(result, scope, dependencies))
				return either2, either.ErrorFailure{}
			})
			switch field2.Tag {
			case variant.EitherUnitLeft:
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				server.FailImpl(call, 13, exception.Message(field2.Left.Exception))
			default: // variant.EitherUnitRight
				var result variant.EitherUnit
				if err.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
					result = variant.LeftEitherUnit(variant.Exception{Exception: err.Left})
				} else {
					result = either4
				}
				switch result.Tag {
				case variant.EitherUnitLeft:
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
					server.FailImpl(call, 13, exception.Message(result.Left.Exception))
				default: // variant.EitherUnitRight
				}
			}
		}
	}, server2)
}

//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:809:14
func MergeStreams(capacity int, strategy strom.HaltStrategy, streams []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure), scope scope_internal.Scope, r internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:809:14
	c := chanffi.NewBufferedChanImpl[strom.Signal](capacity)
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
	mapped := make([]func(internal.Room) (scope_internal.ScopeStopped, variant.ExceptionFailure), len(streams))
	exitRecord := stromscope.Exit{Tag: stromscope.ExitInterrupted}
	for slot, stream := range streams {
		source := slot
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		mapped[slot] = func(r2 internal.Room) (scope_internal.ScopeStopped, variant.ExceptionFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			producerScope, err := ForkScope(scope, r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return scope_internal.ScopeStopped{}, err
			}
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/MVar.purs:35:19
			mVar := mvar.NewEmptyMVarImpl[struct{}]()
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			_, err = AddFinalizer(scope, func(_ stromscope.Exit, _ internal.Room) variant.EitherUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:769:33
				mvar.ReadMVarImpl(mVar)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:769:33
				return variant.RightEitherUnit(struct{}{})
			}, r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return scope_internal.ScopeStopped{}, err
			}
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:52:43
			ioffi.ForkIOImpl(func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:771:16
				io.Finally(func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
					either3, err3 := internalffi.CatchResultsImpl(func(x exception.Error) (variant.EitherBoolean, either.ErrorFailure) {
						return variant.EitherBoolean{}, either.ErrorFailure{Failed: true, Left: x}
					}, func() (variant.EitherBoolean, either.ErrorFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						f, err2 := stream(producerScope, r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						var either2 variant.EitherBoolean
						if err2.Failed {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							either2 = variant.LeftEitherBoolean(err2.Left)
						} else {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
							either2 = variant.EitherBooleanFromResults(MergeStreamsHandOverAll(source, c, producerScope, f, r2))
						}
						return either2, either.ErrorFailure{}
					})
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
					var result variant.EitherBoolean
					if err3.Failed {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						result = variant.LeftEitherBoolean(variant.Exception{Exception: err3.Left})
					} else {
						result = either3
					}
					if result.Tag == variant.EitherBooleanRight && result.Right {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
						_, err2 := MergeStreamsHandOver(c, producerScope, strom.Finished(source), r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						if !err2.Failed {
							CloseScope(stromscope.Exit{Tag: stromscope.ExitEnded}, producerScope, r2)
						}
					} else if result.Tag == variant.EitherBooleanRight && !result.Right {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						CloseScope(exitRecord, producerScope, r2)
					} else {
						failure := result.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
						boolean := context.IsCancelledImpl(producerScope.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
						var field stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
						if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
							field = exitRecord
						} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
							field = stromscope.Failed(failure)
						}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
						switch field.Tag {
						case stromscope.ExitInterrupted:
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							CloseScope(field, producerScope, r2)
						default:
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							_, err2 := MergeStreamsHandOver(c, producerScope, strom.FailedWith(source, failure), r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
							if !err2.Failed {
								CloseScope(field, producerScope, r2)
							}
						}
					}
				}, func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:771:73
					mvar.PutMVarImpl(mVar, struct{}{})
				})
			})
			return scope_internal.ScopeStopped{Scope: producerScope, Stopped: mVar}, variant.ExceptionFailure{}
		}
	}
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
	array_, err4 := MergeStreamsFrom(streams, r, mapped, 0, []scope_internal.ScopeStopped{})
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	if err4.Failed {
		return nil, err4
	}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:811:14
	ioRef := ioref.New(len(streams))
	return func(r2 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:811:14
		return MergeStreamsReceiveNext(strategy, scope, c, array_, ioRef, r2)
	}, variant.ExceptionFailure{}
}

//line server.go:627
func ForkScope(parent scope_internal.Scope, r internal.Room) (scope_internal.Scope, variant.ExceptionFailure) {
	parentContext := parent.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:109:31
	tuple2 := context.WithCancelImpl(func(context2 context.Context) func(context.CancelContext) tuple.TupleContext {
		return func(cancelContext context.CancelContext) tuple.TupleContext {
			return tuple.TupleContext{First: context2, Second: cancelContext}
		}
	}, parentContext)
	finalizers := ioref.New(Open(0, []scope_internal.Registered{}))
	detachFromParent := ioref.New(func() {
	})
	scope := scope_internal.Scope{Context: tuple2.First, CancelContext: tuple2.Second, Finalizers: finalizers, DetachFromParent: detachFromParent}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	i, err := AddFinalizer(parent, func(_ stromscope.Exit, room internal.Room) variant.EitherUnit {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		return variant.EitherUnitFromResults(CloseScope(stromscope.Exit{Tag: stromscope.ExitInterrupted}, scope, room))
	}, r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	if err.Failed {
		return scope_internal.Scope{}, err
	}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:134:3
	scope.DetachFromParent.Store(func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
		ioref.AtomicModifyIORefImpl(FstUnit, func(tuple TupleUnit) struct{} {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
			SndUnit(tuple)
			return struct{}{}
		}, func(state scope_internal.Finalizers) TupleUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:87
			var result scope_internal.Finalizers
			switch state.Tag {
			case scope_internal.FinalizersOpen:
				kept := []scope_internal.Registered{}
				for _, x := range state.Open1 {
					if x.First != i {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
						kept = append(kept, x)
					}
				}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
				result = Open(state.Open0, kept)
			default: // scope_internal.FinalizersClosed
				result = scope_internal.Finalizers{Tag: scope_internal.FinalizersClosed}
			}
			return TupleUnit{First: result, Second: struct{}{}}
		}, parent.Finalizers)
	})
	return scope, variant.ExceptionFailure{}
}

//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
func AddFinalizer(scope scope_internal.Scope, finalize func(stromscope.Exit, internal.Room) variant.EitherUnit, r internal.Room) (int, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
	maybe2 := ioref.AtomicModifyIORefImpl(FstMaybeFinalizerKey, SndMaybeFinalizerKey, func(value scope_internal.Finalizers) TupleMaybeInt {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:145:3
		switch value.Tag {
		case scope_internal.FinalizersOpen:
			key := value.Open0
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:146:28
			return TupleMaybeInt{First: Open(key+1, array.Snoc(value.Open1, scope_internal.Registered{First: key, Second: finalize})), Second: maybe.JustInt(key)}
		default: // scope_internal.FinalizersClosed
			return TupleMaybeInt{First: scope_internal.Finalizers{Tag: scope_internal.FinalizersClosed}, Second: maybe.NoInt}
		}
	}, scope.Finalizers)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch maybe2.Tag {
	case maybe.TagJust:
		return maybe2.Just, variant.ExceptionFailure{}
	default: // maybe.TagNothing
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		either := finalize(stromscope.Exit{Tag: stromscope.ExitInterrupted}, r)
		var i int
		var err variant.ExceptionFailure
		switch either.Tag {
		case variant.EitherUnitLeft:
			err = variant.ExceptionFailure{Failed: true, Left: either.Left}
		default: // variant.EitherUnitRight
			i, err = -1, variant.ExceptionFailure{}
		}
		return i, err
	}
}

//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:164:12
func CloseScope(exit stromscope.Exit, scope scope_internal.Scope, r internal.Room) (struct{}, variant.ExceptionFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:164:12
	finalizers := ioref.AtomicModifyIORefImpl(FstFinalizersRoomNoRequest, SndFinalizersRoomNoRequest, func(state scope_internal.Finalizers) TupleFinalizersRoomNoRequest {
		return TupleFinalizersRoomNoRequest{First: scope_internal.Finalizers{Tag: scope_internal.FinalizersClosed}, Second: state}
	}, scope.Finalizers)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch finalizers.Tag {
	case scope_internal.FinalizersClosed:
		return struct{}{}, variant.ExceptionFailure{}
	default: // scope_internal.FinalizersOpen
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:168:7
		context.CancelImpl(scope.CancelContext)
		effect := scope.DetachFromParent.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		effect()
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:171:24
		elements := array.Reverse(finalizers.Open1)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:171:24
		array_, err := CloseScopeFrom(exit, r, elements, 0, []variant.EitherUnit{})
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return struct{}{}, err
		}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
		found := variant.NoEitherVariantExceptionUnit
		for _, either := range array_ {
			var result bool
			switch either.Tag {
			case variant.EitherUnitLeft:
				result = true
			default: // variant.EitherUnitRight
				result = false
			}
			if result {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
				found = variant.JustEitherVariantExceptionUnit(either)
				break
			}
		}
		if found.Tag == maybe.TagJust && found.Just.Tag == variant.EitherUnitLeft {
			return struct{}{}, variant.ExceptionFailure{Failed: true, Left: found.Just.Left}
		}
		return struct{}{}, variant.ExceptionFailure{}
	}
}

//line server.go:759
func ForkScope2(parent scope_internal.Scope, r internal.Room) (scope_internal.Scope, variant.ExceptionFailure) {
	x := func() scope_internal.Scope {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:121:12
		return NewScope(parent.Context)
	}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	scope := x()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	detachFromParent := scope.DetachFromParent
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
	finalizersRecordRoomRecord := scope_internal.Finalizers{Tag: scope_internal.FinalizersClosed}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:140:12
	maybe2 := ioref.AtomicModifyIORefImpl(FstMaybeFinalizerKey, SndMaybeFinalizerKey, func(value scope_internal.Finalizers) TupleMaybeInt {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:145:3
		switch value.Tag {
		case scope_internal.FinalizersOpen:
			key := value.Open0
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:146:28
			return TupleMaybeInt{First: Open(key+1, array.Snoc(value.Open1, scope_internal.Registered{First: key, Second: func(ignored stromscope.Exit, room internal.Room) variant.EitherUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:146:28
				return variant.EitherUnitFromResults(ForkScopeFinalize(scope, ignored, room))
			}})), Second: maybe.JustInt(key)}
		default: // scope_internal.FinalizersClosed
			return TupleMaybeInt{First: finalizersRecordRoomRecord, Second: maybe.NoInt}
		}
	}, parent.Finalizers)
	switch maybe2.Tag {
	case maybe.TagJust:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:134:3
		detachFromParent.Store(func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
			ioref.AtomicModifyIORefImpl(FstUnit, func(tuple TupleUnit) struct{} {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
				SndUnit(tuple)
				return struct{}{}
			}, func(state scope_internal.Finalizers) TupleUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:87
				var result scope_internal.Finalizers
				switch state.Tag {
				case scope_internal.FinalizersOpen:
					kept := []scope_internal.Registered{}
					for _, x2 := range state.Open1 {
						if x2.First != maybe2.Just {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
							kept = append(kept, x2)
						}
					}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
					result = Open(state.Open0, kept)
				default: // scope_internal.FinalizersClosed
					result = finalizersRecordRoomRecord
				}
				return TupleUnit{First: result, Second: struct{}{}}
			}, parent.Finalizers)
		})
		return scope, variant.ExceptionFailure{}
	default: // maybe.TagNothing
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		_, err := ForkScopeFinalize(scope, stromscope.Exit{Tag: stromscope.ExitInterrupted}, r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return scope_internal.Scope{}, err
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:134:3
		detachFromParent.Store(func() {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
			ioref.AtomicModifyIORefImpl(FstUnit, func(tuple TupleUnit) struct{} {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:46
				SndUnit(tuple)
				return struct{}{}
			}, func(state scope_internal.Finalizers) TupleUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:152:87
				var result scope_internal.Finalizers
				switch state.Tag {
				case scope_internal.FinalizersOpen:
					kept := []scope_internal.Registered{}
					for _, x2 := range state.Open1 {
						if x2.First != -1 {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
							kept = append(kept, x2)
						}
					}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:155:32
					result = Open(state.Open0, kept)
				default: // scope_internal.FinalizersClosed
					result = finalizersRecordRoomRecord
				}
				return TupleUnit{First: result, Second: struct{}{}}
			}, parent.Finalizers)
		})
		return scope, variant.ExceptionFailure{}
	}
}

//line server.go:854
func FstMaybeFinalizerKey(tuple TupleMaybeInt) scope_internal.Finalizers {
	return tuple.First
}

func SndMaybeFinalizerKey(tuple TupleMaybeInt) maybe.Int {
	return tuple.Second
}

//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:109:31
func NewScope(parentContext context.Context) scope_internal.Scope {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:109:31
	tuple2 := context.WithCancelImpl(func(context2 context.Context) func(context.CancelContext) tuple.TupleContext {
		return func(cancelContext context.CancelContext) tuple.TupleContext {
			return tuple.TupleContext{First: context2, Second: cancelContext}
		}
	}, parentContext)
	finalizers := ioref.New(Open(0, []scope_internal.Registered{}))
	detachFromParent := ioref.New(func() {
	})
	return scope_internal.Scope{Context: tuple2.First, CancelContext: tuple2.Second, Finalizers: finalizers, DetachFromParent: detachFromParent}
}

//line server.go:877
func FstFinalizersRoomNoRequest(tuple TupleFinalizersRoomNoRequest) scope_internal.Finalizers {
	return tuple.First
}

func SndFinalizersRoomNoRequest(tuple TupleFinalizersRoomNoRequest) scope_internal.Finalizers {
	return tuple.Second
}

func FstUnit(tuple TupleUnit) scope_internal.Finalizers {
	return tuple.First
}

func SndUnit(tuple TupleUnit) {
}

func For_StopProducer(b []scope_internal.ScopeStopped, room internal.Room) (struct{}, variant.ExceptionFailure) {
	return variant.EitherUnitResults(array.Foldr(func(x scope_internal.ScopeStopped, b2 func(internal.Room) variant.EitherUnit) func(internal.Room) variant.EitherUnit {
		return func(r internal.Room) variant.EitherUnit {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:181:39
			context.CancelImpl(x.Scope.CancelContext)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:802:3
			mvar.ReadMVarImpl(x.Stopped)
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:55:57
			either := b2(r)
			switch either.Tag {
			case variant.EitherUnitLeft:
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:55:57
				return variant.LeftEitherUnit(either.Left)
			default: // variant.EitherUnitRight
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:55:57
				return variant.RightEitherUnit(either.Right)
			}
		}
	}, func(_ internal.Room) variant.EitherUnit {
		return variant.RightEitherUnit(struct{}{})
	}, b)(room))
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func FoldPulledStepFF(codec codec.Codec, call server.Call, scope scope_internal.Scope, pull func(internal.Room) (strom.StepXY, variant.ExceptionFailure), accumulated struct{}, r internal.Room) (struct{}, variant.ExceptionFailure) {
	for {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		context2 := scope.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(context2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if boolean {
			return struct{}{}, variant.ExceptionFailure{Failed: true, Left: stromscope.Interrupted}
		}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err := pull(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return struct{}{}, err
		}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
		switch step.Tag {
		case strom.StepXYEnd:
			return accumulated, variant.ExceptionFailure{}
		default: // strom.StepXYEmit
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			unit, err2 := FoldPulledStepFFForEachFrom(step.Emit, codec, call, 0, r)
			if err2.Failed {
				return struct{}{}, err2
			}
			accumulated = unit
			continue
		}
	}
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func DrawPullInnerOrNext(scope scope_internal.Scope, pullOuter func(internal.Room) (strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY, variant.ExceptionFailure), queued ioref.IORef[TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY], current ioref.IORef[func(internal.Room) (strom.StepXY, variant.ExceptionFailure)], r internal.Room) (strom.StepXY, variant.ExceptionFailure) {
	for {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:530:14
		maybe := current.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if maybe != nil {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			step, err := maybe(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepXY{}, err
			}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			switch step.Tag {
			case strom.StepXYEnd:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:536:13
				current.Store((func(internal.Room) (strom.StepXY, variant.ExceptionFailure))(nil))
				continue
			default:
				return step, variant.ExceptionFailure{}
			}
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:540:27
		tuple := queued.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		chunk := tuple.First
		index := tuple.Second
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if index >= 0 && index < len(chunk) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:543:13
			queued.Store(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY{First: chunk, Second: index + 1})
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			b, err := ForkScope2(scope, r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepXY{}, err
			}
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			f, err := chunk[index](b, r)
			if err.Failed {
				failure := err.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
				boolean := context.IsCancelledImpl(b.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
					result = stromscope.Exit{Tag: stromscope.ExitInterrupted}
				} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
					result = stromscope.Failed(failure)
				}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				_, err2 := CloseScope(result, b, r)
				if err2.Failed {
					return strom.StepXY{}, err2
				}
				return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: failure}
			}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:546:13
			current.Store(func(r2 internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
				step, err2 := f(r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				return DrawPullInnerOrNextJoin(b, r2, step, err2)
			})
			continue
		} else {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			step, err := pullOuter(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepXY{}, err
			}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			switch step.Tag {
			case strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEnd:
				return strom.StepXY{Tag: strom.StepXYEnd}, variant.ExceptionFailure{}
			default: // strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEmit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:553:17
				queued.Store(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY{First: step.Emit, Second: 0})
				continue
			}
		}
	}
}

//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
func DrawPullInnerOrNextJoin(b scope_internal.Scope, r internal.Room, step strom.StepXY, err variant.ExceptionFailure) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
	if err.Failed {
		failure := err.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(b.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
			result = stromscope.Exit{Tag: stromscope.ExitInterrupted}
		} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
			result = stromscope.Failed(failure)
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		_, err2 := CloseScope(result, b, r)
		if err2.Failed {
			return strom.StepXY{}, err2
		}
		return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: failure}
	}
	if !err.Failed && step.Tag == strom.StepXYEnd {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		_, err2 := CloseScope(stromscope.Exit{Tag: stromscope.ExitEnded}, b, r)
		var step2 strom.StepXY
		var err3 variant.ExceptionFailure
		if err2.Failed {
			err3 = err2
		} else {
			step2, err3 = strom.StepXY{Tag: strom.StepXYEnd}, variant.ExceptionFailure{}
		}
		return step2, err3
	}
	return step, variant.ExceptionFailure{}
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
func DrawGo(outbox chanffi.Chan[tuple.TupleInt], k int, value *internal.MapInt) *internal.MapInt {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
	if value == internal.SharedMapIntLeaf {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
		return internal.NodeMapInt(1, 1, k, outbox, internal.SharedMapIntLeaf, internal.SharedMapIntLeaf)
	}
	mk := value.Node2
	mv := value.Node3
	ml := value.Node4
	mr := value.Node5
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
	if k < mk {
		return internal.UnsafeBalancedNodeInt(mk, mv, DrawGo(outbox, k, ml), mr)
	}
	if k == mk {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
		return internal.NodeMapInt(value.Node0, value.Node1, k, outbox, ml, mr)
	}
	return internal.UnsafeBalancedNodeInt(mk, mv, ml, DrawGo(outbox, k, mr))
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:495:3
func DrawGo2(k int, value *internal.MapInt) *internal.MapInt {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:495:3
	if value == internal.SharedMapIntLeaf {
		return internal.SharedMapIntLeaf
	}
	mk := value.Node2
	mv := value.Node3
	ml := value.Node4
	mr := value.Node5
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:498:7
	if k < mk {
		return internal.UnsafeBalancedNodeInt(mk, mv, DrawGo2(k, ml), mr)
	}
	if k == mk {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:501:15
		if ml == internal.SharedMapIntLeaf {
			return mr
		}
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:767:5
		scrutinee := internal.UnsafeSplitLast(ml.Node2, ml.Node3, ml.Node4, ml.Node5)
		return internal.UnsafeBalancedNodeInt(scrutinee.First, scrutinee.Second, scrutinee.Third, mr)
	}
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:500:15
	return internal.UnsafeBalancedNodeInt(mk, mv, ml, DrawGo2(k, mr))
}

//line server.go:1124
func DrawFrom(room ioref.IORef[internal.Canvas], chunk []shared.XY, index int, field0 []struct{}) ([]struct{}, variant.ExceptionFailure) {
	field0 = slices.Clip(field0)
	for {
		if index < len(chunk) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:404:36
			scrutinee := chunk[index]
			x := scrutinee.X
			y := scrutinee.Y
//line ../../../../../../src/Draw/Server.purs:48:13
			canvas := ioref.AtomicModifyIORefImpl(internal.FstCanvas, internal.SndCanvas, func(canvas2 internal.Canvas) internal.TupleCanvas {
				updated := canvas2
//line ../../../../../../src/Draw/Server.purs:49:19
				updated.Cells = set.InsertGo(tuple.TupleInt{First: x, Second: y}, canvas2.Cells)
				return internal.TupleCanvas{First: updated, Second: updated}
			}, room)
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:167:21
			DrawFromGo(x, y, canvas.Outboxes, func() {
			})
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:167:21
			index, field0 = index+1, append(field0, struct{}{})
			continue
		}
		return field0, variant.ExceptionFailure{}
	}
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:169:26
func DrawFromGo(x int, y int, m_ *internal.MapInt, z_ func()) {
	for {
		xCaptured := x
		yCaptured := y
		z_Captured := z_
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:169:26
		if m_ == internal.SharedMapIntLeaf {
			z_Captured()
			return
		}
		v := m_.Node3
		l := m_.Node4
		r := m_.Node5
		x, y, m_, z_ = xCaptured, yCaptured, l, func() {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			chanffi.TrySendImpl(v, tuple.TupleInt{First: xCaptured, Second: yCaptured})
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:172:27
			DrawFromGo(xCaptured, yCaptured, r, z_Captured)
		}
	}
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func DrawPullInnerOrNext2(scope scope_internal.Scope, pullOuter func(internal.Room) (strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt, variant.ExceptionFailure), queued ioref.IORef[TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt], current ioref.IORef[func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure)], r internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure) {
	for {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:530:14
		maybe := current.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if maybe != nil {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			step, err := maybe(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepTupleIntInt{}, err
			}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			switch step.Tag {
			case strom.StepTupleIntIntEnd:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:536:13
				current.Store((func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure))(nil))
				continue
			default:
				return step, variant.ExceptionFailure{}
			}
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:540:27
		tuple := queued.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		chunk := tuple.First
		index := tuple.Second
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if index >= 0 && index < len(chunk) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:543:13
			queued.Store(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt{First: chunk, Second: index + 1})
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			b, err := ForkScope2(scope, r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepTupleIntInt{}, err
			}
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
			f, err := chunk[index](b, r)
			if err.Failed {
				failure := err.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
				boolean := context.IsCancelledImpl(b.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
					result = stromscope.Exit{Tag: stromscope.ExitInterrupted}
				} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
					result = stromscope.Failed(failure)
				}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
				_, err2 := CloseScope(result, b, r)
				if err2.Failed {
					return strom.StepTupleIntInt{}, err2
				}
				return strom.StepTupleIntInt{}, variant.ExceptionFailure{Failed: true, Left: failure}
			}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:546:13
			current.Store(func(r2 internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure) {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
				step, err2 := f(r2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				return DrawPullInnerOrNextJoin2(b, r2, step, err2)
			})
			continue
		} else {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			step, err := pullOuter(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			if err.Failed {
				return strom.StepTupleIntInt{}, err
			}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			switch step.Tag {
			case strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEnd:
				return strom.StepTupleIntInt{Tag: strom.StepTupleIntIntEnd}, variant.ExceptionFailure{}
			default: // strom_internal.StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEmit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:553:17
				queued.Store(TupleArrayFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt{First: step.Emit, Second: 0})
				continue
			}
		}
	}
}

//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
func DrawPullInnerOrNextJoin2(b scope_internal.Scope, r internal.Room, step strom.StepTupleIntInt, err variant.ExceptionFailure) (strom.StepTupleIntInt, variant.ExceptionFailure) {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
	if err.Failed {
		failure := err.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(b.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		var result stromscope.Exit
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		if boolean || exception.Message(failure.Exception) == "the stream was interrupted" {
			result = stromscope.Exit{Tag: stromscope.ExitInterrupted}
		} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
			result = stromscope.Failed(failure)
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:8
		_, err2 := CloseScope(result, b, r)
		if err2.Failed {
			return strom.StepTupleIntInt{}, err2
		}
		return strom.StepTupleIntInt{}, variant.ExceptionFailure{Failed: true, Left: failure}
	}
	if !err.Failed && step.Tag == strom.StepTupleIntIntEnd {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
		_, err2 := CloseScope(stromscope.Exit{Tag: stromscope.ExitEnded}, b, r)
		var step2 strom.StepTupleIntInt
		var err3 variant.ExceptionFailure
		if err2.Failed {
			err3 = err2
		} else {
			step2, err3 = strom.StepTupleIntInt{Tag: strom.StepTupleIntIntEnd}, variant.ExceptionFailure{}
		}
		return step2, err3
	}
	return step, variant.ExceptionFailure{}
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func MergeStreamsReceiveNext(strategy strom.HaltStrategy, scope scope_internal.Scope, handoff chanffi.Chan[strom.Signal], producers []scope_internal.ScopeStopped, running ioref.IORef[int], r internal.Room) (strom.StepXY, variant.ExceptionFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
	stepRecordXY := strom.StepXY{Tag: strom.StepXYEnd}
	for {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:815:21
		i := running.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if i <= 0 {
			return stepRecordXY, variant.ExceptionFailure{}
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:818:19
		context2 := scope.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:818:19
		maybe2 := selectffi.Select2Impl(selectffi.OnReceiveImpl(chanffi.ReceiveEndImpl(handoff), func(value strom.Signal, open bool) strom.MaybeSignalNoRequestXY {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
			var result strom.MaybeSignalNoRequestXY
			if open {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:62
				result = strom.JustSignalNoRequestXY(value)
			} else {
				result = strom.NoSignalNoRequestXY
			}
			return result
		}), selectffi.OnReceiveImpl(context.DoneImpl(context2), func(_ struct{}, _ bool) strom.MaybeSignalNoRequestXY {
			return strom.NoSignalNoRequestXY
		}))
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if maybe2.Tag == maybe.TagNothing {
			return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: stromscope.Interrupted}
		}
		if maybe2.Tag == maybe.TagJust && maybe2.Just.Tag == strom.SignalPulled {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			return strom.EmitStepXY(maybe2.Just.Field1), variant.ExceptionFailure{}
		}
		if maybe2.Tag == maybe.TagJust && maybe2.Just.Tag == strom.SignalFailedWith {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:823:11
			running.Store(0)
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			_, err := For_StopProducer(producers, r)
			if err.Failed {
				return strom.StepXY{}, err
			}
			return strom.StepXY{}, variant.ExceptionFailure{Failed: true, Left: maybe2.Just.FailedWith1}
		}
		source := maybe2.Just.Field0
		var result2 bool
		switch strategy {
		case strom.HaltStrategyHaltBoth:
			result2 = false
		case strom.HaltStrategyHaltEither:
			result2 = true
		case strom.HaltStrategyHaltLeft:
			result2 = source == 0
		default: // strom.HaltStrategyHaltRight
			result2 = source == 1
		}
		if result2 || i == 1 {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:828:15
			running.Store(0)
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			_, err := For_StopProducer(producers, r)
			if err.Failed {
				return strom.StepXY{}, err
			}
			return stepRecordXY, variant.ExceptionFailure{}
		}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:832:15
		running.Store(i - 1)
	}
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func MergeStreamsHandOverAll(source int, c chanffi.Chan[strom.Signal], producerScope scope_internal.Scope, pull func(internal.Room) (strom.StepXY, variant.ExceptionFailure), r internal.Room) (bool, variant.ExceptionFailure) {
	for {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err := pull(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return false, err
		}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
		switch step.Tag {
		case strom.StepXYEnd:
			return true, variant.ExceptionFailure{}
		default: // strom.StepXYEmit
			context2 := producerScope.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:39:16
			cancelled := context.IsCancelledImpl(context2)
			var boolean bool
			if cancelled {
				boolean = false
			} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:41:8
				boolean = selectffi.Select2Impl(selectffi.OnSendImpl(c, strom.Pulled(source, step.Emit), func() bool {
					return true
				}), selectffi.OnReceiveImpl(context.DoneImpl(context2), func(_ struct{}, _ bool) bool {
					return false
				}))
			}
			if boolean {
				continue
			}
			return false, variant.ExceptionFailure{}
		}
	}
}

//line server.go:1408
func MergeStreamsHandOver(c chanffi.Chan[strom.Signal], producerScope scope_internal.Scope, signal strom.Signal, room internal.Room) (struct{}, variant.ExceptionFailure) {
	context2 := producerScope.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Handoff.purs:39:16
	cancelled := context.IsCancelledImpl(context2)
	if !cancelled {
		selectffi.Select2Impl(selectffi.OnSendImpl(c, signal, func() bool {
			return true
		}), selectffi.OnReceiveImpl(context.DoneImpl(context2), func(_ struct{}, _ bool) bool {
			return false
		}))
	}
	return struct{}{}, variant.ExceptionFailure{}
}

//line server.go:1423
func MergeStreamsFrom(streams []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure), r internal.Room, elements []func(internal.Room) (scope_internal.ScopeStopped, variant.ExceptionFailure), index int, field0 []scope_internal.ScopeStopped) ([]scope_internal.ScopeStopped, variant.ExceptionFailure) {
	field0 = slices.Clip(field0)
	for {
		if index < len(streams) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:55:57
			scopeStopped, err := elements[index](r)
			if err.Failed {
				return nil, err
			}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:55:57
			index, field0 = index+1, append(field0, scopeStopped)
			continue
		} else {
			return field0, variant.ExceptionFailure{}
		}
	}
}

//line server.go:1442
func CloseScopeFrom(exit stromscope.Exit, r internal.Room, elements []scope_internal.Registered, index int, field0 []variant.EitherUnit) ([]variant.EitherUnit, variant.ExceptionFailure) {
	field0 = slices.Clip(field0)
	for {
		exitCaptured := exit
		rCaptured := r
		if index < len(elements) {
			scrutinee := elements[index]
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			either3, err := internalffi.CatchResultsImpl(func(x exception.Error) (variant.EitherUnit, either.ErrorFailure) {
				return variant.EitherUnit{}, either.ErrorFailure{Failed: true, Left: x}
			}, func() (variant.EitherUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:22
				either2 := scrutinee.Second(exitCaptured, rCaptured)
				return either2, either.ErrorFailure{}
			})
			var result variant.EitherUnit
			if err.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
				result = variant.LeftEitherUnit(variant.Exception{Exception: err.Left})
			} else {
				result = either3
			}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			exit, r, index, field0 = exitCaptured, rCaptured, index+1, append(field0, result)
			continue
		}
		return field0, variant.ExceptionFailure{}
	}
}

//line server.go:1473
func ForkScopeFinalize(child scope_internal.Scope, ignored stromscope.Exit, room internal.Room) (struct{}, variant.ExceptionFailure) {
	return CloseScope(stromscope.Exit{Tag: stromscope.ExitInterrupted}, child, room)
}

func FoldPulledStepFFForEachFrom(chunk []shared.XY, codec codec.Codec, call server.Call, index int, r internal.Room) (struct{}, variant.ExceptionFailure) {
	for {
		if index < len(chunk) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			response := chunk[index]
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:68:32
			bytes := FoldPulledStepFFForEachFromAttempt(response, codec, 0)
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:68:32
			boolean := server.SendImpl(call, bytes)
			if !boolean {
				return struct{}{}, variant.ExceptionFailure{Failed: true, Left: stromscope.Interrupted}
			}
			index += 1
			continue
		}
		return struct{}{}, variant.ExceptionFailure{}
	}
}

//line server.go:1497
const (
	FoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromRetryStep = iota
	FoldPulledStepFFForEachFromAttemptStep
)

func FoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromRetryOrFoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromAttempt(step int, response shared.XY, codec codec.Codec, writer2 writer.Writer, capacity int) bytes.Bytes {
	for {
		switch step {
		case FoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromRetryStep:
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:76:17
			tooSmall := writer.CapacityImpl(writer2)
			step, capacity = FoldPulledStepFFForEachFromAttemptStep, 2*tooSmall
		default: // FoldPulledStepFFForEachFromAttemptStep
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:71:15
			writer3 := writer.AcquireImpl(capacity)
			codec.Write(writer3, response)
			overflowed := writer.OverflowedImpl(writer3)
			if overflowed {
				step, writer2 = FoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromRetryStep, writer3
				continue
			}
//line ../../../../../p/yoga-protobuf/30c1fa1b8673b15445f93cddaf2217af43a7b20d/src/Yoga/Protobuf/Writer.purs:74:42
			return writer.FinishImpl(writer3)
		}
	}
}

//line server.go:1525
func FoldPulledStepFFForEachFromAttempt(response shared.XY, codec codec.Codec, capacity int) bytes.Bytes {
	var writer2 writer.Writer
	return FoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromRetryOrFoldPulledRecordRoomRecordRecordXYUnitStepFFForEachFromAttempt(FoldPulledStepFFForEachFromAttemptStep, response, codec, writer2, capacity)
}

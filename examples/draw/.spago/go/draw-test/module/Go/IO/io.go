package io

import (
	"github.com/i-am-the-slime/iris-go/exception"
	internal "iris.local/draw-test/Go_IO_Internal_ffi"
)

//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:139:13
func Finally(action func(), finalizer func()) {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:139:13
	internal.CatchImpl(func(error_ exception.Error) struct{} {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:145:53
		finalizer()
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:145:64
		panic(error_)
	}, func() struct{} {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:139:13
		action()
		return struct{}{}
	})
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:140:3
	finalizer()
}

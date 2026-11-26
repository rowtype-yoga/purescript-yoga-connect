package effectaff

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Either"
)

type Fiber struct {
	IsSuspended func() bool
	Join        func(func(either.EitherUnit)) func()
	Kill        func(exception.Error) func(func(either.EitherUnit)) func()
	OnComplete  func(either.OnComplete) func()
	Run         func()
}

func NonCanceler(ignored exception.Error) {
}

//line ../../../../../p/aff-8.0.0/src/Effect/Aff.purs:109:3
func ThrowError(e exception.Error) {
//line ../../../../../p/aff-8.0.0/src/Effect/Aff.purs:109:3
	panic(e)
}

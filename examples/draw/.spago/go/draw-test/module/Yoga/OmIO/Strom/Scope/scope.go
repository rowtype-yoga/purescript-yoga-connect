package scope

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Variant"
)

type Exit struct {
	Failed variant.Exception
	Tag    uint8
}

const (
	ExitEnded = iota
	ExitFailed
	ExitInterrupted
)

func Failed(failed variant.Exception) Exit {
	return Exit{Tag: ExitFailed, Failed: failed}
}

var Interrupted variant.Exception = variant.Exception{Exception: Interruption}

var Interruption = InitializeInterruption()

//line ../../../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
func InitializeInterruption() exception.Error {
//line ../../../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
	return exception.NewError("the stream was interrupted")
}

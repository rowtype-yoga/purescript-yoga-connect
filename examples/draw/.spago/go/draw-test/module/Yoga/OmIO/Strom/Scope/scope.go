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

//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
func InitializeInterruption() exception.Error {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
	return exception.NewError("the stream was interrupted")
}

package scope

import "github.com/i-am-the-slime/iris-go/exception"

var Interruption = InitializeInterruption()

//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
func InitializeInterruption() exception.Error {
//line ../../../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:209:16
	return exception.NewError("the stream was interrupted")
}

package config

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Test/Spec/Tree"
)

var DefaultConfig = InitializeDefaultConfig()

//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Config.purs:36:3
func InitializeDefaultConfig() tree.Config {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Config.purs:36:3
	return tree.Config{Slow: 75.0, Timeout: maybe.JustNumber(2000.0), Exit: true, FailFast: false, FilterTree: tree.Identity}
}

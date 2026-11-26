package key

import (
	key "iris.local/draw/Bubbletea_Key_ffi"
	"iris.local/draw/Data/Maybe"
)

type KeyCode struct {
	Char   rune
	Field0 int
	Tag    uint8
}

const (
	KeyCodeChar = iota
	KeyCodeEnter
	KeyCodeTab
	KeyCodeEsc
	KeyCodeSpace
	KeyCodeBackspace
	KeyCodeDelete
	KeyCodeUp
	KeyCodeDown
	KeyCodeLeft
	KeyCodeRight
	KeyCodeHome
	KeyCodeEnd
	KeyCodePgUp
	KeyCodePgDown
	KeyCodeInsert
	KeyCodeF
	KeyCodeOther
)

type Key struct {
	Keystroke string
	Text      string
	Code      KeyCode
	Alt       bool
	Ctrl      bool
	Hyper     bool
	Meta      bool
	Repeat    bool
	Shift     bool
	Super     bool
}

type MaybeKeyCode struct {
	Just KeyCode
	Tag  uint8
}

var NoKeyCode = MaybeKeyCode{Tag: maybe.TagNothing}

type Tuple struct {
	First  KeyCode
	Second int
}

func JustKeyCode(just KeyCode) MaybeKeyCode {
	return MaybeKeyCode{Tag: maybe.TagJust, Just: just}
}

func Char(char rune) KeyCode {
	return KeyCode{Tag: KeyCodeChar, Char: char}
}

func Other(field0 int) KeyCode {
	return KeyCode{Tag: KeyCodeOther, Field0: field0}
}

func F(field0 int) KeyCode {
	return KeyCode{Tag: KeyCodeF, Field0: field0}
}

var SpecialKeys = InitializeSpecialKeys()

//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
func InitializeSpecialKeys() []Tuple {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built := make([]Tuple, len(array)+15)
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[0] = Tuple{First: KeyCode{Tag: KeyCodeEnter}, Second: key.KeyEnter}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[1] = Tuple{First: KeyCode{Tag: KeyCodeTab}, Second: key.KeyTab}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[2] = Tuple{First: KeyCode{Tag: KeyCodeEsc}, Second: key.KeyEsc}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[3] = Tuple{First: KeyCode{Tag: KeyCodeSpace}, Second: key.KeySpace}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[4] = Tuple{First: KeyCode{Tag: KeyCodeBackspace}, Second: key.KeyBackspace}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[5] = Tuple{First: KeyCode{Tag: KeyCodeDelete}, Second: key.KeyDelete}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[6] = Tuple{First: KeyCode{Tag: KeyCodeUp}, Second: key.KeyUp}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[7] = Tuple{First: KeyCode{Tag: KeyCodeDown}, Second: key.KeyDown}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[8] = Tuple{First: KeyCode{Tag: KeyCodeLeft}, Second: key.KeyLeft}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[9] = Tuple{First: KeyCode{Tag: KeyCodeRight}, Second: key.KeyRight}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[10] = Tuple{First: KeyCode{Tag: KeyCodeHome}, Second: key.KeyHome}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[11] = Tuple{First: KeyCode{Tag: KeyCodeEnd}, Second: key.KeyEnd}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[12] = Tuple{First: KeyCode{Tag: KeyCodePgUp}, Second: key.KeyPgUp}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[13] = Tuple{First: KeyCode{Tag: KeyCodePgDown}, Second: key.KeyPgDown}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[14] = Tuple{First: KeyCode{Tag: KeyCodeInsert}, Second: key.KeyInsert}
	for index, n := range array {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
		built[index+15] = Tuple{First: F(n), Second: key.KeyF1 + n - 1}
	}
	return built
}

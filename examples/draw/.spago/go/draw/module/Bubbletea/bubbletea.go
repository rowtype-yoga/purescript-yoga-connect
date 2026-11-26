package bubbletea

import (
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/prim"
	"iris.local/draw/Bubbletea/Key"
	bubbletea "iris.local/draw/Bubbletea_ffi"
	datamaybe "iris.local/draw/Data/Maybe"
)

type MouseAction uint8

const (
	MouseActionClick MouseAction = iota
	MouseActionRelease
	MouseActionWheel
	MouseActionMotion
)

type MouseButton uint8

const (
	MouseButtonNoButton MouseButton = iota
	MouseButtonLeftButton
	MouseButtonMiddleButton
	MouseButtonRightButton
	MouseButtonWheelUp
	MouseButtonWheelDown
	MouseButtonWheelLeft
	MouseButtonWheelRight
	MouseButtonBackButton
	MouseButtonForwardButton
)

type MouseEvent struct {
	X      int
	Y      int
	Action MouseAction
	Alt    bool
	Button MouseButton
	Ctrl   bool
	Shift  bool
}

type Snapshot struct {
	Snapshot struct{}
}

type MaybeMouseButton struct {
	Tag  uint8
	Just MouseButton
}

var NoMouseButton = MaybeMouseButton{Tag: datamaybe.TagNothing}

type Tuple struct {
	Second int
	First  MouseButton
}

func JustMouseButton(just MouseButton) MaybeMouseButton {
	return MaybeMouseButton{Tag: datamaybe.TagJust, Just: just}
}

//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:227:19
func KeyOf(native bubbletea.NativeMsg) key.Key {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:227:19
	modifiers := bubbletea.KeyModifiersImpl(native)
	rune_ := bubbletea.KeyCodeImpl(native)
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:153:17
	findMap := direct.FindMap(key.NoKeyCode, func(maybe key.MaybeKeyCode) bool {
		return maybe.Tag == datamaybe.TagJust
	}, func(tuple key.Tuple) key.MaybeKeyCode {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:153:51
		if tuple.Second == rune_ {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:153:73
			return key.JustKeyCode(tuple.First)
		}
		return key.NoKeyCode
	}, key.SpecialKeys)
	var code key.KeyCode
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:153:17
	keyCode := key.Other(rune_)
	if findMap.Tag == datamaybe.TagJust {
		code = findMap.Just
	} else if findMap.Tag == datamaybe.TagNothing && (rune_ >= 33 && rune_ < 65536) {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:155:47
		if rune_ >= prim.ToCharCode(prim.BottomChar) && rune_ <= prim.ToCharCode(prim.TopChar) {
//line ../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
			code = key.Char(prim.FromCharCode(rune_))
		} else {
			code = keyCode
		}
	} else {
		code = keyCode
	}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:229:3
	return key.Key{Code: code, Text: bubbletea.KeyTextImpl(native), Keystroke: bubbletea.KeystrokeImpl(native), Ctrl: modifiers&bubbletea.ModCtrl != 0, Alt: modifiers&bubbletea.ModAlt != 0, Shift: modifiers&bubbletea.ModShift != 0, Meta: modifiers&bubbletea.ModMeta != 0, Super: modifiers&bubbletea.ModSuper != 0, Hyper: modifiers&bubbletea.ModHyper != 0, Repeat: bubbletea.KeyRepeatImpl(native)}
}

//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:264:19
func MouseOf(action MouseAction, native bubbletea.NativeMsg) MouseEvent {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:264:19
	modifiers := bubbletea.MouseModifiersImpl(native)
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:267:23
	code := bubbletea.MouseButtonImpl(native)
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:306:17
	findMap := direct.FindMap(NoMouseButton, func(maybe MaybeMouseButton) bool {
		return maybe.Tag == datamaybe.TagJust
	}, func(tuple Tuple) MaybeMouseButton {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:306:68
		if tuple.Second == code {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:306:90
			return JustMouseButton(tuple.First)
		}
		return NoMouseButton
	}, Buttons)
	var result MouseButton
	switch findMap.Tag {
	case datamaybe.TagNothing:
		result = MouseButtonNoButton
	default: // datamaybe.TagJust
		result = findMap.Just
	}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:306:17
	return MouseEvent{X: bubbletea.MouseXImpl(native), Y: bubbletea.MouseYImpl(native), Button: result, Action: action, Ctrl: modifiers&bubbletea.ModCtrl != 0, Alt: modifiers&bubbletea.ModAlt != 0, Shift: modifiers&bubbletea.ModShift != 0}
}

//line bubbletea.go:130
var Buttons []Tuple = []Tuple{{First: MouseButtonLeftButton, Second: bubbletea.MouseLeft}, {First: MouseButtonMiddleButton, Second: bubbletea.MouseMiddle}, {First: MouseButtonRightButton, Second: bubbletea.MouseRight}, {First: MouseButtonWheelUp, Second: bubbletea.MouseWheelUp}, {First: MouseButtonWheelDown, Second: bubbletea.MouseWheelDown}, {First: MouseButtonWheelLeft, Second: bubbletea.MouseWheelLeft}, {First: MouseButtonWheelRight, Second: bubbletea.MouseWheelRight}, {First: MouseButtonBackButton, Second: bubbletea.MouseBackward}, {First: MouseButtonForwardButton, Second: bubbletea.MouseForward}}

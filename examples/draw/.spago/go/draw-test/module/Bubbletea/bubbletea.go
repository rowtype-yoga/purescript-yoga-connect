package bubbletea

import (
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/prim"
	"iris.local/draw-test/Bubbletea/Key"
	bubbletea "iris.local/draw-test/Bubbletea_ffi"
	datamaybe "iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/shared"
)

type Snapshot struct {
	Snapshot struct{}
}

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

type Event struct {
	Message struct{}
	FieldA0 key.Key
	FieldB0 string
	Native  bubbletea.NativeMsg
	Mouse   MouseEvent
	Resize  shared.Size
	Tag     uint8
}

const (
	EventMessage = iota
	EventKeyPress
	EventKeyRelease
	EventMouse
	EventResize
	EventPaste
	EventClipboard
	EventFocus
	EventBlur
	EventNative
)

type Tuple struct {
	Second int
	First  MouseButton
}

type MaybeMouseButton struct {
	Tag  uint8
	Just MouseButton
}

var NoMouseButton = MaybeMouseButton{Tag: datamaybe.TagNothing}

func JustMouseButton(just MouseButton) MaybeMouseButton {
	return MaybeMouseButton{Tag: datamaybe.TagJust, Just: just}
}

func Mouse(mouse MouseEvent) Event {
	return Event{Tag: EventMouse, Mouse: mouse}
}

func ToNative(value Event) bubbletea.NativeMsg {
	switch value.Tag {
	case EventKeyPress:
		pressed := value.FieldA0
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:198:23
		var result int
		if pressed.Ctrl {
			result = bubbletea.ModCtrl
		} else {
			result = 0
		}
		var result2 int
		if pressed.Alt {
			result2 = bubbletea.ModAlt
		} else {
			result2 = 0
		}
		var result3 int
		if pressed.Shift {
			result3 = bubbletea.ModShift
		} else {
			result3 = 0
		}
		var result4 int
		if pressed.Meta {
			result4 = bubbletea.ModMeta
		} else {
			result4 = 0
		}
		var result5 int
		if pressed.Super {
			result5 = bubbletea.ModSuper
		} else {
			result5 = 0
		}
		var result6 int
		if pressed.Hyper {
			result6 = bubbletea.ModHyper
		} else {
			result6 = 0
		}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:198:23
		return bubbletea.KeyPressNative(key.ToRune(pressed.Code), pressed.Text, result|result2|result3|result4|result5|result6, pressed.Repeat)
	case EventKeyRelease:
		released := value.FieldA0
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:199:26
		var result int
		if released.Ctrl {
			result = bubbletea.ModCtrl
		} else {
			result = 0
		}
		var result2 int
		if released.Alt {
			result2 = bubbletea.ModAlt
		} else {
			result2 = 0
		}
		var result3 int
		if released.Shift {
			result3 = bubbletea.ModShift
		} else {
			result3 = 0
		}
		var result4 int
		if released.Meta {
			result4 = bubbletea.ModMeta
		} else {
			result4 = 0
		}
		var result5 int
		if released.Super {
			result5 = bubbletea.ModSuper
		} else {
			result5 = 0
		}
		var result6 int
		if released.Hyper {
			result6 = bubbletea.ModHyper
		} else {
			result6 = 0
		}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:199:26
		return bubbletea.KeyReleaseNative(key.ToRune(released.Code), released.Text, result|result2|result3|result4|result5|result6)
	case EventMouse:
		event := value.Mouse
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:200:18
		var result int
		switch event.Action {
		case MouseActionClick:
			result = bubbletea.KindMouseClick
		case MouseActionRelease:
			result = bubbletea.KindMouseRelease
		case MouseActionWheel:
			result = bubbletea.KindMouseWheel
		default: // MouseActionMotion
			result = bubbletea.KindMouseMotion
		}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:200:73
		findMap := direct.FindMap(datamaybe.NoInt, func(maybe datamaybe.Int) bool {
			return maybe.Tag == datamaybe.TagJust
		}, func(tuple Tuple) datamaybe.Int {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:309:71
			if tuple.First == event.Button {
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:309:95
				return datamaybe.JustInt(tuple.Second)
			}
			return datamaybe.NoInt
		}, Buttons)
		var result2 int
		switch findMap.Tag {
		case datamaybe.TagNothing:
			result2 = bubbletea.MouseNone
		default: // datamaybe.TagJust
			result2 = findMap.Just
		}
		var result3 int
		if event.Ctrl {
			result3 = bubbletea.ModCtrl
		} else {
			result3 = 0
		}
		var result4 int
		if event.Alt {
			result4 = bubbletea.ModAlt
		} else {
			result4 = 0
		}
		var result5 int
		if event.Shift {
			result5 = bubbletea.ModShift
		} else {
			result5 = 0
		}
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:200:73
		return bubbletea.MouseNative(result, event.X, event.Y, result2, result3|result4|result5)
	case EventResize:
		size := value.Resize
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:201:18
		return bubbletea.ResizeNative(size.Width, size.Height)
	case EventPaste:
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:202:17
		return bubbletea.PasteNative(value.FieldB0)
	case EventClipboard:
//line ../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:203:21
		return bubbletea.ClipboardNative(value.FieldB0)
	case EventFocus:
		return bubbletea.FocusNative
	case EventBlur:
		return bubbletea.BlurNative
	case EventMessage:
		return bubbletea.NoNative
	default: // EventNative
		return value.Native
	}
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

//line bubbletea.go:315
var Buttons []Tuple = []Tuple{{First: MouseButtonLeftButton, Second: bubbletea.MouseLeft}, {First: MouseButtonMiddleButton, Second: bubbletea.MouseMiddle}, {First: MouseButtonRightButton, Second: bubbletea.MouseRight}, {First: MouseButtonWheelUp, Second: bubbletea.MouseWheelUp}, {First: MouseButtonWheelDown, Second: bubbletea.MouseWheelDown}, {First: MouseButtonWheelLeft, Second: bubbletea.MouseWheelLeft}, {First: MouseButtonWheelRight, Second: bubbletea.MouseWheelRight}, {First: MouseButtonBackButton, Second: bubbletea.MouseBackward}, {First: MouseButtonForwardButton, Second: bubbletea.MouseForward}}

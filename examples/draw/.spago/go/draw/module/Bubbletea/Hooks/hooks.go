package hooks

import (
	tea "iris.local/draw/Bubbletea"
	"iris.local/draw/Bubbletea/Key"
	hooks "iris.local/draw/Bubbletea_Hooks_ffi"
	bubbletea "iris.local/draw/Bubbletea_ffi"
	"iris.local/draw/Data/Void"
	"iris.local/draw/shared"
)

type Event struct {
	Message ActionUnit
	FieldA0 key.Key
	FieldB0 string
	Native  bubbletea.NativeMsg
	Mouse   tea.MouseEvent
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

type Scope struct {
	Listeners hooks.Registry[func(Event) ActionUnit]
	Watchers  hooks.Registry[ActionUnit]
}

type Env struct {
	Outbox hooks.Outbox[bubbletea.Cmd[ActionUnit]]
	Scope  Scope
	Theme  shared.Theme
}

type ActionUnit struct {
	Action func(Env)
}

type Self struct {
	Emit  func(void.Void) ActionUnit
	Props ActionUnit
}

type ActionView struct {
	Action func(Env) bubbletea.ViewType
}

func Message(message ActionUnit) Event {
	return Event{Tag: EventMessage, Message: message}
}

func KeyPress(fieldA0 key.Key) Event {
	return Event{Tag: EventKeyPress, FieldA0: fieldA0}
}

func KeyRelease(fieldA0 key.Key) Event {
	return Event{Tag: EventKeyRelease, FieldA0: fieldA0}
}

func Mouse(mouse tea.MouseEvent) Event {
	return Event{Tag: EventMouse, Mouse: mouse}
}

func Resize(resize shared.Size) Event {
	return Event{Tag: EventResize, Resize: resize}
}

func Paste(fieldB0 string) Event {
	return Event{Tag: EventPaste, FieldB0: fieldB0}
}

func Clipboard(fieldB0 string) Event {
	return Event{Tag: EventClipboard, FieldB0: fieldB0}
}

func Native(native bubbletea.NativeMsg) Event {
	return Event{Tag: EventNative, Native: native}
}

var Quit ActionUnit = ActionUnit{Action: func(env Env) {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:458:30
	env.Outbox.Push(bubbletea.Quit[ActionUnit]())
}}

//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:435:18
func NewScope() Scope {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:435:18
	listeners := hooks.NewRegistry[func(Event) ActionUnit]()
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:435:18
	watchers := hooks.NewRegistry[ActionUnit]()
	return Scope{Listeners: listeners, Watchers: watchers}
}

//line hooks.go:106
func AbsurdSpin(void void.Void) ActionUnit {
	for {
	}
}

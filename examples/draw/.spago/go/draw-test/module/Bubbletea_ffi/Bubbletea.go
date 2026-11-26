package Bubbletea_ffi

import (
	"context"
	"image/color"
	"io"
	"os/exec"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

// Bubble Tea calls Init, Update and View on a goroutine of its own and runs
// every command on another. Update and View are PureScript functions it calls
// directly. Commands are go-io IO actions, which are made to run on any
// goroutine; their messages come back the way Bubble Tea's own do, as what the
// command returns. No event loop is involved.

type NativeMsg = tea.Msg

type NativeCmd = tea.Cmd

type ViewType = tea.View

type NativeRunning[S any] = *running[S]

// Cmd is a tree of commands. Leaves that produce a PureScript message take
// the function that turns the message into the program's own, so that `map`
// stays typed: the tree is converted to a tea.Cmd only once it reaches the
// program, where the message type is known. Leaves also get the program's
// context, done once the program has ended.
type Cmd[M any] struct {
	kind     cmdKind
	leaf     func(ctx context.Context, deliver func(M) tea.Msg) tea.Cmd
	native   tea.Cmd
	children []Cmd[M]
}

type cmdKind uint8

const (
	noCmd cmdKind = iota
	leafCmd
	nativeCmd
	batchCmd
	sequenceCmd
)

// message carries a PureScript message through Bubble Tea.
type message[M any] struct{ value M }

func None[M any]() Cmd[M] { return Cmd[M]{} }

func BatchImpl[M any](cmds []Cmd[M]) Cmd[M] { return Cmd[M]{kind: batchCmd, children: cmds} }

func SequenceImpl[M any](cmds []Cmd[M]) Cmd[M] {
	return Cmd[M]{kind: sequenceCmd, children: cmds}
}

func MapCmd[A, B any](f func(A) B, cmd Cmd[A]) Cmd[B] {
	switch cmd.kind {
	case leafCmd:
		leaf := cmd.leaf
		return Cmd[B]{kind: leafCmd, leaf: func(ctx context.Context, deliver func(B) tea.Msg) tea.Cmd {
			return leaf(ctx, func(a A) tea.Msg { return deliver(f(a)) })
		}}
	case nativeCmd:
		return Cmd[B]{kind: nativeCmd, native: cmd.native}
	case batchCmd, sequenceCmd:
		children := make([]Cmd[B], len(cmd.children))
		for i, child := range cmd.children {
			children[i] = MapCmd(f, child)
		}
		return Cmd[B]{kind: cmd.kind, children: children}
	}
	return Cmd[B]{}
}

func FromNativeImpl[M any](_ []M, native tea.Cmd) Cmd[M] { return FromNative[M](native) }

func FromNative[M any](native tea.Cmd) Cmd[M] {
	if native == nil {
		return Cmd[M]{}
	}
	return Cmd[M]{kind: nativeCmd, native: native}
}

func leaf[M any](build func(ctx context.Context, deliver func(M) tea.Msg) tea.Cmd) Cmd[M] {
	return Cmd[M]{kind: leafCmd, leaf: build}
}

func Emit[M any](value M) Cmd[M] {
	return leaf(func(_ context.Context, deliver func(M) tea.Msg) tea.Cmd {
		return func() tea.Msg { return deliver(value) }
	})
}

// PerformWith runs the IO action on the command's goroutine with the
// program's context. Once the program has ended its message goes nowhere.
func PerformWith[M any](action func(context.Context) M) Cmd[M] {
	return leaf(func(ctx context.Context, deliver func(M) tea.Msg) tea.Cmd {
		return func() tea.Msg {
			value := action(ctx)
			if ctx.Err() != nil {
				return nil
			}
			return deliver(value)
		}
	})
}

// MillisecondsToNextImpl is how long until the clock next reaches a whole
// multiple of the interval, as tea.Every counts it.
func MillisecondsToNextImpl(milliseconds float64) float64 {
	now := time.Now()
	interval := duration(milliseconds)
	return float64(now.Truncate(interval).Add(interval).Sub(now)) / float64(time.Millisecond)
}

func ExecImpl[M any](program string, arguments []string, failed func(string) M, finished M) Cmd[M] {
	return leaf(func(_ context.Context, deliver func(M) tea.Msg) tea.Cmd {
		return tea.ExecProcess(exec.Command(program, arguments...), func(err error) tea.Msg {
			if err != nil {
				return deliver(failed(err.Error()))
			}
			return deliver(finished)
		})
	})
}

func duration(milliseconds float64) time.Duration {
	return time.Duration(milliseconds * float64(time.Millisecond))
}

func Quit[M any]() Cmd[M]              { return FromNative[M](tea.Quit) }
func Interrupt[M any]() Cmd[M]         { return FromNative[M](tea.Interrupt) }
func Suspend[M any]() Cmd[M]           { return FromNative[M](tea.Suspend) }
func ClearScreen[M any]() Cmd[M]       { return FromNative[M](tea.ClearScreen) }
func RequestWindowSize[M any]() Cmd[M] { return FromNative[M](tea.RequestWindowSize) }
func ReadClipboard[M any]() Cmd[M]     { return FromNative[M](tea.ReadClipboard) }

func PrintlnImpl(line string) tea.Cmd { return tea.Println(line) }

func SetClipboardImpl(text string) tea.Cmd { return tea.SetClipboard(text) }

func toTea[M any](cmd Cmd[M], ctx context.Context, deliver func(M) tea.Msg) tea.Cmd {
	switch cmd.kind {
	case leafCmd:
		return cmd.leaf(ctx, deliver)
	case nativeCmd:
		return cmd.native
	case batchCmd, sequenceCmd:
		children := make([]tea.Cmd, 0, len(cmd.children))
		for _, child := range cmd.children {
			if converted := toTea(child, ctx, deliver); converted != nil {
				children = append(children, converted)
			}
		}
		if cmd.kind == batchCmd {
			return tea.Batch(children...)
		}
		return tea.Sequence(children...)
	}
	return nil
}

func deliver[M any](value M) tea.Msg { return message[M]{value} }

// The program

// model is what Bubble Tea runs: the PureScript model, or nothing for a
// program that keeps its state itself, and the functions that change and
// show it.
type model[S, M any] struct {
	state   S
	init    Cmd[M]
	update  func(tea.Msg, S) (S, Cmd[M])
	view    func(S) tea.View
	ctx     context.Context
	running *running[S]
}

func (m *model[S, M]) Init() tea.Cmd { return toTea(m.init, m.ctx, deliver[M]) }

func (m *model[S, M]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	state, cmd := m.update(msg, m.state)
	m.state = state
	return m, toTea(cmd, m.ctx, deliver[M])
}

func (m *model[S, M]) View() tea.View {
	frame := m.view(m.state)
	m.running.remember(frame.Content)
	return frame
}

type running[S any] struct {
	program *tea.Program
	ended   chan struct{}
	final   S
	err     error
	mu      sync.Mutex
	frame   string
}

func (r *running[S]) remember(frame string) {
	r.mu.Lock()
	r.frame = frame
	r.mu.Unlock()
}

// StartAppImpl runs an Elm-style program: Bubble Tea keeps the model and calls
// the pure update and view with it. first and second take apart what update
// returns.
func StartAppImpl[S, M, T any](
	width int,
	height int,
	first func(T) S,
	second func(T) Cmd[M],
	initial T,
	update func(tea.Msg) func(S) T,
	view func(S) tea.View,
) *running[S] {
	return start(width, height, &model[S, M]{
		state: first(initial),
		init:  second(initial),
		update: func(msg tea.Msg, state S) (S, Cmd[M]) {
			next := update(msg)(state)
			return first(next), second(next)
		},
		view: view,
	})
}

// StartProgramImpl runs a program that keeps its state itself: update and
// view are IO actions, run on Bubble Tea's goroutine.
func StartProgramImpl[S, M any](
	width int,
	height int,
	nothing S,
	init Cmd[M],
	update func(tea.Msg) Cmd[M],
	view func() tea.View,
) *running[S] {
	return start(width, height, &model[S, M]{
		state:  nothing,
		init:   init,
		update: func(msg tea.Msg, state S) (S, Cmd[M]) { return state, update(msg) },
		view:   func(S) tea.View { return view() },
	})
}

// start runs the program on a goroutine of its own. Once it ends, its context
// is cancelled, which stops the commands that wait with it, and ended closes.
// Headless runs (width > 0) read no terminal input and draw nothing.
func start[S, M any](width int, height int, m *model[S, M]) *running[S] {
	ctx, cancel := context.WithCancel(context.Background())
	r := &running[S]{ended: make(chan struct{}), final: m.state}
	m.ctx, m.running = ctx, r
	r.program = tea.NewProgram(m, options(width, height)...)
	go func() {
		defer close(r.ended)
		defer cancel()
		final, err := r.program.Run()
		if finished, ok := final.(*model[S, M]); ok {
			r.final = finished.state
		}
		r.err = err
	}()
	return r
}

func options(width int, height int) []tea.ProgramOption {
	if width <= 0 {
		return nil
	}
	return []tea.ProgramOption{
		tea.WithInput(nil),
		tea.WithOutput(io.Discard),
		tea.WithWindowSize(width, height),
		tea.WithoutSignalHandler(),
	}
}

// WaitImpl waits for the program to end and gives its last model, or what
// went wrong.
func WaitImpl[S, R any](failed func(string) R, finished func(S) R, r *running[S]) R {
	<-r.ended
	if r.err != nil {
		return failed(r.err.Error())
	}
	return finished(r.final)
}

func EndedImpl[S any](r *running[S]) <-chan struct{} { return r.ended }

// SendImpl waits until the program takes the message, or has ended.
func SendImpl[S, M any](r *running[S], value M) { r.program.Send(message[M]{value}) }

func SendNativeImpl[S any](r *running[S], msg tea.Msg) {
	if msg != nil {
		r.program.Send(msg)
	}
}

// StopImpl asks the program to quit. Unlike Program.Quit it also works on a
// program that hasn't got going yet.
func StopImpl[S any](r *running[S]) { r.program.Send(tea.QuitMsg{}) }

func LastFrameImpl[S any](r *running[S]) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.frame
}

func StripAnsi(text string) string { return ansi.Strip(text) }

// Reading Bubble Tea's messages

const (
	KindOther int = iota
	KindKeyPress
	KindKeyRelease
	KindMouseClick
	KindMouseRelease
	KindMouseWheel
	KindMouseMotion
	KindResize
	KindPaste
	KindFocus
	KindBlur
	KindClipboard
)

// KindImpl tells Bubble Tea's messages apart; MessageImpl finds PureScript's.
func KindImpl(msg tea.Msg) int {
	switch msg.(type) {
	case tea.KeyPressMsg:
		return KindKeyPress
	case tea.KeyReleaseMsg:
		return KindKeyRelease
	case tea.MouseClickMsg:
		return KindMouseClick
	case tea.MouseReleaseMsg:
		return KindMouseRelease
	case tea.MouseWheelMsg:
		return KindMouseWheel
	case tea.MouseMotionMsg:
		return KindMouseMotion
	case tea.WindowSizeMsg:
		return KindResize
	case tea.PasteMsg:
		return KindPaste
	case tea.FocusMsg:
		return KindFocus
	case tea.BlurMsg:
		return KindBlur
	case tea.ClipboardMsg:
		return KindClipboard
	}
	return KindOther
}

func MessageImpl[M any, R any](missing R, found func(M) R, msg tea.Msg) R {
	if carried, ok := msg.(message[M]); ok {
		return found(carried.value)
	}
	return missing
}

func keyOf(msg tea.Msg) tea.Key {
	switch key := msg.(type) {
	case tea.KeyPressMsg:
		return tea.Key(key)
	case tea.KeyReleaseMsg:
		return tea.Key(key)
	}
	return tea.Key{}
}

func KeyCodeImpl(msg tea.Msg) int      { return int(keyOf(msg).Code) }
func KeyTextImpl(msg tea.Msg) string   { return keyOf(msg).Text }
func KeystrokeImpl(msg tea.Msg) string { return keyOf(msg).Keystroke() }
func KeyModifiersImpl(msg tea.Msg) int { return int(keyOf(msg).Mod) }
func KeyRepeatImpl(msg tea.Msg) bool   { return keyOf(msg).IsRepeat }

func mouseOf(msg tea.Msg) tea.Mouse {
	if mouse, ok := msg.(tea.MouseMsg); ok {
		return mouse.Mouse()
	}
	return tea.Mouse{}
}

func MouseXImpl(msg tea.Msg) int         { return mouseOf(msg).X }
func MouseYImpl(msg tea.Msg) int         { return mouseOf(msg).Y }
func MouseButtonImpl(msg tea.Msg) int    { return int(mouseOf(msg).Button) }
func MouseModifiersImpl(msg tea.Msg) int { return int(mouseOf(msg).Mod) }

func WidthImpl(msg tea.Msg) int {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		return size.Width
	}
	return 0
}

func HeightImpl(msg tea.Msg) int {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		return size.Height
	}
	return 0
}

func ContentImpl(msg tea.Msg) string {
	switch content := msg.(type) {
	case tea.PasteMsg:
		return content.Content
	case tea.ClipboardMsg:
		return content.Content
	}
	return ""
}

// Building Bubble Tea's messages, for components and scripted runs

func KeyPressNative(code int, text string, modifiers int, repeat bool) tea.Msg {
	return tea.KeyPressMsg(tea.Key{Code: rune(code), Text: text, Mod: tea.KeyMod(modifiers), IsRepeat: repeat})
}

func KeyReleaseNative(code int, text string, modifiers int) tea.Msg {
	return tea.KeyReleaseMsg(tea.Key{Code: rune(code), Text: text, Mod: tea.KeyMod(modifiers)})
}

func MouseNative(kind int, x int, y int, button int, modifiers int) tea.Msg {
	mouse := tea.Mouse{X: x, Y: y, Button: tea.MouseButton(button), Mod: tea.KeyMod(modifiers)}
	switch kind {
	case KindMouseRelease:
		return tea.MouseReleaseMsg(mouse)
	case KindMouseWheel:
		return tea.MouseWheelMsg(mouse)
	case KindMouseMotion:
		return tea.MouseMotionMsg(mouse)
	}
	return tea.MouseClickMsg(mouse)
}

func ResizeNative(width int, height int) tea.Msg {
	return tea.WindowSizeMsg{Width: width, Height: height}
}

func PasteNative(text string) tea.Msg { return tea.PasteMsg{Content: text} }

func ClipboardNative(text string) tea.Msg { return tea.ClipboardMsg{Content: text} }

func FocusNative() tea.Msg { return tea.FocusMsg{} }

func BlurNative() tea.Msg { return tea.BlurMsg{} }

func NoNative() tea.Msg { return nil }

const (
	ModShift int = int(tea.ModShift)
	ModAlt   int = int(tea.ModAlt)
	ModCtrl  int = int(tea.ModCtrl)
	ModMeta  int = int(tea.ModMeta)
	ModHyper int = int(tea.ModHyper)
	ModSuper int = int(tea.ModSuper)
)

const (
	MouseNone       int = int(tea.MouseNone)
	MouseLeft       int = int(tea.MouseLeft)
	MouseMiddle     int = int(tea.MouseMiddle)
	MouseRight      int = int(tea.MouseRight)
	MouseWheelUp    int = int(tea.MouseWheelUp)
	MouseWheelDown  int = int(tea.MouseWheelDown)
	MouseWheelLeft  int = int(tea.MouseWheelLeft)
	MouseWheelRight int = int(tea.MouseWheelRight)
	MouseBackward   int = int(tea.MouseBackward)
	MouseForward    int = int(tea.MouseForward)
)

// Views

func ViewImpl(content string) tea.View { return tea.NewView(content) }

func AltScreen(view tea.View) tea.View {
	view.AltScreen = true
	return view
}

func WindowTitleImpl(title string, view tea.View) tea.View {
	view.WindowTitle = title
	return view
}

func BackgroundImpl(background color.Color, view tea.View) tea.View {
	view.BackgroundColor = background
	return view
}

func ForegroundImpl(foreground color.Color, view tea.View) tea.View {
	view.ForegroundColor = foreground
	return view
}

// MouseModeImpl: 1 reports clicks, wheel and drags, 2 every motion too.
func MouseModeImpl(mode int, view tea.View) tea.View {
	switch mode {
	case 1:
		view.MouseMode = tea.MouseModeCellMotion
	case 2:
		view.MouseMode = tea.MouseModeAllMotion
	default:
		view.MouseMode = tea.MouseModeNone
	}
	return view
}

func ReportFocus(view tea.View) tea.View {
	view.ReportFocus = true
	return view
}

func CursorImpl(x int, y int, view tea.View) tea.View {
	view.Cursor = tea.NewCursor(x, y)
	return view
}

func ViewContentImpl(view tea.View) string { return view.Content }

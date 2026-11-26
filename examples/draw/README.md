# Draw

The smallest thing that's fun with more than one person: a canvas in your terminal that everyone
shares. Click or drag to paint, and whatever anyone paints shows up for everyone. That's it, no
colors, no names, no cursors. It's here to read in one go, all PureScript, compiled to Go with
[googoo](https://github.com/i-am-the-slime/iris-go-backend), with
[Bubble Tea](https://github.com/i-am-the-slime/purescript-bubbletea) for the terminal.

## Running it

You need googoo 0.1.18-go.6 or newer and Go 1.27, and a terminal that reports the mouse (iTerm2,
kitty, WezTerm, Ghostty, Terminal.app all do).

```
cd examples/draw
googoo run -p draw --main Draw.Server
```

Then a client in another terminal, and one more in a third:

```
googoo run -p draw --main Draw.Client
googoo run -p draw --main Draw.Client
```

Drag in one and watch it appear in the other. `q` quits. A client that just started shows the
canvas once you've painted your first cell.

## How it works

[Service.purs](src/Draw/Service.purs) is the whole service, as a type: one call, `draw`, that
streams cells both ways. [Server.purs](src/Draw/Server.purs) keeps the painted cells and everyone's
outbox in an `IORef`; when you join it sends you what's there, and every cell you send goes into
everyone's outbox, yours too. [Client.purs](src/Draw/Client.purs) is a Bubble Tea component: mouse
clicks and drags go into the call's request stream, and every cell coming back lands in a `Set`
that gets drawn as blocks.

The spec starts a server and two clients without a terminal and drags the mouse in them:

```
googoo run -p draw-test
```

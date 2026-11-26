# Live cursors

A canvas in your terminal that everyone shares. You see everyone else's mouse as a little colored
arrow with their name on it, whatever anyone draws shows up for everyone, and what anyone types
appears in a bubble by their arrow, letter by letter, as they type it. It's one gRPC call
per person, streaming both ways, all PureScript, compiled to Go with
[googoo](https://github.com/i-am-the-slime/iris-go-backend), with
[Bubble Tea](https://github.com/i-am-the-slime/purescript-bubbletea) for the terminal.

```
                     on it ▌
                    ▲bob
     ████████
           ███             ▲cleo
             ██
● ada (you)  ● bob  ● cleo  ■ 1–6 color · click or drag draws · enter says · q quits
```

The whole service is a type, in [src/Cursors/Service.purs](src/Cursors/Service.purs):

```purescript
type Canvas = Service "canvas.v1.Canvas"
  { join :: BidiStream Gesture Scene (nameTaken :: String, invalidArgument :: String)
  }

type Gesture = Named "Gesture" (OneOf
  ( hello :: 1 := Message { name :: 1 := String }
  , moved :: 2 := Point
  , drew :: 3 := Cell
  , typing :: 4 := Message { text :: 1 := String }
  , said :: 5 := Message { text :: 1 := String }
  ))
```

You say hello with your name, then send where your mouse is, what you draw and what you're typing.
You get back a `welcome` with the canvas as it is (and what everyone's typing right now), then
whatever happens: cursors moving, cells drawn, bubbles changing, people coming and going.

## Running it

You need googoo 0.1.18-go.5 or newer (`brew install spago i-am-the-slime/tap/googoo`) and Go 1.27,
and a terminal that reports the mouse (iTerm2, kitty, WezTerm, Ghostty, Terminal.app all do).

Start the server:

```
cd examples/live-cursors
googoo run -p live-cursors --main Cursors.Server
```

Then a client in another terminal, with your name:

```
googoo run -p live-cursors --main Cursors.Client -- ada
```

and another one or two in more terminals, with other names:

```
googoo run -p live-cursors --main Cursors.Client -- bob
googoo run -p live-cursors --main Cursors.Client -- cleo
```

Move the mouse in one and watch the arrow move in the others. Click or drag to draw, press 1 to 6
for another color, `q`, Esc or ctrl+c to leave. Leaving takes your arrow off everyone's canvas.

Press Enter to say something. Everyone sees it in a bubble above your arrow while you type it, a
letter at a time, Backspace too; your own line is in the status line. While you're typing, every
key types, so `q` and 1 to 6 are just letters. Enter sends it: the bubble stays for a few seconds,
fades and goes. Esc gives up on it and takes the bubble away.

The server listens on 50052; `-- 50060` picks another port, and the client takes the address
after the name: `-- ada http://127.0.0.1:50060`.

Stop the server while clients are open and they say they're joining again; start it again and
they're back (on an empty canvas, the server keeps it in memory).

A name that's taken gets turned down with `nameTaken`, an error of the service's own (other gRPC
clients see `FAILED_PRECONDITION` and a `NameTaken` detail), and the client says so:

```
● ada  joining again, ada is here already, pick another name
```

## How it works

[The server](src/Cursors/Server.purs) keeps everyone and the canvas in an `IORef`. Each person's
call is their gestures coming in, merged with their own outbox of scenes going out:

```purescript
join :: Handler Joining (room :: Room)
join = handle do
  { requests: gestures, room } <- ask
  pure (Strom.bracketExit (arrive # liftIO) (leave room) \visit -> scenesFor room visit gestures)
```

`bracketExit` is what makes leaving work: however the call ends (you quit, your terminal dies,
your laptop goes to sleep), `leave` runs, takes you off the canvas and tells everyone else.

A mouse sends a lot of moves. They don't go out one by one: the server notes the latest position
of each cursor and every 25 milliseconds sends everyone one `cursors` scene with all of them, and
one `drawn` with the cells drawn since. So 40 updates a second at most, however fast you wiggle.
Someone whose connection can't keep up misses scenes instead of holding everyone else up.

Typing goes the same way: every keystroke sends the whole line so far, the server keeps the
latest line of each person, and each tick one `speech` scene has the bubbles that changed and the
names whose bubble is gone. The server also counts down a line that was said, three seconds and
one fading, so someone joining late sees the same bubbles as everyone else.

[The client](src/Cursors/Client.purs) is a Bubble Tea component with hooks. The mouse comes in as
`Tea.Mouse` events (`Tea.mouse Tea.AllMotion` asks the terminal for every move, not only drags),
which go into a channel that is the call's request stream. Scenes come back on the call's
goroutine and go to Bubble Tea's through `useCommand`, the same way the
[ticker dashboard](https://github.com/i-am-the-slime/purescript-bubbletea/tree/main/examples/ticker-dashboard)
does it.

## Other clients

`googoo run -q -p live-cursors --main Cursors.Service > canvas.proto` prints the service as a
`.proto`, so anything that speaks gRPC can join. With
[grpcurl](https://github.com/fullstorydev/grpcurl), say hello and park a cursor for two seconds:

```
(echo '{"hello":{"name":"curl"}}'; echo '{"moved":{"x":12,"y":3}}'; sleep 2) | \
  grpcurl -plaintext -import-path . -proto canvas.proto -max-time 2 -d @ \
  127.0.0.1:50052 canvas.v1.Canvas/Join
```

and `curl` shows up in everyone's terminal until grpcurl gives up.

## Specs

```
googoo run -p live-cursors-test
```

They start a server and two or three clients without a terminal, move and drag with mouse events,
and look at what the others' last frames show.

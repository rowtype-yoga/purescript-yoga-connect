# Live cursors

A canvas in your terminal that everyone shares. You see everyone else's mouse as a little colored
arrow with their name on it, and whatever anyone draws shows up for everyone. It's one gRPC call
per person, streaming both ways, all PureScript, compiled to Go with
[googoo](https://github.com/i-am-the-slime/iris-go-backend), with
[Bubble Tea](https://github.com/i-am-the-slime/purescript-bubbletea) for the terminal.

```
                    ▲bob
     ████████
           ███             ▲cleo
             ██
● ada (you)  ● bob  ● cleo  ■ 1–6 color · click or drag draws · q quits
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
  ))
```

You say hello with your name, then send where your mouse is and what you draw. You get back a
`welcome` with the canvas as it is, then whatever happens: cursors moving, cells drawn, people
coming and going.

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

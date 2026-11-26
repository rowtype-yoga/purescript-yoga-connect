# purescript-yoga-connect

gRPC for PureScript, from a PureScript type ([Services as types](#services-as-types)) or from your
`.proto` files. Servers and clients are derived from one service row, and on Go (through
[googoo](https://github.com/i-am-the-slime/iris-go-backend)) each method compiles down to
straight-line Go: decode, your handler, encode.

```proto
service Greeter {
  rpc SayHello(HelloRequest) returns (HelloReply);
  rpc ListGreetings(ListRequest) returns (stream Greeting);
  rpc GreetEveryone(stream HelloRequest) returns (HelloReply);
  rpc Chat(stream ChatMessage) returns (stream ChatMessage);
}
```

`protoc-gen-purescript` from
[yoga-protobuf](https://github.com/rowtype-yoga/purescript-yoga-protobuf) turns that into

```purescript
type Greeter =
  ( sayHello :: Connect.Unary HelloRequest HelloReply
  , listGreetings :: Connect.ServerStream ListRequest Greeting
  , greetEveryone :: Connect.ClientStream HelloRequest HelloReply
  , chat :: Connect.BidiStream ChatMessage ChatMessage
  )

greeter :: Connect.Service Greeter
```

`greeter` carries each method's route and message codecs. The messages are plain records, so
the codecs travel as a value rather than as instances.

## Services as types

You don't need a `.proto` at all. Write the service as a PureScript type and the server, the client
and the `.proto` all come from it:

```purescript
import Yoga.Connect.Om

type Ticker = Service "ticker.v1.Ticker"
  { watch :: ServerStream Subscribe Quote (notFound :: String)
  , trade :: BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)
  }

type Subscribe = Message { symbols :: 1 := Array String }
type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }
```

The numbers are the protocol buffer field numbers, so they stay put when you rename a field. The
last row of a method is how it can fail, gRPC status codes by name. Handlers `ask` for the request
and whatever else they need, like routes in yoga-go-http-om:

```purescript
main = do
  server <- createServer
  registerServiceWith @Ticker { market } { watch, trade } server
  launchIO_ (listen { host: Host "127.0.0.1", port: Port 50051 } server)

watch :: Handler (ServerStream Subscribe Quote (notFound :: String)) (market :: Market)
watch = handle do
  { request: { symbols }, market } <- ask
  listings <- traverse (listingIn market) symbols # maybe (reject @"notFound" "we don't list that") pure
  pure (Strom.mergeAll (listings <#> quotesOf))
```

Handlers and clients only ever see plain records: a `Quote` is
`{ symbol :: String, price :: Number, average :: Number }` to them.

| Method | `handle`s |
| --- | --- |
| `Unary req res err` | `OmIO { request :: req \| ctx } err res` |
| `ServerStream req res err` | `OmIO { request :: req \| ctx } err (Strom (Record ctx) err res)` |
| `ClientStream req res err` | `OmIO { requests :: Strom (Record ctx) err req \| ctx } err res` |
| `BidiStream req res err` | `OmIO { requests :: Strom (Record ctx) err req \| ctx } err (Strom (Record ctx) err res)` |

An error doesn't have to be a gRPC code. Any label works, with any value a field can have:

```purescript
type Playing = BidiStream Move Event
  ( notYourTurn :: String
  , columnFull :: Int
  , nameTaken :: WithCode "alreadyExists" String
  , notFound :: String
  )
```

`notFound :: String` is plain `NOT_FOUND` with its message. The others are `FAILED_PRECONDITION`
(or the code `WithCode` says) with their value as an error detail, a `ColumnFull` message in the
`.proto`, so `reject @"columnFull" 3` on the server is `columnFull` with `3` on the client, and
grpcurl shows the detail.

A client is a record of calls, and fails with the method's own errors, so a `NOT_FOUND` from the
server is your `notFound` with its message:

```purescript
let ticker = client @Ticker connection
ticker.watch { symbols: [ "ACME" ] } # Strom.take 3 # Strom.forEach (log <<< _.symbol)
```

and `printProto @Ticker` gives you the `.proto`, for grpcurl or for clients in other languages. The
encoding is worked out from the types, with nothing generated, and it's the same bytes protoc's
code writes, so clients generated from a `.proto` talk to these servers and the other way round.

What a field can be:

| PureScript | protocol buffers |
| --- | --- |
| `String`, `Boolean`, `Number`, `Bytes` | `string`, `bool`, `double`, `bytes` |
| `Int` | `int32` |
| `Int64`, `UInt32`, `UInt64` | `int64`, `uint32`, `uint64` |
| a sum type of constructors without arguments, with `Generic` | an `enum`, the first constructor is 0 |
| `Message { … }` | a nested message |
| `Array a` | `repeated a`, packed if it's a number |
| `Maybe a` | `optional a` |
| `OneOf (circle :: 4 := Circle, square :: 5 := Square)`, no number of its own | a `oneof`, a `Variant` to handlers |

A method's messages are called `WatchRequest` and `WatchResponse` in the `.proto`, and nested ones
after their field. `Named "Quote" (Message { … })` names one yourself. A method can take or answer
something that isn't a message, `Unary {} (Array Room) ()`, and it goes as the field `value = 1`.

[examples/ticker](examples/ticker) and [examples/live-cursors](examples/live-cursors) are built
this way.

## Installing

It's all in the [Go package set](https://github.com/rowtype-yoga/purescript-go-package-sets), so your
`spago.yaml` only needs the set and the packages:

```yaml
package:
  name: my-service
  dependencies:
    - yoga-connect-go
    - yoga-strom-io
workspace:
  packageSet:
    url: https://raw.githubusercontent.com/rowtype-yoga/purescript-go-package-sets/go-0.1.18-20261126.6/packages.json
```

You need googoo 0.1.18-go.5 or newer (`brew install spago i-am-the-slime/tap/googoo`) and Go 1.27.
`protoc-gen-purescript` comes with yoga-protobuf, its README says how to build it.

## Server

```purescript
main :: Effect Unit
main = do
  server <- createServer
  serve greeter { sayHello, listGreetings, greetEveryone, chat } server
  launchIO_ do
    bound <- listen { host: Host "127.0.0.1", port: Port 50051 } server
    log ("listening on " <> show bound.port)

sayHello :: forall ctx. HelloRequest -> OmIO ctx (invalidArgument :: String) HelloReply
sayHello { name }
  | name == "" = reject @"invalidArgument" "say who you are"
  | otherwise = pure { message: "Hello, " <> name <> "!" }

listGreetings :: forall ctx. ListRequest -> Strom ctx () Greeting
listGreetings { name, count } = Strom.range 1 (count + 1) # map \index -> { index, text: name }

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # map _.name # Strom.runCollect
  pure { message: "Hello, " <> String.joinWith " and " names <> "!" }

-- | Ends when the client is done talking; the heartbeat stops with it.
chat :: forall ctx. Strom ctx () ChatMessage -> Strom ctx () ChatMessage
chat incoming = Strom.mergeHaltLeft (incoming # map answer) heartbeats
```

| Method | Handler |
| --- | --- |
| `Unary req res` | `req -> OmIO ctx err res` |
| `ServerStream req res` | `req -> Strom ctx err res` |
| `ClientStream req res` | `Strom ctx err req -> OmIO ctx err res` |
| `BidiStream req res` | `Strom ctx err req -> Strom ctx err res` |

- Handlers run in `OmIO` (from yoga-go-http-om) on the call's own goroutine. No `Aff`.
- `serveWith context` hands them a context they `ask` for.
- A handler's error row names the status codes it rejects with: `notFound`, `permissionDenied`
  and so on, one label per gRPC code, each with its message. `reject @"notFound" "no such user"`
  answers with `NOT_FOUND`. Any other exception is `INTERNAL`, and a request that doesn't decode
  is `INVALID_ARGUMENT`.
- The server runs on connect-go, so the same port answers gRPC, gRPC-Web and Connect, over HTTP/2
  without TLS (h2c) and HTTP/1.1.

## Client

```purescript
talkTo :: Connection -> OmIO Unit (connectError :: ConnectError) Unit
talkTo connection = do
  let greeterClient = client greeter connection
  reply <- greeterClient.sayHello { name: "Ada" }
  greeterClient.listGreetings { name: "Ada", count: 3 } # Strom.forEach \{ text } -> say text
  everyone <- greeterClient.greetEveryone (Strom.fromArray [ { name: "Ada" }, { name: "Grace" } ])
  greeterClient.chat chatter # Strom.forEach \{ from, text } -> say (from <> ": " <> text)
```

`connect "http://127.0.0.1:50051"` speaks gRPC over HTTP/2. A failed call throws
`connectError :: ConnectError`, which has the code and the message.

## Streams

Streams are `Yoga.OmIO.Strom`, from `packages/yoga-strom-io`: pulled in chunks, like ZIO's
`ZStream`, in `OmIO`. Nothing is read before someone pulls, so a slow client slows the handler
down: `Send` blocks while HTTP/2 flow control says the client can't take more.

A streaming call lives as long as its stream:

- A client that's had enough cancels the call. `take 3` of a server stream ends it after the third
  message, and the server's stream is interrupted: its goroutines stop and its finalizers run.
- `Strom.timeout` around a client's stream is the call's deadline, and the server sees it.
- When a client cancels or its deadline passes, the handler's stream is interrupted the same way.

Unary and client-stream calls take theirs from a Go context: `clientWithin deadline greeter connection`
makes them with one, say from go-io's `withTimeout`, and they fail with `DeadlineExceeded` once it
passes. `client` makes them without a deadline.

```purescript
listGreetings { name, count } = Strom.every (Milliseconds 20.0)
  # Strom.take count
  # map greeting
  # Strom.ensuring (log "listGreetings is done, however it ended" # liftIO)

greeterClient.listGreetings { name: "Grace", count: 1000 } # Strom.take 3 # Strom.forEach say
```

[examples/log-tailer](examples/log-tailer) shows the streams without gRPC: three logs merged,
parsed, cut into windows, looked up four at a time with `mapPar`, reopened with `catchAll` when one
is rotated, and closed by their finalizers once `take` has enough. Next to it, the same program in
plain Go.

## Packages

- `yoga-connect`: the service types, messages as types and their protocol buffer codecs, the
  `.proto` printer and the status codes, for both sides and any backend
- `yoga-connect-go`: the server and the client on Go
- `yoga-strom-io`: the streams, which don't need gRPC

A browser client would be another package next to `yoga-connect-go`, derived from the same rows.

## Examples

[examples/draw](examples/draw): the smallest of them, to read in one sitting. A shared canvas in the
terminal: one bidi call each, the server passes every painted cell on to everyone.

[examples/ticker](examples/ticker): a live stock ticker. The server merges a price stream per
symbol and fills trades over a bidi stream, turning bad ones down with `FAILED_PRECONDITION`; the
client watches with a deadline and trades. Its README holds it up against the same server in
grpc-go.

[examples/live-cursors](examples/live-cursors): a canvas in the terminal that everyone shares, with
Bubble Tea. Everyone sees everyone else's mouse and drawing, one bidi call each, cursor moves
sent out together every 25 milliseconds.

## Demo

`demo/proto/greeter.proto` is generated into `demo/greeter`. The server and the client are
`demo/server` and `demo/client`.

```
googoo run --go -p greeter-server
googoo run --go -p greeter-client
```

Any gRPC client works too:

```
grpcurl -plaintext -proto demo/proto/greeter.proto \
  -d '{"from":"me","text":"hi"} {"from":"me","text":"bye"}' \
  127.0.0.1:50051 greeter.v1.Greeter/Chat
```

## Specs

```
googoo run --go -p yoga-connect-test
googoo run --go -p yoga-strom-io-test
```

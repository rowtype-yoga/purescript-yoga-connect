# purescript-yoga-connect

gRPC for PureScript, from your `.proto` files. Servers and clients are derived from one service
row, and on Go (through [Iris](https://github.com/i-am-the-slime/iris-go-backend)) each method
compiles down to straight-line Go: decode, your handler, encode.

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

## Installing

None of it is on the registry yet, so it all comes from GitHub. In your `spago.yaml`:

```yaml
workspace:
  extraPackages:
    yoga-connect:
      git: https://github.com/rowtype-yoga/purescript-yoga-connect.git
      ref: v0.1.1
      subdir: packages/yoga-connect
    yoga-connect-go:
      git: https://github.com/rowtype-yoga/purescript-yoga-connect.git
      ref: v0.1.1
      subdir: packages/yoga-connect-go
    yoga-strom-io:
      git: https://github.com/rowtype-yoga/purescript-yoga-connect.git
      ref: v0.1.1
      subdir: packages/yoga-strom-io
    yoga-protobuf:
      git: https://github.com/rowtype-yoga/purescript-yoga-protobuf.git
      ref: v0.1.0
    yoga-go-http-om:
      git: https://github.com/rowtype-yoga/purescript-yoga-go-http-om.git
      ref: v0.1.1
    go-io:
      git: https://github.com/i-am-the-slime/purescript-go.git
      ref: v0.2.0
      subdir: packages/go-io
```

You need Iris 0.1.18-go.2 or newer (`brew install spago i-am-the-slime/tap/iris`) and Go 1.27.
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

- `yoga-connect`: the service types and status codes, for both sides and any backend
- `yoga-connect-go`: the server and the client on Go
- `yoga-strom-io`: the streams, which don't need gRPC

A browser client would be another package next to `yoga-connect-go`, derived from the same rows.

## Examples

[examples/ticker](examples/ticker): a live stock ticker. The server merges a price stream per
symbol and fills trades over a bidi stream, turning bad ones down with `FAILED_PRECONDITION`; the
client watches with a deadline and trades. Its README holds it up against the same server in
grpc-go.

## Demo

`demo/proto/greeter.proto` is generated into `demo/greeter`. The server and the client are
`demo/server` and `demo/client`.

```
iris run --go -p greeter-server
iris run --go -p greeter-client
```

Any gRPC client works too:

```
grpcurl -plaintext -proto demo/proto/greeter.proto \
  -d '{"from":"me","text":"hi"} {"from":"me","text":"bye"}' \
  127.0.0.1:50051 greeter.v1.Greeter/Chat
```

## Specs

```
iris run --go -p yoga-connect-test
iris run --go -p yoga-strom-io-test
```

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
listGreetings { name, count } = Strom.range 1 count # map \index -> { index, text: name }

greetEveryone :: forall ctx. Strom ctx () HelloRequest -> OmIO ctx () HelloReply
greetEveryone people = do
  names <- people # map _.name # Strom.collect
  pure { message: "Hello, " <> String.joinWith " and " names <> "!" }

chat :: forall ctx. Strom ctx () ChatMessage -> Strom ctx () ChatMessage
chat incoming = Strom.merge (incoming # map answer) heartbeats
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

`Yoga.Connect.Strom` is yoga-om-strom's `Strom`, running in `OmIO` on the call's goroutine instead
of in `Aff`: `range`, `fromArray`, `iterate`, `unfold`, `repeatOm`, `every`, `map`, `mapM`,
`filter`, `scan`, `take`, `takeWhile`, `drop`, `concatMap`, `merge`, `<>`, and `forEach`, `foldl`,
`collect`, `drain` to run one.

A stream hands each element over and waits until it's taken. A server stream sends each message as
it comes, and `Send` blocks while HTTP/2 flow control says the client can't take more, so a slow
client slows the handler down and nothing piles up. `take` stops the stream it takes from, and
`merge` runs both sides on threads of their own.

## Packages

- `yoga-connect`: the service types and status codes, for both sides and any backend
- `yoga-connect-go`: the server, the client and `Strom` on Go

A browser client would be another package next to `yoga-connect-go`, derived from the same rows.

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
```

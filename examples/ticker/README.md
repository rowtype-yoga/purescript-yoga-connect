# Ticker

A live stock ticker over gRPC, in PureScript, compiled to one Go binary with
[googoo](https://github.com/i-am-the-slime/iris-go-backend).

The service is a PureScript type. The server, the client and the `.proto` all come from it:

```purescript
type Ticker = Service "ticker.v1.Ticker"
  { watch :: ServerStream Subscribe Quote (notFound :: String)
  , trade :: BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)
  }

type Quote = Message { symbol :: 1 := String, price :: 2 := Number, average :: 3 := Number }
```

- `watch` merges a random walk per symbol into one stream of quotes, each with its moving average
- `trade` fills your trades as they come in, and turns you down with `FAILED_PRECONDITION` when you
  sell what you don't have
- a client that hangs up, or whose deadline passes, stops the server's streams, and their
  finalizers say so

The service is [src/Ticker/Service.purs](src/Ticker/Service.purs), the whole server is
[src/Ticker/Server.purs](src/Ticker/Server.purs), the client is
[src/Ticker/Client.purs](src/Ticker/Client.purs).

## Running it

You need googoo 0.1.18-go.5 or newer (`brew install spago i-am-the-slime/tap/googoo`) and Go 1.27.

```
cd examples/ticker
googoo run -p ticker --main Ticker.Server
```

and in another terminal

```
googoo run -p ticker --main Ticker.Client
```

```
watching ACME, GLOBEX and INITECH for two seconds
ACME      121.02  ▼  avg  121.40
GLOBEX     47.88  ▼  avg   48.14
ACME      121.19  ▼  avg  121.36
INITECH     9.87  ▲  avg    9.77
ACME      121.27  ▼  avg  121.34
GLOBEX     47.91  ▼  avg   48.09
…
trading
→ buy 10 ACME at 121.40
← bought 10 ACME at 121.40, holding 10
→ buy 25 INITECH at 9.80
← bought 25 INITECH at 9.80, holding 25
→ sell 5 ACME at 122.10
← sold 5 ACME at 122.10, holding 5
→ sell 40 INITECH at 9.90
refused: failed_precondition: you hold 25 INITECH, so you can't sell 40
done
```

Meanwhile the server tells you how each call ended:

```
ticker listening on 127.0.0.1:50051
watching ACME, GLOBEX, INITECH: cancelled, and its finalizer ran
trading, holding 5 ACME, 25 INITECH: failed, and its finalizer ran
```

The client watched for two seconds (`Strom.timeout`), which cancelled the call, which interrupted
all three price streams on the server.

## The interesting bits

The numbers in a message are the protocol buffer field numbers, so a field keeps its number when
you rename it. Handlers and clients never see them: a `Quote` is
`{ symbol :: String, price :: Number, average :: Number }` to them, and the encoding is worked out
from the type, no protoc and no generated code.

The compiler checks that the server has a handler for every method, of the right kind, and that
the context has what each handler asks for:

```purescript
registerServiceWith @Ticker { market } { watch, trade } server
```

A handler `ask`s for the request and its context. Watching is a merge of one stream per symbol:

```purescript
watch :: Handler Watching (market :: Market)
watch = handle do
  { request: { symbols }, market } <- ask
  let wanted = if Array.null symbols then market <#> _.symbol else symbols
  listings <- traverse (listingIn market) wanted # maybe (reject @"notFound" ("we list " <> symbolsIn market)) pure
  pure (Strom.mergeAll (listings <#> quotesOf) # Strom.onExit (sayHowItEnded ("watching " <> joinWith ", " wanted)))
```

and each of those is a tick, a random walk and a moving average:

```purescript
quotesOf { symbol, opening, volatility, pace, seed } = Strom.every pace
  # Strom.scan (\walk _ -> step walk) (Seed seed /\ opening)
  # map snd
  # Strom.scan (\(_ /\ average) price -> price /\ smooth average price) (opening /\ opening)
  # map \(price /\ average) -> { symbol, price: toCents price, average: toCents average }
```

The last row of a method is the list of status codes its handler can answer with.
`reject @"notFound"` only compiles because `notFound` is in it, and the client gets `NOT_FOUND`:

```purescript
type Trading = BidiStream Trade Fill (notFound :: String, failedPrecondition :: String)

trade :: Handler Trading (market :: Market)
trade = handle do
  { requests: trades, market } <- ask
  pure (Strom.bracketExit openBook closeBook \book -> trades # Strom.mapM (fill market book))
```

`bracketExit` closes the book however the session ends: the client is done, a trade is refused, or
the client hangs up.

The client gets those as errors of its own, `failedPrecondition` with the server's message, next
to `connectError` for everything else. A deadline is just a stream that stops:

```purescript
let ticker = client @Ticker connection
ticker.watch { symbols: [ "ACME", "GLOBEX", "INITECH" ] }
  # Strom.timeout (Milliseconds 2000.0)
  # Strom.forEach (say <<< describeQuote)
```

## Compared with grpc-go

I wrote the same server with grpc-go and protoc-gen-go-grpc, so it streams the very same prices.
Leaving out the code protoc generates for Go:

| | PureScript | Go |
| --- | --- | --- |
| the whole server | 129 lines, and 38 for the service | 165 lines |
| `watch` (merging the symbols, moving averages, cancellation) | 17 lines | 62 lines |
| `trade` (positions, rejections, the closing log line) | 27 lines | 25 lines |

The difference is in `watch`. In Go, merging streams means a goroutine per symbol, a channel, a
`WaitGroup`, and a `select` on `ctx.Done()` at every place one could wait:

<details>
<summary>Watch in grpc-go</summary>

```go
func (server) Watch(req *pb.Subscribe, stream grpc.ServerStreamingServer[pb.Quote]) error {
	wanted := req.Symbols
	if len(wanted) == 0 {
		for _, l := range market {
			wanted = append(wanted, l.symbol)
		}
	}
	var listings []listing
	for _, symbol := range wanted {
		l, ok := listingOf(symbol)
		if !ok {
			return status.Error(codes.NotFound, "we list ACME, GLOBEX, INITECH, UMBRELLA")
		}
		listings = append(listings, l)
	}
	ctx, cancel := context.WithCancel(stream.Context())
	quotes := make(chan *pb.Quote)
	var feeds sync.WaitGroup
	for _, l := range listings {
		feeds.Add(1)
		go func() {
			defer feeds.Done()
			feed(ctx, l, quotes)
		}()
	}
	defer func() {
		cancel()
		feeds.Wait()
		log.Printf("watching %s: cancelled, and its finalizer ran", strings.Join(wanted, ", "))
	}()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case quote := <-quotes:
			if err := stream.Send(quote); err != nil {
				return err
			}
		}
	}
}

// One symbol's random walk with its moving average, until the context is done.
func feed(ctx context.Context, l listing, quotes chan<- *pb.Quote) {
	ticker := time.NewTicker(l.pace)
	defer ticker.Stop()
	seed, price, average := l.seed, l.opening, l.opening
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		price *= 1 + l.volatility*wobble(seed)
		seed = nextSeed(seed)
		average += 0.2 * (price - average)
		select {
		case quotes <- &pb.Quote{Symbol: l.symbol, Price: cents(price), Average: cents(average)}:
		case <-ctx.Done():
			return
		}
	}
}
```

</details>

Get the order of `cancel()` and `feeds.Wait()` wrong and a client that hangs up leaves the handler
stuck forever. In PureScript, `mergeAll` owns those goroutines, and closing the call closes them.

## Any gRPC client works

The `.proto` comes from the type too:

```
googoo run -q -p ticker --main Ticker.Service > ticker.proto
```

```proto
service Ticker {
  rpc Trade(stream TradeRequest) returns (stream TradeResponse);
  rpc Watch(WatchRequest) returns (stream WatchResponse);
}

message TradeRequest {
  string symbol = 1;
  int32 quantity = 2;
  double limit = 3;
}
…
```

Each method gets a request and a response message of its own. The names don't go over the wire,
only the field numbers do, so a client generated from the `.proto` this example used to have
(with `Subscribe` and `Quote`) still works.

With [grpcurl](https://github.com/fullstorydev/grpcurl):

```
grpcurl -plaintext -import-path . -proto ticker.proto -max-time 2 \
  -d '{"symbols":["ACME","UMBRELLA"]}' 127.0.0.1:50051 ticker.v1.Ticker/Watch
```

```
{
  "symbol": "ACME",
  "price": 121.02,
  "average": 121.4
}
…
ERROR:
  Code: DeadlineExceeded
  Message: context deadline exceeded
```

and the server logs `watching ACME, UMBRELLA: cancelled, and its finalizer ran`.

```
grpcurl -plaintext -import-path . -proto ticker.proto \
  -d '{"symbol":"ACME","quantity":3,"limit":120.5} {"symbol":"ACME","quantity":-4,"limit":121}' \
  127.0.0.1:50051 ticker.v1.Ticker/Trade
```

```
{
  "symbol": "ACME",
  "quantity": 3,
  "price": 120.5,
  "position": 3
}
ERROR:
  Code: FailedPrecondition
  Message: you hold 3 ACME, so you can't sell 4
```

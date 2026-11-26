# log-tailer

Three logs read at the same time, merged, parsed into typed records, cut into windows of 50000
lines, summed up, and handed to a slow on-call lookup, four windows at a time. It stops after 12
windows or 2 seconds, whichever comes first, and every log says when it was closed.

```purescript
summaryOfTheLogs = Strom.mergeAll (tailFollowingRotations <$> services)
  # Strom.filterMap parseLine
  # Strom.grouped 50_000
  # Strom.take 12
  # map summarize
  # Strom.mapPar 4 findWhoIsOnCall
  # Strom.tap (printWindow >>> liftIO)
  # Strom.scan addUp { lines: 0, errors: 0, slowest: Milliseconds 0.0 }
  # Strom.timeout (Milliseconds 2000.0)
```

That's the whole program, the rest of [`src/Main.purs`](src/Main.purs) is what each step means.

What it shows:

- **Each log on a goroutine of its own.** `mergeAll` pulls every log on its own goroutine and hands
  the chunks over as they come.
- **Backpressure.** Nobody reads a line before it's needed. The logs are endless, and when `take`
  has its 12 windows (600000 lines), the three of them together have read 617792, only a few
  chunks of 4096 more.
- **Finalizers.** Each log is opened with `Strom.bracketExit`, which is told how the stream ended,
  so the log says whether it was rotated or nobody needs more. They run on the log's own goroutine,
  as soon as the merge is stopped, before the last windows are even printed.
- **Typed errors.** `billing.log` gets rotated every 100000 lines, which fails its stream with
  `rotated :: Service`, a case of the stream's error row, next to `exception`. `catchAll` sees it
  and opens the log again. Anything else still fails the stream, and `described` has to handle
  both cases, or it doesn't compile.
- **`mapPar`.** The on-call lookup takes 40 ms. `mapPar 4` runs four at a time, and the windows
  still come out in order.
- **Chunks.** `summarize` gets a window as an `Array`. The mean and the slowest latency are
  `Array.foldl (+)` and `Array.foldl max` over an `Array Number`, which googoo turns into SIMD loops
  with four vector accumulators, on arm64 (NEON) and WebAssembly. Elsewhere they're plain loops.

## Running it

You need [googoo](https://github.com/i-am-the-slime/iris-go-backend) 0.1.18-go.5 or later and Go
1.27. The `spago.yaml` uses the Go package set, so there's nothing else to set up:

```bash
cd examples/log-tailer
googoo run --go
```

## Sample output

```
opened billing.log
opened api.log
opened auth.log
50000 lines, 883 errors, mostly api.log (paging Ada), mean 187.9 ms, slowest 3550.0 ms
50000 lines, 787 errors, mostly api.log (paging Ada), mean 174.4 ms, slowest 3550.0 ms
50000 lines, 336 errors, mostly billing.log (paging Edsger), mean 108.3 ms, slowest 2900.0 ms
closed billing.log after 100000 lines, billing.log was rotated
opened billing.log
50000 lines, 843 errors, mostly billing.log (paging Edsger), mean 182.0 ms, slowest 3550.0 ms
50000 lines, 1431 errors, mostly billing.log (paging Edsger), mean 268.1 ms, slowest 3550.0 ms
50000 lines, 1012 errors, mostly api.log (paging Ada), mean 206.3 ms, slowest 3550.0 ms
50000 lines, 880 errors, mostly api.log (paging Ada), mean 187.9 ms, slowest 3550.0 ms
50000 lines, 509 errors, mostly api.log (paging Ada), mean 134.1 ms, slowest 3550.0 ms
50000 lines, 370 errors, mostly auth.log (paging Grace), mean 113.3 ms, slowest 3500.0 ms
closed billing.log after 100000 lines, billing.log was rotated
opened billing.log
50000 lines, 1360 errors, mostly billing.log (paging Edsger), mean 257.3 ms, slowest 3550.0 ms
closed api.log after 196608 lines, nobody needs more
closed billing.log after 12288 lines, nobody needs more
closed auth.log after 208896 lines, nobody needs more
50000 lines, 1279 errors, mostly auth.log (paging Grace), mean 245.8 ms, slowest 3550.0 ms
50000 lines, 755 errors, mostly api.log (paging Ada), mean 169.1 ms, slowest 3550.0 ms
In total 600000 lines and 10445 errors, the slowest took 3550.0 ms, 1791045 lines a second
```

The merge mixes the logs differently each run, so your windows will differ.

Change `Strom.take 12` to `Strom.take 1200` and the 2 second `timeout` stops it instead, with the
same `nobody needs more` for each log.

## The same in Go

[`go/main.go`](go/main.go) is the same program written by hand, with the standard library, and it
prints the same lines:

```bash
cd go && go run .
```

| | PureScript | Go |
| --- | --- | --- |
| lines, without blank ones | 142 | 208 |
| of those, imports | 28 | 8 |

The Go version is faster, about 1.5 times on this machine, though it's noisy at a third of a second.

What's longer in Go is the plumbing. Each stage is a goroutine with a channel in and a channel out,
and every send needs a `select` on `ctx.Done()`, or a stage that stops leaves the one before it
stuck. Keeping `mapPar`'s order takes a channel of channels:

```go
results := make(chan chan window, 4)
go func() {
	defer close(results)
	for lines := range windows {
		result := make(chan window, 1)
		select {
		case results <- result:
		case <-ctx.Done():
			return
		}
		go func() {
			w := summarize(lines)
			time.Sleep(40 * time.Millisecond)
			result <- w
		}()
	}
}()
```

and following a rotation is a loop around a function whose deferred closure looks at a named
result to know why it's closing:

```go
for tailOnce(ctx, svc, batches) == errRotated {
}
```

In the PureScript version that's `Strom.mapPar 4 findWhoIsOnCall` and `Strom.catchAll reopenIfRotated`.

## The Go googoo writes

`googoo go out` puts the Go for this program in `out/`. 0.1.18-go.3 made it quite a bit easier to
read than 0.1.18-go.2 did:

| | 0.1.18-go.2 | 0.1.18-go.3 |
| --- | --- | --- |
| `switch either` | 56 | 35 |
| average function name | 40.6 characters | 25.9 characters |
| function names of 30 characters or more | 83 | 33 |
| longest function name | 138 characters | 115 characters |

Failures come back as `(value, err)` now instead of an `Either` to switch on. The long names are
the Strom pipeline's specialisations, which still get long.

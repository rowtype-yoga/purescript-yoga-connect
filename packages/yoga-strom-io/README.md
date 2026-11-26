# yoga-strom-io

Streams in `OmIO` on Go, after ZIO's `ZStream`. You pull them in chunks. They fail with the errors
of your Om row, clean up after themselves, and stop when Go's `context.Context` says so.

```purescript
import Yoga.OmIO.Strom as Strom

Strom.range 0 1000000
  # map (_ * 3)
  # Strom.take 10
  # Strom.forEach send
```

Through googoo that's a loop over chunks of 4096. Each chunk is mapped with plain `for` loops (SIMD
kernels for numbers), and `send` runs in a `for` loop over the chunk, with no closure per element.

## How it works

Running a stream opens it in a `Scope`, which gives a pull. Each pull answers with the next chunk,
says the stream has ended, or fails. Nothing happens between pulls, so a slow reader slows the
stream down.

Once the reader is done, because it has what it wants, it failed, or its context was cancelled,
the scope closes. Goroutines the stream started stop at their next wait, and finalizers run with
how the stream ended: `Ended`, `Failed` or `Interrupted`.

```purescript
Strom.bracket openFile closeFile readLines
  # Strom.take 3
```

`take` closes what it takes from as soon as it has enough, so `closeFile` runs right after the third
line, not when the whole run is over.

## What's in it

- making: `empty`, `succeed`, `fail`, `fromArray`, `fromChunks`, `fromOm`, `fromChan`,
  `fromHandoff`, `fromPull`, `range`, `iterate`, `unfold`, `unfoldM`, `repeat`, `repeatOm`,
  `repeatOmWith` (a `Schedule`), `every`
- resources: `bracket`, `bracketExit`, `acquireRelease`, `ensuring`, `onExit`
- changing: `map`, `mapChunks`, `mapM`, `mapIO`, `filter`, `filterMap`, `tap`, `scan`, `take`,
  `takeWhile`, `takeUntil`, `drop`, `dropWhile`, `concatMap`/`flatMap`/`>>=`, `<>`, `concat`, `zip`,
  `zipWith`, `grouped`, `rechunk`
- failing: `catchAll`, `orElse`, `retry` (a `Schedule`), `timeout`
- at the same time, each side on a goroutine of its own: `buffer`, `merge`, `mergeAll`,
  `mergeWith`, `mergeHaltEither`/`Left`/`Right`, `race`, `raceAll`, `zipPar`, `zipLatest`,
  `mapPar`, `mapParUnordered`, `forkDrain`
- running: `forEach`, `forEachWithin` (a `context.Context`, such as a gRPC call's), `runFold`,
  `runFoldM`, `runCollect`, `runDrain`, `runHead`, `runLast`

Names follow [yoga-om-strom](https://github.com/rowtype-yoga/purescript-yoga-om-strom) where it has
them, so the two can meet later. Unlike yoga-om-strom, `race` keeps the first stream to give
something and stops the other.

## Good to know

- `mapM` and `tap` work a chunk at a time. `Strom.fromArray [ a, b ] # Strom.mapM slowly` finishes
  both elements before passing either on. To pace elements, zip them with `every`.
- Wait with `Strom.every`, `repeatOmWith` and `Strom.sleep scope`, not with go-io's `threadDelay`: those
  stop as soon as the stream is interrupted.
- Keep records out of a stream's own state, such as a `scan`'s. googoo boxes a record that goes
  through go-io's `IORef`, so a `/\` or a data type is faster there.

## Specs and benchmark

```
googoo run --go -p yoga-strom-io-test
googoo run --go -p yoga-strom-io-bench
```

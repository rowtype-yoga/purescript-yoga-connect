// The same program in plain Go, for comparison: three tails merged, parsed, cut into windows of
// 50000 lines, an on-call lookup for 4 windows at a time with the windows kept in order, running
// totals, and a stop after 12 windows or 2 seconds.
package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type service int

const (
	api service = iota
	auth
	billing
)

var names = []string{"api", "auth", "billing"}
var onCall = []string{"Ada", "Grace", "Edsger"}

type logLine struct {
	service service
	isError bool
	latency float64
	route   string
}

type window struct {
	lines, errors int
	mean, slowest float64
	noisiest      service
}

var errRotated = errors.New("rotated")

const chunkSize = 4096

func main() {
	started := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	batches := make(chan []string, 1)
	var tails sync.WaitGroup
	for svc := range names {
		tails.Add(1)
		go func(svc service) {
			defer tails.Done()
			for tailOnce(ctx, svc, batches) == errRotated {
			}
		}(service(svc))
	}
	go func() {
		tails.Wait()
		close(batches)
	}()

	windows := make(chan []logLine)
	go func() {
		defer close(windows)
		var current []logLine
		for batch := range batches {
			for _, text := range batch {
				if line, ok := parseLine(text); ok {
					current = append(current, line)
				}
				if len(current) == 50000 {
					select {
					case windows <- current:
					case <-ctx.Done():
						return
					}
					current = nil
				}
			}
		}
	}()

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

	totalLines, totalErrors, slowest, seen := 0, 0, 0.0, 0
	for result := range results {
		var w window
		select {
		case w = <-result:
		case <-ctx.Done():
		}
		if ctx.Err() != nil {
			break
		}
		fmt.Printf("%d lines, %d errors, mostly %s.log (paging %s), mean %.1f ms, slowest %.1f ms\n",
			w.lines, w.errors, names[w.noisiest], onCall[w.noisiest], w.mean, w.slowest)
		totalLines += w.lines
		totalErrors += w.errors
		slowest = max(slowest, w.slowest)
		if seen++; seen == 12 {
			break
		}
	}
	cancel()
	tails.Wait()
	elapsed := time.Since(started).Seconds()
	fmt.Printf("In total %d lines and %d errors, the slowest took %.1f ms, %d lines a second\n",
		totalLines, totalErrors, slowest, int(float64(totalLines)/elapsed))
}

// tailOnce reads one service's log until it's rotated or nobody needs more.
func tailOnce(ctx context.Context, svc service, out chan<- []string) (err error) {
	file := names[svc] + ".log"
	fmt.Println("opened " + file)
	linesRead := 0
	defer func() {
		reason := "nobody needs more"
		if err == errRotated {
			reason = file + " was rotated"
		}
		fmt.Printf("closed %s after %d lines, %s\n", file, linesRead, reason)
	}()
	for index := 0; ; {
		batch := make([]string, 0, chunkSize)
		for ; len(batch) < chunkSize; index++ {
			if svc == billing && index == 100000 {
				return errRotated
			}
			linesRead = index + 1
			batch = append(batch, lineNumbered(svc, index))
		}
		select {
		case out <- batch:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func parseLine(text string) (logLine, bool) {
	fields := strings.Split(text, " ")
	if len(fields) != 4 {
		return logLine{}, false
	}
	svc := -1
	for i, name := range names {
		if name == fields[0] {
			svc = i
		}
	}
	switch fields[1] {
	case "INFO", "WARN", "ERROR":
	default:
		return logLine{}, false
	}
	latency, err := strconv.ParseFloat(strings.TrimSuffix(fields[2], "ms"), 64)
	if svc < 0 || err != nil || !strings.HasSuffix(fields[2], "ms") {
		return logLine{}, false
	}
	return logLine{service(svc), fields[1] == "ERROR", latency, fields[3]}, true
}

func summarize(lines []logLine) window {
	w := window{lines: len(lines)}
	errorsFrom := make([]int, len(names))
	total := 0.0
	for _, line := range lines {
		total += line.latency
		w.slowest = max(w.slowest, line.latency)
		if line.isError {
			w.errors++
			errorsFrom[line.service]++
		}
	}
	w.mean = total / float64(len(lines))
	for svc, count := range errorsFrom {
		if count > errorsFrom[w.noisiest] {
			w.noisiest = service(svc)
		}
	}
	return w
}

// lineNumbered is line index of a made-up log. Every service has bad stretches, with more errors
// and slower answers, and errors are slowest of all.
func lineNumbered(svc service, index int) string {
	name := names[svc]
	noise := (index%10007*7919 + len(name)) % 10007
	inABadStretch := (index/30000+len(name))%3 == 0
	errorEvery, spread := 150, 200
	if inABadStretch {
		errorEvery, spread = 25, 600
	}
	isError := noise%errorEvery == 0
	level := "INFO"
	if isError {
		level = "ERROR"
	} else if noise%7 == 0 {
		level = "WARN"
	}
	latency := noise % spread
	if isError {
		latency += noise % 3000
	}
	routes := []string{"/checkout", "/login", "/cart", "/search"}
	return name + " " + level + " " + strconv.Itoa(latency) + "ms " + routes[noise%4]
}

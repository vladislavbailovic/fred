package main

import (
	"context"
	"fred/internal"
	"fred/pkg/feed"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opts := parseArgs(os.Args[1:])

	out := ConsolePrinter{}
	exitCode := printSources(ctx, opts, out)
	os.Exit(exitCode)
}

// TODO if something goes wrong here, we hang. Not good.
func printSources(ctx context.Context, opts options, out internal.Printer) int {
	rsp := make(chan *feed.Source)
	done := 0
	urls := getSourceUrls(out)

	if len(urls) == 0 {
		return 1
	}

	for _, url := range urls {
		go getSource(ctx, url, rsp, out)
	}

	for {
		src, ok := <-rsp
		if !ok {
			break
		}

		done++
		if src != nil { // src can be nil pointer
			renderSource(src, opts, out)
		}
		if done == len(urls) {
			close(rsp)
			break
		}
	}

	return 0
}

func getSource(ctx context.Context, raw string, resp chan *feed.Source, out internal.Printer) {
	src := feed.NewSource(raw, out)
	if src != nil { // can be nil pointer
		src.Load(ctx, out)
	}
	resp <- src
}

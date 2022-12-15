package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := ConsolePrinter{}
	result := getSources(ctx, out)

	for idx, r := range result {
		out.Debug("%d) %q posts: %d", idx+1, r.Title, len(r.Articles))
	}
}

func getSources(ctx context.Context, out Printer) []*Source {
	rsp := make(chan *Source)
	result := make([]*Source, 0, len(urls))
	done := 0

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
			result = append(result, src)
		}
		if done == len(urls) {
			close(rsp)
			break
		}
	}
	return result
}

func getSource(ctx context.Context, raw string, resp chan *Source, out Printer) {
	printer := ConsolePrinter{}
	src := NewSource(raw, printer)
	if src != nil { // can be nil pointer
		src.Load(ctx, printer)
	}
	resp <- src
}

var urls []string = []string{
	"https://aws.amazon.com/blogs/architecture/feed/",
	// "https://aws.amazon.com/blogs/compute/feed/",
	// "https://aws.amazon.com/blogs/containers/feed/",
	// "https://aws.amazon.com/blogs/security/feed/",
	// "https://aws.amazon.com/blogs/developer/feed/",
	// "https://aws.amazon.com/blogs/devops/feed/",
	// "https://appliedgo.net/index.xml",
	// "https://dave.cheney.net/feed",
	"https://eli.thegreenplace.net/feeds/all.atom.xml",
	// "https://golangcode.com/index.xml",
	// "https://ieftimov.com/index.xml",
	// "https://research.swtch.com/feed.atom",
	// "https://scene-si.org/index.xml",
	// "https://utcc.utoronto.ca/~cks/space/blog/?atom",
}

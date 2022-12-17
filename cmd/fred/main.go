package main

import (
	"context"
	"fred/internal"
	"fred/pkg/feed"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := ConsolePrinter{}
	printSources(ctx, out)
}

func renderSource(src *feed.Source, printer internal.Printer) {
	var out strings.Builder
	for _, src := range src.Articles {
		out.Grow(len(src.GetTitle()) + len(src.GetLink()) + len(src.GetBrief()) + 7)
		out.WriteByte('[')
		out.WriteString(src.GetTitle())
		out.WriteByte(']')
		out.WriteByte('(')
		out.WriteString(src.GetLink())
		out.WriteByte(')')

		if len(src.GetTopics()) > 0 {
			topics := strings.Join(src.GetTopics(), ":")
			out.Grow(len(topics) + 3)
			out.WriteByte(' ')
			out.WriteByte(':')
			out.WriteString(topics)
			out.WriteByte(':')
		}
		out.WriteByte('\n')

		out.WriteString(src.GetBrief())
		out.WriteByte('\n')

		out.WriteByte('\n')
	}
	printer.Out(out.String())
}

func printSources(ctx context.Context, out internal.Printer) {
	rsp := make(chan *feed.Source)
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
			renderSource(src, out)
		}
		if done == len(urls) {
			close(rsp)
			break
		}
	}
}

func getSource(ctx context.Context, raw string, resp chan *feed.Source, out internal.Printer) {
	src := feed.NewSource(raw, out)
	if src != nil { // can be nil pointer
		src.Load(ctx, out)
	}
	resp <- src
}

var urls []string = []string{
	"https://aws.amazon.com/blogs/architecture/feed/",
	"https://aws.amazon.com/blogs/compute/feed/",
	"https://aws.amazon.com/blogs/containers/feed/",
	"https://aws.amazon.com/blogs/security/feed/",
	"https://aws.amazon.com/blogs/developer/feed/",
	"https://aws.amazon.com/blogs/devops/feed/",
	"https://appliedgo.net/index.xml",
	"https://dave.cheney.net/feed",
	"https://eli.thegreenplace.net/feeds/all.atom.xml",
	"https://golangcode.com/index.xml",
	"https://ieftimov.com/index.xml",
	"https://research.swtch.com/feed.atom",
	"https://scene-si.org/index.xml",
	"https://utcc.utoronto.ca/~cks/space/blog/?atom",
}

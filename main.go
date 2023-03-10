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

	out := internal.NullPrinter{}
	result := getSources(ctx, out)

	render(result, out)
}

func render(result []*feed.Source, printer internal.Printer) {
	var out strings.Builder
	for _, r := range result {
		for _, src := range r.Feed.GetArticles() {
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
	}
	printer.Debug(out.String())
}

func getSources(ctx context.Context, out internal.Printer) []*feed.Source {
	rsp := make(chan *feed.Source)
	result := make([]*feed.Source, 0, len(urls))
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

func getSource(ctx context.Context, raw string, resp chan *feed.Source, out internal.Printer) {
	src := feed.NewSource(raw, out)
	if src != nil { // can be nil pointer
		src.Load(ctx, out)
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

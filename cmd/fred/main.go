package main

import (
	"context"
	"fred/internal"
	"fred/pkg/feed"
	"os"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opts := parseArgs(os.Args[1:])

	out := ConsolePrinter{}
	printSources(ctx, opts, out)
}

func renderSource(src *feed.Source, opts options, printer internal.Printer) {
	if src.Feed == nil { // Feed can be nil pointer
		return
	}
	var out strings.Builder
	before := time.Duration(-1*opts.days*24) * time.Hour
	for _, src := range src.Feed.GetArticles() {
		if opts.days > 0 {
			if src.GetDate().Before(time.Now().Add(before)) {
				continue
			}
		}

		topics := src.GetTopics()
		var topicsStr string
		if len(src.GetTopics()) > 0 {
			topicsStr = strings.Join(topics, ":")
			out.Grow(len(topicsStr) + 3)
		}
		if len(opts.topics) > 0 {
			skip := true
			for _, topic := range opts.topics {
				if strings.Contains(topicsStr, topic) {
					skip = false
					break
				}
			}
			if skip {
				continue
			}
		}

		title := src.GetTitle()
		link := src.GetLink()
		brief := src.GetBrief()

		out.Grow(len(title) + len(link) + len(brief) + 7)
		out.WriteByte('[')
		out.WriteString(title)
		out.WriteByte(']')
		out.WriteByte('(')
		out.WriteString(link)
		out.WriteByte(')')

		if len(topicsStr) > 0 {
			out.WriteByte(' ')
			out.WriteByte(':')
			out.WriteString(topicsStr)
			out.WriteByte(':')
		}
		out.WriteByte('\n')

		out.WriteString(brief)
		out.WriteByte('\n')

		out.WriteByte('\n')
	}
	printer.Out(out.String())
}

func printSources(ctx context.Context, opts options, out internal.Printer) {
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
			renderSource(src, opts, out)
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

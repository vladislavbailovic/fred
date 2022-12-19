package main

import (
	"fred/internal"
	"fred/pkg/feed"
	"strings"
	"time"
)

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

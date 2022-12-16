package main

import (
	"fred/internal"
	"fred/pkg/data"
	"fred/pkg/feed"
	"testing"
)

func Benchmark_Render(b *testing.B) {
	sources := []*feed.Source{
		&feed.Source{},
		&feed.Source{},
	}
	out := &ConsolePrinter{}
	PrintDebug = false

	sources[0].Parse(internal.GetTestFile("atom.xml"))
	sources[1].Parse(internal.GetTestFile("rss.xml"))

	for i := 0; i < b.N; i++ {
		renderSource(sources[0], out)
		renderSource(sources[1], out)
	}
}

func Benchmark_SanitizeCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data.SanitizeCategory("A weird thing (with braces)")
	}
}

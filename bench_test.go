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
	out := &internal.ConsolePrinter{}
	internal.PrintDebug = false

	sources[0].Parse(internal.GetTestFile("atom.xml"))
	sources[1].Parse(internal.GetTestFile("rss.xml"))

	for i := 0; i < b.N; i++ {
		render(sources, out)
	}
}

func Benchmark_SanitizeCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data.SanitizeCategory("A weird thing (with braces)")
	}
}

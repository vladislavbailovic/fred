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
	out := &internal.NullPrinter{}

	sources[0].Parse(internal.GetTestFile("atom.xml"))
	sources[1].Parse(internal.GetTestFile("rss.xml"))

	for i := 0; i < b.N; i++ {
		renderSource(sources[0], options{}, out)
		renderSource(sources[1], options{}, out)
	}
}

func Benchmark_SanitizeCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data.SanitizeCategory("A weird thing (with braces)")
	}
}

func Benchmark_Parse_Atom(b *testing.B) {
	source := &feed.Source{}
	for i := 0; i < b.N; i++ {
		source.Parse(internal.GetTestFile("atom.xml"))

	}
}

func Benchmark_Parse_RSS(b *testing.B) {
	source := &feed.Source{}
	for i := 0; i < b.N; i++ {
		source.Parse(internal.GetTestFile("rss.xml"))
	}
}

func Benchmark_ParseAndRender(b *testing.B) {
	out := &internal.NullPrinter{}

	for i := 0; i < b.N; i++ {
		sources := []*feed.Source{
			&feed.Source{},
			&feed.Source{},
		}
		sources[0].Parse(internal.GetTestFile("atom.xml"))
		sources[1].Parse(internal.GetTestFile("rss.xml"))

		renderSource(sources[0], options{}, out)
		renderSource(sources[1], options{}, out)
	}
}

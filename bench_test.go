package main

import "testing"

func Benchmark_Render(b *testing.B) {
	sources := []*Source{
		&Source{},
		&Source{},
	}
	out := &ConsolePrinter{}
	PrintDebug = false

	sources[0].parse(GetTestFile("atom.xml"))
	sources[1].parse(GetTestFile("rss.xml"))

	for i := 0; i < b.N; i++ {
		render(sources, out)
	}
}

func Benchmark_SanitizeCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sanitizeCategory("A weird thing (with braces)")
	}
}

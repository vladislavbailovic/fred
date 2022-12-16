package main

import "strings"

/// According to https://developer.mozilla.org/en-US/docs/Web/HTML/Block-level_elements
var blockLevelTags []string = []string{
	"address",
	"article",
	"aside",
	"blockquote",
	"details",
	"dialog",
	"div",
	"dd", "dl", "dt",
	"fieldset",
	"figcaption",
	"figure",
	"footer",
	"form",
	"h1", "h2", "h3", "h4", "h5", "h6",
	"header",
	"hgroup",
	"hr",
	"main",
	"nav", "ol", "ul", "li",
	"p",
	"pre",
	"section",
	"table",
}

func isBlockTag(which string) bool {
	for _, tag := range blockLevelTags {
		if tag == which {
			return true
		}
	}
	return false
}

func stripHtmlTags(raw string) string {
	src := []rune(raw)
	var out strings.Builder

	var inTag bool
	var tagStart int
	for i, c := range src {
		if '<' == c {
			inTag = true
			tagStart = i
			continue
		}

		if inTag && ' ' == c {
			if isBlockTag(string(src[tagStart:i])) {
				out.WriteByte('\n')
			}
			continue
		}

		if inTag && '>' == c {
			inTag = false
			continue
		}

		if !inTag {
			out.WriteRune(c)
		}
	}

	return out.String()
}

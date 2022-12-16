package main

import (
	"strings"
	"unicode"
)

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
	for _, c := range src {
		if '<' == c {
			inTag = true
			continue
		}

		if '>' == c {
			inTag = false
			continue
		}

		if !inTag {
			out.WriteRune(c)
		}
	}

	return out.String()
}

func sanitizeCategory(raw string) string {
	return strings.ToLower(
		strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsNumber(r) || '-' == r {
				return r
			}
			return '-'
		}, raw))
}

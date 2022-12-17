package data

import (
	"html"
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

func StripHtmlTags(raw string) string {
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

	return html.UnescapeString(strings.ReplaceAll(out.String(), "_", "`_`"))
}

func SanitizeCategory(raw string) string {
	replaced := false
	var out strings.Builder
	out.Grow(len(raw))
	for _, c := range strings.ToLower(raw) {
		if unicode.IsLetter(c) || unicode.IsNumber(c) || '-' == c {
			out.WriteRune(c)
			if replaced {
				replaced = false
			}
		} else {
			if !replaced {
				out.WriteByte('-')
				replaced = true
			}
		}
	}
	return strings.Trim(out.String(), "-")
}

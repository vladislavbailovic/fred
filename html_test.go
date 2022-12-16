package main

import "testing"

func Test_isBlockTag(t *testing.T) {
	suite := map[string]bool{
		"a":      false,
		"span":   false,
		"strong": false,
		"i":      false,
		"code":   false,
		"div":    true,
		"p":      true,
		"ul":     true,
		"li":     true,
		"h1":     true,
		"h3":     true,
		"h5":     true,
		"pre":    true,
	}
	for tag, expected := range suite {
		t.Run(tag, func(t *testing.T) {
			actual := isBlockTag(tag)
			if actual != expected {
				t.Errorf("expected %q block level result: %v, got %v",
					tag, expected, actual)
			}
		})
	}
}

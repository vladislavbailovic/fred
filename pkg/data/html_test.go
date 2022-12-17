package data

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

func Test_stripHtmlTags(t *testing.T) {
	suite := map[string]struct {
		test string
		want string
	}{
		"no tags": {
			test: "whatever this is a test",
			want: "whatever this is a test",
		},
		"random angle braces": {
			test: "whatever > this < is >> a << test",
			want: "whatever  this  a ",
		},
		"inline tags": {
			test: "<b>whatever</b>! <i>New stuff here</i>",
			want: "whatever! New stuff here",
		},
		"block level tags": {
			test: "<p>paragraph1</p><p>paragraph2</p>",
			want: "paragraph1paragraph2",
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := StripHtmlTags(test.test)
			if test.want != got {
				t.Errorf("expected %q, got %q", test.want, got)
			}
		})
	}
}

func Test_SanitizeCategory(t *testing.T) {
	suite := map[string]string{
		"AWS Lambda":                  "aws-lambda",
		"A weird thing (with braces)": "a-weird-thing-with-braces",
		"Something & The other thing": "something-the-other-thing",
		"test-test":                   "test-test",
	}
	for test, expected := range suite {
		t.Run(test, func(t *testing.T) {
			actual := SanitizeCategory(test)
			if actual != expected {
				t.Errorf("expected %q, got %q", expected, actual)
			}
		})
	}
}

func Test_isStopWord(t *testing.T) {
	suite := map[string]bool{
		"is":   true,
		"your": true,
		"my":   true,
		"go":   false,
		"aws":  false,
	}
	for test, want := range suite {
		got := isStopWord(test)
		if got != want {
			t.Errorf("expected %q stopword status to be %v", test, want)
		}
	}
}

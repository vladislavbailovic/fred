package feed

import (
	"fred/internal"
	"fred/pkg/data"
	"net/url"
	"strings"
	"testing"
	"time"
)

func Test_Parse_Atom(t *testing.T) {
	source := &Source{}
	buffer := internal.GetTestFile("atom.xml")

	if err := source.Parse(buffer); err != nil {
		t.Error(err)
	}
	if len(source.Articles) != 10 {
		t.Errorf("invalid articles count: %d", len(source.Articles))
	}

	for _, a := range source.Articles {
		validateArticle(a, t)
	}
}

func Test_Parse_RSS(t *testing.T) {
	source := &Source{}
	buffer := internal.GetTestFile("rss.xml")

	if err := source.Parse(buffer); err != nil {
		t.Error(err)
	}
	if len(source.Articles) != 20 {
		t.Errorf("invalid articles count: %d", len(source.Articles))
	}

	for _, a := range source.Articles {
		validateArticle(a, t)
	}
}

func validateArticle(a data.Article, t *testing.T) {
	if a.Title == "" {
		t.Error("expected article to have a title")
	}

	if a.Link == "" {
		t.Error("expected article to have a link")
	}
	if lnk, err := url.Parse(a.Link); err != nil {
		t.Errorf("article link should be an URL: %v", err)
	} else if lnk.String() != a.Link {
		t.Errorf("article URL: expected %q, got %q", a.Link, lnk.String())
	}

	if len(a.Topics) == 0 {
		t.Error("expected article topics")
	}

	if a.Brief == "" {
		t.Error("expected brief")
	}
	if strings.Contains(a.Brief, "<") && strings.Contains(a.Brief, ">") {
		t.Error("brief should not contain tags")
		t.Log(a.Title, "\n", a.Brief, "\n")
	}

	if !a.Date.Before(time.Now().Add(-24 * time.Hour)) {
		t.Errorf("expected valid date: %q", a.Date.String())
	}

	if a.Origin.Title == "" {
		t.Error("expected origin title")
	}
	if a.Origin.Link == nil {
		t.Error("expected origin URL")
	}
}

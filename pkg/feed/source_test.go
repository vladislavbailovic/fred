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
	if a.GetTitle() == "" {
		t.Error("expected article to have a title")
	}

	if a.GetLink() == "" {
		t.Error("expected article to have a link")
	}
	if lnk, err := url.Parse(a.GetLink()); err != nil {
		t.Errorf("article link should be an URL: %v", err)
	} else if lnk.String() != a.GetLink() {
		t.Errorf("article URL: expected %q, got %q", a.GetLink(), lnk.String())
	}

	if len(a.GetTopics()) == 0 {
		t.Error("expected article topics")
	}

	if a.GetBrief() == "" {
		t.Error("expected brief")
	}
	if strings.Contains(a.GetBrief(), "<") && strings.Contains(a.GetBrief(), ">") {
		t.Error("brief should not contain tags")
		t.Log(a.GetTitle(), "\n", a.GetBrief(), "\n")
	}

	if !a.GetDate().Before(time.Now().Add(-24 * time.Hour)) {
		t.Errorf("expected valid date: %q", a.GetDate().String())
	}

	if a.GetOrigin().Title == "" {
		t.Error("expected origin title")
	}
	if a.GetOrigin().Link == nil {
		t.Error("expected origin URL")
	}
}

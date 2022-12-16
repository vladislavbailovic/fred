package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func Test_parse_Atom(t *testing.T) {
	source := &Source{}
	buffer := GetTestFile("atom.xml")

	if err := source.parse(buffer); err != nil {
		t.Error(err)
	}
	if len(source.Articles) != 10 {
		t.Errorf("invalid articles count: %d", len(source.Articles))
	}

	for _, a := range source.Articles {
		validateArticle(a, t)
	}
}

func Test_parse_RSS(t *testing.T) {
	source := &Source{}
	buffer := GetTestFile("rss.xml")

	if err := source.parse(buffer); err != nil {
		t.Error(err)
	}
	if len(source.Articles) != 20 {
		t.Errorf("invalid articles count: %d", len(source.Articles))
	}

	for _, a := range source.Articles {
		validateArticle(a, t)
	}
}

func validateArticle(a Article, t *testing.T) {
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

	if !a.Date.ts.Before(time.Now().Add(-24 * time.Hour)) {
		t.Errorf("expected valid date: %q", a.Date.String())
	}

	if a.Origin.Title == "" {
		t.Error("expected origin title")
	}
	if a.Origin.Link == nil {
		t.Error("expected origin URL")
	}
}

func GetTestFilePath(relpath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading test file: %v", err)
		return ""
	}
	pth := filepath.Join(cwd, "testdata", relpath)
	return pth
}

func GetTestFile(relpath string) []byte {
	pth := GetTestFilePath(relpath)
	buffer, err := os.ReadFile(pth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no such test file: %s: %v", pth, err)
		return []byte{}
	}
	return buffer
}

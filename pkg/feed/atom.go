package feed

import (
	"encoding/xml"
	"fred/pkg/data"
	"net/url"
)

type Atom struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Link    Link     `xml:"link"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	XMLName    xml.Name   `xml:"entry"`
	Title      string     `xml:"title"`
	Content    string     `xml:"content"`
	Link       Link       `xml:"link"`
	Summary    string     `xml:"summary"`
	Published  string     `xml:"published"`
	Categories []Category `xml:"category"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

// TODO: use it or lose it
type Raw struct {
	Raw string `xml:",innerxml"`
}

func (x *Atom) GetArticles() []data.Article {
	origin := data.Origin{Title: x.Title}
	if lnk, err := url.Parse(x.Link.Href); err == nil {
		origin.Link = lnk
	}
	articles := make([]data.Article, 0, len(x.Entries))
	for _, e := range x.Entries {
		articles = append(articles, data.Article{
			Title:  e.Title,
			Link:   e.Link.Href,
			Topics: e.GetCategories(),
			Brief:  e.GetBrief(),
			Date:   data.ParseDate(e.Published),
			Origin: origin,
		})
	}
	return articles
}

// TODO: fallback to categories detection from the title
func (e *Entry) GetCategories() []string {
	categories := make([]string, 0, len(e.Categories))
	for _, c := range e.Categories {
		if "" == c.Term {
			continue
		}
		categories = append(categories, data.SanitizeCategory(c.Term))
	}
	return categories
}

func (e *Entry) GetBrief() string {
	var brief string
	if e.Summary != "" {
		brief = e.Summary
	} else {
		brief = e.Content
	}
	// TODO: process HTML entities
	// TODO: strip newlines surrounding the brief
	// TODO: lodash escape somehow
	return data.StripHtmlTags(brief)
}

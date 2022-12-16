package main

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss">"channel"`
	Title   string   `xml:"channel>title"`
	Items   []Item   `xml:"channel>item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Categories  []string `xml:"category"`
	PubDate     string   `xml:"pubDate"`
}

func (e *Item) GetCategories() []string {
	categories := make([]string, 0, len(e.Categories))
	for _, c := range e.Categories {
		if "" == c {
			continue
		}
		categories = append(categories, sanitizeCategory(c))
	}
	return categories
}

func (x *RSS) GetArticles() []Article {
	articles := make([]Article, 0, len(x.Items))
	for _, e := range x.Items {
		articles = append(articles, Article{
			Title:  e.Title,
			Link:   e.Link,
			Topics: e.GetCategories(),
			Brief:  stripHtmlTags(e.Description),
			Date:   ParseDate(e.PubDate),
		})
	}
	return articles
}

type Atom struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
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

type Raw struct {
	Raw string `xml:",innerxml"`
}

func (x *Atom) GetArticles() []Article {
	articles := make([]Article, 0, len(x.Entries))
	for _, e := range x.Entries {
		articles = append(articles, Article{
			Title:  e.Title,
			Link:   e.Link.Href,
			Topics: e.GetCategories(),
			Brief:  e.GetBrief(),
			Date:   ParseDate(e.Published),
		})
	}
	return articles
}

func (e *Entry) GetCategories() []string {
	categories := make([]string, 0, len(e.Categories))
	for _, c := range e.Categories {
		if "" == c.Term {
			continue
		}
		categories = append(categories, sanitizeCategory(c.Term))
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
	return stripHtmlTags(brief)
}

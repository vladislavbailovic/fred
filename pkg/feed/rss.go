package feed

import (
	"encoding/xml"
	"fred/pkg/data"
	"net/url"
)

type RSS struct {
	XMLName xml.Name `xml:"rss">"channel"` // TODO: what the actual fuck?
	Title   string   `xml:"channel>title"`
	Link    string   `xml:"channel>link"`
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
		categories = append(categories, data.SanitizeCategory(c))
	}
	return categories
}

func (x *RSS) GetArticles() []data.Article {
	origin := data.Origin{Title: x.Title}
	if lnk, err := url.Parse(x.Link); err == nil {
		origin.Link = lnk
	}
	articles := make([]data.Article, 0, len(x.Items))
	for _, e := range x.Items {
		articles = append(articles, data.Article{
			Title:  e.Title,
			Link:   e.Link,
			Topics: e.GetCategories(),
			Brief:  data.StripHtmlTags(e.Description),
			Date:   data.ParseDate(e.PubDate),
			Origin: origin,
		})
	}
	return articles
}

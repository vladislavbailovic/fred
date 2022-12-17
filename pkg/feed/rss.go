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
	Items   []*Item  `xml:"channel>item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Categories  []string `xml:"category"`
	PubDate     string   `xml:"pubDate"`

	date   *data.Date  `xml:"-"`
	origin data.Origin `xml:"-"`
}

func (x *Item) GetTitle() string       { return x.Title }
func (x *Item) GetLink() string        { return x.Link }
func (x *Item) GetBrief() string       { return data.StripHtmlTags(x.Description) }
func (x *Item) GetTopics() []string    { return x.GetCategories() }
func (x *Item) GetDate() *data.Date    { return x.date }
func (x *Item) GetOrigin() data.Origin { return x.origin }

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
	for _, a := range x.Items {
		e := a
		e.origin = origin
		e.date = data.ParseDate(e.PubDate)
		articles = append(articles, e)
	}
	return articles
}

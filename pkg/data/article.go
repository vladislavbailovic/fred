package data

import "net/url"

type Article interface {
	GetTitle() string
	GetDate() *Date
	GetLink() string
	GetBrief() string
	GetTopics() []string
	GetOrigin() Origin
}

type Origin struct {
	Title string
	Link  *url.URL
}

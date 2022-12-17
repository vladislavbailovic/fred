package data

import "net/url"

/*
type Article struct {
	Title  string
	Date   Date
	Link   string
	Brief  string
	Topics []string
	Origin Origin
}
*/

type Article interface {
	GetTitle() string
	GetDate() Date
	GetLink() string
	GetBrief() string
	GetTopics() []string
	GetOrigin() Origin
}

type Origin struct {
	Title string
	Link  *url.URL
}

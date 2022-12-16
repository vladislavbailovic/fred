package main

import "net/url"

type Article struct {
	Title  string
	Date   Date
	Link   string
	Brief  string
	Topics []string
	Origin Origin
}

type Origin struct {
	Title string
	Link  *url.URL
}

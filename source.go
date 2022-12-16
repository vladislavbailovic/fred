package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"time"
)

var RequestTimeout time.Duration = 5

type Source struct {
	URL      *url.URL
	Articles []Article
}

func NewSource(rawUrl string, out Printer) *Source {
	url, err := url.Parse(rawUrl)
	if err != nil {
		out.Error(err, "parsing raw URL")
		return nil
	}
	return &Source{URL: url}
}

func (x *Source) Load(ctx context.Context, out Printer) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.URL.String(), nil)
	if err != nil {
		out.Error(err, "creating new request for %s", x.URL)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		out.Error(err, "fetching %s", x.URL)
		return
	}

	if resp.StatusCode != http.StatusOK {
		out.Error(err, "Fetching %s, status not OK (%v)", x.URL, resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		out.Error(err, "reading body %s", x.URL)
		return
	}

	if err := x.parse(buffer); err != nil {
		out.Error(err, "unmarshaling feed XML: %s", x.URL)
	}
}

func (x *Source) parse(buffer []byte) error {
	r := RSS{}
	if rssErr := xml.Unmarshal(buffer, &r); rssErr != nil {
		f := Atom{}
		if atomErr := xml.Unmarshal(buffer, &f); atomErr != nil {
			return atomErr
		} else {
			x.Articles = f.GetArticles()
		}
	} else {
		x.Articles = r.GetArticles()
	}

	return nil
}

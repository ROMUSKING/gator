package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	rss := RSSFeed{}
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return &RSSFeed{}, err
	}
	rss = unescapeString(rss)

	return &rss, err
}

func unescapeString(rss RSSFeed) RSSFeed {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for i, _ := range rss.Channel.Item {

		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
	return rss
}

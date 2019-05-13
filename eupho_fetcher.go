package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/naoki-kishi/feeder"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

// EuphoFetcher is ...
type euphoFetcher struct {
	URL string
}

//NewEuphoFetcher is ...
func NewEuphoFetcher(url string) feeder.Fetcher {
	return &euphoFetcher{URL: url}
}

// Fetch is ...
func (fetcher *euphoFetcher) Fetch() (*feeder.Items, error) {
	resp, err := http.Get(fetcher.URL)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "Failed to get response from anime-eupho.com.")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "Failed to create document from response body.")
	}
	defer resp.Body.Close()

	items := extractNewsToItems(doc)
	return &feeder.Items{items}, nil
}

func extractNewsToItems(doc *goquery.Document) []*feeder.Item {
	items := []*feeder.Item{}

	doc.Find(".newsContentList").Each(func(_ int, s *goquery.Selection) {
		title := s.Find(".newsTitle").Text()
		// 末端などに改行・スペースが含まれることがあるので整形
		title = strings.Trim(title, " ")
		title = strings.Trim(title, "\n")

		link := s.Find(".newsTitle").AttrOr("href", "/news/")

		// 日付を取得後、日付オブジェクトに変換
		dateText := s.Find(".date").Text()

		const DATE_LAYOUT = "2006.01.02"
		date, err := time.Parse(DATE_LAYOUT, dateText)
		if err != nil {
			log.Fatal(err)
			return
		}

		body, err := s.Find(".content").Html()
		if err != nil {
			log.Fatal(err)
			return
		}

		const SITE_URL = "http://anime-eupho.com"
		i := &feeder.Item{
			Title:       title,
			Link:        &feeder.Link{Href: SITE_URL + link}, // 相対リンクを変換
			Created:     &date,
			Id:          link,
			Description: body,
		}

		items = append(items, i)
	})

	return items
}

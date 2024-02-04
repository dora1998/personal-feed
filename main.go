package main

import (
	"log"
	"os"
	"time"

	"github.com/naoki-kishi/feeder"
)

func main() {
	feed := fetch()

	rss, err := feed.ToRSS()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("./feed/rss.xml", rss)

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("./feed/atom.xml", atom)

	json, err := feed.ToJSON()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("./feed/json.json", json)
}

func fetch() *feeder.Feed {
	euphoFetcher := NewEuphoFetcher("http://anime-eupho.com/news/")

	// Fetch data using goroutine.
	items := feeder.Crawl(euphoFetcher)

	feed := &feeder.Feed{
		Title:       "Personal feeds",
		Link:        &feeder.Link{Href: "http://personal-feed.minoru.dev/rss"},
		Description: "Personal feeds.",
		Author: &feeder.Author{
			Name:  "Minoru Takeuchi",
			Email: "me@minoru.dev"},
		Created: time.Now(),
		Items:   *items,
	}

	return feed
}

func writeFeed(path string, body string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(([]byte)(body))

	return nil
}

package main

import (
	"fmt"
	"nuntium/feed"
	"nuntium/formatter"
	"time"

	"github.com/go-co-op/gocron"
)

var task = func() {
	for feedTitle, feedURL := range feed.GetURLs() {
		news, err := feed.Fetch(feedURL)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(news) == 0 {
			continue
		}

		fmt.Println(formatter.FormatFeedNews(feedTitle, news))
	}
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(task)
	s.StartBlocking()
}

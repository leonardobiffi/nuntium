package main

import (
	"fmt"
	"nuntium/feed"
	"nuntium/formatter"
	"nuntium/notifier"
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

		if err = notifier.Send("🤖 Nuntium", formatter.FormatFeedNews(feedTitle, news)); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(task)
	s.StartBlocking()
}

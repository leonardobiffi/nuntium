package main

import (
	"fmt"
	"nuntium/config"
	"nuntium/feed"
	"nuntium/formatter"
	"nuntium/notifier"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

var version = "dev"

// Create a new instance of the logger
var log = logrus.New()

func init() {
	// Log as logfmt instead of the default ASCII formatter.
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

var task = func(cfg *config.Config) {
	for feedTitle, feedURL := range cfg.FeedURLs {
		log.Info("Fetching news from ", feedTitle)
		news, err := feed.Fetch(feedURL, cfg.Schedule)
		if err != nil {
			log.Error(err)
			continue
		}

		if len(news) == 0 {
			log.Info("news not found")
		}

		for _, n := range news {
			log.Info(fmt.Sprintf("Title: %s Time: %s", n.Title, n.Time))
			if cfg.SkipNotification {
				fmt.Println(formatter.FormatFeedNews(feedTitle, n))
				continue
			}

			if err = notifier.Send(formatter.FormatFeedNews(feedTitle, n)); err != nil {
				log.Error(err)
			}
		}
	}
}

func main() {
	log.Info("Nuntium version: ", version)

	cfg, err := config.New()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("Config loaded...")
	for feedTitle := range cfg.FeedURLs {
		log.Info(fmt.Sprintf("Feed: %s", feedTitle))
	}

	if !cfg.SkipNotification {
		notifier.Init()
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(int(cfg.Schedule)).Hours().Do(task, cfg)
	s.StartBlocking()
}

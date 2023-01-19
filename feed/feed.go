package feed

import (
	"fmt"
	"nuntium/entities"
	"nuntium/logger"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v3"
)

// Create a new instance of the logger
var log = logger.New()

type Feeds struct {
	Feeds []Items `yaml:"feeds"`
}

type Items struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func GetURLs(filename string) (feeds map[string]string, err error) {
	feedURLs := make(map[string]string)
	var file []byte

	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		data := os.Getenv("CONFIG_VALUE")
		if data == "" {
			return nil, fmt.Errorf("config file or config value not found")
		}

		file = []byte(data)
		log.Info("Using config value from environment variable")
	} else {
		file, err = os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		log.Info("Using config file")
	}

	var content Feeds
	err = yaml.Unmarshal(file, &content)
	if err != nil {
		return
	}

	for _, feed := range content.Feeds {
		feedURLs[feed.Name] = feed.URL
	}

	return feedURLs, nil
}

func Fetch(feedURL string, diffHours float64) (news []entities.News, err error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		return
	}

	now := time.Now()
	for _, item := range feed.Items {
		var diff time.Duration
		var timePublished time.Time
		if item.UpdatedParsed != nil {
			diff = now.Sub(*item.UpdatedParsed)
			timePublished = *item.UpdatedParsed
		} else {
			diff = now.Sub(*item.PublishedParsed)
			timePublished = *item.PublishedParsed
		}

		if diff.Hours() <= diffHours {
			n := entities.News{
				Title: item.Title,
				Link:  item.Link,
				Time:  timePublished.Format("2006-01-02 15:04:05"),
			}

			news = append(news, n)
		}

		if diff.Hours() > diffHours {
			break
		}
	}

	return
}

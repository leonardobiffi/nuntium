package feed

import (
	"io/ioutil"
	"nuntium/entities"
	"time"

	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v3"
)

type Feeds struct {
	Feeds []Items `yaml:"feeds"`
}

type Items struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func GetURLs(filename string) (feeds map[string]string, err error) {
	feedURLs := make(map[string]string)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
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

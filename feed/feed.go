package feed

import (
	"nuntium/entities"
	"time"

	"github.com/mmcdole/gofeed"
)

func GetURLs() map[string]string {
	feedURLs := make(map[string]string)

	feedURLs["AWS"] = "https://aws.amazon.com/blogs/aws/feed"
	feedURLs["Hashicorp"] = "https://www.hashicorp.com/blog/feed.xml"
	feedURLs["Golang Weekly"] = "https://cprss.s3.amazonaws.com/golangweekly.com.xml"

	return feedURLs
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

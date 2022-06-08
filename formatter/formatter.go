package formatter

import (
	"fmt"
	"nuntium/entities"
	"strings"
)

func FormatFeedNews(title string, news []entities.News) string {
	var body []string
	body = append(body, fmt.Sprintf("📰 %s", title))
	body = append(body, "\n")

	for _, n := range news {
		body = append(body, formatNews(n))
		body = append(body, "\n")
	}

	return strings.Join(body, "\n")
}

func formatNews(news entities.News) string {
	var body []string
	body = append(body, "📌 "+news.Title)
	body = append(body, "🔗 "+news.Link)
	body = append(body, "🕠 "+news.Time)

	return strings.Join(body, "\n")
}

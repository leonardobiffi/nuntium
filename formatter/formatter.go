package formatter

import (
	"fmt"
	"nuntium/entities"
	"strings"
)

func FormatFeedNews(title string, news entities.News) (subject, message string) {
	var body []string
	body = append(body, "\n")
	body = append(body, formatNews(news))
	body = append(body, "\n")

	subject = fmt.Sprintf("📰 %s", title)
	message = strings.Join(body, "\n")

	return
}

func formatNews(news entities.News) string {
	var body []string
	body = append(body, "📌 "+news.Title)
	body = append(body, "🔗 "+news.Link)
	body = append(body, "🕠 "+news.Time)

	return strings.Join(body, "\n")
}

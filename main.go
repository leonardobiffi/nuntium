package main

import (
	"fmt"
	"nuntium/feed"
	"nuntium/formatter"
	"nuntium/notifier"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

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

	// Set the log level
	debugLogLevel, _ := strconv.ParseBool(os.Getenv("DEBUG_LOG_LEVEL"))
	if debugLogLevel {
		log.SetLevel(logrus.DebugLevel)
	}
}

var task = func() {
	notifier.Init()

	skipNotification, _ := strconv.ParseBool(os.Getenv("SKIP_NOTIFICATION"))
	scheduleHours := getSchedule()

	for feedTitle, feedURL := range feed.GetURLs() {
		log.Info("Fetching news from ", feedTitle)
		news, err := feed.Fetch(feedURL, scheduleHours)
		if err != nil {
			log.Error(err)
			continue
		}

		if len(news) == 0 {
			log.Info("news not found")
		}

		for _, n := range news {
			log.Info(fmt.Sprintf("Title: %s Time: %s", n.Title, n.Time))
			if skipNotification {
				log.Debug(formatter.FormatFeedNews(feedTitle, n))
				continue
			}

			if err = notifier.Send(formatter.FormatFeedNews(feedTitle, n)); err != nil {
				log.Error(err)
			}
		}
	}
}

func getSchedule() float64 {
	// scheduleHours is the number of hours to schedule the task
	var scheduleHours float64 = 1

	// Get Schedule Hours from Environment Variable
	scheduleHoursEnv, _ := strconv.ParseFloat(os.Getenv("SCHEDULE_HOURS"), 64)
	if scheduleHoursEnv != 0 {
		scheduleHours = scheduleHoursEnv
	}

	return scheduleHours
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	scheduleHours := getSchedule()
	log.Info("Schedule Hours: ", scheduleHours)

	s.Every(int(scheduleHours)).Minutes().Do(task)
	s.StartBlocking()
}

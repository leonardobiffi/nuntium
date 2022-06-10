package config

import (
	"nuntium/feed"
	"os"
	"strconv"
)

type Config struct {
	FeedURLs         map[string]string
	Schedule         float64
	SkipNotification bool
}

func New() (*Config, error) {
	config := &Config{
		FeedURLs:         make(map[string]string),
		Schedule:         1,
		SkipNotification: false,
	}

	// Get Schedule Hours from Environment Variable
	scheduleHoursEnv, _ := strconv.ParseFloat(os.Getenv("SCHEDULE_HOURS"), 64)
	if scheduleHoursEnv != 0 {
		config.Schedule = scheduleHoursEnv
	}

	// Get skip notification from Environment Variable
	skipNotificationEnv, _ := strconv.ParseBool(os.Getenv("SKIP_NOTIFICATION"))
	if skipNotificationEnv {
		config.SkipNotification = skipNotificationEnv
	}

	// Get Feed URLs from config file
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yml"
	}

	feedURLs, err := feed.GetURLs(configFile)
	if err != nil {
		return nil, err
	}
	config.FeedURLs = feedURLs

	return config, nil
}

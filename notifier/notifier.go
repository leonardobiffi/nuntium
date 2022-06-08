package notifier

import (
	"context"
	"os"
	"strconv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func Init() {
	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, err := telegram.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		panic(err)
	}

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	chatID, _ := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	telegramService.AddReceivers(chatID)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notify.UseServices(telegramService)
}

func Send(subject, message string) error {
	return notify.Send(context.Background(), subject, message)
}

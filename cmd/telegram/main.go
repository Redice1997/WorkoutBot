package main

import (
	"os"
	"time"
	"workout_bot/pkg/bot/telegram"
)

func main() {
	bot := telegram.New(telegram.TelegramBotConfig{
		Token:          os.Getenv("TELEGRAM_API_TOKEN"),
		PollingTimeout: 30 * time.Second,
		MessageOffset:  0,
		DebugMode:      false,
	})
	bot.Run()
}

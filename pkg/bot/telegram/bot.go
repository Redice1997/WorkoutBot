package telegram

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"workout_bot/pkg/bot"
	"workout_bot/pkg/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	api          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
}

func New(cfg TelegramBotConfig) bot.Bot {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(err)
	}
	api.Debug = cfg.DebugMode

	updateConfig := tgbotapi.NewUpdate(cfg.MessageOffset)
	updateConfig.Timeout = int(cfg.PollingTimeout.Seconds())

	return &TelegramBot{
		api:          api,
		updateConfig: updateConfig,
	}
}

type TelegramBotConfig struct {
	Token          string
	PollingTimeout time.Duration
	MessageOffset  int
	DebugMode      bool
}

func (b *TelegramBot) Run() {

	updates := b.api.GetUpdatesChan(b.updateConfig)

	for update := range updates {

		response, err := b.HandleUpdate(update)
		if err != nil {
			panic(err)
		}

		msg := tgbotapi.NewMessage(update.FromChat().ID, response.Text)
		if response.InlineKeyboard != nil {
			markup := make([][]tgbotapi.InlineKeyboardButton, 0)
			for _, row := range response.InlineKeyboard {
				line := make([]tgbotapi.InlineKeyboardButton, 0)
				for _, btn := range row {
					line = append(line, tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Action.String()))
				}
				markup = append(markup, line)
			}
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(markup...)
		}
		if _, err := b.api.Send(msg); err != nil {
			panic(err)
		}
	}
}

func (b *TelegramBot) HandleUpdate(update tgbotapi.Update) (*models.Message, error) {
	switch {
	case update.Message != nil:
		return models.NewStartAction(strconv.FormatInt(update.SentFrom().ID, 10)).Invoke(context.Background(), b)
	case update.CallbackQuery != nil:
		return b.HandleAction(update.CallbackQuery.Data)
	}
	return nil, fmt.Errorf("handle update error")
}

func (b *TelegramBot) HandleAction(actionData string) (msg *models.Message, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("handle action error: %w", err)
		}
	}()
	action, err := models.ParseAction(actionData)
	if err != nil {
		return nil, err
	}
	return action.Invoke(context.Background(), b)
}

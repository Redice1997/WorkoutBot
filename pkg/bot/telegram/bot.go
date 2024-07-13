package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
	"workout_bot/pkg/bot"
	"workout_bot/pkg/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	api          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	storage      models.Storage
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

		if update.SentFrom().IsBot {
			continue
		}

		msg, err := b.HandleUpdate(context.Background(), update)
		if err != nil {
			slog.Error(err.Error())
		}

		if _, err := b.api.Send(msg); err != nil {
			slog.Error(err.Error())
		}
	}
}

func (b *TelegramBot) HandleUpdate(ctx context.Context, update tgbotapi.Update) (msg tgbotapi.MessageConfig, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("handle update error: %w", err)
		}
	}()

	var externalID = strconv.FormatInt(update.SentFrom().ID, 10)
	var userName = update.SentFrom().UserName
	var language models.Language
	if strings.ToLower(update.SentFrom().LanguageCode) == "ru" {
		language = models.RU
	} else {
		language = models.EN
	}

	userCtx := models.NewUserContext(ctx, externalID, userName, language)

	var response *models.Message

	switch {
	case update.Message != nil:
		response, err = b.HandleMessage(userCtx, update.Message)
	case update.CallbackQuery != nil:
		response, err = b.HandleAction(userCtx, update.CallbackQuery.Data)
	}

	msg = tgbotapi.NewMessage(update.FromChat().ID, response.Text)
	if response.IsHtml {
		msg.ParseMode = "html"
	}

	if response.Keyboard != nil {
		if response.IsInlineKeyboard {
			msg.ReplyMarkup = mapInlineKeyboard(response.Keyboard)
		} else {
			msg.ReplyMarkup = mapReplyKeyboard(response.Keyboard)
		}
	}

	return msg, nil
}

const (
	StartCmd    = "start"
	WorkoutsCmd = "workouts"
)

func (b *TelegramBot) HandleMessage(ctx models.UserContext, msg *tgbotapi.Message) (resp *models.Message, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("handle message error: %w", err)
		}
	}()

	if msg.IsCommand() {
		switch msg.Command() {
		case StartCmd:
			return models.NewStartAction().Invoke(ctx, b.storage)
		}
	}

	action, exist := b.storage.GetAction(ctx, strconv.FormatInt(msg.From.ID, 10))
	if exist {
		err := action.Update(msg.Text)
		if err != nil {
			var lang models.Language
			if msg.From.LanguageCode == "ru" {
				lang = models.RU
			} else {
				lang = models.EN
			}
			text, err := b.storage.GetReplica(models.WrongFormatReplica, lang)
			if err != nil {
				return nil, err
			}

			return &models.Message{Text: text}, nil
		}
		return action.Invoke(ctx, b.storage)
	}

	return models.NewStartAction().Invoke(ctx, b.storage)
}

func (b *TelegramBot) HandleAction(ctx models.UserContext, actionData string) (msg *models.Message, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("handle action error: %w", err)
		}
	}()
	action, err := models.ParseAction(actionData)
	if err != nil {
		return nil, err
	}
	return action.Invoke(ctx, b.storage)
}

func notImplementError() (*models.Message, error) {
	return nil, fmt.Errorf("not implemented")
}

func mapInlineKeyboard(keyboard models.Keyboard) tgbotapi.InlineKeyboardMarkup {
	markup := make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, row := range keyboard {
		line := make([]tgbotapi.InlineKeyboardButton, 0)
		for _, btn := range row {
			line = append(line, tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Action.String()))
		}
		markup = append(markup, line)
	}

	return tgbotapi.NewInlineKeyboardMarkup(markup...)
}

func mapReplyKeyboard(keyboard models.Keyboard) tgbotapi.ReplyKeyboardMarkup {
	markup := make([][]tgbotapi.KeyboardButton, 0)

	for _, row := range keyboard {
		line := make([]tgbotapi.KeyboardButton, 0)
		for _, btn := range row {
			line = append(line, tgbotapi.NewKeyboardButton(btn.Text))
		}
		markup = append(markup, line)
	}

	return tgbotapi.NewReplyKeyboard(markup...)
}

package models

type Message struct {
	Text           string
	InlineKeyboard Keyboard
	ReplyKeyboard  Keyboard
}

package models

type Message struct {
	Text             string
	Keyboard         Keyboard
	IsInlineKeyboard bool
	IsHtml           bool
}

func NewMessage(text string, keyboard Keyboard, isInlineKeyboard bool, isHtml bool) *Message {
	return &Message{text, keyboard, isInlineKeyboard, isHtml}
}

type Keyboard [][]Button

type Button struct {
	Text   string
	Action Action
}

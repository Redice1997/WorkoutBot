package models

type Message struct {
	Text             string
	IsHtml           bool
	Keyboard         Keyboard
	IsInlineKeyboard bool
}

func NewMessage(text string, isHtml bool) *Message {
	return &Message{text, isHtml, nil, false}
}

func NewMessageWithKeyboard(text string, isHtml bool, keyboard Keyboard, isInlineKeyboard bool) *Message {
	return &Message{text, isHtml, keyboard, isInlineKeyboard}
}

type Keyboard [][]Button

type Button struct {
	Text   string
	Action Action
}

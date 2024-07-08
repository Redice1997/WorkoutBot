package models

type Button struct {
	Text   string
	Action Action
}

type Keyboard [][]Button

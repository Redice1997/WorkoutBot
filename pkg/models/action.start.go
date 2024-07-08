package models

import (
	"context"
	"fmt"
)

type StartAction struct {
	t          ActionType
	externalID string
}

func NewStartAction(externalID string) *StartAction {
	return &StartAction{
		t:          StartActionType,
		externalID: externalID,
	}
}

func (a *StartAction) Type() ActionType {
	return a.t
}

func (a *StartAction) String() string {
	return fmt.Sprintf("%d{ %s }", a.t, a.externalID)
}

func parseStartActionStruct(s string) (*StartAction, error) {
	var externalID string = s
	return NewStartAction(externalID), nil
}

func (a *StartAction) Invoke(ctx context.Context, handler ActionHandler) (*Message, error) {

	text := `Здарова, дрищ. Я бот, который поможет набрать тебе в силовых и жать, как минимум, сотку от груди. 
	Моя программа сонована на методике из этого <a src="https://www.youtube.com/watch?v=bWnKmO0aj3c">видео</a>.

	Я создам тебе программу тренировок по-умолчанию с целью на 150кг в жиме лёжа, однако, ты сможешь в любой момент настроить параметры под себя.

	Команды:

		/workouts - посмотреть список тренировок
		/start_workout - начать тренировку
		/progress - посмотреть свой прогресс`

	keyboard := Keyboard{{{Text: "/start", Action: NewStartAction(a.externalID)}}}

	return &Message{Text: text, InlineKeyboard: keyboard}, nil
}

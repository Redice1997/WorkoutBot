package models

import (
	"fmt"
)

type StartAction struct {
	t ActionType
}

func NewStartAction() *StartAction {
	return &StartAction{
		t: StartActionType,
	}
}

func (a *StartAction) Type() ActionType {
	return a.t
}

func (a *StartAction) String() string {
	return fmt.Sprintf("%d{ }", a.t)
}

func parseStartAction(s string) (*StartAction, error) {
	return NewStartAction(), nil
}

func (a *StartAction) Invoke(ctx UserContext, storage Storage) (*Message, error) {

	user, err := storage.GetUserByExternalID(ctx, ctx.ExternalID())
	if err != nil {
		return nil, err
	}

	text, err := storage.GetReplica(StartActionReplica, ctx.Language())
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = storage.CreateUser(ctx, &User{
			ID:         0,
			Username:   ctx.Username(),
			ExternalID: ctx.ExternalID(),
			Role:       0,
			Language:   ctx.Language(),
		})
		if err != nil {
			return nil, err
		}
	}

	keyboard := Keyboard{{{Text: text, Action: NewSelectWorkoutsAction(user.ID)}}}

	return NewMessage(text, keyboard, true, true), nil
}

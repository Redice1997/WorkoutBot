package models

import (
	"fmt"
)

type Action interface {
	Type() ActionType
	String() string
	Invoke(ctx UserContext, handler Storage) (*Message, error)
}

type UpdatableAction interface {
	Action
	UpdateParameters(params string) error
}

type ActionType int

const (
	StartActionType ActionType = iota
	SelectWorkoutsActionType
	SelectWorkoutActionType
	StartWorkoutActionType
	ContinueWorkoutActionType
	FinishWorkoutActionType
)

type ActionHandler interface {
}

func ParseAction(input string) (action Action, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse action error: %w", err)
		}
	}()
	var t ActionType
	var s string
	_, err = fmt.Sscanf(input, "%d{ %s }", &t, &s)
	fmt.Printf("'%s'\n", s)
	if err != nil {
		return nil, err
	}
	switch t {
	case StartActionType:
		return parseStartAction(s)
	case SelectWorkoutsActionType:
		return parseSelectWorkoutsAction(s)
	}

	return nil, fmt.Errorf("action is not implemented")
}

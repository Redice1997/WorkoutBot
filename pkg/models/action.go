package models

import (
	"context"
	"fmt"
)

type Action interface {
	Type() ActionType
	String() string
	Invoke(ctx context.Context, handler ActionHandler) (*Message, error)
}

type ActionType int

const (
	StartActionType ActionType = iota
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
		return parseStartActionStruct(s)
	}

	return nil, fmt.Errorf("action is not implemented")
}

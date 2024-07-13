package models

import (
	"fmt"
)

type SelectWorkoutsAction struct {
	t      ActionType
	userID int64
}

func NewSelectWorkoutsAction(userID int64) *SelectWorkoutsAction {
	return &SelectWorkoutsAction{
		t:      SelectWorkoutActionType,
		userID: userID,
	}
}

func (a *SelectWorkoutsAction) Type() ActionType {
	return a.t
}

func (a *SelectWorkoutsAction) String() string {
	return fmt.Sprintf("%d{ %d }", a.t, a.userID)
}

func parseSelectWorkoutsAction(s string) (*SelectWorkoutsAction, error) {
	var userID int64
	_, err := fmt.Sscanf(s, "%d", &userID)
	if err != nil {
		return nil, err
	}

	return NewSelectWorkoutsAction(userID), nil
}

func (a *SelectWorkoutsAction) Invoke(ctx UserContext, storage Storage) (*Message, error) {

	workouts, err := storage.GetWorkoutPrograms(ctx, a.userID)
	if err != nil {
		return nil, err
	}
	if len(workouts) == 0 {
		w, err := storage.CreateWorkoutProgram(ctx, NewDefaultProgram(a.userID, "Bench Press"))
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, *w)
	}

	text, err := storage.GetReplica(SelectWorkoutsActionReplica, ctx.Language())
	if err != nil {
		return nil, err
	}
	createNewText, err := storage.GetReplica(BtnCreateWorkoutReplica, ctx.Language())
	if err != nil {
		return nil, err
	}
	backText, err := storage.GetReplica(BtnReturnReplica, ctx.Language())
	if err != nil {
		return nil, err
	}

	var keyboard = make(Keyboard, 0, 2)
	for _, w := range workouts {
		keyboard = append(keyboard, []Button{{Text: w.Name /*Action: NewSelectWorkoutAction(w.ID)*/}})
	}
	keyboard = append(keyboard, []Button{{Text: createNewText}})
	keyboard = append(keyboard, []Button{{Text: backText, Action: NewStartAction()}})
	return NewMessageWithKeyboard(text, false, keyboard, true), nil
}

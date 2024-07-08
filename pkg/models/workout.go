package models

import (
	"fmt"
	"time"
)

type WorkoutProgram struct {
	ID         int64
	Name       string
	UserID     int64
	CreatedAt  time.Time
	parameters *WorkoutProgramParameters
}

func NewDefaultProgram(userID int64, name string) *WorkoutProgram {
	return &WorkoutProgram{
		Name:       name,
		UserID:     userID,
		CreatedAt:  time.Now(),
		parameters: DefaultParameters(),
	}
}

type WorkoutProgramParameters struct {
	Unit      string
	Base      int
	Target    int
	Step      int
	SetsCycle []int
	RepsCycle []int
	Stage     int
}

func DefaultParameters() *WorkoutProgramParameters {
	return &WorkoutProgramParameters{
		Unit:      "кг",
		Base:      50,
		Target:    150,
		Step:      5,
		SetsCycle: []int{4, 5, 6, 6},
		RepsCycle: []int{8, 6, 4, 2},
		Stage:     1,
	}
}

type Workout struct {
	Name      string
	ProgramID int64
	Current   int
	Sets      int
	Reps      int
	CreatedAt time.Time
}

func (w *WorkoutProgram) Parameters() WorkoutProgramParameters {
	return WorkoutProgramParameters{
		Unit:      w.parameters.Unit,
		Base:      w.parameters.Base,
		Target:    w.parameters.Target,
		Step:      w.parameters.Step,
		RepsCycle: append(make([]int, 0, len(w.parameters.RepsCycle)), w.parameters.RepsCycle...),
		SetsCycle: append(make([]int, 0, len(w.parameters.SetsCycle)), w.parameters.SetsCycle...),
		Stage:     w.parameters.Stage,
	}
}

func (w *WorkoutProgram) Workout() Workout {
	params := w.Parameters()
	return Workout{
		Name:      w.Name,
		ProgramID: w.ID,
		Current:   params.Base + params.Step*(params.Stage-1),
		Sets:      params.SetsCycle[params.Stage-1],
		Reps:      params.RepsCycle[params.Stage-1],
		CreatedAt: time.Now(),
	}
}

func (w *WorkoutProgram) SetUnit(unit string) error {
	w.parameters.Unit = unit
	return nil
}

func (w *WorkoutProgram) SetBase(base int) error {
	if base <= 0 {
		return fmt.Errorf("base must be grater than 0")
	}
	w.parameters.Base = base
	return nil
}

func (w *WorkoutProgram) SetTarget(target int) error {
	if target <= w.parameters.Base {
		return fmt.Errorf("target must be greater than base")
	}
	w.parameters.Target = target
	return nil
}

func (w *WorkoutProgram) SetStep(step int) error {
	if step <= 0 {
		return fmt.Errorf("step must be greater than 0")
	}
	w.parameters.Step = step
	return nil
}

func (w *WorkoutProgram) SetSetsCycle(cycle []int) error {
	err := validateCycle(cycle)
	if err != nil {
		return err
	}
	length := len(cycle)
	sets := append(make([]int, 0, length), cycle...)
	reps := makeLength(w.parameters.RepsCycle, length)
	if w.parameters.Stage > length {
		w.parameters.Stage = length
	}
	w.parameters.SetsCycle = sets
	w.parameters.RepsCycle = reps
	return nil
}

func (w *WorkoutProgram) SetRepsCycle(cycle []int) error {
	err := validateCycle(cycle)
	if err != nil {
		return err
	}
	length := len(cycle)
	reps := append(make([]int, 0, length), cycle...)
	sets := makeLength(w.parameters.SetsCycle, length)
	if w.parameters.Stage > length {
		w.parameters.Stage = length
	}
	w.parameters.SetsCycle = sets
	w.parameters.RepsCycle = reps
	return nil
}

func (w *WorkoutProgram) SetStage(stage int) error {
	if stage < 1 || stage > len(w.parameters.RepsCycle) {
		return fmt.Errorf("stage out of range of cycle")
	}
	w.parameters.Stage = stage
	return nil
}

func validateCycle(cycle []int) error {
	if len(cycle) == 0 {
		return fmt.Errorf("cycle must be not empty")
	}
	for _, item := range cycle {
		if item <= 0 {
			return fmt.Errorf("value in cycle must greater than 0")
		}
	}
	return nil
}

func makeLength(arr []int, length int) []int {
	for len(arr) < length {
		arr = append(arr, 1)
	}
	if len(arr) > length {
		arr = arr[:length]
	}
	return arr
}

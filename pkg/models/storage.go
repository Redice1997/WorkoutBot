package models

import "context"

type Storage interface {
	ActionStorage
	UserStorage
	WorkoutStorage
	ReplicaStorage
}

type ActionStorage interface {
	SetAction(ctx context.Context, externalID string, action UpdatableAction)
	GetAction(ctx context.Context, externalID string) (UpdatableAction, bool)
	ResetAction(ctx context.Context, externalID string)
}

type UserStorage interface {
	GetUserByID(ctx context.Context, userID int64) (*User, error)
	GetUserByExternalID(ctx context.Context, externalID string) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
}

type WorkoutStorage interface {
	CreateWorkoutProgram(ctx context.Context, wp *WorkoutProgram) (*WorkoutProgram, error)
	GetWorkoutProgram(ctx context.Context, workoutprogramID int64) (*WorkoutProgram, error)
	GetWorkoutPrograms(ctx context.Context, userID int64) ([]WorkoutProgram, error)
	UpdateWorkoutProgram(ctx context.Context, wp *WorkoutProgram)
	DeleteWorkoutProgram(ctx context.Context, workoutprogramID int64) error
}

type ReplicaStorage interface {
	GetReplica(label Replica, lang Language) (string, error)
}

type Language string

const (
	RU Language = "RU"
	EN Language = "EN"
)

type Replica int

const (
	StartActionReplica Replica = iota
	WrongFormatReplica
	SelectWorkoutsActionReplica
	BtnReturnReplica
	BtnCreateWorkoutReplica
	SelectWorkoutActionReplica
)

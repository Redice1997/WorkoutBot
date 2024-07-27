package storage

import (
	"context"
	"fmt"
	"workout_bot/pkg/cache"
	"workout_bot/pkg/models"
)

type cacheStorage struct {
	userCache   cache.UserCache
	actionCache cache.ActionCache
	users       []models.User
	workouts    []*models.WorkoutProgram
	replics     models.ReplicaStorage
}

func New() models.Storage {
	return &cacheStorage{
		userCache:   cache.NewUserCache(15, 5),
		actionCache: cache.NewActionCache(15, 5),
		users:       make([]models.User, 0),
		workouts:    make([]*models.WorkoutProgram, 0),
		replics:     NewReplicaStorage(),
	}
}

func (s *cacheStorage) SetAction(ctx context.Context, externalID string, action models.UpdatableAction) {
	s.actionCache.Set(externalID, action)
}

func (s *cacheStorage) GetAction(ctx context.Context, externalID string) (models.UpdatableAction, bool) {
	return s.actionCache.Get(externalID)
}

func (s *cacheStorage) ResetAction(ctx context.Context, externalID string) {
	s.actionCache.Delete(externalID)
}

func (s *cacheStorage) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	id := int(userID)
	if id < 0 || id > len(s.users) {
		return nil, nil
	}
	user := s.users[id]
	return &user, nil
}

func (s *cacheStorage) GetUserByExternalID(ctx context.Context, externalID string) (*models.User, error) {
	if user, ok := s.userCache.Get(externalID); ok {
		s.userCache.Set(externalID, user)
		return user, nil
	}

	for _, user := range s.users {
		if user.ExternalID == externalID {
			s.userCache.Set(externalID, &user)
			return &user, nil
		}
	}

	return nil, nil
}

func (s *cacheStorage) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = int64(len(s.users))
	s.users = append(s.users, *user)
	return &s.users[len(s.users)-1], nil
}

func (s *cacheStorage) CreateWorkoutProgram(ctx context.Context, wp *models.WorkoutProgram) (*models.WorkoutProgram, error) {
	workout := *wp
	workout.ID = int64(len(s.workouts))
	s.workouts = append(s.workouts, &workout)
	return s.workouts[len(s.workouts)-1], nil
}

func (s *cacheStorage) GetWorkoutProgram(ctx context.Context, workoutprogramID int64) (*models.WorkoutProgram, error) {
	id := int(workoutprogramID)
	if id < 0 || id > len(s.workouts) {
		return nil, nil
	}
	return s.workouts[id], nil
}

func (s *cacheStorage) GetWorkoutPrograms(ctx context.Context, userID int64) ([]models.WorkoutProgram, error) {
	workouts := make([]models.WorkoutProgram, 0)
	for _, w := range s.workouts {
		if w != nil && w.UserID == userID {
			workouts = append(workouts, *w)
		}
	}

	return workouts, nil
}

func (s *cacheStorage) UpdateWorkoutProgram(ctx context.Context, wp *models.WorkoutProgram) (*models.WorkoutProgram, error) {
	id := int(wp.ID)
	if id < 0 || id > len(s.workouts) {
		return nil, fmt.Errorf("Cannot update")
	}
	s.workouts[id] = wp
	return s.workouts[id], nil
}

func (s *cacheStorage) DeleteWorkoutProgram(ctx context.Context, workoutprogramID int64) error {
	id := int(workoutprogramID)
	if id < 0 || id > len(s.workouts) {
		return fmt.Errorf("delete out of range")
	}
	s.workouts[id] = nil
	return nil
}

func (s *cacheStorage) GetReplica(label models.Replica, lang models.Language) (string, error) {
	return s.replics.GetReplica(label, lang)
}

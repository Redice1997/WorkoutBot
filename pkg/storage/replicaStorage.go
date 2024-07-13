package storage

import (
	"fmt"
	"workout_bot/pkg/models"
)

type replicaStorage struct {
	replics map[models.Replica]item
}

type item struct {
	ruText string
	enText string
}

func (i item) Get(lang models.Language) string {
	if lang == models.RU {
		return i.ruText
	} else {
		return i.ruText
	}
}

func NewReplicaStorage() models.ReplicaStorage {
	replics := map[models.Replica]item{
		models.StartActionReplica: item{
			ruText: `Я бот, который поможет набрать тебе в силовых и жать, как минимум, сотку от груди. 
	Моя программа сонована на методике из этого <a src="https://www.youtube.com/watch?v=bWnKmO0aj3c">видео</a>.

	Я создам тебе программу тренировок по-умолчанию с целью на 150кг в жиме лёжа, однако, ты сможешь в любой момент настроить параметры под себя.

	Команды:

		&#9/workouts - посмотреть список тренировок
		&#9/start_workout - начать тренировку
		&#9/progress - посмотреть свой прогресс`,
			enText: "",
		},
		models.WrongFormatReplica:          item{"Неверный формат значения", ""},
		models.SelectWorkoutsActionReplica: item{"Твои тренировки", ""},
		models.BtnReturnReplica:            item{"<< Вернуться", ""},
		models.BtnCreateWorkoutReplica:     item{"Создать новую тренировку", ""},
	}

	return &replicaStorage{replics}
}

func (s *replicaStorage) GetReplica(label models.Replica, lang models.Language) (string, error) {
	if r, ok := s.replics[label]; ok {
		return r.Get(lang), nil
	}
	return "", fmt.Errorf("replica not found")
}

package exerciseHistoryService

import (
	"context"
	"errors"
	"gcloud-serverless-gym/internal/core/domain"
	"gcloud-serverless-gym/internal/core/ports"
	"log/slog"
)

type ExerciseHistoryService struct {
	exerciseHistoryRepository ports.ExerciseHistoryRepository
	sessionService            ports.SessionService
	eventPublisher            ports.EventPublisher
}

func New(exerciseHistoryRepository ports.ExerciseHistoryRepository, sessionService ports.SessionService, eventPublisher ports.EventPublisher) *ExerciseHistoryService {
	return &ExerciseHistoryService{
		exerciseHistoryRepository: exerciseHistoryRepository,
		sessionService:            sessionService,
		eventPublisher:            eventPublisher,
	}
}

func (srv *ExerciseHistoryService) GetHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistoryDTO, error) {
	exerciseHistory, err := srv.exerciseHistoryRepository.GetHistoryFor(ctx, exerciseName)

	if err != nil {
		return domain.ExerciseHistoryDTO{}, errors.New("get workout from the repository has failed")
	}

	return exerciseHistory.AsDto(), nil
}

func (srv *ExerciseHistoryService) UpdateHistoryRecordFrom(ctx context.Context, command ports.UpdateExerciseHistoryCommand) {
	sessionData, err := srv.sessionService.Get(ctx, command.SessionId)

	if err != nil {
		return
	}

	var unique []string

	for _, v := range sessionData.Exercises {
		skip := false
		for _, u := range unique {
			if v.Name == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v.Name)
		}
	}

	for _, exerciseName := range unique {
		slog.Info("Processing " + exerciseName)

		exerciseHistory := srv.GetOrCreate(ctx, exerciseName)

		for _, sessionExercise := range sessionData.Exercises {
			if exerciseName == sessionExercise.Name {
				exerciseHistory.History = append(exerciseHistory.History, domain.ExerciseHistoryRecord{
					Date:   sessionData.Date,
					Set:    sessionExercise.Set,
					Reps:   sessionExercise.Reps,
					Weight: sessionExercise.Weight,
				})
			}
		}

		srv.exerciseHistoryRepository.UpdateHistoryRecord(ctx, exerciseHistory)

		srv.eventPublisher.PublishExerciseUpdatedEvent(ctx, exerciseHistory)
	}
}

func (srv *ExerciseHistoryService) GetOrCreate(ctx context.Context, exerciseName string) domain.ExerciseHistory {
	existingHistoryRecord, err := srv.exerciseHistoryRepository.GetHistoryFor(ctx, exerciseName)

	if err != nil {
		existingHistoryRecord, _ := srv.exerciseHistoryRepository.CreateHistoryFor(ctx, exerciseName)

		return existingHistoryRecord
	}

	return existingHistoryRecord
}

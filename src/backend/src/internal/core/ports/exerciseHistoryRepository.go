package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type UpdateExerciseHistoryCommand struct {
	SessionId string `json:"sessionId"`
}

type ExerciseHistoryRepository interface {
	GetHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistory, error)
	CreateHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistory, error)
	UpdateHistoryRecord(ctx context.Context, history domain.ExerciseHistory) error
}

type ExerciseHistoryService interface {
	GetHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistoryDTO, error)
	UpdateHistoryRecordFrom(ctx context.Context, command UpdateExerciseHistoryCommand)
}

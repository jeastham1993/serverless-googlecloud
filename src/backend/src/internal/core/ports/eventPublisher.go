package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type EventPublisher interface {
	PublishExerciseUpdatedEvent(ctx context.Context, e domain.ExerciseHistory)
}

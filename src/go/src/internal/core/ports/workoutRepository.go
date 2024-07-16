package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type CreateWorkoutCommand struct {
	Name      string                         `json:"name"`
	Exercises []CreateWorkoutCommandExercise `json:"exercises"`
}

type CreateWorkoutCommandExercise struct {
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}

type WorkoutRepository interface {
	Get(ctx context.Context, id string) (domain.Workout, error)
	Save(ctx context.Context, workout domain.Workout) error
	Exists(ctx context.Context, name string) (bool, error)
}

type WorkoutService interface {
	List(ctx context.Context) []domain.WorkoutDTO
	Get(ctx context.Context, id string) (domain.WorkoutDTO, error)
	Create(ctx context.Context, ommand CreateWorkoutCommand) (domain.WorkoutDTO, error)
	AddExerciseTo(ctx context.Context, orkout domain.Workout, exerciseName string) (domain.WorkoutDTO, error)
}

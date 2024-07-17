package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type CreateSessionCommand struct {
	Exercises []CreateSessionCommandExercise `json:"exercises"`
	Name      string                         `json:"name"`
}

type CreateSessionCommandExercise struct {
	Name   string  `json:"name"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

type CreateSessionFromWorkoutCommand struct {
	WorkoutId string `json:"workoutId"`
	Name      string `json:"name"`
}

type DuplicateSessionCommand struct {
	SessionId string `json:"sessionId"`
	Name      string `json:"name"`
}

type FinishSessionCommand struct {
	SessionId string `json:"sessionId"`
}

type SessionRepository interface {
	List(ctx context.Context) ([]domain.Session, error)
	Get(ctx context.Context, id string) (domain.Session, error)
	Save(ctx context.Context, session domain.Session) error
	Update(ctx context.Context, session domain.Session) error
}

type SessionService interface {
	List(ctx context.Context) []domain.SessionDTO
	Get(ctx context.Context, id string) (domain.SessionDTO, error)
	Create(ctx context.Context, command CreateSessionCommand) (domain.SessionDTO, error)
	Update(ctx context.Context, session domain.SessionDTO) (domain.SessionDTO, error)
	CreateSessionFromWorkout(ctx context.Context, command CreateSessionFromWorkoutCommand) (domain.SessionDTO, error)
	DuplicateSession(ctx context.Context, command DuplicateSessionCommand) (domain.SessionDTO, error)
	FinishSession(ctx context.Context, command FinishSessionCommand) (domain.SessionDTO, error)
}

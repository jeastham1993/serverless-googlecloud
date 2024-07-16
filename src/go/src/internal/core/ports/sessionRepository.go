package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type CreateSessionCommand struct {
	Exercises []CreateSessionCommandExercise `json:"exercises"`
}

type CreateSessionCommandExercise struct {
	Name   string  `json:"name"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

type SessionRepository interface {
	Get(ctx context.Context, id string) (domain.Session, error)
	Save(ctx context.Context, session domain.Session) error
}

type SessionService interface {
	List(ctx context.Context) []domain.SessionDTO
	Get(ctx context.Context, id string) (domain.SessionDTO, error)
	Create(ctx context.Context, command CreateSessionCommand) (domain.SessionDTO, error)
}

package ports

import (
	"gcloud-serverless-gym/internal/core/domain"

	"github.com/gin-gonic/gin"
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
	Get(ctx *gin.Context, id string) (domain.Session, error)
	Save(ctx *gin.Context, session domain.Session) error
}

type SessionService interface {
	List(ctx *gin.Context) []domain.SessionDTO
	Get(ctx *gin.Context, id string) (domain.SessionDTO, error)
	Create(ctx *gin.Context, command CreateSessionCommand) (domain.SessionDTO, error)
}

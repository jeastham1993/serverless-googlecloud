package ports

import (
	"gcloud-serverless-gym/internal/core/domain"

	"github.com/gin-gonic/gin"
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
	Get(ctx *gin.Context, id string) (domain.Workout, error)
	Save(ctx *gin.Context, workout domain.Workout) error
	Exists(ctx *gin.Context, name string) (bool, error)
}

type WorkoutService interface {
	List(ctx *gin.Context) []domain.WorkoutDTO
	Get(ctx *gin.Context, id string) (domain.WorkoutDTO, error)
	Create(cctx *gin.Context, ommand CreateWorkoutCommand) (domain.WorkoutDTO, error)
	AddExerciseTo(wctx *gin.Context, orkout domain.Workout, exerciseName string) (domain.WorkoutDTO, error)
}

package services

import (
	"errors"
	"gcloud-serverless-gym/internal/core/domain"
	"gcloud-serverless-gym/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type service struct {
	workoutRepository ports.WorkoutRepository
}

func New(workoutRepository ports.WorkoutRepository) *service {
	return &service{
		workoutRepository: workoutRepository,
	}
}

func (srv *service) Get(ctx *gin.Context, id string) (domain.WorkoutDTO, error) {
	workout, err := srv.workoutRepository.Get(ctx, id)

	if err != nil {
		return domain.WorkoutDTO{}, errors.New("get workout from the repository has failed")
	}

	return workout.AsDto(), nil
}

func (srv *service) Create(ctx *gin.Context, command ports.CreateWorkoutCommand) (domain.WorkoutDTO, error) {
	if len(command.Name) < 3 {
		return domain.WorkoutDTO{}, errors.New("name must have a length of at least 3")
	}

	exists, err := srv.workoutRepository.Exists(ctx, command.Name)

	if exists == true || err != nil {
		return domain.WorkoutDTO{}, errors.New("Workout exists")
	}

	workout := domain.NewWorkout(command.Name)

	for e := range command.Exercises {
		exercise := command.Exercises[e]
		workout.Exercises = append(workout.Exercises, domain.Exercise{Name: exercise.Name, Sets: exercise.Sets, Reps: exercise.Reps})
	}

	if err := srv.workoutRepository.Save(ctx, workout); err != nil {
		return domain.WorkoutDTO{}, errors.New("Create new workout has failed")
	}

	return workout.AsDto(), nil
}

func (srv *service) AddExerciseTo(ctx *gin.Context, workout domain.Workout, exerciseName string) (domain.WorkoutDTO, error) {
	return domain.WorkoutDTO{}, nil
}

func (srv *service) List(ctx *gin.Context) []domain.WorkoutDTO {
	var workouts [0]domain.WorkoutDTO

	return workouts[:]
}

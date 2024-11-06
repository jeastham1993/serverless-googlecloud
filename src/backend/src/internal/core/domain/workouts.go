package domain

import (
	"github.com/google/uuid"
)

type Workout struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Exercises []Exercise `json:"exercises"`
}

func NewWorkout(name string) Workout {
	workout := Workout{Id: uuid.New().String(), Name: name, Exercises: []Exercise{}}
	return workout
}

func (w *Workout) AsDto() WorkoutDTO {
	var exercises []ExerciseDTO

	for e := range w.Exercises {
		exercise := w.Exercises[e]

		exercises = append(exercises, ExerciseDTO{Name: exercise.Name, Sets: exercise.Sets, Reps: exercise.Reps})
	}

	workoutDto := WorkoutDTO{Id: w.Id, Name: w.Name, Exercises: exercises}

	return workoutDto
}

type WorkoutDTO struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Exercises []ExerciseDTO `json:"exercises"`
}

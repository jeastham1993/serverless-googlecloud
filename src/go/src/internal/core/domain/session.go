package domain

import (
	"time"
)

type Session struct {
	Id        string            `json:"id"`
	Date      time.Time         `json:"date"`
	Exercises []SessionExercise `json:"exercises"`
}

func NewSession() Session {
	timeNow := time.Now().Local()

	workout := Session{Id: timeNow.Format("2006-01-02 15"), Date: timeNow, Exercises: []SessionExercise{}}
	return workout
}

func (w *Session) AsDto() SessionDTO {
	var exercises []SessionExerciseDTO

	for e := range w.Exercises {
		exercise := w.Exercises[e]

		exercises = append(exercises, SessionExerciseDTO{Name: exercise.Name, Set: exercise.Set, Reps: exercise.Reps, Weight: exercise.Weight})
	}

	workoutDto := SessionDTO{Id: w.Id, Date: w.Date, Exercises: exercises}

	return workoutDto
}

type SessionDTO struct {
	Id        string               `json:"id"`
	Date      time.Time            `json:"date"`
	Exercises []SessionExerciseDTO `json:"exercises"`
}

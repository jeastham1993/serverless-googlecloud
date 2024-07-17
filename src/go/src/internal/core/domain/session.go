package domain

import (
	"time"
)

type Session struct {
	Id        string            `json:"id"`
	Date      time.Time         `json:"date"`
	Exercises []SessionExercise `json:"exercises"`
	Status    string            `json:"status"`
}

func NewSession(name string) Session {
	timeNow := time.Now().Local()

	workout := Session{Id: name, Date: timeNow, Exercises: []SessionExercise{}}
	return workout
}

func NewSessionFrom(name string, session Session) Session {
	timeNow := time.Now().Local()

	workout := Session{Id: name, Date: timeNow, Exercises: session.Exercises}
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

func (w *Session) Finished() {
	w.Status = "FINISHED"
}

type SessionDTO struct {
	Id        string               `json:"id"`
	Date      time.Time            `json:"date"`
	Exercises []SessionExerciseDTO `json:"exercises"`
}

package domain

import "time"

type ExerciseHistory struct {
	Name    string                  `json:"name"`
	History []ExerciseHistoryRecord `json:"history"`
}

type ExerciseHistoryRecord struct {
	Date   time.Time `json:"date"`
	Set    int       `json:"set"`
	Weight float64   `json:"weight"`
	Reps   int       `json:"reps"`
}

type ExerciseHistoryDTO struct {
	Name    string                     `json:"name"`
	History []ExerciseHistoryRecordDTO `json:"history"`
}

type ExerciseHistoryRecordDTO struct {
	Date   time.Time `json:"date"`
	Set    int       `json:"set"`
	Weight float64   `json:"weight"`
	Reps   int       `json:"reps"`
}

func (w *ExerciseHistory) AsDto() ExerciseHistoryDTO {
	var exercises []ExerciseHistoryRecordDTO

	for e := range w.History {
		exercise := w.History[e]

		exercises = append(exercises, ExerciseHistoryRecordDTO{Date: exercise.Date, Set: exercise.Set, Reps: exercise.Reps, Weight: exercise.Weight})
	}

	workoutDto := ExerciseHistoryDTO{Name: w.Name, History: exercises}

	return workoutDto
}

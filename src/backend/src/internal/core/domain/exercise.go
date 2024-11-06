package domain

type Exercise struct {
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}
type ExerciseDTO struct {
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}

type ExerciseUpdatedEventV1 struct {
	ExerciseName string                     `json:"ExerciseName"`
	History      []ExerciseHistoryRecordDTO `json:"History"`
	TraceId      string                     `json:"TraceId"`
}

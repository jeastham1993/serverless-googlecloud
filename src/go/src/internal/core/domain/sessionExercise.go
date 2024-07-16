package domain

type SessionExercise struct {
	Name   string  `json:"string"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

type SessionExerciseDTO struct {
	Name   string  `json:"string"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

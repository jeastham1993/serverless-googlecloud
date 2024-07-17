package domain

type SessionExercise struct {
	Name   string  `json:"name"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

type SessionExerciseDTO struct {
	Name   string  `json:"name"`
	Set    int     `json:"set"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

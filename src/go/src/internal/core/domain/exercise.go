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

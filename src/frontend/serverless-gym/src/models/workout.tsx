export interface Workout {
    id: string,
    name: string,
    exercises: Exercise[]
}

export interface Exercise {
    name: string,
    reps: number,
    sets: number
}
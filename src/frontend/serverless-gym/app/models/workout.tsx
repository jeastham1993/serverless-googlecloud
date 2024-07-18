interface Workout {
    id: string,
    name: string,
    exercises: Exercise[]
}

interface Exercise {
    name: string,
    reps: number,
    sets: number
}
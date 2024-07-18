export interface Session {
    id: string,
    date: string,
    exercises: SessionExercise[]
}

export interface SessionExercise {
    name: string,
    set: number,
    reps: number,
    weight: number
}
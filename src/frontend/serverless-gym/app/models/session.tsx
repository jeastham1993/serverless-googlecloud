interface Session {
    id: string,
    date: string,
    exercises: SessionExercise[]
}

interface SessionExercise {
    name: string,
    set: number,
    reps: number,
    weight: number
}
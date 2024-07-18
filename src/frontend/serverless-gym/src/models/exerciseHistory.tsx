export interface ExerciseHistroy {
  name: string;
  history: ExerciseHistoryRecord[];
}

export interface ExerciseHistoryRecord {
  date: string;
  set: number;
  weight: number;
  reps: number;
}

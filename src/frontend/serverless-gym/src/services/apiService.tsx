import { NavigateFunction } from "react-router-dom";
import { api } from "../axiosConfig";
import { Workout } from "../models/workout";
import { Session } from "../models/session";
import { ExerciseHistroy } from "../models/exerciseHistory";

export class ApiService {
  private router: NavigateFunction;
  constructor(router: NavigateFunction) {
    this.router = router;
  }

  async getWorkouts(): Promise<Workout[]> {
    const response = await api.get<Workout[]>("/workout");

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }

  async startSessionFromWorkout(
    workoutId: string,
    newSessionName: string
  ): Promise<Session> {
    const response = await api.post<Session>("/session/from", {
      workoutId: workoutId,
      name: newSessionName,
    });

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }

  async duplicateSession(sessionId: string, newSessionName: string) {
    const response = await api.post<Session>("/session/duplicate", {
      sessionId,
      name: newSessionName,
    });

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }

  async saveWorkout(workout: Workout) {
    const response = await api.post("/workout", workout);

    if (response.status === 401) {
      this.router("/login");
    }
  }

  async getExerciseHistory(exerciseName: string): Promise<ExerciseHistroy> {
    const response = await api.get<ExerciseHistroy>(`/history/${exerciseName}`);

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }

  async saveSession(session: Session) {
    const response = await api.put(`/session/${session.id}`, session);

    if (response.status === 401) {
      this.router("/login");
    }
  }

  async finishSession(session: Session) {
    const response = await api.post(`/session/finish`, {
      sessionId: session.id,
    });

    if (response.status === 401) {
      this.router("/login");
    }
  }

  async listSessions() {
    const response = await api.get<Session[]>("/session");

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }

  async getSession(sessionId: string) {
    const response = await api.get<Session>(`/session/${sessionId}`);

    if (response.status === 401) {
      this.router("/login");
    }

    return response.data;
  }
}

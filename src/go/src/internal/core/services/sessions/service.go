package sessionService

import (
	"context"
	"errors"
	"gcloud-serverless-gym/internal/core/domain"
	"gcloud-serverless-gym/internal/core/ports"
	"log/slog"
)

type SessionService struct {
	sessionRepository ports.SessionRepository
	workoutService    ports.WorkoutService
}

func New(sessionRepository ports.SessionRepository, workoutService ports.WorkoutService) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
		workoutService:    workoutService,
	}
}

func (srv *SessionService) Get(ctx context.Context, id string) (domain.SessionDTO, error) {
	session, err := srv.sessionRepository.Get(ctx, id)

	if err != nil {
		return domain.SessionDTO{}, errors.New("get workout from the repository has failed")
	}

	return session.AsDto(), nil
}

func (srv *SessionService) Create(ctx context.Context, command ports.CreateSessionCommand) (domain.SessionDTO, error) {
	session := domain.NewSession()

	for e := range command.Exercises {
		exercise := command.Exercises[e]
		session.Exercises = append(session.Exercises, domain.SessionExercise{Name: exercise.Name, Set: exercise.Set, Reps: exercise.Reps, Weight: exercise.Weight})
	}

	if err := srv.sessionRepository.Save(ctx, session); err != nil {
		return domain.SessionDTO{}, errors.New("Create new workout has failed")
	}

	return session.AsDto(), nil
}

func (srv *SessionService) CreateSessionFromWorkout(ctx context.Context, command ports.CreateSessionFromWorkoutCommand) (domain.SessionDTO, error) {
	workout, err := srv.workoutService.Get(ctx, command.WorkoutId)

	if err != nil {
		slog.Error("Workout not found")
		return domain.SessionDTO{}, err
	}

	session := domain.NewSession()

	for e := range workout.Exercises {
		exercise := workout.Exercises[e]

		for i := 1; i <= exercise.Sets; i++ {
			session.Exercises = append(session.Exercises, domain.SessionExercise{
				Name:   exercise.Name,
				Set:    i,
				Reps:   exercise.Reps,
				Weight: 0,
			})
		}
	}

	srv.sessionRepository.Save(ctx, session)

	return session.AsDto(), nil
}

func (srv *SessionService) Update(ctx context.Context, session domain.SessionDTO) (domain.SessionDTO, error) {
	existingSession, err := srv.sessionRepository.Get(ctx, session.Id)

	if err != nil {
		slog.Error(err.Error())
		return domain.SessionDTO{}, err
	}

	var newSessions []domain.SessionExercise

	for e := range session.Exercises {
		newExercise := session.Exercises[e]

		newSessions = append(newSessions, domain.SessionExercise{
			Name:   newExercise.Name,
			Weight: newExercise.Weight,
			Reps:   newExercise.Reps,
			Set:    newExercise.Set,
		})
	}

	existingSession.Exercises = newSessions

	srv.sessionRepository.Update(ctx, existingSession)

	return existingSession.AsDto(), nil
}

func (srv *SessionService) List(ctx context.Context) []domain.SessionDTO {
	sessions, err := srv.sessionRepository.List(ctx)
	var sessionList []domain.SessionDTO

	if err != nil {
		return sessionList
	}

	for w := range sessions {
		sessionList = append(sessionList, sessions[w].AsDto())
	}

	return sessionList
}

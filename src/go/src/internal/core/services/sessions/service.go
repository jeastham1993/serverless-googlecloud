package sessionService

import (
	"errors"
	"gcloud-serverless-gym/internal/core/domain"
	"gcloud-serverless-gym/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type SessionService struct {
	sessionRepository ports.SessionRepository
}

func New(sessionRepository ports.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

func (srv *SessionService) Get(ctx *gin.Context, id string) (domain.SessionDTO, error) {
	session, err := srv.sessionRepository.Get(ctx, id)

	if err != nil {
		return domain.SessionDTO{}, errors.New("get workout from the repository has failed")
	}

	return session.AsDto(), nil
}

func (srv *SessionService) Create(ctx *gin.Context, command ports.CreateSessionCommand) (domain.SessionDTO, error) {
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

func (srv *SessionService) List(ctx *gin.Context) []domain.SessionDTO {
	var workouts [0]domain.SessionDTO

	return workouts[:]
}

package sessionrepo

import (
	"log/slog"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"

	"gcloud-serverless-gym/internal/core/domain"
)

type FirestoreSessionRepository struct {
	collection *firestore.CollectionRef
}

func NewFirestoreRepository(client *firestore.Client) *FirestoreSessionRepository {
	workouts := client.Collection("Sessions")
	return &FirestoreSessionRepository{collection: workouts}
}

func (repo *FirestoreSessionRepository) Get(ctx *gin.Context, id string) (domain.Session, error) {
	workoutData, err := repo.collection.Doc(id).Get(ctx)

	if err != nil {
		slog.Error(err.Error())
		return domain.Session{}, err
	}

	var workout domain.Session

	if err := workoutData.DataTo(&workout); err != nil {
		slog.Error("Failure converting data to Session struct")
		slog.Error(err.Error())

		return domain.Session{}, err
	}

	return workout, nil
}

func (repo *FirestoreSessionRepository) Save(ctx *gin.Context, workout domain.Session) error {
	workoutDoc := repo.collection.Doc(workout.Id)

	_, err := workoutDoc.Create(ctx, workout)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Created sucessfully")

	return nil
}

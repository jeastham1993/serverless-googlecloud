package sessionrepo

import (
	"context"
	"log/slog"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"gcloud-serverless-gym/internal/core/domain"
)

type FirestoreSessionRepository struct {
	collection *firestore.CollectionRef
}

func NewFirestoreRepository(client *firestore.Client) *FirestoreSessionRepository {
	workouts := client.Collection("Sessions")
	return &FirestoreSessionRepository{collection: workouts}
}

func (repo *FirestoreSessionRepository) Get(ctx context.Context, id string) (domain.Session, error) {
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

func (repo *FirestoreSessionRepository) Save(ctx context.Context, workout domain.Session) error {
	workoutDoc := repo.collection.Doc(workout.Id)

	_, err := workoutDoc.Create(ctx, workout)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Created sucessfully")

	return nil
}

func (repo *FirestoreSessionRepository) Update(ctx context.Context, workout domain.Session) error {
	workoutDoc := repo.collection.Doc(workout.Id)

	_, err := workoutDoc.Set(ctx, map[string]interface{}{
		"Exercises": workout.Exercises,
	}, firestore.MergeAll)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Created sucessfully")

	return nil
}

func (repo *FirestoreSessionRepository) List(ctx context.Context) ([]domain.Session, error) {
	var sessions []domain.Session

	docsIter := repo.collection.Documents(ctx)
	defer docsIter.Stop()

	for {
		doc, err := docsIter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		var session domain.Session
		if err := doc.DataTo(&session); err != nil {
			slog.Error(err.Error())
			continue
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

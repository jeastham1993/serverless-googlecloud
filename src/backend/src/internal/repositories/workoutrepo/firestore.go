package workoutrepo

import (
	"context"
	"log/slog"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"gcloud-serverless-gym/internal/core/domain"
)

type FirestoreWorkoutRepository struct {
	collection *firestore.CollectionRef
}

func NewFirestoreRepository(client *firestore.Client) *FirestoreWorkoutRepository {
	workouts := client.Collection("Workouts")
	return &FirestoreWorkoutRepository{collection: workouts}
}

func (repo *FirestoreWorkoutRepository) Get(ctx context.Context, id string) (domain.Workout, error) {
	workoutData, err := repo.collection.Doc(id).Get(ctx)

	if err != nil {
		slog.Error(err.Error())
		return domain.Workout{}, err
	}

	var workout domain.Workout

	if err := workoutData.DataTo(&workout); err != nil {
		slog.Error("Failure converting data to Workout struct")
		slog.Error(err.Error())

		return domain.Workout{}, err
	}

	return workout, nil
}

func (repo *FirestoreWorkoutRepository) Save(ctx context.Context, workout domain.Workout) error {
	workoutDoc := repo.collection.Doc(workout.Id)

	_, err := workoutDoc.Create(ctx, workout)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Created sucessfully")

	return nil
}

func (repo *FirestoreWorkoutRepository) Exists(ctx context.Context, name string) (bool, error) {
	query := repo.collection.Where("Name", "==", name)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		slog.Error(err.Error())
		return true, err
	}

	if len(docs) > 0 {
		slog.Info("Workout already exists")
		return true, nil
	}

	return false, nil
}

func (repo *FirestoreWorkoutRepository) List(ctx context.Context) ([]domain.Workout, error) {
	var workouts []domain.Workout

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

		var workout domain.Workout
		if err := doc.DataTo(&workout); err != nil {
			slog.Error(err.Error())
			continue
		}
		workouts = append(workouts, workout)
	}

	return workouts, nil
}

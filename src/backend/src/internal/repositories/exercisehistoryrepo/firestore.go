package historyrepo

import (
	"context"
	"log/slog"

	"cloud.google.com/go/firestore"

	"gcloud-serverless-gym/internal/core/domain"
)

type FirestoreExerciseHistoryRepository struct {
	collection *firestore.CollectionRef
}

func NewFirestoreRepository(client *firestore.Client) *FirestoreExerciseHistoryRepository {
	exerciseHistory := client.Collection("ExerciseHistory")
	return &FirestoreExerciseHistoryRepository{collection: exerciseHistory}
}

func (repo *FirestoreExerciseHistoryRepository) GetHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistory, error) {
	exerciseHistoryDoc, err := repo.collection.Doc(exerciseName).Get(ctx)

	if err != nil {
		slog.Error(err.Error())
		return domain.ExerciseHistory{}, err
	}

	var exerciseHistory domain.ExerciseHistory

	if err := exerciseHistoryDoc.DataTo(&exerciseHistory); err != nil {
		slog.Error("Failure converting data to Exercise History struct")
		slog.Error(err.Error())

		return domain.ExerciseHistory{}, err
	}

	return exerciseHistory, nil
}

func (repo *FirestoreExerciseHistoryRepository) UpdateHistoryRecord(ctx context.Context, history domain.ExerciseHistory) error {
	exerciseHistoryDoc := repo.collection.Doc(history.Name)

	_, err := exerciseHistoryDoc.Set(ctx, map[string]interface{}{
		"Name":    history.Name,
		"History": history.History,
	}, firestore.MergeAll)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Updated sucessfully")

	return nil
}

func (repo *FirestoreExerciseHistoryRepository) CreateHistoryFor(ctx context.Context, exerciseName string) (domain.ExerciseHistory, error) {

	historyRecord := domain.ExerciseHistory{
		Name:    exerciseName,
		History: []domain.ExerciseHistoryRecord{},
	}

	historyRecordDoc := repo.collection.Doc(historyRecord.Name)

	_, err := historyRecordDoc.Create(ctx, historyRecord)

	if err != nil {
		slog.Error(err.Error())
		return historyRecord, nil
	}

	slog.Info("Created sucessfully")

	return historyRecord, nil
}

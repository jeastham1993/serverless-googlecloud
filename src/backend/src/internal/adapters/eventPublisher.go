package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"gcloud-serverless-gym/internal/core/domain"
	"log/slog"
	"os"
	"strconv"

	pubsub "cloud.google.com/go/pubsub"
	"github.com/DataDog/datadog-go/v5/statsd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type PubSubEventPublisher struct {
	client       *pubsub.Client
	statsDClient *statsd.Client
}

func NewPubSubEventPublisher(ctx context.Context) *PubSubEventPublisher {
	c, _ := pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECT_ID"))

	client, _ := statsd.New("127.0.0.1:8125",
		statsd.WithTags([]string{"env:prod"}),
	)

	return &PubSubEventPublisher{client: c, statsDClient: client}
}

func (srv *PubSubEventPublisher) PublishExerciseUpdatedEvent(ctx context.Context, e domain.ExerciseHistory, newRecords []domain.ExerciseHistoryRecord) {
	span, ctx := tracer.StartSpanFromContext(ctx, "start.publishEvent")
	defer span.Finish()

	topic := srv.client.Topic(os.Getenv("EXERCISE_UPDATED_TOPIC_ID"))
	defer topic.Stop()

	history := []domain.ExerciseHistoryRecordDTO{}

	for index := range newRecords {
		exercise_span, _ := tracer.StartSpanFromContext(ctx, "exercise")
		exercise_span.SetTag("exercise.name", e.Name)
		exercise_span.SetTag("exercise.date", newRecords[index].Date)
		exercise_span.SetTag("exercise.set", newRecords[index].Set)
		exercise_span.SetTag("exercise.weight", newRecords[index].Weight)
		exercise_span.SetTag("exercise.reps", newRecords[index].Reps)

		historyRecord := domain.ExerciseHistoryRecordDTO{
			Date:   newRecords[index].Date,
			Set:    newRecords[index].Set,
			Weight: newRecords[index].Weight,
			Reps:   newRecords[index].Reps,
		}

		history = append(history, historyRecord)

		exercise_span.Finish()
	}

	evt := domain.ExerciseUpdatedEventV1{
		ExerciseName: e.Name,
		History:      history,
		TraceId:      strconv.FormatUint(span.Context().TraceID(), 10),
	}

	json, _ := json.Marshal(evt)

	slog.Info("Publishing message:")
	slog.Info(string(json))

	r := topic.Publish(ctx, &pubsub.Message{
		Data: json,
	})

	msgId, err := r.Get(ctx)

	slog.Info(fmt.Sprintf("MessageId is '%s'", msgId))
	span.SetTag("messaging.pubSubMsgId", msgId)

	if err != nil {
		span.SetTag("error", true)
		slog.Error(err.Error())
	}
}

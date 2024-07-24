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
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type PubSubEventPublisher struct {
	client *pubsub.Client
}

func NewPubSubEventPublisher(ctx context.Context) *PubSubEventPublisher {
	c, _ := pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECT_ID"))

	return &PubSubEventPublisher{client: c}
}

func (srv *PubSubEventPublisher) PublishExerciseUpdatedEvent(ctx context.Context, e domain.ExerciseHistory) {
	span, ctx := tracer.StartSpanFromContext(ctx, "start.publishEvent")
	defer span.Finish()

	topic := srv.client.Topic(os.Getenv("EXERCISE_UPDATED_TOPIC_ID"))
	defer topic.Stop()

	evt := domain.ExerciseUpdatedEventV1{
		ExerciseName: e.Name,
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

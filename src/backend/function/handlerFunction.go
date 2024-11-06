package handlerFunction

import (
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("ExerciseUpdatedHandlerFunction", handleExerciseUpdatedEvent)
}

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func handleExerciseUpdatedEvent(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	name := string(msg.Message.Data) // Automatically decoded from base64.

	log.Printf(name)
	return nil
}

// [END functions_cloudevent_pubsub]

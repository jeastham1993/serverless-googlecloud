package handlerFunction

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
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
	client, err := statsd.New("127.0.0.1:8125",
		statsd.WithTags([]string{"env:prod"}),
	)

	if err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	messageData := string(msg.Message.Data) // Automatically decoded from base64.

	exerciseData := ExerciseHistoryDTO{}
	err = json.Unmarshal([]byte(messageData), &exerciseData)

	if err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	log.Printf(messageData)

	for item := range exerciseData.History {
		client.Event(&statsd.Event{
			SourceTypeName: "serverless-gym",
			Title:          exerciseData.Name,
			Text:           exerciseData.Name,
			AlertType:      "info",
			Timestamp:      exerciseData.History[item].Date,
		})
	}
	return nil
}

type ExerciseHistoryDTO struct {
	Name    string                     `json:"name"`
	History []ExerciseHistoryRecordDTO `json:"history"`
}

type ExerciseHistoryRecordDTO struct {
	Date   time.Time `json:"date"`
	Set    int       `json:"set"`
	Weight float64   `json:"weight"`
	Reps   int       `json:"reps"`
}

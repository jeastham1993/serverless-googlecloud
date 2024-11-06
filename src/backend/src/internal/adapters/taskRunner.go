package adapters

import (
	"context"
	"fmt"
	"gcloud-serverless-gym/internal/core/domain"
	"log/slog"
	"net/http"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type CloudTasksRunner struct {
	client *cloudtasks.Client
}

func NewCloudTaskRunner(ctx context.Context) *CloudTasksRunner {
	c, _ := cloudtasks.NewClient(ctx)

	return &CloudTasksRunner{client: c}
}

func (srv *CloudTasksRunner) StartHistoryUpdateFor(ctx context.Context, session domain.Session) {
	span, ctx := tracer.StartSpanFromContext(ctx, "start.historyUpdate")
	defer span.Finish()

	url := fmt.Sprintf("%s/history", os.Getenv("APP_URL"))

	h := http.Header{}

	_ = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(h))

	httpRequestHeaders := make(map[string]string)

	for name, values := range h {
		slog.Info(name)
		slog.Info(values[0])

		httpRequestHeaders[name] = values[0]
	}

	httpRequestHeaders["Authorization"] = os.Getenv("API_KEY")

	req := &taskspb.CreateTaskRequest{
		Parent: os.Getenv("QUEUE_NAME"),
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        url,
					Headers:    httpRequestHeaders,
				},
			},
		},
	}

	message := fmt.Sprintf("{\"sessionId\": \"%s\"}", session.Id)

	req.Task.GetHttpRequest().Body = []byte(message)

	createTaskRespone, err := srv.client.CreateTask(ctx, req)

	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(createTaskRespone.Name)
}

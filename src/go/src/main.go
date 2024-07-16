package main

import (
	"context"
	sessionService "gcloud-serverless-gym/internal/core/services/sessions"
	services "gcloud-serverless-gym/internal/core/services/workouts"
	"gcloud-serverless-gym/internal/handlers"
	"gcloud-serverless-gym/internal/repositories/sessionrepo"
	"gcloud-serverless-gym/internal/repositories/workoutrepo"
	"os"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"

	"log/slog"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	ctx := context.Background()

	configureObservability()
	firestoreClient, err := configureGoogleClients(ctx)

	if err != nil {
		slog.Error("Failure starting service due to Google client configuration error. Exiting...")
		panic("Failure starting service due to Google client configuration error. Exiting...")
	}

	defer firestoreClient.Close()

	firestoreRepository := workoutrepo.NewFirestoreRepository(firestoreClient)
	workoutService := services.New(firestoreRepository)
	workoutHandler := handlers.NewWorkoutHTTPHandler(workoutService)

	sessionRepository := sessionrepo.NewFirestoreRepository(firestoreClient)
	sessionService := sessionService.New(sessionRepository)
	sessionHandler := handlers.NewSessionHTTPHandler(sessionService)

	router := gin.New()
	router.Use(gintrace.Middleware(os.Getenv("DD_SERVICE")))
	router.GET("/workout/:id", workoutHandler.Get)
	router.POST("/workout", workoutHandler.Post)
	router.GET("/session/:id", sessionHandler.Get)
	router.POST("/session", sessionHandler.Post)

	router.Run(":8080")

	slog.Info("Running on port 8080")

	defer tracer.Stop()
}

func configureGoogleClients(ctx context.Context) (*firestore.Client, error) {
	si := grpctrace.StreamClientInterceptor(grpctrace.WithServiceName("firestore"))

	c, err := firestore.NewClient(ctx, os.Getenv("GCLOUD_PROJECT_ID"),
		option.WithGRPCDialOption(grpc.WithStreamInterceptor(si)),
	)

	if err != nil {
		slog.Error(err.Error())

		return nil, err
	}

	return c, nil
}

func configureObservability() {
	tracer.Start(
		tracer.WithEnv(os.Getenv("DD_ENV")),
		tracer.WithService(os.Getenv("DD_SERVICE")),
		tracer.WithServiceVersion(os.Getenv("DD_VERSION")),
	)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

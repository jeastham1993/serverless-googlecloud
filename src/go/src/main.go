package main

import (
	"context"
	"gcloud-serverless-gym/internal/adapters"
	exerciseHistoryService "gcloud-serverless-gym/internal/core/services/exerciseHistory"
	sessionService "gcloud-serverless-gym/internal/core/services/sessions"
	services "gcloud-serverless-gym/internal/core/services/workouts"
	"gcloud-serverless-gym/internal/handlers"
	historyrepo "gcloud-serverless-gym/internal/repositories/exercisehistoryrepo"
	"gcloud-serverless-gym/internal/repositories/sessionrepo"
	"gcloud-serverless-gym/internal/repositories/workoutrepo"
	"os"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log/slog"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	configureObservability()
	firestoreClient, err := configureGoogleClients(&gin.Context{})

	if err != nil {
		slog.Error("Failure starting service due to Google client configuration error. Exiting...")
		panic("Failure starting service due to Google client configuration error. Exiting...")
	}

	defer firestoreClient.Close()

	taskRunner := adapters.NewCloudTaskRunner(&gin.Context{})

	firestoreRepository := workoutrepo.NewFirestoreRepository(firestoreClient)
	workoutService := services.New(firestoreRepository)
	workoutHandler := handlers.NewWorkoutHTTPHandler(workoutService)

	sessionRepository := sessionrepo.NewFirestoreRepository(firestoreClient)
	sessionService := sessionService.New(sessionRepository, workoutService, taskRunner)
	sessionHandler := handlers.NewSessionHTTPHandler(sessionService)

	historyRepository := historyrepo.NewFirestoreRepository(firestoreClient)
	historyService := exerciseHistoryService.New(historyRepository, sessionService)
	historyHandlers := handlers.NewExerciseHistoryHTTPHandler(historyService)

	router := gin.New()

	router.Use(gintrace.Middleware(os.Getenv("DD_SERVICE")))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/workout", workoutHandler.List)
	router.POST("/workout", workoutHandler.Post)
	router.GET("/workout/:id", workoutHandler.Get)

	router.GET("/session", sessionHandler.List)
	router.POST("/session", sessionHandler.Post)
	router.GET("/session/:id", sessionHandler.Get)
	router.PUT("/session/:id", sessionHandler.Update)
	router.POST("/session/from", sessionHandler.PostFromWorkout)
	router.POST("/session/finish", sessionHandler.Finish)

	router.GET("/history/:name", historyHandlers.Get)
	router.POST("/history", historyHandlers.CreateFor)

	router.Run(":8080")

	slog.Info("Running on port 8080")

	defer tracer.Stop()
}

func configureGoogleClients(ctx context.Context) (*firestore.Client, error) {
	si := grpctrace.StreamClientInterceptor(
		grpctrace.WithServiceName("firestore"),
	)

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

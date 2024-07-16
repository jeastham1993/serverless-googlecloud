package handlers

import (
	"gcloud-serverless-gym/internal/core/ports"

	"github.com/gin-gonic/gin"

	"log/slog"
)

type WorkoutHTTPHandler struct {
	workoutService ports.WorkoutService
}

func NewWorkoutHTTPHandler(workoutService ports.WorkoutService) *WorkoutHTTPHandler {
	return &WorkoutHTTPHandler{
		workoutService: workoutService,
	}
}

func (hdl *WorkoutHTTPHandler) Get(c *gin.Context) {
	workout, err := hdl.workoutService.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		slog.Error("failure retrieving data from workout service")
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, workout)
}

func (hdl *WorkoutHTTPHandler) Post(c *gin.Context) {
	var command ports.CreateWorkoutCommand

	if err := c.BindJSON(&command); err != nil {
		slog.Error(err.Error())
		return
	}

	workout, err := hdl.workoutService.Create(c.Request.Context(), command)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, workout)
}

package handlers

import (
	"gcloud-serverless-gym/internal/core/ports"

	"github.com/gin-gonic/gin"

	"log/slog"
)

type ExerciseHistoryHTTPHandler struct {
	exerciseHistoryService ports.ExerciseHistoryService
}

func NewExerciseHistoryHTTPHandler(exerciseHistoryService ports.ExerciseHistoryService) *ExerciseHistoryHTTPHandler {
	return &ExerciseHistoryHTTPHandler{
		exerciseHistoryService: exerciseHistoryService,
	}
}

func (hdl *ExerciseHistoryHTTPHandler) Get(c *gin.Context) {
	exerciseHistory, err := hdl.exerciseHistoryService.GetHistoryFor(c.Request.Context(), c.Param("name"))
	if err != nil {
		slog.Error("failure retrieving data from exerciseHistory service")
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, exerciseHistory)
}

func (hdl *ExerciseHistoryHTTPHandler) CreateFor(c *gin.Context) {
	var command ports.UpdateExerciseHistoryCommand

	if err := c.BindJSON(&command); err != nil {
		slog.Error(err.Error())
		return
	}

	hdl.exerciseHistoryService.UpdateHistoryRecordFrom(c.Request.Context(), command)

	c.Status(201)
}

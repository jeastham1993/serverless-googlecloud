package handlers

import (
	"gcloud-serverless-gym/internal/core/domain"
	"gcloud-serverless-gym/internal/core/ports"

	"github.com/gin-gonic/gin"

	"log/slog"
)

type SessionHTTPHandler struct {
	sessionService ports.SessionService
}

func NewSessionHTTPHandler(sessionService ports.SessionService) *SessionHTTPHandler {
	return &SessionHTTPHandler{
		sessionService: sessionService,
	}
}

func (hdl *SessionHTTPHandler) Get(c *gin.Context) {
	session, err := hdl.sessionService.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		slog.Error("failure retrieving data from session service")
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

func (hdl *SessionHTTPHandler) Post(c *gin.Context) {
	var command ports.CreateSessionCommand

	if err := c.BindJSON(&command); err != nil {
		slog.Error(err.Error())
		return
	}

	session, err := hdl.sessionService.Create(c.Request.Context(), command)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

func (hdl *SessionHTTPHandler) PostFromWorkout(c *gin.Context) {
	var command ports.CreateSessionFromWorkoutCommand

	if err := c.BindJSON(&command); err != nil {
		slog.Error(err.Error())
		return
	}

	session, err := hdl.sessionService.CreateSessionFromWorkout(c.Request.Context(), command)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

func (hdl *SessionHTTPHandler) Update(c *gin.Context) {
	var session domain.SessionDTO

	if err := c.BindJSON(&session); err != nil {
		slog.Error(err.Error())
		return
	}

	session.Id = c.Param("id")

	session, err := hdl.sessionService.Update(c.Request.Context(), session)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

func (hdl *SessionHTTPHandler) Finish(c *gin.Context) {
	var command ports.FinishSessionCommand

	if err := c.BindJSON(&command); err != nil {
		slog.Error(err.Error())
		return
	}

	session, err := hdl.sessionService.FinishSession(c.Request.Context(), command)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

func (hdl *SessionHTTPHandler) List(c *gin.Context) {
	workouts := hdl.sessionService.List(c.Request.Context())

	c.JSON(200, workouts)
}

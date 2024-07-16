package handlers

import (
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
	session, err := hdl.sessionService.Get(c, c.Param("id"))
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

	session, err := hdl.sessionService.Create(c, command)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, session)
}

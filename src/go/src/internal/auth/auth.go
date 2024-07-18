package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const (
	authorizationHeader = "Authorization"
	userId              = "USER_ID"
	subject             = "SUBJECT"
)

var (
	firebaseConfigFile = os.Getenv("FIREBASE_CONFIG_FILE")
)

func InitAuth() (*auth.Client, error) {
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: os.Getenv("GCLOUD_PROJECT_ID"),
	})
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	}

	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		slog.Error(errAuth.Error())
		return nil, errors.Wrap(errAuth, "error initializing firebase auth (creating client)")
	}

	slog.Info("Auth client setup ok")

	return client, nil
}

func ValidateFirebaseToken(client *auth.Client, token string) (*auth.Token, error) {
	tokenInfo, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}

	return tokenInfo, nil
}

func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := tracer.StartSpanFromContext(c, "auth")
		defer span.Finish()

		authHeader := c.Request.Header.Get(authorizationHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		tokenInfo, err := ValidateFirebaseToken(client, token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		if tokenInfo.UID != os.Getenv("ALLOWED_USER_ID") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		c.Set(userId, tokenInfo.UID)
		c.Next()
	}
}

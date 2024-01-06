package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"task-management-with-gin/configs"
	"task-management-with-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type AuthenticatedUser struct {
	Id    uint
	Email string
}

func extractBearerToken(r *http.Request) (string, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	// Check if the header has the "Bearer" prefix
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	// Extract the token excluding the "Bearer " prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token, nil
}

func AuthGuard(ctx *gin.Context) {
	token, err := extractBearerToken(ctx.Request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	sub, err := utils.ValidateJwtToken(token, configs.JwtSecret)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	var userData AuthenticatedUser
	mapErr := mapstructure.Decode(sub, &userData)
	if mapErr != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "cannot get user data"})
		return
	}

	ctx.Set("user", userData)
	ctx.Next()

}

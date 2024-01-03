package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"task-management-with-gin/configs"
	"task-management-with-gin/utils"

	"github.com/gin-gonic/gin"
)

type AuthenticatedUser struct {
	utils.TokenPayload
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
	var token string

	token, err := extractBearerToken(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	sub, err := utils.ValidateJwtToken(token, configs.JwtSecret)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	userData, ok := sub.(map[string]interface{})
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "cannot get user data"})
		return
	}

	id, idOK := userData["Id"].(float64)
	email, emailOK := userData["Email"].(string)

	if !idOK || !emailOK {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "invalid data format"})
		return
	}

	var authenticatedUser AuthenticatedUser
	authenticatedUser.Id = uint(id)
	authenticatedUser.Email = email

	ctx.Set("user", authenticatedUser)

	ctx.Next()
}

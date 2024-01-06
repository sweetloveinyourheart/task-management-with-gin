package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(ctx *gin.Context, errorMessage string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":   errorMessage,
		"success": false,
	})
}

func UnauthorizedResponse(ctx *gin.Context, errorMessage string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"error":   errorMessage,
		"success": false,
	})
}

func ForbiddenResponse(ctx *gin.Context, errorMessage string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"error":   errorMessage,
		"success": false,
	})
}

func NotFoundResponse(ctx *gin.Context, errorMessage string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error":   errorMessage,
		"success": false,
	})
}

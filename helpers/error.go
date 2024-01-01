package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleValidationError(ctx *gin.Context, validationError validator.ValidationErrors) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "Validation error",
		"details": map[string]string{
			"field": validationError[0].Field(),
			"tag":   validationError[0].Tag(),
		},
	})
}

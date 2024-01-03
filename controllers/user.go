package controllers

import (
	"net/http"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserController represents the controller for handling user-related operations
type UserController struct {
	UserService services.IUserService
}

// NewUserController creates a new instance of UserController
func NewUserController() *UserController {
	services := services.NewUserService()

	return &UserController{
		UserService: services,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	newUserData := dtos.RegisterDTO{}
	bindErr := ctx.ShouldBindJSON(&newUserData)
	helpers.ErrorPanic(bindErr)

	success, err := c.UserService.CreateNewUser(newUserData)
	if err != nil {
		// Check if it's a validation error
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		// If it's not a validation error, handle it as a general error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": success,
		})
		return
	}

	// Send a success response with status code 201
	ctx.JSON(http.StatusCreated, gin.H{"success": success})
}

func (c *UserController) SignIn(ctx *gin.Context) {
	signInData := dtos.SignInDTO{}
	bindErr := ctx.ShouldBindJSON(&signInData)
	helpers.ErrorPanic(bindErr)

	tokens, err := c.UserService.SignIn(signInData)
	if err != nil {
		// Check if it's a validation error
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  tokens,
	})
}

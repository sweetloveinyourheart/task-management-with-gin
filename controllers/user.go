package controllers

import (
	"net/http"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/helpers/exceptions"
	"task-management-with-gin/middlewares"
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
		exceptions.BadRequestResponse(ctx, err.Error())
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

		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  tokens,
	})
}

func (c *UserController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refresh-token")

	tokens, err := c.UserService.RefreshToken(refreshToken)
	if err != nil {
		// Check if it's a validation error
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  tokens,
	})
}

func (c *UserController) GetUserProfile(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		exceptions.UnauthorizedResponse(ctx, "Unauthorized")
		return
	}

	authUser, ok := user.(middlewares.AuthenticatedUser)
	if !ok {
		exceptions.UnauthorizedResponse(ctx, "Invalid user type")
		return
	}

	profile, err := c.UserService.GetUserProfile(authUser.Id)
	if err != nil {
		exceptions.ForbiddenResponse(ctx, "Error getting user profile")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  profile,
	})
}

func (c *UserController) UpdateUserProfile(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		exceptions.UnauthorizedResponse(ctx, "Unauthorized")
		return
	}

	// Get the authentication user
	authUser, ok := user.(middlewares.AuthenticatedUser)
	if !ok {
		exceptions.UnauthorizedResponse(ctx, "Invalid user type")
		return
	}

	userData := dtos.UpdateProfileDTO{}
	err := ctx.ShouldBindJSON(&userData)
	helpers.ErrorPanic(err)

	if err != nil {
		// Check if it's a validation error
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	success, err := c.UserService.UpdateUserProfile(authUser.Id, userData)
	if !success {
		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	// Send a success response with status code 201
	ctx.JSON(http.StatusCreated, gin.H{"success": success})
}

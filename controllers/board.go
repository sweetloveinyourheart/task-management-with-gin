package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/middlewares"
	"task-management-with-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BoardController represents the controller for handling user-related operations
type BoardController struct {
	BoardService services.IBoardService
}

// NewBoardController creates a new instance of BoardController
func NewBoardController() *BoardController {
	services := services.NewBoardService()

	return &BoardController{
		BoardService: services,
	}
}

func (c *BoardController) NewBoard(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   fmt.Errorf("unauthorized").Error(),
			"success": false,
		})
		return
	}

	authUser, ok := user.(middlewares.AuthenticatedUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   fmt.Errorf("invalid user type").Error(),
			"success": false,
		})
		return
	}

	newBoard := dtos.NewBoardDTO{}
	bindErr := ctx.ShouldBindJSON(&newBoard)
	helpers.ErrorPanic(bindErr)

	success, err := c.BoardService.CreateNewBoard(authUser.Id, newBoard)
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

func (c *BoardController) NewBoardMembers(ctx *gin.Context) {
	boardId := ctx.Query("board-id")
	if boardId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   fmt.Errorf("board-id is required").Error(),
			"success": false,
		})
		return
	}

	// Convert string to uint
	boardIdUint, err := strconv.ParseUint(boardId, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Printf("Error converting string to uint: %v\n", err)
		return
	}

	members := dtos.AddBoardMembers{}
	bindErr := ctx.ShouldBindJSON(&members)
	helpers.ErrorPanic(bindErr)

	success, err := c.BoardService.AddBoardMember(uint(boardIdUint), members)
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

package controllers

import (
	"net/http"
	"strconv"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/helpers/exceptions"
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
	// Verify user authentication
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

	// Bind JSON request to newBoard DTO
	var newBoard dtos.NewBoardDTO
	if bindErr := ctx.ShouldBindJSON(&newBoard); bindErr != nil {
		helpers.ErrorPanic(bindErr)
		return
	}

	// Call service to create a new board
	success, err := c.BoardService.CreateNewBoard(authUser.Id, newBoard)
	if err != nil {
		// Handle validation errors
		if validationError, ok := err.(validator.ValidationErrors); ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		// Handle general errors
		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	// Send a success response with status code 201
	ctx.JSON(http.StatusCreated, gin.H{"success": success})
}

func (c *BoardController) NewBoardMembers(ctx *gin.Context) {
	// Extract boardId from query parameters
	boardId := ctx.Query("board-id")
	if boardId == "" {
		exceptions.BadRequestResponse(ctx, "board-id is required")
		return
	}

	// Parse and validate boardId
	boardIdUint, err := strconv.ParseUint(boardId, 10, 64)
	if err != nil {
		exceptions.BadRequestResponse(ctx, "Invalid board-id format")
		return
	}

	// Bind JSON request to members DTO
	var members dtos.AddBoardMembers
	if bindErr := ctx.ShouldBindJSON(&members); bindErr != nil {
		helpers.ErrorPanic(bindErr)
		return
	}

	// Call service to add board members
	success, err := c.BoardService.AddBoardMember(uint(boardIdUint), members)
	if err != nil {
		// Handle validation errors
		if validationError, ok := err.(validator.ValidationErrors); ok {
			helpers.HandleValidationError(ctx, validationError)
			return
		}

		// Handle general errors
		exceptions.BadRequestResponse(ctx, err.Error())
		return
	}

	// Send a success response with status code 201
	ctx.JSON(http.StatusCreated, gin.H{"success": success})
}

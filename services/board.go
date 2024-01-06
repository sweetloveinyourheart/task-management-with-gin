package services

import (
	"fmt"
	"task-management-with-gin/configs"
	"task-management-with-gin/dtos"
	"task-management-with-gin/models"
	"task-management-with-gin/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IBoardService interface {
	CreateNewBoard(creatorId uint, newBoardData dtos.NewBoardDTO) (bool, error)
	AddBoardMember(boardId uint, members dtos.AddBoardMembers) (bool, error)
}

type BoardService struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

func NewBoardService() IBoardService {
	validate := utils.GetValidator()
	db := configs.GetDB()

	return &BoardService{
		Validate: validate,
		Db:       db,
	}
}

func (b *BoardService) CreateNewBoard(creatorId uint, newBoardData dtos.NewBoardDTO) (bool, error) {
	validatorError := b.Validate.Struct(newBoardData)
	if validatorError != nil {
		return false, validatorError
	}

	// Create a new board
	newBoard := models.Board{
		Title:  newBoardData.Title,
		UserId: creatorId,
	}

	// Retrieve the creator user from the database
	var creatorUser models.User
	if err := b.Db.First(&creatorUser, creatorId).Error; err != nil {
		return false, err
	}

	// Append the creator to the Members slice
	newBoard.Members = append(newBoard.Members, creatorUser)

	// Save the new board to the database
	if err := b.Db.Create(&newBoard).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (b *BoardService) AddBoardMember(boardId uint, members dtos.AddBoardMembers) (bool, error) {
	validatorError := b.Validate.Struct(members)
	if validatorError != nil {
		return false, validatorError
	}

	var users []models.User
	findMemErr := b.Db.Find(&users, members.Members).Error
	if findMemErr != nil {
		return false, fmt.Errorf("cannot add members")
	}

	var board models.Board
	findBoardErr := b.Db.First(&board, boardId).Error
	if findBoardErr != nil {
		return false, fmt.Errorf("cannot add members")
	}

	board.Members = append(board.Members, users...)

	// Save the updated board to the database
	if saveErr := b.Db.Save(&board).Error; saveErr != nil {
		return false, saveErr
	}

	return true, nil
}

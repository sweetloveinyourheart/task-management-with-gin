package services

import (
	"errors"
	"task-management-with-gin/configs"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/models"
	"task-management-with-gin/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateNewUser(userData dtos.RegisterDTO) (bool, error)
}

type UserService struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

func NewUserService() IUserService {
	db := configs.GetDB()
	validate := utils.Validate

	return &UserService{
		Validate: validate,
		Db:       db,
	}
}

var user models.User

func (u *UserService) CreateNewUser(userData dtos.RegisterDTO) (bool, error) {
	existedUser := u.Db.Where("email = ?", userData.Email).First(&user)
	if existedUser.RowsAffected > 0 {
		// User already exists, construct and return an error
		return false, errors.New("user already exists")
	}

	validatorError := u.Validate.Struct(userData)
	if validatorError != nil {
		return false, validatorError
	}

	// Hash the password
	hashedPassword, hashErr := utils.HashPassword(userData.Password)
	if hashErr != nil {
		helpers.ErrorPanic(hashErr)
	}

	newUser := models.User{
		FullName: &userData.FullName,
		Email:    userData.Email,
		Password: hashedPassword,
	}

	createResult := u.Db.Create(&newUser)
	if createResult.Error != nil {
		helpers.ErrorPanic(createResult.Error)
	}

	return true, nil
}

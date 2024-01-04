package services

import (
	"database/sql"
	"fmt"
	"task-management-with-gin/configs"
	"task-management-with-gin/dtos"
	"task-management-with-gin/helpers"
	"task-management-with-gin/models"
	"task-management-with-gin/responses"
	"task-management-with-gin/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateNewUser(userData dtos.RegisterDTO) (bool, error)
	SignIn(userData dtos.SignInDTO) (responses.SignInResponse, error)
	GetUserProfile(userId uint) (responses.UserProfile, error)
	UpdateUserProfile(userId uint, userData dtos.UpdateProfileDTO) (bool, error)
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

func (u *UserService) CreateNewUser(userData dtos.RegisterDTO) (bool, error) {
	validatorError := u.Validate.Struct(userData)
	if validatorError != nil {
		return false, validatorError
	}

	var user models.User

	existedUser := u.Db.Where("email = ?", userData.Email).First(&user)
	if existedUser.RowsAffected > 0 {
		// User already exists, construct and return an error
		return false, fmt.Errorf("user already exists")
	}

	// Hash the password
	hashedPassword, hashErr := utils.HashPassword(userData.Password)
	if hashErr != nil {
		helpers.ErrorPanic(hashErr)
	}

	newUser := models.User{
		FullName: sql.NullString{String: userData.FullName, Valid: true},
		Email:    userData.Email,
		Password: hashedPassword,
		Username: utils.GenerateRandomUsername("user_"),
	}

	createResult := u.Db.Create(&newUser)
	if createResult.Error != nil {
		helpers.ErrorPanic(createResult.Error)
	}

	return true, nil
}

func (u *UserService) SignIn(userData dtos.SignInDTO) (responses.SignInResponse, error) {
	validatorError := u.Validate.Struct(userData)
	if validatorError != nil {
		return responses.SignInResponse{}, validatorError
	}

	var user models.User

	userQuery := u.Db.Where("email = ?", userData.Email).First(&user)
	if userQuery.RowsAffected == 0 {
		return responses.SignInResponse{}, fmt.Errorf("email or password is not valid")
	}

	isValidPassword := utils.CheckPasswordHash(userData.Password, user.Password)
	if !isValidPassword {
		return responses.SignInResponse{}, fmt.Errorf("email or password is not valid")
	}

	payload := utils.TokenPayload{
		Id:    user.ID,
		Email: user.Email,
	}

	// Generate Token
	accessToken, errAccessToken := utils.GenerateToken(payload, configs.JwtSecret, 15*time.Minute)
	helpers.ErrorPanic(errAccessToken)

	refreshToken, errRefreshToken := utils.GenerateToken(payload, configs.JwtSecret, 24*time.Hour)
	helpers.ErrorPanic(errRefreshToken)

	return responses.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserService) GetUserProfile(userId uint) (responses.UserProfile, error) {
	var user models.User

	userQuery := u.Db.Where("id = ?", userId).First(&user)
	if userQuery.RowsAffected == 0 {
		return responses.UserProfile{}, fmt.Errorf("missing Authorization header")
	}

	return responses.UserProfile{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName.String,
	}, nil
}

func (u *UserService) UpdateUserProfile(userId uint, userData dtos.UpdateProfileDTO) (bool, error) {
	var user models.User

	userQuery := u.Db.Where("id = ?", userId).First(&user)
	if userQuery.RowsAffected == 0 {
		return false, fmt.Errorf("missing Authorization header")
	}

	fmt.Println(userData.Username, userData.FullName)

	if userData.Username != "" {
		user.Username = userData.Username
	}

	if userData.FullName != "" {
		user.FullName = sql.NullString{String: userData.FullName, Valid: true}
	}

	saveResult := u.Db.Save(&user)
	if saveResult.Error != nil {
		helpers.ErrorPanic(saveResult.Error)
		return false, nil
	}

	return true, nil
}

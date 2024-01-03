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
	SignIn(userData dtos.SignInDTO) (SignInResponse, error)
}

type UserService struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

type SignInResponse struct {
	AccessToken  string
	RefreshToken string
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
	validatorError := u.Validate.Struct(userData)
	if validatorError != nil {
		return false, validatorError
	}

	existedUser := u.Db.Where("email = ?", userData.Email).First(&user)
	if existedUser.RowsAffected > 0 {
		// User already exists, construct and return an error
		return false, errors.New("user already exists")
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

func (u *UserService) SignIn(userData dtos.SignInDTO) (SignInResponse, error) {
	validatorError := u.Validate.Struct(userData)
	if validatorError != nil {
		return SignInResponse{}, validatorError
	}

	userQuery := u.Db.Where("email = ?", userData.Email).First(&user)
	if userQuery.RowsAffected == 0 {
		return SignInResponse{}, errors.New("email or password is not valid")
	}

	isValidPassword := utils.CheckPasswordHash(userData.Password, user.Password)
	if !isValidPassword {
		return SignInResponse{}, errors.New("email or password is not valid")
	}

	payload := utils.TokenPayload{
		Id:    user.ID,
		Email: user.Email,
	}

	// Generate Token
	accessToken, errAccessToken := utils.GenerateToken(payload, configs.JwtSecret, 15*60)
	helpers.ErrorPanic(errAccessToken)

	refreshToken, errRefreshToken := utils.GenerateToken(payload, configs.JwtSecret, 3600)
	helpers.ErrorPanic(errRefreshToken)

	return SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

func GetValidator() *validator.Validate {
	return validate
}

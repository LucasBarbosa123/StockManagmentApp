package utilities_controllers

import (
	"github.com/go-playground/validator/v10"
)

func InitValidator() validator.Validate {
	var validate *validator.Validate
	validate = validator.New()

	return *validate
}

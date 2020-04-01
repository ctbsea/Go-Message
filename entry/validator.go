package entry

import (
	"github.com/go-playground/validator/v10"
)

func InitValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(UserStructLevelValidation, LoginParams{})
	validate.RegisterStructValidation(UserRegisterStructLevelValidation, RegisterParams{})
	return validate
}

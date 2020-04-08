package entry

import (
	"github.com/ctbsea/Go-Message/entry/entryRequest"
	"github.com/go-playground/validator/v10"
)

func InitValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(entryRequest.UserStructLevelValidation, entryRequest.LoginParams{})
	validate.RegisterStructValidation(entryRequest.UserRegisterStructLevelValidation, entryRequest.RegisterParams{})
	return validate
}

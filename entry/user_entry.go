package entry

import "github.com/go-playground/validator/v10"

type LoginParams struct {
	UserName       string `json:"userName" validate:"required"`
	UserPass       string `json:"userPass" validate:"required"`
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(LoginParams)
	if len(user.UserName) > 20 && len(user.UserPass) > 50 {
		sl.ReportError(user.UserName, "用户", "fname", "fnameorlname", "")
		sl.ReportError(user.UserPass, "密码", "lname", "fnameorlname", "")
	}
}


type RegisterParams struct {
	UserName       string `json:"userName" validate:"required"`
	UserPass       string `json:"userPass" validate:"required"`
}

func UserRegisterStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(RegisterParams)
	if len(user.UserName) > 20 && len(user.UserPass) > 50 {
		sl.ReportError(user.UserName, "用户", "fname", "fnameorlname", "")
		sl.ReportError(user.UserPass, "密码", "lname", "fnameorlname", "")
	}
}
package entryRequest

import "github.com/go-playground/validator/v10"

type LoginParams struct {
	UserName string `json:"userName"`
	UserPass string `json:"userPass"`
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(LoginParams)
	if len(user.UserName) < 6 {
		sl.ReportError(user.UserName, "userName", "UserName", "用户名需6~20位", "")
	}
	if len(user.UserPass) < 6 {
		sl.ReportError(user.UserPass, "userPass", "UserPass", "用户密码错误", "")
	}
}

type RegisterParams struct {
	UserName string `json:"userName" validate:"required"`
	UserPass string `json:"userPass" validate:"required"`
}

func UserRegisterStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(RegisterParams)
	if len(user.UserName) < 6 || len(user.UserName) > 20 {
		sl.ReportError(user.UserName, "userName", "UserName", "用户名需6~20位", "")
	}
	if len(user.UserPass) < 6 {
		sl.ReportError(user.UserPass, "userPass", "UserPass", "用户密码错误", "")
	}
}

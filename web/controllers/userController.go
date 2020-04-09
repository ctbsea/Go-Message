package controllers

import (
	"fmt"
	"github.com/ctbsea/Go-Message/entry"
	"github.com/ctbsea/Go-Message/entry/entryRequest"
	"github.com/ctbsea/Go-Message/services"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
)

type UserController struct {
	Service  *services.Service
	Validate *validator.Validate
}

func (u *UserController) Login(ctx iris.Context) {
	var loginParams entryRequest.LoginParams
	requestParams , err := entryRequest.RequestParams(ctx, u.Validate, &loginParams)
	fmt.Println(requestParams , err)
	if err.Code != 0 {
		ctx.JSON(entry.Response{
			err.Code,
			err.Msg,
			nil,
		})
		return
	}
	loginParams = requestParams.(entryRequest.LoginParams)
	params := make(map[string]string)
	params["user_name"] = loginParams.UserName
	params["user_pass"] = loginParams.UserPass
	params["login_ip"] = ctx.Request().RemoteAddr
	res, code := u.Service.UserService.Login(params)
	if code != entry.SUCCESS {
		ctx.JSON(entry.Response{
			code,
			"",
			nil,
		})
		return
	}
	ctx.JSON(entry.Response{
		entry.SUCCESS,
		"",
		res,
	})
}

func (u *UserController) Register(ctx iris.Context) entry.Response {
	var registerParams entryRequest.RegisterParams
	if err := ctx.ReadJSON(&registerParams); err != nil {
		return entry.Response{
			entry.INVAILD_PARAM,
			err.Error(),
			nil,
		}
	}
	err := u.Validate.Struct(registerParams)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return entry.Response{
				entry.INVAILD_PARAM,
				err.Error(),
				nil,
			}
		}
	}
	params := make(map[string]string)
	params["user_name"] = registerParams.UserName
	params["user_pass"] = registerParams.UserPass
	res, code := u.Service.UserService.Register(params)
	if code != entry.SUCCESS {
		return entry.Response{
			code,
			"",
			nil,
		}
	}
	return entry.Response{
		entry.SUCCESS,
		"",
		res,
	}
}

package controllers

import (
	"github.com/ctbsea/Go-Message/entry"
	"github.com/ctbsea/Go-Message/services"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
)

type MemberController struct {
	Service  *services.Service
	Validate *validator.Validate
}

func (u *MemberController) Detail(ctx iris.Context) entry.Response {
	return entry.Response{
		entry.SUCCESS,
		"ok",
		nil,
	}
}

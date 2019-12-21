package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type User struct {
}

func (u *User) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login", "Login")
}

func (u *User) Login(ctx iris.Context) {
	fmt.Println(1)
}

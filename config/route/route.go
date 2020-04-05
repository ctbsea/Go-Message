package route

import (
	"github.com/ctbsea/Go-Message/services"
	"github.com/ctbsea/Go-Message/web/controllers"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func Router(app *iris.Application, service *services.Service, validate *validator.Validate) {
	app.PartyFunc("/user" , func(r iris.Party) {
		user := controllers.UserController{Service: service, Validate: validate}
		r.Post("/login", hero.Handler(user.Login))
		r.Post("/register", hero.Handler(user.Register))
	})
}

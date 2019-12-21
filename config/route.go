package config

import (
	"github.com/ctbsea/Go-Message/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Router(app *iris.Application) {
	mvc.Configure(app.Party("/user"), mvcUser)

}

func mvcUser(app *mvc.Application) {
	//app.Router.Use();
	app.Handle(new(controllers.User))
}

package main

import (
	"context"
	"github.com/ctbsea/Go-Message/config"
	"github.com/ctbsea/Go-Message/config/db"
	"github.com/ctbsea/Go-Message/config/route"
	"github.com/ctbsea/Go-Message/datamodels"
	"github.com/ctbsea/Go-Message/entry"
	"github.com/ctbsea/Go-Message/gateway"
	"github.com/ctbsea/Go-Message/repositories"
	"github.com/ctbsea/Go-Message/services"
	log2 "github.com/ctbsea/Go-Message/util/log"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
	"go.uber.org/dig"
	"log"
	_ "net/http/pprof"
	"time"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(App)
	container.Provide(config.InitAllConfig)
	container.Provide(db.InitDb)
	container.Provide(datamodels.InitModels)
	container.Provide(repositories.InitRep)
	container.Provide(services.InitService)
	container.Provide(entry.InitValidator)
	container.Provide(services.NewUserService)
	return container
}

func App() *iris.Application {
	app := iris.New()
	return app
}

func run(
	app *iris.Application,
	service *services.Service,
	validate *validator.Validate,
	config2 config.Config) {
	//日志
	r, close := log2.NewRequestLogger(config2)
	defer close()
	app.Use(r)
	//限速器
	app.Use(gateway.NewLimiter(config2.GateWay.LimiterOneSec))
	//路由
	route.Router(app, service, validate)
	//性能日志
	//gateway.NewPprof(app)
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	})
	ipPort := config2.Run.IP + ":" + config2.Run.Port
	app.Run(iris.Addr(ipPort), iris.WithoutInterruptHandler)
}

func main() {
	container := BuildContainer()
	err := container.Invoke(run)
	if err != nil {
		log.Fatal(err)
	}
}

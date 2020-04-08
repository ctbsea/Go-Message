package gateway

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/ctbsea/Go-Message/config"
	"github.com/ctbsea/Go-Message/util/log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GateWay(app *iris.Application, config config.Config) ([]context.Handler, []func()(error) ) {
	Gate := []context.Handler{}
	deferFunc := []func()(error){}
	//限速器
	Gate = append(Gate, NewLimiter(config.GateWay.LimiterOneSec))
	//日志
	r, close := log.NewRequestLogger(config)
	Gate = append(Gate, r)
	deferFunc = append(deferFunc, close)
	app.Use(r)
	if config.Env.Env == "dev" {
		//性能日志
		NewPprof(app, config)
		//文档
		ApiDoc(app)
		Gate = append(Gate, irisyaag.New())
	}
	return Gate, deferFunc
}

package globalMiddle

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/ctbsea/Go-Message/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go.uber.org/zap"
)

func GateWay(app *iris.Application, config config.Config, zapLogger *zap.SugaredLogger) ([]context.Handler, []func() error) {
	gate := []context.Handler{}
	deferFunc := []func() error{}
	//限速器
	gate = append(gate, NewLimiter(config.GateWay.LimiterOneSec))
	//日志
	gate = append(gate, AccessLog(zapLogger))
	deferFunc = append(deferFunc, zapLogger.Sync)
	if config.Env.Env == "dev" {
		//性能日志
		NewPprof(app, config)
		//文档
		ApiDoc()
		gate = append(gate, irisyaag.New())
	}
	return gate, deferFunc
}

package globalMiddle

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/ctbsea/Go-Message/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go.uber.org/zap"
)

func GateWay(app *iris.Application, config config.Config, zapLogger *zap.Logger) ([]context.Handler, []func() error) {
	gate := []context.Handler{}
	deferFunc := []func() error{}
	//限速器
	gate = append(gate, NewLimiter(config.GateWay.LimiterOneSec))
	//日志
	gate = append(gate, AccessLog(zapLogger))
	deferFunc = append(deferFunc, zapLogger.Sync)
	if config.Env.OpenDoc  {
		//文档
		ApiDoc()
		gate = append(gate, irisyaag.New())
	}
	if config.Env.OpenPprof {
		//性能日志
		NewPprof(app, config)
	}
	return gate, deferFunc
}

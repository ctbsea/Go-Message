package gateway

import (
	"github.com/ctbsea/Go-Message/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/handlerconv"
	"net/http/pprof"
)

func NewPprof(app *iris.Application, config config.Config) {
	//测试环境
	if config.Env.Env == "dev" {
		cmdlineHandler := handlerconv.FromStd(pprof.Index)
		profileHandler := handlerconv.FromStd(pprof.Profile)
		symbolHandler := handlerconv.FromStd(pprof.Symbol)
		goroutineHandler := handlerconv.FromStd(pprof.Handler("goroutine"))
		heapHandler := handlerconv.FromStd(pprof.Handler("heap"))
		threadcreateHandler := handlerconv.FromStd(pprof.Handler("threadcreate"))
		debugBlockHandler := handlerconv.FromStd(pprof.Handler("block"))
		traceHandler := handlerconv.FromStd(pprof.Trace)
		mutexHandler := handlerconv.FromStd(pprof.Handler("mutex"))

		app.Any("/debug/pprof/", cmdlineHandler)
		app.Any("/debug/pprof/cmdline", cmdlineHandler)
		app.Any("/debug/pprof/profile", profileHandler)
		app.Any("/debug/pprof/trace", traceHandler)
		app.Any("/debug/pprof/symbol", symbolHandler)
		app.Any("/debug/pprof/heap", heapHandler)
		app.Any("/debug/pprof/threadcreate", threadcreateHandler)
		app.Any("/debug/pprof/block", debugBlockHandler)
		app.Any("/debug/pprof/goroutine", goroutineHandler)
		app.Any("/debug/pprof/mutex", mutexHandler)
	}
}

package gateway

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
)

func NewPprof(app *iris.Application)  {
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)
	app.Any("/debug/pprof/cmdline", p)
	app.Any("/debug/pprof/profile", p)
	app.Any("/debug/pprof/trace", p)
	app.Any("/debug/pprof/symbol", p)
}
package gateway

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
)

func NewPprof(app *iris.Application)  {
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)
}
package globalMiddle

import (
	"github.com/ctbsea/Go-Message/util/logger"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go.uber.org/zap"
	"strings"
)

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func AccessLog(zapLogger *zap.SugaredLogger) context.Handler {
	c := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	}
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	return logger.NewAccessLog(zapLogger, c)
}

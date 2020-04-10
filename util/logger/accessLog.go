// Package logger provides request logging via middleware. See _examples/http_request/request-logger
package logger

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"

	"github.com/kataras/iris/context"
)

type requestLoggerMiddleware struct {
	logger *zap.SugaredLogger
	config Config
}

func NewAccessLog(logger *zap.SugaredLogger, cfg ...Config) context.Handler {
	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	c.buildSkipper()
	l := &requestLoggerMiddleware{config: c, logger: logger}

	return l.ServeHTTP
}

func (l *requestLoggerMiddleware) ServeHTTP(ctx context.Context) {
	if l.config.skip != nil {
		if l.config.skip(ctx) {
			ctx.Next()
			return
		}
	}
	//all except latency to string
	var status, ip, method, path string
	var latency time.Duration
	var startTime, endTime time.Time
	startTime = time.Now()

	ctx.Next()

	//no time.Since in order to format it well after
	endTime = time.Now()
	latency = endTime.Sub(startTime)

	if l.config.Status {
		status = strconv.Itoa(ctx.GetStatusCode())
	}

	if l.config.IP {
		ip = ctx.RemoteAddr()
	}

	if l.config.Method {
		method = ctx.Method()
	}

	if l.config.Path {
		if l.config.Query {
			path = ctx.Request().URL.RequestURI()
		} else {
			path = ctx.Path()
		}
	}

	var message interface{}
	if ctxKeys := l.config.MessageContextKeys; len(ctxKeys) > 0 {
		for _, key := range ctxKeys {
			msg := ctx.Values().Get(key)
			if message == nil {
				message = msg
			} else {
				message = fmt.Sprintf(" %v %v", message, msg)
			}
		}
	}
	var headerMessage interface{}
	if headerKeys := l.config.MessageHeaderKeys; len(headerKeys) > 0 {
		for _, key := range headerKeys {
			msg := ctx.GetHeader(key)
			if headerMessage == nil {
				headerMessage = msg
			} else {
				headerMessage = fmt.Sprintf(" %v %v", headerMessage, msg)
			}
		}
	}

	// no new line, the framework's logger is responsible how to render each log.
	line := fmt.Sprintf("%v %4v %s %s %s", status, latency, ip, method, path)
	if message != nil {
		line += fmt.Sprintf(" %v", message)
	}

	if headerMessage != nil {
		line += fmt.Sprintf(" %v", headerMessage)
	}
	fmt.Println(line)
	l.logger.Info(line)
}

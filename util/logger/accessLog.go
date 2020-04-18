// Package logger provides request logging via middleware. See _examples/http_request/request-logger
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
	"time"

	"github.com/kataras/iris/context"
)

type requestLoggerMiddleware struct {
	logger *zap.Logger
	config Config
}

func NewAccessLog(logger *zap.Logger, cfg ...Config) context.Handler {
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
	latency = endTime.Sub(startTime) * 1000

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

	var message []interface{}
	if ctxKeys := l.config.MessageContextKeys; len(ctxKeys) > 0 {
		for _, key := range ctxKeys {
			message = append(message, ctx.Values().Get(key))
		}
	}
	var headerMessage []interface{}
	if headerKeys := l.config.MessageHeaderKeys; len(headerKeys) > 0 {
		for _, key := range headerKeys {
			headerMessage = append(headerMessage, ctx.GetHeader(key))
		}
	}

	filed := []zapcore.Field{
		zap.String("status", status),
		zap.Duration("latency", latency),
		zap.String("ip", ip),
		zap.String("method", method),
		zap.String("path", path),
		zap.Any("message", message),
		zap.Any("headerMessage", headerMessage),
	}
	l.logger.Info("AccessLog", filed...)
}

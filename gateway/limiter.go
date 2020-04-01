package util

import (
	"context"
	"github.com/kataras/iris"
	"golang.org/x/time/rate"
	"time"
)

func NewLimiter(c iris.Context)  {
	limit := rate.Every(100 * time.Millisecond)
	l := rate.NewLimiter(limit, 1)
	for {
		l.Wait(c.(context.Context))
	}
}
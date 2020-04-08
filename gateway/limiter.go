package gateway

//优化成全局以及单路由
import (
	"fmt"
	"github.com/ctbsea/Go-Message/entry/entryReturn"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris/context"
	"time"
)

func NewLimiter(num float64) context.Handler {
	fmt.Printf("gateway-Limiter : start limiter max %f  every one sec \n", num)
	lmt := tollbooth.NewLimiter(num, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetMessage("You have reached maximum request limit.")
	lmt.SetMessageContentType("application/json; charset=utf-8")
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	return limitHandler(lmt)
}

func limitHandler(l *limiter.Limiter) context.Handler {
	return func(ctx context.Context) {
		httpError := tollbooth.LimitByRequest(l, ctx.ResponseWriter(), ctx.Request())
		if httpError != nil {
			entryReturn.CtxResException(ctx,
				&entryReturn.BaseStruct{httpError.StatusCode, httpError.Message, nil})
			return
		}
		ctx.Next()
	}
}

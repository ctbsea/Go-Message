package entryReturn

import (
	"encoding/json"
	"github.com/kataras/iris"
)

type BaseStruct struct {
	Code int
	Msg  string
	Data interface{}
}

func Res(code int, msg string, data interface{}) *BaseStruct {
	return &BaseStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

//中断ctx 并返回一般作为异常处理
func CtxResException(ctx iris.Context, res *BaseStruct) {
	ctx.ContentType("application/json")
	ctx.StatusCode(200)
	str, _ := json.Marshal(res)
	ctx.WriteString(string(str))
	ctx.StopExecution()
	return
}

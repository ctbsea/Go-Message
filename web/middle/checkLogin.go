package middle

import (
	"github.com/ctbsea/Go-Message/entry"
	"github.com/ctbsea/Go-Message/entry/entryReturn"
	"github.com/ctbsea/Go-Message/util/jwtfuc"
	"github.com/kataras/iris"
)

func CheckLogin(ctx iris.Context) {
	token := ctx.GetHeader(jwtfuc.JWT_KEY)
	loginInfo, code := jwtfuc.Check(token)
	if code != entry.SUCCESS {
		entryReturn.CtxResException(ctx , entryReturn.Res(code , "" ,nil))
		return
	}
	ctx.Values().Set("loginInfo", loginInfo)
	ctx.Next()
}
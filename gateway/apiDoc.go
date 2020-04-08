package gateway

import (
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

func ApiDoc(app *iris.Application) {
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Doc",
		DocPath:  "./web/apiDoc.html",
		BaseUrls: map[string]string{"Production": "/", "Staging": ""},
	})
}

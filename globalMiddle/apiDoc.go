package globalMiddle

import (
	"github.com/betacraft/yaag/yaag"
)

func ApiDoc() {
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Doc",
		DocPath:  "./web/apiDoc.html",
		BaseUrls: map[string]string{"Production": "/", "Staging": ""},
	})
}

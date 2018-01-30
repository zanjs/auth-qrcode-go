package middleware

import (
	"github.com/kataras/iris"
)

// Before is
func Before(ctx iris.Context) {

	if ctx.Method() == "OPTIONS" {
		ctx.WriteString("ok")
		return
	}

	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)

	// for the sake of simplicity, in order see the logs at the ./_today_.txt
	println(ctx.URLParams())
	ctx.Application().Logger().Info("Request path: " + ctx.Path())
	ctx.Next()
}

// After is
func After(ctx iris.Context) {
	println("After the indexHandler or contactHandler")
}

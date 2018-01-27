package handler

import (
	"github.com/kataras/iris"
)

// IndexHandler is
func IndexHandler(ctx iris.Context) {
	ResponseJSON(ctx, "hello")

	println("Inside indexHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")
	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after" handler registered via `Done`.
}

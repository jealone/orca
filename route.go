package http

import (
	"github.com/valyala/fasthttp"
)

type Routes func(Router)

func Route(routes Routes) *FasthttpRoute {
	root := NewRouter()
	root.MethodNotAllowed = notAllowed
	root.NotFound = notFound
	root.PanicHandler = panicHandler
	routes(root)
	return root
}

func panicHandler(ctx *fasthttp.RequestCtx, rcp interface{}) {
	badRequestResponse(ctx, fasthttp.StatusInternalServerError)
	ctx.Logger().Printf("panic:%s", rcp)
}

//405
func notAllowed(ctx *fasthttp.RequestCtx) {
	badRequestResponse(ctx, fasthttp.StatusMethodNotAllowed)
}

//404
func notFound(ctx *fasthttp.RequestCtx) {
	badRequestResponse(ctx, fasthttp.StatusNotFound)
}

func badRequestResponse(ctx *fasthttp.RequestCtx, statusCode int) {
	message := fasthttp.StatusMessage(statusCode)
	ctx.Error(message, statusCode)
	ctx.Logger().Printf(message)
}

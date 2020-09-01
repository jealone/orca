package orca

func panicHandler(ctx *RequestCtx, rcp interface{}) {
	badRequestResponse(ctx, StatusMethodNotAllowed)
	ctx.Logger().Printf("panic:%s", rcp)
}

//405
func notAllowed(ctx *RequestCtx) {
	badRequestResponse(ctx, StatusMethodNotAllowed)
}

//404
func notFound(ctx *RequestCtx) {
	badRequestResponse(ctx, StatusNotFound)
}

package orca

func Middleware(middleware MiddlewareHandler, handler Handler) Handler {
	if nil == middleware {
		return handler
	}
	return func(ctx *RequestCtx) {
		if err := middleware(ctx); nil != err {
			return
		}
		handler(ctx)
	}
}

func LambdaMiddleware(handler MiddlewareHandler, h ...MiddlewareHandler) MiddlewareHandler {

	var handlers []MiddlewareHandler
	for _, m := range h {
		if nil != m {
			handlers = append(handlers, m)
		}
	}

	if 0 == len(handlers) {
		return handler
	}

	var middleware []MiddlewareHandler

	if nil != handler {
		middleware = append(middleware, handler)
	}

	middleware = append(middleware, handlers...)

	return func(ctx *RequestCtx) error {
		for _, m := range middleware {
			if err := m(ctx); nil != err {
				return err
			}
		}
		return nil
	}

}

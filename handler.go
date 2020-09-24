package orca

type lambda func(handler Handler, handlers []Handler) Handler

func BeforeHandler(handler Handler, handlers ...Handler) Handler {
	return lambdaHandler(before)(handler, handlers)
}

func AfterHandler(handler Handler, handlers ...Handler) Handler {
	return lambdaHandler(after)(handler, handlers)
}

func after(handler Handler, handlers []Handler) Handler {
	if 0 == len(handlers) {
		return handler
	} else if nil == handlers[0] {
		return handler
	} else if nil == handler {
		return handlers[0]
	} else {
		return func(ctx *RequestCtx) {
			handler(ctx)
			handlers[0](ctx)
		}
	}
}

func before(handler Handler, handlers []Handler) Handler {
	if 0 == len(handlers) {
		return handler
	} else if nil == handlers[0] {
		return handler
	} else if nil == handler {
		return handlers[0]
	} else {
		return func(ctx *RequestCtx) {
			handlers[0](ctx)
			handler(ctx)
		}
	}
}

func lambdaHandler(l lambda) lambda {
	return func(handler Handler, handlers []Handler) Handler {
		if handlers == nil || len(handlers) == 0 {
			return handler
		} else {
			return lambdaHandler(l)(l(handler, handlers), handlers[1:])
		}
	}
}

//func BeforeFilter(handler Handler, filters ...Handler) Handler {
//	if nil == handler && 0 == len(filters) {
//		return nil
//	} else if 0 == len(filters) {
//		return handler
//	} else if nil == handler {
//		return unfold(filters...)
//	} else {
//		return func(ctx *RequestCtx) {
//			for _, filter := range filters {
//				filter(ctx)
//			}
//			handler(ctx)
//		}
//	}
//}
//
//func AfterFilter(handler Handler, filters ...Handler) Handler {
//	if nil == handler && 0 == len(filters) {
//		return nil
//	} else if nil == handler {
//		return unfold(filters...)
//	} else if 0 == len(filters) {
//		return handler
//	} else {
//		return func(ctx *RequestCtx) {
//			handler(ctx)
//			for _, filter := range filters {
//				filter(ctx)
//			}
//		}
//	}
//}
//
//func unfold(filters ...Handler) Handler {
//	if 0 == len(filters) {
//		return nil
//	}
//	return func(ctx *RequestCtx) {
//		for _, filter := range filters {
//			filter(ctx)
//		}
//	}
//}

package orca

//type lambda func(filter Filter, filters []Filter) Filter
//
//func BeforeFilter(handler Filter, filters ...Filter) Filter {
//	return lambdaFilter(before)(handler, filters)
//}
//
//func AfterFilter(handler Filter, filters ...Filter) Filter {
//	return lambdaFilter(after)(handler, filters)
//}
//
//func after(filter Filter, filters []Filter) Filter {
//	return func(ctx *RequestCtx) {
//		if nil != filter {
//			filter(ctx)
//		}
//		if nil != filters[0] {
//			filters[0](ctx)
//		}
//	}
//}
//
//func before(filter Filter, filters []Filter) Filter {
//	return func(ctx *RequestCtx) {
//		if nil != filters[0] {
//			filters[0](ctx)
//		}
//		if nil != filter {
//			filter(ctx)
//		}
//	}
//}
//
//func lambdaFilter(l lambda) lambda {
//	return func(handler Filter, filters []Filter) Filter {
//		if filters == nil || len(filters) == 0 {
//			return handler
//		} else {
//			return lambdaFilter(l)(l(handler, filters), filters[1:])
//		}
//	}
//}

func BeforeFilter(handler Filter, filters ...Filter) Filter {
	return func(ctx *RequestCtx) {
		for _, filter := range filters {
			filter(ctx)
		}
		handler(ctx)
	}
}

func AfterFilter(handler Filter, filters ...Filter) Filter {
	return func(ctx *RequestCtx) {
		handler(ctx)
		for _, filter := range filters {
			filter(ctx)
		}
	}
}

func unfold(filters ...Filter) Filter {
	return func(ctx *RequestCtx) {
		for _, filter := range filters {
			filter(ctx)
		}
	}
}

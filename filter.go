package orca

type lambda func(filter Filter, filters []Filter) Filter

func BeforeFilter(handler Filter, filters ...Filter) Filter {
	return lambdaFilter(before)(handler, filters)
}

func AfterFilter(handler Filter, filters ...Filter) Filter {
	return lambdaFilter(after)(handler, filters)
}

func after(filter Filter, filters []Filter) Filter {
	if 0 == len(filters) {
		return filter
	} else if nil == filters[0] {
		return filter
	} else if nil == filter {
		return filters[0]
	} else {
		return func(ctx *RequestCtx) {
			filter(ctx)
			filters[0](ctx)
		}
	}
}

func before(filter Filter, filters []Filter) Filter {
	if 0 == len(filters) {
		return filter
	} else if nil == filters[0] {
		return filter
	} else if nil == filter {
		return filters[0]
	} else {
		return func(ctx *RequestCtx) {
			filters[0](ctx)
			filter(ctx)
		}
	}
}

func lambdaFilter(l lambda) lambda {
	return func(handler Filter, filters []Filter) Filter {
		if filters == nil || len(filters) == 0 {
			return handler
		} else {
			return lambdaFilter(l)(l(handler, filters), filters[1:])
		}
	}
}

//func BeforeFilter(handler Filter, filters ...Filter) Filter {
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
//func AfterFilter(handler Filter, filters ...Filter) Filter {
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
//func unfold(filters ...Filter) Filter {
//	if 0 == len(filters) {
//		return nil
//	}
//	return func(ctx *RequestCtx) {
//		for _, filter := range filters {
//			filter(ctx)
//		}
//	}
//}

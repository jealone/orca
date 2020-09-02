package orca

var (
	afterFilters  = lambdaFilter(after)
	beforeFilters = lambdaFilter(before)
)

type lambda func(filter Filter, filters []Filter) Filter

func BeforeFilter(handler Filter, filters ...Filter) Filter {
	return beforeFilters(handler, filters)
}

func AfterFilter(handler Filter, filters ...Filter) Filter {
	return afterFilters(handler, filters)
}

func after(filter Filter, filters []Filter) Filter {
	return func(ctx *RequestCtx) {
		filter(ctx)
		filters[0](ctx)
	}
}

func before(filter Filter, filters []Filter) Filter {
	return func(ctx *RequestCtx) {
		filters[0](ctx)
		filter(ctx)
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

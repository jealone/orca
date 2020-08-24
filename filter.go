package orca

import (
	"github.com/valyala/fasthttp"
)

type Filter fasthttp.RequestHandler

func BeforeFilter(handler fasthttp.RequestHandler, filters ...Filter) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		for _, filter := range filters {
			filter(ctx)
		}
		handler(ctx)
	}
}

func AfterFilter(handler fasthttp.RequestHandler, filters ...Filter) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		for _, filter := range filters {
			filter(ctx)
		}
	}
}

func AccessLogFilter(ctx *fasthttp.RequestCtx) {
	accessLog(ctx)
}

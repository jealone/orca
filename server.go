package orca

import (
	"github.com/valyala/fasthttp"
)

type RouterAdapter interface {
	Handler(*fasthttp.RequestCtx)
}

func NewHttpServer(r RouterAdapter, Logger fasthttp.Logger) *fasthttp.Server {
	// @todo 通过配置实例化server
	return &fasthttp.Server{
		Handler: AfterFilter(r.Handler, AccessLogFilter),
		Logger:  Logger,
	}
}

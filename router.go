package orca

import (
	"github.com/valyala/fasthttp"
)

type Router interface {
	GET(path string, handler fasthttp.RequestHandler)
	HEAD(path string, handler fasthttp.RequestHandler)
	OPTIONS(path string, handler fasthttp.RequestHandler)
	POST(path string, handler fasthttp.RequestHandler)
	PATCH(path string, handler fasthttp.RequestHandler)
	DELETE(path string, handler fasthttp.RequestHandler)
	ANY(path string, handler fasthttp.RequestHandler)
	Handle(method, path string, handler fasthttp.RequestHandler)
	Group(path string) Router
}

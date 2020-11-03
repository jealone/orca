package orca

import (
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v3"
)

type (
	RequestCtx     = fasthttp.RequestCtx
	Handler        = fasthttp.RequestHandler
	Server         = fasthttp.Server
	RequestHandler = fasthttp.RequestHandler
	Logger         = fasthttp.Logger
	HttpClient     = fasthttp.Client

	Args     = fasthttp.Args
	URI      = fasthttp.URI
	Request  = fasthttp.Request
	Response = fasthttp.Response

	RequestHeader = fasthttp.RequestHeader

	YamlNode = yaml.Node
)

type MiddlewareHandler func(*RequestCtx) error

const (
	StatusInternalServerError = fasthttp.StatusInternalServerError
	StatusNotFound            = fasthttp.StatusNotFound
	StatusMethodNotAllowed    = fasthttp.StatusMethodNotAllowed
)

func badRequestResponse(ctx *RequestCtx, statusCode int) {
	message := fasthttp.StatusMessage(statusCode)
	ctx.Error(message, statusCode)
	ctx.Logger().Printf(message)
}

func errorHandler(ctx *RequestCtx, err error) {
	ctx.Error(fasthttp.StatusMessage(StatusInternalServerError), StatusInternalServerError)
	ctx.Logger().Printf(err.Error())
}

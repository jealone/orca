package orca

import (
	"github.com/valyala/fasthttp"
)

type (
	RequestCtx     = fasthttp.RequestCtx
	Filter         = fasthttp.RequestHandler
	Server         = fasthttp.Server
	RequestHandler = fasthttp.RequestHandler
	Logger         = fasthttp.Logger
)

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

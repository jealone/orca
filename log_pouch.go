package orca

import (
	"io"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	AccessLogger = zap.Logger
	Field        = zap.Field
	Rotater      = lumberjack.Logger
)

func log(logger *AccessLogger, ctx *RequestCtx) {
	now := time.Now()
	statusCode := ctx.Response.StatusCode()

	fields := []Field{
		zap.Uint64("ID", ctx.ID()),
		zap.ByteString("Host", ctx.Host()),
		zap.String("Remote", ctx.RemoteAddr().String()),
		zap.String("Local", ctx.LocalAddr().String()),
		zap.ByteString("Method", ctx.Method()),
		zap.ByteString("URI", ctx.RequestURI()),
		zap.ByteString("AuthUid", ctx.Request.Header.Peek("auth_uid")),
		zap.ByteString("Referer", ctx.Referer()),
		zap.ByteString("UA", ctx.Request.Header.UserAgent()),
		zap.Duration("RequestElapse", now.Sub(ctx.Time())),
		zap.Int("Status", ctx.Response.StatusCode()),
	}

	if statusCode < 400 {
		logger.Info(
			fasthttp.StatusMessage(ctx.Response.StatusCode()),
			fields...,
		)
	} else {
		logger.Error(
			fasthttp.StatusMessage(ctx.Response.StatusCode()),
			fields...,
		)
	}
}

func newLogger(writer io.Writer) *AccessLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writerSync := zapcore.AddSync(writer)
	core := zapcore.NewCore(encoder, writerSync, zapcore.InfoLevel)
	return zap.New(core)
}

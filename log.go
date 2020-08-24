package orca

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	accessLogger *zap.Logger
)

func accessLog(ctx *fasthttp.RequestCtx) {
	log(accessLogger, ctx)
}

func log(logger *zap.Logger, ctx *fasthttp.RequestCtx) {
	now := time.Now()
	statusCode := ctx.Response.StatusCode()

	fields := []zap.Field{
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
	//logger.Info(
	//
	//	zap.Uint64("ID", ctx.ID()),
	//	zap.ByteString("Host", ctx.Host()),
	//	zap.String("Remote", ctx.RemoteAddr().String()),
	//	zap.String("Local", ctx.LocalAddr().String()),
	//	zap.ByteString("Method", ctx.Method()),
	//	zap.ByteString("URI", ctx.RequestURI()),
	//	zap.ByteString("AuthUid", ctx.Request.Header.Peek("auth_uid")),
	//	zap.ByteString("Referer", ctx.Referer()),
	//	zap.ByteString("UA", ctx.Request.Header.UserAgent()),
	//	zap.Duration("RequestElapse", now.Sub(ctx.Time())),
	//	zap.Int("Status", ctx.Response.StatusCode()),
	//)
}

func NewLogger(dir string) {
	abs, err := filepath.Abs(dir)

	if nil != err {
		// 创建目录
		fmt.Printf("log dir error:%s\n", err)
		os.Exit(0)
	}

	accessLogFile := filepath.Join(abs, "orcaAccess.log")

	c, rotater := newFileNotify(accessLogFile, syscall.SIGHUP)

	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("recover_panic:%s", e)
			}
		}()
		for {
			select {
			case <-c:
				rotater.Rotate()
			}
		}

	}()

	accessLogger = newLogger(rotater)
}

func newFileNotify(logfile string, signals ...os.Signal) (chan os.Signal, *lumberjack.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("recover_panic:%s", e)
			}
		}()
		ticker := time.Tick(time.Second * 10)
		for {
			select {
			case <-ticker:
				_, err := os.Stat(logfile)
				if nil != err && !os.IsExist(err) {
					c <- syscall.SIGHUP
				}
			}
		}
	}()

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    100,
		MaxBackups: 4,
		MaxAge:     7,
		Compress:   false,
	}
	return c, lumberJackLogger
}

func newLogger(writer io.Writer) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writerSync := zapcore.AddSync(writer)
	core := zapcore.NewCore(encoder, writerSync, zapcore.DebugLevel)
	//return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	return zap.New(core)
}

func init() {
	NewLogger("./logs")
}

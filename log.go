package orca

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	accessLogger *AccessLogger
)

func accessLog(ctx *RequestCtx) {
	log(accessLogger, ctx)
}

func NewLogger(config *AccessLogConfig) {

	c, rotater := newFileNotify(config, syscall.SIGHUP)

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

func newFileNotify(config *AccessLogConfig, signals ...os.Signal) (chan os.Signal, *Rotater) {
	c := make(chan os.Signal, 1)
	file := config.GetLogfile()
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
				_, err := os.Stat(file)
				if nil != err && !os.IsExist(err) {
					c <- syscall.SIGHUP
				}
			}
		}
	}()

	rotater := &Rotater{
		Filename:   file,
		MaxSize:    config.GetMaxSize(),
		MaxBackups: config.GetMaxBackups(),
		MaxAge:     config.GetMaxAge(),
		Compress:   config.GetCompress(),
	}

	return c, rotater
}

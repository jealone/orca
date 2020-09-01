package orca

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	accessLogger *AccessLogger
)

func accessLog(ctx *RequestCtx) {
	log(accessLogger, ctx)
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

func newFileNotify(logfile string, signals ...os.Signal) (chan os.Signal, *Rotater) {
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

	rotater := &Rotater{
		Filename:   logfile,
		MaxSize:    100,
		MaxBackups: 4,
		MaxAge:     7,
		Compress:   false,
	}
	return c, rotater
}

func init() {
	NewLogger("./logs")
}

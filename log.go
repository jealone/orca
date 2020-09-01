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
	accessLogger         *AccessLogger
	defaultAccessLogName = "logs/access.log"
)

func accessLog(ctx *RequestCtx) {
	log(accessLogger, ctx)
}

func NewLogger(path string) {

	if "" == path {
		path = defaultAccessLogName
	}

	abs, err := filepath.Abs(path)

	if nil != err {
		// 创建目录
		fmt.Printf("log path error:%s\n", err)
		os.Exit(0)
	}

	info, err := os.Stat(abs)

	if nil != err {
		fmt.Printf("log path error:%s\n", err)
		os.Exit(0)
	}

	var accessLogFile string

	if info.IsDir() {
		accessLogFile = filepath.Join(abs, defaultAccessLogName)
	} else {
		accessLogFile = abs
	}

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

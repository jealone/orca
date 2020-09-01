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

func logFileCheck(path string) {
	abs, err := filepath.Abs(path)
	if nil != err {
		panic(fmt.Sprintf("get abs for file(%s) error: %s", path, err))

	}

	info, err := os.Stat(abs)

	if nil != err {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(abs), os.ModePerm)
			if nil != err {
				panic(fmt.Sprintf("log file stat error: %s", err))

			}
		} else {
			panic(fmt.Sprintf("log file stat error: %s", err))

		}
	}

	if info.IsDir() {
		panic(fmt.Sprintf("the specific log file (%s) is directory", abs))
	}

}

func newFileNotify(config *AccessLogConfig, signals ...os.Signal) (chan os.Signal, *Rotater) {
	file := config.GetLogfile()

	c := make(chan os.Signal, 1)
	logFileCheck(file)
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

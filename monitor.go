package orca

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	locker     sync.RWMutex
	statusCode = 503
)

func Healthy(status int) {
	locker.Lock()
	defer locker.Unlock()
	statusCode = status
}

func healthyCheck() int {
	locker.RLock()
	defer locker.RUnlock()
	return statusCode
}

func Monitor(conf *KeepaliveConfig, server *Server) <-chan struct{} {

	s := conf.GetSignal()
	c := make(chan os.Signal, len(s))
	done := make(chan struct{})

	signal.Notify(c, s...)

	go func() {
		<-c
		interval := conf.GetInterval()
		fmt.Println("terminate...")
		Healthy(conf.GetUnhealthy())
		fmt.Printf("change alive status(%d)\n", healthyCheck())
		time.Sleep(interval)
		err := server.Shutdown()
		if nil != err {
			fmt.Println("stop fatal")
		} else {
			fmt.Println("graceful stop...")
		}
		done <- struct{}{}
	}()

	go func() {
		err := KeepaliveServer(*conf)
		fmt.Printf("keepalive server err:%s\n", err)
	}()

	go func() {
		time.Sleep(conf.GetInterval())
		Healthy(conf.GetHealthy())
	}()

	return done
}

func KeepaliveServer(conf KeepaliveConfig) error {
	return fasthttp.ListenAndServe(conf.GetAddr(), func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case conf.GetPath():
			ctx.SetStatusCode(healthyCheck())
			_, _ = ctx.WriteString(conf.GetMsg())
			return
		default:
			ctx.Error("Forbidden Request", fasthttp.StatusForbidden)
		}
	})
}

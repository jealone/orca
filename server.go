package orca

import "fmt"

type RouterAdapter interface {
	Handler(*RequestCtx)
}

func NewHttpServe(c ServerConfig, Logger Logger, r RouterAdapter, options ...func(filter Filter) Filter) error {

	h := r.Handler

	for _, option := range options {
		h = option(h)
	}

	server := &Server{
		Handler:      h,
		Logger:       Logger,
		ErrorHandler: errorHandler,

		// 通过配置获取
		GetOnly:                      c.GetMethodOnly(),
		DisablePreParseMultipartForm: c.GetDisablePreParseMultipartForm(),
		ReduceMemoryUsage:            c.GetReduceMemoryUsage(),
		LogAllErrors:                 c.GetLogAllErrors(),

		TCPKeepalive:       c.GetTcp().GetTCPKeepalive(),
		TCPKeepalivePeriod: c.GetTcp().GetTCPKeepalivePeriod(),

		Concurrency:                        c.GetConn().GetConcurrency(),
		SleepWhenConcurrencyLimitsExceeded: c.GetConn().GetSleepWhenConcurrencyLimitsExceeded(),
		DisableKeepalive:                   c.GetConn().GetDisableKeepalive(),
		MaxConnsPerIP:                      c.GetConn().GetMaxConnsPerIP(),
		IdleTimeout:                        c.GetConn().GetIdleTimeout(),
		ReadTimeout:                        c.GetConn().GetReadTimeout(),
		WriteTimeout:                       c.GetConn().GetWriteTimeout(),

		ReadBufferSize:  c.GetBuffer().GetReadBufferSize(),
		WriteBufferSize: c.GetBuffer().GetWriteBufferSize(),

		Name:                          c.GetHeader().GetServer(),
		NoDefaultServerHeader:         c.GetHeader().GetNoDefaultServerHeader(),
		DisableHeaderNamesNormalizing: c.GetHeader().GetDisableHeaderNamesNormalizing(),
		NoDefaultDate:                 c.GetHeader().GetNoDefaultDate(),
		NoDefaultContentType:          c.GetHeader().GetNoDefaultContentType(),

		MaxRequestsPerConn: c.GetRequest().GetMaxRequestsPerConn(),
		MaxRequestBodySize: c.GetRequest().GetMaxRequestBodySize(),
	}

	fmt.Printf("http serve at %s\n", c.GetTcp().GetAddr())

	return server.ListenAndServe(c.GetTcp().GetAddr())
}

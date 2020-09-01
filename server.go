package orca

type RouterAdapter interface {
	Handler(*RequestCtx)
}

func NewHttpServer(c ServerConfig, r RouterAdapter, Logger Logger) *Server {
	return &Server{
		Handler:      AfterFilter(r.Handler, AccessLogFilter),
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
}

package orca

type RouterAdapter interface {
	Handler(*RequestCtx)
}

func NewHttpServer(r RouterAdapter, Logger Logger) *Server {
	// @todo 通过配置实例化server
	return &Server{
		Handler: AfterFilter(r.Handler, AccessLogFilter),
		Logger:  Logger,
		ErrorHandler: errorHandler,
	}
}

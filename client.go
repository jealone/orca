package orca

func NewClient(options ...Option) *HttpClient {
	c := &HttpClient{}
	for _, opt := range options {
		opt.apply(c)
	}
	return c
}

type Option interface {
	apply(*HttpClient)
}

type optionFunc func(*HttpClient)

func (f optionFunc) apply(c *HttpClient) {
	f(c)
}

func ApplyConfig(c *HttpClientConfig) Option {
	return optionFunc(func(client *HttpClient) {
		if nil != c {
			client.Name = c.GetName()
			client.NoDefaultUserAgentHeader = c.GetNoDefaultUserAgentHeader()
			client.MaxConnsPerHost = c.GetMaxConnsPerHost()
			client.MaxIdleConnDuration = c.GetMaxIdleConnDuration()
			client.MaxConnDuration = c.GetMaxConnDuration()
			client.MaxIdemponentCallAttempts = c.GetMaxIdemponentCallAttempts()
			client.ReadBufferSize = c.GetReadBufferSize()
			client.WriteBufferSize = c.GetWriteBufferSize()
			client.ReadTimeout = c.GetReadTimeout()
			client.WriteTimeout = c.GetWriteTimeout()
			client.MaxResponseBodySize = c.GetMaxResponseBodySize()
			client.DisableHeaderNamesNormalizing = c.GetDisableHeaderNamesNormalizing()
			client.DisablePathNormalizing = c.GetDisablePathNormalizing()
			client.MaxConnWaitTimeout = c.GetMaxConnWaitTimeout()
		}
	})
}

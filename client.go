package orca

import (
	"sync"

	"github.com/valyala/fasthttp"
)

func createPool(f func() interface{}) func() *sync.Pool {
	var (
		once sync.Once
		pool *sync.Pool
	)
	return func() *sync.Pool {
		once.Do(func() {
			pool = &sync.Pool{
				New: f,
			}
		})
		return pool
	}
}

func CreateClient(options ...Option) func() *Client {
	var (
		once   sync.Once
		client *Client
	)
	return func() *Client {
		once.Do(func() {
			client = &Client{
				http:     &HttpClient{},
				executor: exec,
				request:  request,
			}
			client = client.WithOptions(options...)
		})
		return client
	}
}

type Client struct {
	http     *HttpClient
	executor func(HttpClientHandler, *HttpClient, *Request, *Response)
	request  func(HttpClientRequestHandler, *HttpClient)
}

func (c *Client) clone() *Client {
	copied := *c
	return &copied
}

func (c *Client) WithOptions(opts ...Option) *Client {
	client := c.clone()
	for _, opt := range opts {
		opt.apply(client)
	}
	return client
}

func (c *Client) Do(handler HttpClientHandler) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)
	c.executor(handler, c.http, request, response)
}

func (c *Client) Request(handler HttpClientRequestHandler) {
	c.request(handler, c.http)
}

type HttpClientHandler func(client *HttpClient, request *Request, response *Response) error

type HttpClientRequestHandler func(client *HttpClient) (status int, err error)

type Option interface {
	apply(*Client)
}

type optionFunc func(*Client)

func (f optionFunc) apply(c *Client) {
	f(c)
}

func ApplyConfig(c *HttpClientConfig) Option {
	return optionFunc(func(client *Client) {
		if nil != c {
			client.http = &HttpClient{
				Name:                          c.GetName(),
				NoDefaultUserAgentHeader:      c.GetNoDefaultUserAgentHeader(),
				MaxConnsPerHost:               c.GetMaxConnsPerHost(),
				MaxIdleConnDuration:           c.GetMaxIdleConnDuration(),
				MaxConnDuration:               c.GetMaxConnDuration(),
				MaxIdemponentCallAttempts:     c.GetMaxIdemponentCallAttempts(),
				ReadBufferSize:                c.GetReadBufferSize(),
				WriteBufferSize:               c.GetWriteBufferSize(),
				ReadTimeout:                   c.GetReadTimeout(),
				WriteTimeout:                  c.GetWriteTimeout(),
				MaxResponseBodySize:           c.GetMaxResponseBodySize(),
				DisableHeaderNamesNormalizing: c.GetDisableHeaderNamesNormalizing(),
				DisablePathNormalizing:        c.GetDisablePathNormalizing(),
				MaxConnWaitTimeout:            c.GetMaxConnWaitTimeout(),
			}
		}
	})
}

func AddExecutor(e func(HttpClientHandler, *HttpClient, *Request, *Response)) Option {
	return optionFunc(func(client *Client) {
		if nil != e {
			client.executor = e
		}
	})
}

func AddRequest(r func(HttpClientRequestHandler, *HttpClient)) Option {
	return optionFunc(func(client *Client) {
		if nil != r {
			client.request = r
		}
	})
}

func exec(handler HttpClientHandler, client *HttpClient, request *Request, response *Response) {
	_ = handler(client, request, response)
}

func request(handler HttpClientRequestHandler, client *HttpClient) {
	_, _ = handler(client)
}

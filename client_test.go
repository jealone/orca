package orca

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

var defaultClient = &Client{
	http:     &HttpClient{},
	executor: exec,
	request:  request,
}

type fatal interface {
	Fatal(...interface{})
}

func initDefaultClient(t fatal) {

	conf, err := testClientConfig("tests/client.yml")

	if nil != err {
		t.Fatal(err)
	}

	defaultClient = NewClient(ApplyConfig(conf))
}

func TestClient_Get(t *testing.T) {

	defaultClient.Request(func(client *HttpClient) (status int, err error) {
		status, _, err = client.Get(nil, "http://api.bee.to/sso/")
		return
	})
}

func TestClient_Do(t *testing.T) {
	defaultClient.Do(func(client *HttpClient, request *Request, response *Response) error {
		request.SetRequestURI("http://api.bee.to/sso/")
		err := client.Do(request, response)
		if nil != err {
			t.Fatal(err)
		}
		return err
	})
}

func BenchmarkClient_Do(b *testing.B) {
	initDefaultClient(b)

	b.ResetTimer()
	b.ReportAllocs()
	var (
		err    error
		status int
	)
	for i := 0; i < b.N; i++ {
		defaultClient.Do(func(client *HttpClient, req *Request, resp *Response) error {
			req.SetRequestURI("http://api.bee.to/sso/")
			err = client.Do(req, resp)
			status = resp.StatusCode()
			return err
		})
	}

	fmt.Println(status)

}

func BenchmarkClientDo(b *testing.B) {

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI("http://api.bee.to/sso/")
		//_ = defaultClient.http.Do(req, resp)
		_ = fasthttp.Do(req, resp)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}
}

func BenchmarkClient_Request(b *testing.B) {
	initDefaultClient(b)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		defaultClient.Request(func(client *HttpClient) (status int, err error) {
			status, _, err = client.Get(nil, "http://api.bee.to/sso/")
			return
		})
	}
}

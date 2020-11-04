package orca

import (
	"testing"

	"github.com/valyala/fasthttp"
)

type fatal interface {
	Fatal(...interface{})
}

func initDefaultClientConfig(t fatal) *HttpClientConfig {

	conf, err := testClientConfig("tests/client.yml")

	if nil != err {
		t.Fatal(err)
	}

	return conf
}

func DoTest(do func(*Request, *Response) error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	_ = do(request, response)
}

func DoRequestTest(do func() (int, error)) {
	_, _ = do()
}

func TestClient_Get(t *testing.T) {

	//defaultClient.Request(func(client *HttpClient) (status int, err error) {
	//	status, _, err = client.Get(nil, "http://api.bee.to/sso/")
	//	return
	//})

	conf := initDefaultClientConfig(t)

	c := NewClient(ApplyConfig(conf))

	DoRequestTest(func() (int, error) {
		status, _, err := c.Get(nil, "http://api.bee.to/sso/")
		return status, err
	})

}

func TestClient_Do(t *testing.T) {
	conf := initDefaultClientConfig(t)

	c := NewClient(ApplyConfig(conf))

	DoTest(func(request *Request, response *Response) error {
		request.SetRequestURI("http://api.bee.to/sso/")
		err := c.Do(request, response)
		if nil != err {
			t.Fatal(err)
		}
		return err
	})

}

func BenchmarkClient_Do(b *testing.B) {

	//conf := initDefaultClientConfig(b)

	//c := NewClient(ApplyConfig(conf))

	c := NewClient()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DoTest(func(request *Request, response *Response) error {
			request.SetRequestURI("http://api.bee.to/sso/")
			return c.Do(request, response)
			//status = response.StatusCode()
		})
	}

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
	conf := initDefaultClientConfig(b)

	c := NewClient(ApplyConfig(conf))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		DoRequestTest(func() (status int, err error) {
			status, _, err = c.Get(nil, "http://api.bee.to/sso/")
			return
		})
	}
}

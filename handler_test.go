package orca

import (
	"bytes"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestHandler(t *testing.T) {
	type args struct {
		ctx     *RequestCtx
		Handler Handler
	}

	var buf bytes.Buffer

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "before",
			args: args{
				Handler: BeforeHandler(func(ctx *RequestCtx) {
					buf.WriteString("handler")
				}, func(ctx *RequestCtx) {
					buf.WriteString("pre handler=>")
				}),
			},
			want: []byte("pre handler=>handler"),
		},
		{
			name: "after",
			args: args{
				Handler: AfterHandler(func(ctx *RequestCtx) {
					buf.WriteString("handler")
				}, func(ctx *RequestCtx) {
					buf.WriteString("=>after handler")
				}),
			},
			want: []byte("handler=>after handler"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.args.Handler(tt.args.ctx)

			if !bytes.Equal(tt.want, buf.Bytes()) {
				t.Errorf("%s filter test Failed\n", tt.name)
				t.Errorf("want: %s", tt.want)
				t.Errorf("token: %s", buf.Bytes())
			}
		})

	}
}

func BenchmarkFilter(b *testing.B) {
	b.ReportAllocs()
	c := &RequestCtx{}

	h := AfterHandler(func(ctx *RequestCtx) {
	}, func(ctx *RequestCtx) {
	})

	for i := 0; i < b.N; i++ {
		h(c)
		c.Response.Reset()
	}
}

func BenchmarkUnfoldFilter(b *testing.B) {
	b.ReportAllocs()
	c := &RequestCtx{}

	h := AfterHandler(nil, func(ctx *fasthttp.RequestCtx) {
	})

	for i := 0; i < b.N; i++ {
		h(c)
		c.Response.Reset()
	}
}

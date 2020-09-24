package orca

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestMiddleware(t *testing.T) {
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
			name: "pass",
			args: args{
				Handler: Middleware(LambdaMiddleware(nil, func(ctx *RequestCtx) error {
					buf.WriteString("pass")
					return nil
				}), func(ctx *RequestCtx) {
					buf.WriteString("=>handler")
				}),
			},
			want: []byte("pass=>handler"),
		},
		{
			name: "abort",
			args: args{
				Handler: Middleware(func(ctx *RequestCtx) error {
					buf.WriteString("abort")
					return fmt.Errorf("abort")
				}, func(ctx *RequestCtx) {
					buf.WriteString("=>handler")
				}),
			},
			want: []byte("abort"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.args.Handler(tt.args.ctx)

			if !bytes.Equal(tt.want, buf.Bytes()) {
				t.Errorf("%s middleware test Failed\n", tt.name)
				t.Errorf("want: %s", tt.want)
				t.Errorf("token: %s", buf.Bytes())
			}
		})

	}
}

func BenchmarkMiddleware(b *testing.B) {
	b.ReportAllocs()
	c := &RequestCtx{}

	h := Middleware(LambdaMiddleware(func(ctx *RequestCtx) error {
		return nil
	}), func(ctx *RequestCtx) {
	})

	for i := 0; i < b.N; i++ {
		h(c)
		c.Response.Reset()
	}
}

func BenchmarkLambdaMiddleware(b *testing.B) {
	b.ReportAllocs()
	c := &RequestCtx{}

	err := errors.New("test")

	h := LambdaMiddleware(func(ctx *RequestCtx) error {
		return nil
	}, func(ctx *RequestCtx) error {
		return nil
	}, func(ctx *RequestCtx) error {
		return err
	})

	for i := 0; i < b.N; i++ {
		_ = h(c)
		c.Response.Reset()
	}
}

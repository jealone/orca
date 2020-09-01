package orca

import (
	"bytes"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		ctx     *RequestCtx
		Handler Filter
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
				Handler: BeforeFilter(func(ctx *RequestCtx) {
					buf.WriteString("handler\n")
				}, func(ctx *RequestCtx) {
					buf.WriteString("pre handler\n")
				}),
			},
			want: []byte("pre handler\nhandler\n"),
		},
		{
			name: "after",
			args: args{
				Handler: AfterFilter(func(ctx *RequestCtx) {
					buf.WriteString("handler\n")
				}, func(ctx *RequestCtx) {
					buf.WriteString("after handler\n")
				}),
			},
			want: []byte("handler\nafter handler\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.args.Handler(tt.args.ctx)

			if !bytes.Equal(tt.want, buf.Bytes()) {
				t.Errorf("%s filter test Failed\n", tt.name)
			}
		})

	}

}

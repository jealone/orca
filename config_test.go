package orca

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		path string
	}

	sc := &ServerConfig{}

	tests := []struct {
		name string
		args args
		want *ServerConfig
	}{
		{
			name: "config_default",
			args: args{
				path: "tests/orca_default.yml",
			},
			want: sc,
		},
		{
			name: "config_miss",
			args: args{
				path: "tests/orca_miss.yml",
			},
			want: sc,
		},
		{
			name: "config",
			args: args{
				path: "tests/orca.yml",
			},
			want: &ServerConfig{
				SystemConfig: SystemConfig{
					GetOnly:                      false,
					DisablePreParseMultipartForm: false,
					ReduceMemoryUsage:            false,
					LogAllErrors:                 false,
				},
				Tcp: TcpConfig{
					TCPKeepalive:       false,
					TCPKeepalivePeriod: 0,
				},
				Conn: ConnConfig{
					Concurrency:                        0,
					SleepWhenConcurrencyLimitsExceeded: 0,
					DisableKeepalive:                   false,
					MaxConnsPerIP:                      0,
					ReadTimeout:                        0,
					WriteTimeout:                       0,
					IdleTimeout:                        0,
				},
				Buffer: BufferConfig{
					ReadBufferSize:  0,
					WriteBufferSize: 0,
				},
				Header: HeaderConfig{
					Name:                          "orca",
					DisableHeaderNamesNormalizing: false,
					NoDefaultServerHeader:         false,
					NoDefaultDate:                 false,
					NoDefaultContentType:          false,
				},
				Request: RequestConfig{
					MaxRequestsPerConn: 0,
					MaxRequestBodySize: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := filepath.Abs(tt.args.path)

			if nil != err {
				t.Errorf("file path（%s）error: %s", tt.args.path, err)
			}

			file, err := os.Open(path)
			if nil != err {
				t.Errorf("open file（%s）error: %s", path, err)
			}

			decoder := yaml.NewDecoder(file)

			got, err := ParseConfig(decoder)

			if nil != err {
				t.Errorf("parse config error: %s", err)
			}

			err = file.Close()

			if nil != err {
				t.Errorf("close config file error: %s", err)
			}

			if !reflect.DeepEqual(got.SystemConfig, tt.want.SystemConfig) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(*got.GetTcp(), *tt.want.GetTcp()) {
				t.Errorf("TcpConfig() = %v, want %v", got.GetTcp(), tt.want.GetTcp())
			}

			if !reflect.DeepEqual(*got.GetConn(), *tt.want.GetConn()) {
				t.Errorf("ConnConfig() = %v, want %v", got.GetConn(), tt.want.GetConn())
			}

			if !reflect.DeepEqual(*got.GetBuffer(), *tt.want.GetBuffer()) {
				t.Errorf("BufferConfig() = %v, want %v", got.GetBuffer(), tt.want.GetBuffer())
			}

			if !reflect.DeepEqual(*got.GetHeader(), *tt.want.GetHeader()) {
				t.Errorf("HeaderConfig() = %v, want %v", got.GetHeader(), tt.want.GetHeader())
			}

			if !reflect.DeepEqual(*got.GetRequest(), *tt.want.GetRequest()) {
				t.Errorf("RequestConfig() = %v, want %v", got.GetRequest(), tt.want.GetRequest())
			}

			if !reflect.DeepEqual(*got.GetAccessLog(), *tt.want.GetAccessLog()) {
				t.Errorf("AccessLogConfig() = %v, want %v", got.GetAccessLog(), tt.want.GetAccessLog())
			}

		})
	}

}

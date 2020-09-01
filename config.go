package orca

import (
	"errors"
	"time"
)

var (
	errorEmptyDecoder = errors.New("decoder is nil")
)

type ConfigDecoder interface {
	Decode(v interface{}) (err error)
}

func ParseConfig(decoder ConfigDecoder) (*ServerConfig, error) {

	if nil == decoder {
		return nil, errorEmptyDecoder
	}

	serverConf := &ServerConfig{}
	err := decoder.Decode(serverConf)
	if nil != err {
		return nil, err
	}
	return serverConf, nil
}

type ServerConfig struct {
	SystemConfig

	Tcp       TcpConfig       `yaml:"tcp"`
	Conn      ConnConfig      `yaml:"conn"`
	Buffer    BufferConfig    `yaml:"buffer"`
	Header    HeaderConfig    `yaml:"header"`
	Request   RequestConfig   `yaml:"request"`
	AccessLog AccessLogConfig `yaml:"access_log"`
}

func (c *ServerConfig) GetTcp() *TcpConfig {
	return &c.Tcp
}

func (c *ServerConfig) GetConn() *ConnConfig {
	return &c.Conn
}

func (c *ServerConfig) GetBuffer() *BufferConfig {
	return &c.Buffer
}

func (c *ServerConfig) GetHeader() *HeaderConfig {
	return &c.Header
}

func (c *ServerConfig) GetRequest() *RequestConfig {
	return &c.Request
}

func (c *ServerConfig) GetAccessLog() *AccessLogConfig {
	return &c.AccessLog
}

type AccessLogConfig struct {
	Logfile    string `yaml:"logfile"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

func (c *AccessLogConfig) GetLogfile() string {
	if "" == c.Logfile {
		return "logs/access.log"

	}
	return c.Logfile
}

func (c *AccessLogConfig) GetMaxSize() int {
	return c.MaxSize
}

func (c *AccessLogConfig) GetMaxAge() int {
	return c.MaxAge
}

func (c *AccessLogConfig) GetMaxBackups() int {
	return c.MaxBackups
}

func (c *AccessLogConfig) GetCompress() bool {
	return c.Compress
}

type SystemConfig struct {
	DisableAccessLog             bool `yaml:"disable_access_log"`
	GetOnly                      bool `yaml:"get_only"`
	DisablePreParseMultipartForm bool `yaml:"disable_multipart_parse"`
	ReduceMemoryUsage            bool `yaml:"reduce_memory_usage"`
	LogAllErrors                 bool `yaml:"all_errors"`
}

func (c *SystemConfig) GetMethodOnly() bool {
	return c.GetOnly
}

func (c *SystemConfig) GetDisablePreParseMultipartForm() bool {
	return c.DisablePreParseMultipartForm
}

func (c *SystemConfig) GetReduceMemoryUsage() bool {
	return c.ReduceMemoryUsage
}

func (c *SystemConfig) GetLogAllErrors() bool {
	return c.LogAllErrors
}

func (c *SystemConfig) GetDisableAccessLog() bool {
	return c.DisableAccessLog
}

type TcpConfig struct {
	TCPKeepalive       bool   `yaml:"tcp_keepalive"`
	TCPKeepalivePeriod int    `yaml:"tcp_keepalive_interval"`
	Addr               string `yaml:"addr"`
}

func (c *TcpConfig) GetTCPKeepalive() bool {
	return c.TCPKeepalive
}

func (c *TcpConfig) GetTCPKeepalivePeriod() time.Duration {
	return time.Duration(c.TCPKeepalivePeriod) * time.Millisecond
}

func (c *TcpConfig) GetAddr() string {
	if "" == c.Addr {
		return ":8080"
	}
	return c.Addr
}

type HeaderConfig struct {
	Name                          string `yaml:"server"`
	DisableHeaderNamesNormalizing bool   `yaml:"disable_header_names_normalizing"`
	NoDefaultServerHeader         bool   `yaml:"no_default_server_header"`
	NoDefaultDate                 bool   `yaml:"no_default_date"`
	NoDefaultContentType          bool   `yaml:"no_default_content_type"`
}

func (c *HeaderConfig) GetServer() string {
	return c.Name
}

func (c *HeaderConfig) GetDisableHeaderNamesNormalizing() bool {
	return c.DisableHeaderNamesNormalizing
}

func (c *HeaderConfig) GetNoDefaultServerHeader() bool {
	return c.NoDefaultServerHeader
}

func (c *HeaderConfig) GetNoDefaultDate() bool {
	return c.NoDefaultDate
}
func (c *HeaderConfig) GetNoDefaultContentType() bool {
	return c.NoDefaultContentType
}

type BufferConfig struct {
	ReadBufferSize  int `yaml:"read_buffer_size"`
	WriteBufferSize int `yaml:"write_buffer_size"`
}

func (c *BufferConfig) GetReadBufferSize() int {
	return c.ReadBufferSize
}

func (c *BufferConfig) GetWriteBufferSize() int {
	return c.WriteBufferSize
}

type ConnConfig struct {
	Concurrency                        int  `yaml:"concurrency"`
	SleepWhenConcurrencyLimitsExceeded int  `yaml:"concurrency_limits_wait"`
	DisableKeepalive                   bool `yaml:"disable_keepalive"`
	MaxConnsPerIP                      int  `yaml:"max_connections"`
	ReadTimeout                        int  `yaml:"read_timeout"`
	WriteTimeout                       int  `yaml:"write_timeout"`
	IdleTimeout                        int  `yaml:"idle_timeout"`
}

func (c *ConnConfig) GetConcurrency() int {
	return c.Concurrency
}

func (c *ConnConfig) GetSleepWhenConcurrencyLimitsExceeded() time.Duration {
	return time.Duration(c.SleepWhenConcurrencyLimitsExceeded) * time.Millisecond
}

func (c *ConnConfig) GetDisableKeepalive() bool {
	return c.DisableKeepalive
}

func (c *ConnConfig) GetMaxConnsPerIP() int {
	return c.MaxConnsPerIP
}

func (c *ConnConfig) GetIdleTimeout() time.Duration {
	return time.Duration(c.IdleTimeout) * time.Millisecond
}

func (c *ConnConfig) GetReadTimeout() time.Duration {
	return time.Duration(c.ReadTimeout) * time.Millisecond
}

func (c *ConnConfig) GetWriteTimeout() time.Duration {
	return time.Duration(c.WriteTimeout) * time.Millisecond
}

type RequestConfig struct {
	MaxRequestsPerConn int `yaml:"max_requests"`
	MaxRequestBodySize int `yaml:"max_request_body_size"`
}

func (c *RequestConfig) GetMaxRequestsPerConn() int {
	return c.MaxRequestsPerConn
}

func (c *RequestConfig) GetMaxRequestBodySize() int {
	return c.MaxRequestBodySize
}

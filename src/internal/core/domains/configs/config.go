package configs

type AppConfig struct {
	HTTPServerConfig *HTTPServerConfig
	LogLevel         string
	IsProd           bool
}

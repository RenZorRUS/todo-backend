package configs

type AppConfig struct {
	HTTPServerConfig *HTTPServerConfig
	DBConfig         *DBConfig
	LogLevel         string
	IsProd           bool
}

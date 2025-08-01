package consts

const (
	// Application environment variables.
	AppMode = "APP_MODE"

	// HTTP server environment variables.
	HTTPServerPort = "HTTP_SERVER_PORT"
	HTTPServerHost = "HTTP_SERVER_HOST"

	// Database environment variables.
	DBDriver   = "DB_DRIVER"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBURL      = "DB_URL"
	DBSSLMode  = "DB_SSL_MODE"

	// Database connection pool environment variables.
	DBMaxConnLifeTime       = "DB_MAX_CONN_LIFETIME"
	DBMaxConnLifeTimeJitter = "DB_MAX_CONN_LIFE_TIME_JITTER"
	DBMaxConnIdleTime       = "DB_MAX_CONN_IDLE_TIME"
	DBMaxConns              = "DB_MAX_CONNS"
	DBMinConns              = "DB_MIN_CONNS"
	DBMinIdleConns          = "DB_MIN_IDLE_CONNS"
	DBHealthCheckPeriod     = "DB_HEALTH_CHECK_PERIOD"
	DBPoolBuildTimeout      = "DB_POOL_BUILD_TIMEOUT"

	// Logger environment variables.
	LogLevel = "LOG_LEVEL"
)

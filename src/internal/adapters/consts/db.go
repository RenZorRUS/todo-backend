package consts

import (
	"runtime"
	"time"
)

const (
	// Default database connection pool values.
	DBUrlTemplate                  = "%s://%s:%s@%s:%s/%s?sslmode=%s"
	DBHostDefault                  = "0.0.0.0"
	DBPortDefault                  = "5432"
	DBNameDefault                  = "postgres"
	DBUserDefault                  = "postgres"
	DBPasswordDefault              = "postgres"
	DBDriverDefault                = "postgres"
	DBSSLModeDefault               = "disable"
	DBMaxConnLifeTimeDefault       = 30 * time.Minute
	DBMaxConnLifeTimeJitterDefault = 5 * time.Minute
	DBMaxConnIdleTimeDefault       = 10 * time.Minute
	DBHealthCheckPeriodDefault     = 30 * time.Second
	DBPoolBuildTimeoutDefault      = 5 * time.Second

	// Default queries values.
	LimitDefault  = 25
	OffsetDefault = 0
)

var (
	DBMaxConnsDefault     = max(2, runtime.NumCPU()-1)  //nolint:gochecknoglobals,mnd
	DBMinConnsDefault     = 3                           //nolint:gochecknoglobals
	DBMinIdleConnsDefault = max(1, DBMaxConnsDefault/2) //nolint:gochecknoglobals,mnd
)

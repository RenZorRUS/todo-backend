package configs

import (
	"time"
)

type DBConfig struct {
	Driver                string
	Host                  string
	Port                  string
	User                  string
	Password              string
	DBName                string
	ConnUrl               string
	SSLMode               string
	MaxConnLifeTime       time.Duration
	MaxConnLifeTimeJitter time.Duration
	MaxConnIdleTime       time.Duration
	HealthCheckPeriod     time.Duration
	PoolBuildTimeout      time.Duration
	MaxConns              int32
	MinConns              int32
	MinIdleConns          int32
}

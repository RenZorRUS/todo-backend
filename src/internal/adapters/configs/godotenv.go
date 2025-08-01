package configs

import (
	"errors"
	"fmt"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/RenZorRUS/todo-backend/src/internal/utils"
	"github.com/joho/godotenv"
)

type GoDotEnvLoader struct{}

func NewGoDotEnvLoader() *GoDotEnvLoader {
	return &GoDotEnvLoader{}
}

func (loader *GoDotEnvLoader) Load(appMode string) (*configs.AppConfig, error) {
	envFilePath, err := utils.GetEnvFilePath(appMode)
	if err != nil {
		return nil, err
	}

	envConfig, err := godotenv.Read(envFilePath)
	if err != nil {
		return nil, err
	}

	return buildAppConfig(appMode, envConfig)
}

func buildAppConfig(appMode string, envConfig EnvConfig) (*configs.AppConfig, error) {
	dbConfig, err := buildDBConfig(envConfig)
	if err != nil {
		return nil, err
	}

	return &configs.AppConfig{
		HTTPServerConfig: buildHTTPServerConfig(envConfig),
		DBConfig:         dbConfig,
		LogLevel:         envConfig.GetOrDefault(consts.LogLevel, consts.LogLevelDefault),
		IsProd:           utils.IsProdMode(appMode),
	}, nil
}

func buildHTTPServerConfig(envConfig EnvConfig) *configs.HTTPServerConfig {
	return &configs.HTTPServerConfig{
		Port: envConfig.GetOrDefault(consts.HTTPServerPort, consts.HTTPServerPortDefault),
		Host: envConfig.GetOrDefault(consts.HTTPServerHost, consts.HTTPServerHostDefault),
	}
}

func buildDBConfig(envConfig EnvConfig) (*configs.DBConfig, error) { //nolint:funlen
	var allErrs error

	driver := envConfig.GetOrDefault(consts.DBDriver, consts.DBDriverDefault)
	host := envConfig.GetOrDefault(consts.DBHost, consts.DBHostDefault)
	port := envConfig.GetOrDefault(consts.DBPort, consts.DBPortDefault)
	user := envConfig.GetOrDefault(consts.DBUser, consts.DBUserDefault)
	password := envConfig.GetOrDefault(consts.DBPassword, consts.DBPasswordDefault)
	dbName := envConfig.GetOrDefault(consts.DBName, consts.DBNameDefault)
	sslMode := envConfig.GetOrDefault(consts.DBSSLMode, consts.DBSSLModeDefault)

	dbUrlDefault := fmt.Sprintf(
		consts.DBUrlTemplate,
		driver, user, password, host, port, dbName, sslMode,
	)
	dbUrl := envConfig.GetOrDefault(consts.DBURL, dbUrlDefault)

	maxConnLifeTime, err := envConfig.GetOrDefaultDuration(
		consts.DBMaxConnLifeTime,
		consts.DBMaxConnLifeTimeDefault,
	)
	allErrs = errors.Join(allErrs, err)

	maxConnLifeTimeJitter, err := envConfig.GetOrDefaultDuration(
		consts.DBMaxConnLifeTimeJitter,
		consts.DBMaxConnLifeTimeJitterDefault,
	)
	allErrs = errors.Join(allErrs, err)

	maxConnIdleTime, err := envConfig.GetOrDefaultDuration(
		consts.DBMaxConnIdleTime,
		consts.DBMaxConnIdleTimeDefault,
	)
	allErrs = errors.Join(allErrs, err)

	healthCheckPeriod, err := envConfig.GetOrDefaultDuration(
		consts.DBHealthCheckPeriod,
		consts.DBHealthCheckPeriodDefault,
	)
	allErrs = errors.Join(allErrs, err)

	poolBuildTimeout, err := envConfig.GetOrDefaultDuration(
		consts.DBPoolBuildTimeout,
		consts.DBPoolBuildTimeoutDefault,
	)
	allErrs = errors.Join(allErrs, err)

	maxConns, err := envConfig.GetOrDefaultInt(consts.DBMaxConns, consts.DBMaxConnsDefault)
	allErrs = errors.Join(allErrs, err)

	minConns, err := envConfig.GetOrDefaultInt(consts.DBMinConns, consts.DBMinConnsDefault)
	allErrs = errors.Join(allErrs, err)

	minIdleConns, err := envConfig.GetOrDefaultInt(
		consts.DBMinIdleConns,
		consts.DBMinIdleConnsDefault,
	)

	allErrs = errors.Join(allErrs, err)
	if allErrs != nil {
		return nil, allErrs
	}

	return &configs.DBConfig{
		Driver:                driver,
		Host:                  host,
		Port:                  port,
		User:                  user,
		Password:              password,
		DBName:                dbName,
		ConnUrl:               dbUrl,
		SSLMode:               sslMode,
		MaxConnLifeTime:       maxConnLifeTime,
		MaxConnLifeTimeJitter: maxConnLifeTimeJitter,
		MaxConnIdleTime:       maxConnIdleTime,
		HealthCheckPeriod:     healthCheckPeriod,
		PoolBuildTimeout:      poolBuildTimeout,
		MaxConns:              int32(maxConns),     //nolint:gosec
		MinConns:              int32(minConns),     //nolint:gosec
		MinIdleConns:          int32(minIdleConns), //nolint:gosec
	}, nil
}

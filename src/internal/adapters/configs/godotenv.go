package configs

import (
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/RenZorRUS/todo-backend/src/internal/utils"
	"github.com/joho/godotenv"
)

type (
	GoDotEnvLoader struct{}
)

func NewGoDotEnvLoader() *GoDotEnvLoader {
	return &GoDotEnvLoader{}
}

func (loader *GoDotEnvLoader) Load() (*configs.AppConfig, error) {
	isProd := utils.IsAppInProd()
	envFilePath := utils.GetEnvFilePath(isProd)

	envConfig, err := godotenv.Read(envFilePath)
	if err != nil {
		return nil, err
	}

	return buildAppConfig(isProd, envConfig), nil
}

func buildAppConfig(isProd bool, envConfig EnvConfig) *configs.AppConfig {
	return &configs.AppConfig{
		HTTPServerConfig: buildHTTPServerConfig(envConfig),
		LogLevel:         envConfig.GetOrDefault(consts.LogLevel, consts.LogLevelDefault),
		IsProd:           isProd,
	}
}

func buildHTTPServerConfig(envConfig EnvConfig) *configs.HTTPServerConfig {
	return &configs.HTTPServerConfig{
		Port: envConfig.GetOrDefault(consts.HTTPServerPort, consts.HTTPServerPortDefault),
		Host: envConfig.GetOrDefault(consts.HTTPServerHost, consts.HTTPServerHostDefault),
	}
}

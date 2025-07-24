package configs

import (
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

	return buildAppConfig(appMode, envConfig), nil
}

func buildAppConfig(appMode string, envConfig EnvConfig) *configs.AppConfig {
	return &configs.AppConfig{
		HTTPServerConfig: buildHTTPServerConfig(envConfig),
		LogLevel:         envConfig.GetOrDefault(consts.LogLevel, consts.LogLevelDefault),
		IsProd:           utils.IsProdMode(appMode),
	}
}

func buildHTTPServerConfig(envConfig EnvConfig) *configs.HTTPServerConfig {
	return &configs.HTTPServerConfig{
		Port: envConfig.GetOrDefault(consts.HTTPServerPort, consts.HTTPServerPortDefault),
		Host: envConfig.GetOrDefault(consts.HTTPServerHost, consts.HTTPServerHostDefault),
	}
}

package configs

import (
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
)

type AppConfigLoader interface {
	Load() (*configs.AppConfig, error)
}

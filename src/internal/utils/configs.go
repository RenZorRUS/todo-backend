package utils

import (
	"os"
	"strings"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
)

func IsAppInProd() bool {
	appMode, isExist := os.LookupEnv(consts.AppMode)

	if !isExist {
		return false
	}

	appModeLower := strings.ToLower(appMode)

	return appModeLower == "prod" || appModeLower == "production"
}

func GetEnvFilePath(isProd bool) string {
	if isProd {
		return ".env.prod"
	}

	return ".env"
}

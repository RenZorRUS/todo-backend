package utils

import (
	"os"
	"strings"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
)

func GetAppMode() string {
	appMode, isExist := os.LookupEnv(consts.AppMode)
	if !isExist {
		return consts.DevMode
	}

	return appMode
}

func IsProdMode(appMode string) bool {
	appModeLower := strings.ToLower(appMode)
	return appModeLower == consts.ProdMode || appModeLower == "production"
}

func GetEnvFilePath(appMode string) (string, error) {
	switch strings.ToLower(appMode) {
	case consts.ProdMode, "production":
		return consts.ProdEnvFile, nil
	case consts.DevMode, "development":
		return consts.DevEnvFile, nil
	case consts.TestMode:
		return consts.TestEnvFile, nil
	default:
		return "", errs.ErrUnknownAppMode
	}
}

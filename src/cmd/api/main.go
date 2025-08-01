package main

import (
	"log"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/configs"
	controllers "github.com/RenZorRUS/todo-backend/src/internal/adapters/controllers/http/std"
	// "github.com/RenZorRUS/todo-backend/src/internal/adapters/databases/postgres/driver".
	httpstd "github.com/RenZorRUS/todo-backend/src/internal/adapters/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/serializers"
	"github.com/RenZorRUS/todo-backend/src/internal/utils"
)

func main() {
	appMode := utils.GetAppMode()
	envLoader := configs.NewGoDotEnvLoader()

	appConfig, err := envLoader.Load(appMode)
	if err != nil {
		log.Fatal("Failed to load application configuration, ", err.Error())
	}

	logger, _ := loggers.NewZerolog(appConfig)
	baseLogger := logger.GetBaseLogger()
	errLog := log.New(baseLogger, "", 0)

	json, _ := serializers.NewJSONSerializer(logger)
	rw, _ := controllers.NewJSONResponseWriter(logger, json)

	// dbPool, err := driver.NewPgxPool(appConfig.DBConfig)

	router, _ := httpstd.BuildAppServerMux(rw)
	server, _ := httpstd.NewHTTPServer(appConfig.HTTPServerConfig, router, errLog)

	logger.Fatal(server.Run(), "HTTP server stopped.")
}

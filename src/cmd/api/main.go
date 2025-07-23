package main

import (
	"log"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/configs"
	httpstd "github.com/RenZorRUS/todo-backend/src/internal/adapters/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/serializers"
)

func main() {
	envLoader := configs.NewGoDotEnvLoader()

	appConfig, err := envLoader.Load()
	if err != nil {
		log.Fatal("Failed to load application configuration, ", err.Error())
	}

	logger := loggers.New(appConfig)
	errorLog := log.New(logger.GetBaseLogger(), "", 0)

	jsonSerializer := serializers.NewJSONSerializer(logger)
	router := httpstd.BuildAppServerMux(logger, jsonSerializer)
	server := httpstd.NewHTTPServer(
		appConfig.HTTPServerConfig,
		router,
		errorLog,
	)

	logger.Fatal(server.Run(), "HTTP server stopped.")
}

package loggers

import (
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	log *zerolog.Logger
}

func New(appConfig *configs.AppConfig) *ZerologLogger {
	logLevel := convertLogLevel(appConfig.LogLevel)

	if appConfig.IsProd {
		return buildProdLogger(logLevel)
	}

	return buildDevLogger(logLevel)
}

func buildDevLogger(logLevel zerolog.Level) *ZerologLogger {
	buildInfo, _ := debug.ReadBuildInfo()
	consoleWriter := zerolog.ConsoleWriter{ //nolint:exhaustruct
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}

	log := zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Timestamp().
		Int("pid", os.Getpid()).
		Str("v", buildInfo.GoVersion).
		Logger()

	return &ZerologLogger{log: &log}
}

func buildProdLogger(logLevel zerolog.Level) *ZerologLogger {
	buildInfo, _ := debug.ReadBuildInfo()

	log := zerolog.New(os.Stderr).
		Level(logLevel).
		With().
		Timestamp().
		Caller().
		Str("v", buildInfo.GoVersion).
		Logger()

	return &ZerologLogger{log: &log}
}

func convertLogLevel(logLevel string) zerolog.Level {
	switch strings.ToLower(logLevel) {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

func (zl *ZerologLogger) Trace(msg string) {
	zl.log.Trace().Msg(msg)
}

func (zl *ZerologLogger) Debug(msg string) {
	zl.log.Debug().Msg(msg)
}

func (zl *ZerologLogger) Info(msg string) {
	zl.log.Info().Msg(msg)
}

func (zl *ZerologLogger) Warn(msg string) {
	zl.log.Warn().Msg(msg)
}

func (zl *ZerologLogger) Error(err error, msg string) {
	zl.log.Error().Err(err).Msg(msg)
}

func (zl *ZerologLogger) Fatal(err error, msg string) {
	zl.log.Fatal().Err(err).Msg(msg)
}

func (zl *ZerologLogger) Panic(err error, msg string) {
	zl.log.Panic().Err(err).Msg(msg)
}

func (zl *ZerologLogger) GetBaseLogger() *zerolog.Logger {
	return zl.log
}

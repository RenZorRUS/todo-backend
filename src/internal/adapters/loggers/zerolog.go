package loggers

import (
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	log *zerolog.Logger
}

func NewZerolog(appConfig *configs.AppConfig) (*ZerologLogger, error) {
	if appConfig == nil {
		return nil, errs.ErrNilConfig
	}

	logLevel, err := convertLogLevel(appConfig.LogLevel)
	if err != nil {
		return nil, err
	}

	if appConfig.IsProd {
		return buildProdLogger(logLevel), nil
	}

	return buildDevLogger(logLevel), nil
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

func convertLogLevel(logLevel string) (zerolog.Level, error) {
	switch strings.ToLower(logLevel) {
	case "trace":
		return zerolog.TraceLevel, nil
	case "debug":
		return zerolog.DebugLevel, nil
	case "info":
		return zerolog.InfoLevel, nil
	case "warn":
		return zerolog.WarnLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	case "fatal":
		return zerolog.FatalLevel, nil
	case "panic":
		return zerolog.PanicLevel, nil
	default:
		return zerolog.Level(0), errs.ErrUnknownLogLevel
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

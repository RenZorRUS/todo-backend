package loggers

import (
	"errors"
	"os"
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/RenZorRUS/todo-backend/src/internal/tests/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewZerolog_Successfully(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[*configs.AppConfig, *ZerologLogger]{
		{
			Name: "Should return logger for development mode",
			Input: &configs.AppConfig{ //nolint:exhaustruct
				LogLevel: "info",
				IsProd:   false,
			},
			Expected: buildDevLogger(zerolog.InfoLevel),
		},
		{
			Name: "Should return logger for production mode",
			Input: &configs.AppConfig{ //nolint:exhaustruct
				LogLevel: "info",
				IsProd:   true,
			},
			Expected: buildProdLogger(zerolog.InfoLevel),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := NewZerolog(test.Input)

			require.NoError(t, err)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestNewZerolog_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[*configs.AppConfig, error]{
		{
			Name:     "Should return error if config is nil",
			Input:    nil,
			Expected: errs.ErrNilConfig,
		},
		{
			Name: "Should return error if log level is unknown",
			Input: &configs.AppConfig{ //nolint:exhaustruct
				LogLevel: "unknown",
				IsProd:   false,
			},
			Expected: errs.ErrUnknownLogLevel,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := NewZerolog(test.Input)

			require.Error(t, err)
			assert.Nil(t, result)
			assert.Equal(t, test.Expected, err)
		})
	}
}

func TestBuildDevLogger(t *testing.T) {
	t.Parallel()

	config := &configs.AppConfig{ //nolint:exhaustruct
		LogLevel: "info",
		IsProd:   false,
	}

	expected, err := NewZerolog(config)
	require.NoError(t, err)

	result := buildDevLogger(zerolog.InfoLevel)
	require.NotNil(t, result)

	baseLog := result.GetBaseLogger()

	assert.Equal(t, expected, result)
	assert.Equal(t, config.LogLevel, baseLog.GetLevel().String())
}

func TestBuildProdLogger(t *testing.T) {
	t.Parallel()

	config := &configs.AppConfig{ //nolint:exhaustruct
		LogLevel: "info",
		IsProd:   true,
	}

	expected, err := NewZerolog(config)
	require.NoError(t, err)

	result := buildProdLogger(zerolog.InfoLevel)
	require.NotNil(t, result)

	baseLog := result.GetBaseLogger()

	assert.Equal(t, expected, result)
	assert.Equal(t, config.LogLevel, baseLog.GetLevel().String())
}

func TestConvertLogLevel(t *testing.T) { //nolint:funlen
	t.Parallel()

	cases := []tests.TestCase[string, zerolog.Level]{
		{
			Name:     "Should return trace level for lower case",
			Input:    "trace",
			Expected: zerolog.TraceLevel,
		},
		{
			Name:     "Should return trace level for upper case",
			Input:    "TRACE",
			Expected: zerolog.TraceLevel,
		},
		{
			Name:     "Should return debug level for lower case",
			Input:    "debug",
			Expected: zerolog.DebugLevel,
		},
		{
			Name:     "Should return debug level for upper case",
			Input:    "DEBUG",
			Expected: zerolog.DebugLevel,
		},
		{
			Name:     "Should return info level for lower case",
			Input:    "info",
			Expected: zerolog.InfoLevel,
		},
		{
			Name:     "Should return info level for upper case",
			Input:    "INFO",
			Expected: zerolog.InfoLevel,
		},
		{
			Name:     "Should return warn level for lower case",
			Input:    "warn",
			Expected: zerolog.WarnLevel,
		},
		{
			Name:     "Should return warn level for upper case",
			Input:    "WARN",
			Expected: zerolog.WarnLevel,
		},
		{
			Name:     "Should return error level for lower case",
			Input:    "error",
			Expected: zerolog.ErrorLevel,
		},
		{
			Name:     "Should return error level for upper case",
			Input:    "ERROR",
			Expected: zerolog.ErrorLevel,
		},
		{
			Name:     "Should return fatal level for lower case",
			Input:    "fatal",
			Expected: zerolog.FatalLevel,
		},
		{
			Name:     "Should return fatal level for upper case",
			Input:    "FATAL",
			Expected: zerolog.FatalLevel,
		},
		{
			Name:     "Should return panic level for lower case",
			Input:    "panic",
			Expected: zerolog.PanicLevel,
		},
		{
			Name:     "Should return panic level for upper case",
			Input:    "PANIC",
			Expected: zerolog.PanicLevel,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := convertLogLevel(test.Input)

			require.NoError(t, err)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestConvertLogLevel_Fail(t *testing.T) {
	t.Parallel()

	result, err := convertLogLevel("unknown")

	require.Error(t, err)
	assert.Equal(t, errs.ErrUnknownLogLevel, err)
	assert.Equal(t, zerolog.Level(0), result)
}

func TestZerologLogger_Trace_Debug_Info_Warn(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[func(*ZerologLogger), string]{
		{
			Name:     "Should log message with trace level",
			Input:    func(zl *ZerologLogger) { zl.Trace("test") },
			Expected: "{\"level\":\"trace\",\"message\":\"test\"}\n",
		},
		{
			Name:     "Should log message with debug level",
			Input:    func(zl *ZerologLogger) { zl.Debug("test") },
			Expected: "{\"level\":\"debug\",\"message\":\"test\"}\n",
		},
		{
			Name:     "Should log message with info level",
			Input:    func(zl *ZerologLogger) { zl.Info("test") },
			Expected: "{\"level\":\"info\",\"message\":\"test\"}\n",
		},
		{
			Name:     "Should log message with warn level",
			Input:    func(zl *ZerologLogger) { zl.Warn("test") },
			Expected: "{\"level\":\"warn\",\"message\":\"test\"}\n",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			runTestZerologMsg(t, test.Input, test.Expected)
		})
	}
}

func runTestZerologMsg(
	t *testing.T,
	logMsg func(*ZerologLogger),
	expected string,
) {
	t.Helper()

	buffer := mocks.NewZerologBuffer()
	log := zerolog.New(buffer)
	zLog := &ZerologLogger{log: &log}

	logMsg(zLog)

	logs := buffer.Logs()
	result := logs[0]

	assert.Len(t, logs, 1)
	assert.Equal(t, expected, result)
}

func TestZerologLogger_Error_Panic_Fatal(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[func(*ZerologLogger, error, string), string]{
		{
			Name: "Should log message with error level",
			Input: func(zl *ZerologLogger, err error, msg string) {
				zl.Error(err, msg)
			},
			Expected: "{\"level\":\"error\",\"error\":\"test error\",\"message\":\"test message\"}\n",
		},
		{
			Name: "Should log message with panic level",
			Input: func(zl *ZerologLogger, err error, msg string) {
				require.Panics(t, func() { zl.Panic(err, msg) })
			},
			Expected: "{\"level\":\"panic\",\"error\":\"test error\",\"message\":\"test message\"}\n",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			runTestZerologMsgWithErr(t, test.Input, test.Expected)
		})
	}
}

func runTestZerologMsgWithErr(
	t *testing.T,
	logMsg func(*ZerologLogger, error, string),
	expected string,
) {
	t.Helper()

	err := errors.New("test error") //nolint:err113
	msg := "test message"

	buffer := mocks.NewZerologBuffer()
	log := zerolog.New(buffer)
	zLog := &ZerologLogger{log: &log}

	logMsg(zLog, err, msg)

	logs := buffer.Logs()
	result := logs[0]

	assert.Len(t, logs, 1)
	assert.Equal(t, expected, result)
}

func TestZerologLogger_GetBaseLogger(t *testing.T) {
	t.Parallel()

	log := zerolog.New(os.Stderr)
	zLog := &ZerologLogger{log: &log}

	result := zLog.GetBaseLogger()

	assert.Equal(t, &log, result)
}

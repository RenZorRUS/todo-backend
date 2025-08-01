package std

import (
	"log"
	"testing"
	"time"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	httpstd "github.com/RenZorRUS/todo-backend/src/internal/adapters/controllers/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/serializers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPServer_Success(t *testing.T) {
	t.Parallel()

	server := buildHTTPServer(t)

	require.NotNil(t, server)
}

func TestNewHTTPServer_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[func() (*HTTPServer, error), error]{
		{
			Name: "Should return error if config is nil",
			Input: func() (*HTTPServer, error) {
				return NewHTTPServer(nil, nil, nil)
			},
			Expected: errs.ErrNilConfig,
		},
		{
			Name: "Should return error if router is nil",
			Input: func() (*HTTPServer, error) {
				config := &configs.HTTPServerConfig{} //nolint:exhaustruct
				return NewHTTPServer(config, nil, nil)
			},
			Expected: errs.ErrNilRouter,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := test.Input()

			require.Error(t, err)
			assert.Nil(t, result)
			assert.Equal(t, test.Expected, err)
		})
	}
}

func TestHTTPServer_Run_Stop(t *testing.T) {
	t.Parallel()

	server := buildHTTPServer(t)

	time.AfterFunc(1*time.Millisecond, func() {
		require.NoError(t, server.Stop())
	})

	require.Error(t, server.Run())

	time.Sleep(10 * time.Millisecond)
}

func buildHTTPServer(t *testing.T) *HTTPServer {
	t.Helper()

	appConfig := &configs.AppConfig{
		DBConfig: &configs.DBConfig{}, //nolint:exhaustruct // fill DBConfig structure.
		HTTPServerConfig: &configs.HTTPServerConfig{
			Port: consts.HTTPServerPortDefault,
			Host: consts.HTTPServerHostDefault,
		},
		LogLevel: consts.LogLevelDefault,
		IsProd:   false,
	}

	logger, _ := loggers.NewZerolog(appConfig)
	baseLogger := logger.GetBaseLogger()
	errorLog := log.New(baseLogger, "", 0)
	json, _ := serializers.NewJSONSerializer(logger)
	rw, _ := httpstd.NewJSONResponseWriter(logger, json)

	router, _ := BuildAppServerMux(rw)
	server, _ := NewHTTPServer(appConfig.HTTPServerConfig, router, errorLog)

	return server
}

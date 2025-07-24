package std

import (
	"testing"

	httpstd "github.com/RenZorRUS/todo-backend/src/internal/adapters/controllers/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/serializers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildAppServerMux_Success(t *testing.T) {
	t.Parallel()

	appConfig := &configs.AppConfig{LogLevel: "info"} //nolint:exhaustruct
	log, _ := loggers.NewZerolog(appConfig)
	json, _ := serializers.NewJSONSerializer(log)
	rw, _ := httpstd.NewJSONResponseWriter(log, json)

	mux, err := BuildAppServerMux(rw)

	require.NoError(t, err)
	assert.NotNil(t, mux)
}

func TestBuildAppServerMux_Error(t *testing.T) {
	t.Parallel()

	result, err := BuildAppServerMux(nil)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.ErrJSONResponseWriterNotSpecified, err)
}

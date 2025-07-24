package std

import (
	"net/http"
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHealthController_Success(t *testing.T) {
	t.Parallel()

	mockRW := mocks.NewMockJSONResponseWriter()
	controller, err := NewHealthController(mockRW)

	require.NoError(t, err)
	assert.NotNil(t, controller)
}

func TestNewHealthController_Fail(t *testing.T) {
	t.Parallel()

	result, err := NewHealthController(nil)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.ErrJSONResponseWriterNotSpecified, err)
}

func TestHealthController_Register(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()

	controller, _ := NewHealthController(nil)
	controller.Register(mux)
}

func TestHealthController_CheckHealth(t *testing.T) {
	t.Parallel()

	content := map[string]string{"status": "ok"}

	mockHTTPRW := mocks.NewMockHTTPResponseWriter()
	mockJSONRW := mocks.NewMockJSONResponseWriter()
	mockJSONRW.On("WriteJSON", mockHTTPRW, http.StatusOK, content).Once()

	rw, _ := NewHealthController(mockJSONRW)
	rw.CheckHealth(mockHTTPRW, nil)

	mockJSONRW.AssertExpectations(t)
	mockHTTPRW.AssertExpectations(t)
}

package std

import (
	"net/http"
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/serializers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/RenZorRUS/todo-backend/src/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewResponseWriter_Success(t *testing.T) {
	t.Parallel()

	appConfig := &configs.AppConfig{LogLevel: "info"} //nolint:exhaustruct
	log, _ := loggers.NewZerolog(appConfig)
	json, _ := serializers.NewJSONSerializer(log)

	result, err := NewJSONResponseWriter(log, json)

	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNewResponseWriter_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[func() (*JSONResWriter, error), error]{
		{
			Name: "Should return error if logger is nil",
			Input: func() (*JSONResWriter, error) {
				return NewJSONResponseWriter(nil, nil)
			},
			Expected: errs.ErrLogNotSpecified,
		},
		{
			Name: "Should return error if JSON serializer is nil",
			Input: func() (*JSONResWriter, error) {
				appConfig := &configs.AppConfig{LogLevel: "info"} //nolint:exhaustruct
				log, _ := loggers.NewZerolog(appConfig)

				return NewJSONResponseWriter(log, nil)
			},
			Expected: errs.ErrJSONSerializerNotSpecified,
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

func TestWriteJSON_Success(t *testing.T) {
	t.Parallel()

	statusCode := http.StatusOK
	headers := http.Header{}
	data := []byte{}

	mockWriter := mocks.NewMockHTTPResponseWriter()
	mockWriter.On("Header").Return(headers).Once()
	mockWriter.On("WriteHeader", statusCode).Once()
	mockWriter.On("Write", data).Return(5, nil).Once()

	mockLog := mocks.NewLogger()

	mockJson := mocks.NewJSONSerializer()
	mockJson.On("Marshal", data).Return(data, nil).Once()

	rw, _ := NewJSONResponseWriter(mockLog, mockJson)
	rw.WriteJSON(mockWriter, statusCode, data)

	mockWriter.AssertExpectations(t)
	mockLog.AssertExpectations(t)
	mockJson.AssertExpectations(t)

	assert.Equal(t, consts.ApplicationJSON, headers.Get(consts.ContentType))
}

func TestWriteJSONError(t *testing.T) {
	t.Parallel()

	err := errs.ErrNilInput
	errContent := map[string]string{"error": err.Error()}

	headers := http.Header{}
	data := []byte{}

	mockWriter := mocks.NewMockHTTPResponseWriter()
	mockWriter.On("Header").Return(headers).Once()
	mockWriter.On("WriteHeader", http.StatusInternalServerError).Once()
	mockWriter.On("Write", data).Return(5, nil).Once()

	mockLog := mocks.NewLogger()
	mockLog.On("Error", err, "Error serializing response value to JSON.").Once()

	mockJson := mocks.NewJSONSerializer()
	mockJson.On("Marshal", data).Return(nil, err).Once()
	mockJson.On("Marshal", errContent).Return(data, nil).Once()

	rw, _ := NewJSONResponseWriter(mockLog, mockJson)
	rw.WriteJSON(mockWriter, 0, data)

	mockWriter.AssertExpectations(t)
	mockLog.AssertExpectations(t)
	mockJson.AssertExpectations(t)

	assert.Equal(t, consts.ApplicationJSON, headers.Get(consts.ContentType))
}

func TestWriteBytes_Fail(t *testing.T) {
	t.Parallel()

	err := errs.ErrNilInput
	data := []byte{}

	mockWriter := mocks.NewMockHTTPResponseWriter()
	mockWriter.On("Write", data).Return(0, err).Once()

	mockLog := mocks.NewLogger()
	mockLog.On("Error", err, "Error writing response to client.").Once()

	mockJson := mocks.NewJSONSerializer()

	rw, _ := NewJSONResponseWriter(mockLog, mockJson)
	rw.WriteBytes(mockWriter, data)

	mockWriter.AssertExpectations(t)
	mockLog.AssertExpectations(t)
	mockJson.AssertExpectations(t)
}

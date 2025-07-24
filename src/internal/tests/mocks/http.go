package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type (
	MockHTTPResponseWriter struct {
		mock.Mock
	}

	MockJSONResponseWriter struct {
		mock.Mock
	}
)

func NewMockHTTPResponseWriter() *MockHTTPResponseWriter {
	return &MockHTTPResponseWriter{} //nolint:exhaustruct
}

func (m *MockHTTPResponseWriter) Header() http.Header {
	args := m.Called()
	val, _ := args.Get(0).(http.Header)

	return val
}

func (m *MockHTTPResponseWriter) Write(data []byte) (int, error) {
	args := m.Called(data)
	return args.Int(0), args.Error(1)
}

func (m *MockHTTPResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}

func NewMockJSONResponseWriter() *MockJSONResponseWriter {
	return &MockJSONResponseWriter{} //nolint:exhaustruct
}

func (m *MockJSONResponseWriter) WriteJSON(
	res http.ResponseWriter,
	statusCode int,
	value any,
) {
	m.Called(res, statusCode, value)
}

func (m *MockJSONResponseWriter) WriteJSONError(
	res http.ResponseWriter,
	statusCode int,
	err error,
) {
	m.Called(res, statusCode, err)
}

func (m *MockJSONResponseWriter) WriteBytes(
	w http.ResponseWriter,
	bytes []byte,
) {
	m.Called(w, bytes)
}

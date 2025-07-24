package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockJSONSerializer struct {
	mock.Mock
}

func NewJSONSerializer() *MockJSONSerializer {
	return &MockJSONSerializer{} //nolint:exhaustruct
}

func (m *MockJSONSerializer) Marshal(value any) ([]byte, error) {
	args := m.Called(value)
	val, _ := args.Get(0).([]byte)

	return val, args.Error(1)
}

func (m *MockJSONSerializer) Unmarshal(data []byte, value any) error {
	args := m.Called(data, value)
	return args.Error(0)
}

package mocks

import (
	"github.com/stretchr/testify/mock"
)

type (
	Logger struct {
		mock.Mock
	}

	ZlogBuffer struct {
		logs []string
	}
)

func NewLogger() *Logger                      { return &Logger{} } //nolint:exhaustruct
func (l *Logger) Trace(msg string)            { l.Called(msg) }
func (l *Logger) Debug(msg string)            { l.Called(msg) }
func (l *Logger) Info(msg string)             { l.Called(msg) }
func (l *Logger) Warn(msg string)             { l.Called(msg) }
func (l *Logger) Error(err error, msg string) { l.Called(err, msg) }
func (l *Logger) Fatal(err error, msg string) { l.Called(err, msg) }
func (l *Logger) Panic(err error, msg string) { l.Called(err, msg) }

func NewZerologBuffer() *ZlogBuffer  { return &ZlogBuffer{logs: []string{}} }
func (b *ZlogBuffer) Logs() []string { return b.logs }
func (b *ZlogBuffer) Write(msg []byte) (int, error) {
	b.logs = append(b.logs, string(msg))
	return len(msg), nil
}

package errs

import (
	"errors"
)

var (
	ErrNilInput        = errors.New("input cannot be nil")
	ErrNilConfig       = errors.New("config cannot be nil")
	ErrNilRouter       = errors.New("router cannot be nil")
	ErrUnknownLogLevel = errors.New("log level is unknown")
	ErrUnknownAppMode  = errors.New("app mode is unknown")
)

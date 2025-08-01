package errs

import "errors"

var (
	ErrUnknownConversionOutputType = errors.New("unknown conversion output type")
	ErrUintToIntOverflow           = errors.New(
		"uint value is too large to be represented as int (overflow)",
	)
)

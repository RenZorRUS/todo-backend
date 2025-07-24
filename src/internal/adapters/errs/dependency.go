package errs

import "errors"

var (
	ErrJSONSerializerNotSpecified     = errors.New("json serializer is not specified")
	ErrJSONResponseWriterNotSpecified = errors.New("json response writer is not specified")
	ErrLogNotSpecified                = errors.New("logger is not specified")
)

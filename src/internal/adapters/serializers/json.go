package serializers

import (
	"encoding/json"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
)

type JSONSerializer struct {
	log loggers.Logger
}

func NewJSONSerializer(log loggers.Logger) (*JSONSerializer, error) {
	if log == nil {
		return nil, errs.ErrLogNotSpecified
	}

	return &JSONSerializer{log: log}, nil
}

func (s *JSONSerializer) Marshal(value any) ([]byte, error) {
	if value == nil {
		s.log.Error(errs.ErrNilInput, errs.ErrNilInput.Error())
		return nil, errs.ErrNilInput
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		s.log.Error(err, err.Error())
		return nil, err
	}

	return bytes, nil
}

func (s *JSONSerializer) Unmarshal(data []byte, value any) error {
	err := json.Unmarshal(data, value)
	if err != nil {
		s.log.Error(err, err.Error())
		return err
	}

	return nil
}

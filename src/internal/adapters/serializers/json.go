package serializers

import (
	"encoding/json"

	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
)

type JSONSerializer struct {
	log loggers.Logger
}

func NewJSONSerializer(log loggers.Logger) *JSONSerializer {
	return &JSONSerializer{log: log}
}

func (s *JSONSerializer) Marshal(value any) ([]byte, error) {
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

package std

import (
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/serializers"
)

type (
	JSONResWriter struct {
		json serializers.JSONSerializer
		log  loggers.Logger
	}

	JSONResponseWriter interface {
		WriteJSON(res http.ResponseWriter, statusCode int, value any)
		WriteJSONError(res http.ResponseWriter, statusCode int, err error)
		WriteBytes(res http.ResponseWriter, bytes []byte)
	}
)

func NewJSONResponseWriter(
	log loggers.Logger,
	json serializers.JSONSerializer,
) (*JSONResWriter, error) {
	if log == nil {
		return nil, errs.ErrLogNotSpecified
	}

	if json == nil {
		return nil, errs.ErrJSONSerializerNotSpecified
	}

	return &JSONResWriter{json: json, log: log}, nil
}

func (rw *JSONResWriter) WriteJSON(res http.ResponseWriter, statusCode int, value any) {
	bytes, err := rw.json.Marshal(value)
	if err != nil {
		rw.log.Error(err, "Error serializing response value to JSON.")
		rw.WriteJSONError(res, http.StatusInternalServerError, err)

		return
	}

	res.Header().Set(consts.ContentType, consts.ApplicationJSON)
	res.WriteHeader(statusCode)
	rw.WriteBytes(res, bytes)
}

func (rw *JSONResWriter) WriteJSONError(res http.ResponseWriter, statusCode int, err error) {
	content := map[string]string{"error": err.Error()}
	bytes, _ := rw.json.Marshal(content)

	res.Header().Set(consts.ContentType, consts.ApplicationJSON)
	res.WriteHeader(statusCode)
	rw.WriteBytes(res, bytes)
}

func (rw *JSONResWriter) WriteBytes(w http.ResponseWriter, bytes []byte) {
	_, err := w.Write(bytes)
	if err != nil {
		rw.log.Error(err, "Error writing response to client.")
	}
}

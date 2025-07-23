package std

import (
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/serializers"
)

type ResponseWriter struct {
	serializer serializers.JSONSerializer
	log        loggers.Logger
}

func NewResponseWriter(l loggers.Logger, s serializers.JSONSerializer) *ResponseWriter {
	return &ResponseWriter{
		serializer: s,
		log:        l,
	}
}

func (rw ResponseWriter) WriteJSON(res http.ResponseWriter, statusCode int, value any) {
	bytes, err := rw.serializer.Marshal(value)
	if err != nil {
		rw.log.Error(err, "Error serializing response value to JSON.")
		rw.WriteJSONError(res, http.StatusInternalServerError, err)

		return
	}

	res.Header().Set(consts.ContentType, consts.ApplicationJSON)
	res.WriteHeader(statusCode)
	rw.WriteBytes(res, bytes)
}

func (rw ResponseWriter) WriteJSONError(res http.ResponseWriter, statusCode int, err error) {
	content := map[string]string{"error": err.Error()}
	bytes, _ := rw.serializer.Marshal(content)

	res.Header().Set(consts.ContentType, consts.ApplicationJSON)
	res.WriteHeader(statusCode)
	rw.WriteBytes(res, bytes)
}

func (rw ResponseWriter) WriteBytes(w http.ResponseWriter, bytes []byte) {
	_, err := w.Write(bytes)
	if err != nil {
		rw.log.Error(err, "Error writing response to client.")
	}
}

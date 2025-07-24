package std

import (
	"net/http"

	httpstd "github.com/RenZorRUS/todo-backend/src/internal/adapters/controllers/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
)

func BuildAppServerMux(jsonResWriter httpstd.JSONResponseWriter) (*http.ServeMux, error) {
	if jsonResWriter == nil {
		return nil, errs.ErrJSONResponseWriterNotSpecified
	}

	mux := http.NewServeMux()

	healthController, _ := httpstd.NewHealthController(jsonResWriter)
	healthController.Register(mux)

	return mux, nil
}

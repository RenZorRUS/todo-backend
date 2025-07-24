package std

import (
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
)

type HealthController struct {
	rw JSONResponseWriter
}

func NewHealthController(rw JSONResponseWriter) (*HealthController, error) {
	if rw == nil {
		return nil, errs.ErrJSONResponseWriterNotSpecified
	}

	return &HealthController{rw: rw}, nil
}

func (hc *HealthController) Register(appServerMux *http.ServeMux) {
	appServerMux.HandleFunc("GET /health", hc.CheckHealth)
}

func (hc *HealthController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	content := map[string]string{"status": "ok"}
	hc.rw.WriteJSON(w, http.StatusOK, content)
}

package std

import (
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/serializers"
)

type HealthController struct {
	rw *ResponseWriter
}

func NewHealthController(
	log loggers.Logger,
	serializer serializers.JSONSerializer,
) *HealthController {
	return &HealthController{rw: NewResponseWriter(log, serializer)}
}

func (hc *HealthController) Register(appServerMux *http.ServeMux) {
	appServerMux.HandleFunc("GET /health", hc.CheckHealth)
}

func (hc *HealthController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	content := map[string]string{"status": "ok"}
	hc.rw.WriteJSON(w, http.StatusOK, content)
}

package std

import (
	"net/http"

	controllers "github.com/RenZorRUS/todo-backend/src/internal/adapters/controllers/http/std"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/loggers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/serializers"
)

func BuildAppServerMux(
	log loggers.Logger,
	serializer serializers.JSONSerializer,
) *http.ServeMux {
	mux := http.NewServeMux()

	healthController := controllers.NewHealthController(log, serializer)
	healthController.Register(mux)

	return mux
}

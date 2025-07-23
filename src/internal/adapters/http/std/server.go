package std

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
)

type HTTPServer struct {
	server *http.Server
	log    *log.Logger
}

func NewHTTPServer(
	config *configs.HTTPServerConfig,
	appServerMux *http.ServeMux,
	log *log.Logger,
) *HTTPServer {
	server := &http.Server{ //nolint:exhaustruct
		Addr:              fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:           appServerMux,
		ReadTimeout:       consts.ReadTimeout,
		ReadHeaderTimeout: consts.ReadHeaderTimeout,
		WriteTimeout:      consts.WriteTimeout,
		IdleTimeout:       consts.IdleTimeout,
		MaxHeaderBytes:    consts.MaxHeaderBytes,
		ErrorLog:          log,
	}

	return &HTTPServer{server: server, log: log}
}

func (s *HTTPServer) Run() error {
	s.log.Println("HTTP server started and listening on address:", s.server.Addr)
	return s.server.ListenAndServe()
}

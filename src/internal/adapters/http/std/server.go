package std

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
)

type HTTPServer struct {
	server *http.Server
	log    *log.Logger
}

func NewHTTPServer(
	config *configs.HTTPServerConfig,
	appServerMux *http.ServeMux,
	errLog *log.Logger,
) (*HTTPServer, error) {
	if config == nil {
		return nil, errs.ErrNilConfig
	}

	if appServerMux == nil {
		return nil, errs.ErrNilRouter
	}

	server := &http.Server{ //nolint:exhaustruct
		Addr:              fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:           appServerMux,
		ReadTimeout:       consts.ReadTimeout,
		ReadHeaderTimeout: consts.ReadHeaderTimeout,
		WriteTimeout:      consts.WriteTimeout,
		IdleTimeout:       consts.IdleTimeout,
		MaxHeaderBytes:    consts.MaxHeaderBytes,
		ErrorLog:          errLog,
	}

	return &HTTPServer{server: server, log: errLog}, nil
}

func (s *HTTPServer) Run() error {
	s.log.Println("HTTP server started and listening on address:", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop() error {
	return s.server.Shutdown(context.TODO())
}

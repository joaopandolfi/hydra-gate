package server

import (
	"context"
	"net/http"

	"hydra_gate/utils/logger"

	"hydra_gate/config"

	"github.com/gorilla/mux"
)

// Server surgical
type Server struct {
	R      *mux.Router
	Config config.Config
	srv    *http.Server
}

// New server
func New(r *mux.Router, conf config.Config) *Server {
	// Bind to a port and pass our router in
	logger.Info("Server listenning on", conf.Server.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Server.Port,
		WriteTimeout: conf.Server.TimeoutWrite,
		ReadTimeout:  conf.Server.TimeoutRead,
	}

	return &Server{
		R:      r,
		Config: conf,
		srv:    srv,
	}
}

// Start Web server
func (s *Server) Start() {

	var err error
	if s.Config.Server.Debug {
		err = s.srv.ListenAndServe()
	} else {
		err = s.srv.ListenAndServeTLS(s.Config.Server.Security.TLSCert, s.Config.Server.Security.TLSKey)
	}
	if err != nil && err != http.ErrServerClosed {
		logger.Error("Fatal server error", err.Error())
	}
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

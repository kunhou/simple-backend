package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github/kunhou/simple-backend/pkg/servmanager"
)

const DefaultShutdownTimeout = 5 * time.Minute

var _ servmanager.Server = (*Server)(nil)

type Server struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

type Options func(*Server)

func NewServer(e *gin.Engine, cfg *Config, options ...Options) *Server {
	ser := Server{
		shutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    cfg.Addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&ser)
	}

	return &ser
}

func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *Server) {
		server.shutdownTimeout = duration
	}
}

// Start implement servmanager.Server interface
func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown implement servmanager.Server interface
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}

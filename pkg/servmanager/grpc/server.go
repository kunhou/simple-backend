package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github/kunhou/simple-backend/pkg/servmanager"
)

var _ servmanager.Server = (*Server)(nil)

// Server grpc server
type Server struct {
	srv  *grpc.Server
	addr string
}

type Options func(*Server)

// NewServer creates a new server instance with default settings
func NewServer(grpcServer *grpc.Server, cfg *Config, options ...Options) *Server {
	ser := Server{
		srv:  grpcServer,
		addr: cfg.Addr,
	}

	for _, option := range options {
		option(&ser)
	}
	return &ser
}

// Start to start the server and listen on the given address
func (h *Server) Start() (err error) {
	log.Printf("grpc server listening on %s", h.addr)
	lis, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}
	if err = h.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

// Shutdown shuts down the server
func (h *Server) Shutdown() error {
	h.srv.GracefulStop()
	return nil
}

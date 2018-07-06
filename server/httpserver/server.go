package httpserver

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/yowcow/test/grpc-test/service"
)

type Server struct {
	app    pb.DataServer
	sv     *http.Server
	logger *log.Logger
}

func New(logger *log.Logger) *Server {
	s := &Server{
		logger: logger,
	}
	h := http.NewServeMux()
	h.HandleFunc("/hello/", s.GetHello)
	h.HandleFunc("/GetPerson/", s.GetPerson)
	s.sv = &http.Server{
		Handler: h,
	}
	return s
}

func (s *Server) RegisterDataServer(app pb.DataServer) {
	s.app = app
}

func (s *Server) Serve(ln net.Listener) error {
	if err := s.sv.Serve(ln); err != nil {
		return err
	}
	return nil
}

func (s *Server) GracefulStop(ctx context.Context) error {
	if err := s.sv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

package main

import (
	"context"
	"net"
	"net/http"

	_ "net/http/pprof"
)

type PerfServer struct {
	srv *http.Server
}

func NewPerfServer() *PerfServer {
	return &PerfServer{
		srv: &http.Server{Addr: ":6060"},
	}
}

func (s *PerfServer) Listen() error {
	ln, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}

	go s.srv.Serve(ln)
	return nil
}

func (s *PerfServer) Shutdown(ctx context.Context) {
	s.srv.Shutdown(ctx)
}

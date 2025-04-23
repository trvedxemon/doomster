package main

import (
    "net/http"
    "sync"
	"context"
	"time"
)

type EndpointWorker interface {
	Start()
	Stop()
}

type Server struct {
	srv *http.Server
	port int
	endpoint string
}

func (s *Server) Start() {
	httpServerExitDone := &sync.WaitGroup{}

    httpServerExitDone.Add(1)
	s.srv = startHttpServer(s.port, s.endpoint, httpServerExitDone)
}

func (s Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := s.srv.Shutdown(ctx); err != nil {
        // handle err
    }
}
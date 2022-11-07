package code

import (
	"context"
	"net/http"
	"time"
)

const (
	timeTL = time.Second * 15
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunHTTP(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadTimeout:       timeTL,
		ReadHeaderTimeout: timeTL,
		WriteTimeout:      timeTL,
		IdleTimeout:       timeTL,
		MaxHeaderBytes:    1 << 20,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

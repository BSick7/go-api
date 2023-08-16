package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*mux.Router
	Middlewares []mux.MiddlewareFunc
	Logger      *slog.Logger
}

func (s *Server) Register(endpoints ...Endpoint) {
	for _, endpoint := range endpoints {
		endpoint.Register(s.Router)
	}
}

// Use registers middleware on the mux router
// These middlewares are chained in the order they are registered
// If middleware A is registered early than B, then execution will be A => B => route handler
// By default, middlewares are not executed for NotFound or MethodNotAllowed as detected by router
// In order to use middlewares, utilize api.MiddlewaresHandler to wrap a raw handler with these middlewares
func (s *Server) Use(mwf ...mux.MiddlewareFunc) {
	s.Middlewares = append(s.Middlewares, mwf...)
	s.Router.Use(mwf...)
}

func (s *Server) Launch(port int, cancelFn func()) error {
	return s.launch(port, cancelFn, startHttp)
}

func (s *Server) LaunchTLS(port int, certFile, keyFile string, cancelFn func()) error {
	return s.launch(port, cancelFn, startHttps(certFile, keyFile))
}

func (s *Server) launch(port int, cancelFn func(), startFn startFunc) error {
	serverlog := s.Logger.With(slog.String("source", "http-server"))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
		ErrorLog:     log.New(os.Stdout, "[http-server] ", 0),
		Handler:      s,
	}
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-term
		serverlog.Info(fmt.Sprintf("received %s, shutting down", sig))
		cancelFn()
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			serverlog.Error(fmt.Sprintf("server did not shut down: %s", err))
		}
	}()

	serverlog.Info("listening", slog.Int("port", port))
	if err := startFn(server); err != http.ErrServerClosed {
		return err
	}
	serverlog.Info("server shut down")
	return nil
}

type startFunc func(server *http.Server) error

func startHttp(server *http.Server) error {
	return server.ListenAndServe()
}
func startHttps(certFile, keyFile string) startFunc {
	return func(server *http.Server) error {
		return server.ListenAndServeTLS(certFile, keyFile)
	}
}

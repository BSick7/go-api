package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func (s *Server) Register(endpoints ...Endpoint) {
	for _, endpoint := range endpoints {
		endpoint.Register(s.Router)
	}
}

func (s *Server) Launch(port int, cancelFn func()) error {
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
		server.ErrorLog.Printf("received %s, shutting down...\n", sig)
		cancelFn()
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			server.ErrorLog.Printf("server did not shut down: %s\n", err)
		}
	}()

	server.ErrorLog.Printf("listening on :%d\n", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	server.ErrorLog.Printf("server shut down")
	return nil
}

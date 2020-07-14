package api

import (
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
		if err := server.Close(); err != nil {
			server.ErrorLog.Printf("server did not fully close: %s\n", err)
		}
		cancelFn()
	}()

	server.ErrorLog.Printf("listening on :%d\n", port)
	return server.ListenAndServe()
}

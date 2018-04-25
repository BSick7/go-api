package api

import (
	"net/http"
	"strings"

	"github.com/BSick7/go-api/context"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(router *mux.Router, values context.Values) *Server {
	s := &Server{
		router: router,
	}
	s.useCtxValues(values)
	return s
}

func (s *Server) Register(endpoints ...Endpoint) {
	for _, endpoint := range endpoints {
		s.registerSingle(endpoint)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) useCtxValues(values context.Values) {
	if values == nil {
		return
	}
	s.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValues(r.Context(), values)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	})
}

func (s *Server) registerSingle(ep Endpoint) {
	handler := ep.Handler()

	cleanPath := strings.TrimSuffix(ep.Path(), "/")

	// First registration matches without a trailing slash
	s.router.Methods(ep.Method()).
		Path(cleanPath).
		HandlerFunc(handler)

	// Second registration matches with a trailing slash
	s.router.Methods(ep.Method()).
		Path(cleanPath + "/").
		HandlerFunc(handler)
}

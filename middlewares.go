package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Middlewares []mux.MiddlewareFunc

func (m *Middlewares) Use(middleware mux.MiddlewareFunc) {
	*m = append(*m, middleware)
}

func (m Middlewares) Handler(handler http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		handler = m[i](handler)
	}
	return handler
}

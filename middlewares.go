package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Middleware interface {
	Middleware(handler http.Handler) http.Handler
}

type MiddlewareChain struct {
	Middlewares []Middleware
}

func (c *MiddlewareChain) Use(mwf ...mux.MiddlewareFunc) {
	for _, fn := range mwf {
		c.Middlewares = append(c.Middlewares, fn)
	}
}

func (c *MiddlewareChain) Apply(handler http.Handler) http.Handler {
	cur := handler
	for i := len(c.Middlewares) - 1; i >= 0; i-- {
		cur = c.Middlewares[i].Middleware(cur)
	}
	return cur
}

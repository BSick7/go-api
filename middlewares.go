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

// Apply takes an input http.Handler and applies the middleware chain over a single http.Handler
// These middlewares are chained in the order they are registered
// If middleware A is registered earlier than B, then execution will be A => B => handler
func (c *MiddlewareChain) Apply(handler http.Handler) http.Handler {
	cur := handler
	for i := len(c.Middlewares) - 1; i >= 0; i-- {
		cur = c.Middlewares[i].Middleware(cur)
	}
	return cur
}

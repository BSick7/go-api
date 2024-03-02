package api

import "net/http"

type MiddlewareFunc func(http.Handler) http.Handler

func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

type Middleware interface {
	Middleware(handler http.Handler) http.Handler
}

type MiddlewareChain struct {
	Middlewares []Middleware
}

func (c *MiddlewareChain) Use(mwf ...MiddlewareFunc) {
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

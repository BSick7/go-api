package api

import (
	"net/http"
)

// DefaultFallbackBehavior configures the api server with default NotFound and MethodNotAllowed handlers
// These handlers will be wrapped with the middlewares that are configured on the api server
func DefaultFallbackBehavior(apiServer *Server) {
	apiServer.NotFoundHandler = MiddlewaresHandler(apiServer, http.NotFoundHandler())
	apiServer.MethodNotAllowedHandler = MiddlewaresHandler(apiServer, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))
}

// MiddlewaresHandler wraps a raw handler with middlewares
// This can be used to ensure NotFound and MethodNotAllowed are logged
func MiddlewaresHandler(apiServer *Server, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		final := handler
		for i := len(apiServer.Middlewares) - 1; i >= 0; i-- {
			final = apiServer.Middlewares[i].Middleware(final)
		}
		final.ServeHTTP(w, r)
	})
}

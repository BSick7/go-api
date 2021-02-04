package errors

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ReporterMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			container := &Container{}
			newCtx := ContextWithContainer(r.Context(), container)
			handler.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

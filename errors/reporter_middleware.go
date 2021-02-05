package errors

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ReporterMiddleware(onReport func(container *Container, r *http.Request)) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			container := &Container{}
			newCtx := ContextWithContainer(r.Context(), container)
			newRequest := r.WithContext(newCtx)
			handler.ServeHTTP(w, newRequest)
			if onReport != nil {
				onReport(container, newRequest)
			}
		})
	}
}

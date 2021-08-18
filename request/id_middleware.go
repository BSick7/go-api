package request

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	IdHeader = "X-Request-Id"
)

func IdMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(IdHeader)
			if requestId == "" {
				requestId = uuid.New().String()
				r.Header.Set(IdHeader, requestId)
			}
			w.Header().Set(IdHeader, requestId)
			handler.ServeHTTP(w, r)
		})
	}
}

func GetId(r *http.Request) string {
	return r.Header.Get(IdHeader)
}

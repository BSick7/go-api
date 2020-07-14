package gzip

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xi2/httpgzip"
)

func Middleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return httpgzip.NewHandler(handler, httpgzip.DefaultContentTypes)
	}
}

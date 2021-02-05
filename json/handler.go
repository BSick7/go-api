package json

import (
	"net/http"
	"time"

	"github.com/BSick7/go-api/errors"
)

type HandlerFunc func(res *ResponseWriter, req *Request)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			errContainer:   errors.ContainerFromContext(r.Context()),
		}
		req := &Request{Request: r}
		handler(res, req)
	})
}

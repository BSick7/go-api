package jsonapi

import (
	"net/http"
	"time"

	"github.com/svanharmelen/jsonapi"
)

type HandlerFunc func(res *ResponseWriter, req *Request)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonapi.MediaType)
		res := &ResponseWriter{ResponseWriter: w, start: time.Now()}
		req := &Request{Request: r}
		handler(res, req)
	})
}

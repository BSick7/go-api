package jsonapi

import (
	"github.com/BSick7/go-api/errors"
	"github.com/svanharmelen/jsonapi"
	"net/http"
	"time"
)

type HandlerFunc func(res *ResponseWriter, req *Request)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonapi.MediaType)
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			Obscurer:       errors.ObscurerFromContext(r.Context()),
			ErrorCapturer:  errors.CapturerFromContext(r.Context()),
		}
		req := &Request{Request: r}
		handler(res, req)
	})
}

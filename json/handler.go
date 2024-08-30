package json

import (
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/logging"
	"net/http"
	"time"
)

type HandlerFunc func(res *ResponseWriter, req *Request)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			Obscurer:       errors.ObscurerFromContext(r.Context()),
			ErrorCapturer:  errors.CapturerFromContext(r.Context()),
		}
		req := &Request{
			Request: r,
			Logger:  logging.LoggerFromContext(r.Context()),
		}
		handler(res, req)
	})
}

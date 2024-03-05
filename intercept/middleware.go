package intercept

import (
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type ResponseData struct {
	// StatusCode represents the response's status code
	StatusCode int
	// Duration is the total duration of the request/response
	Duration time.Duration
	// Written is the number of bytes written in the response
	Written int64
}

type OnResponseFunc func(r *http.Request, data ResponseData)

func Middleware(onResponses ...OnResponseFunc) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := httpsnoop.CaptureMetrics(handler, w, r)
			for _, onResponse := range onResponses {
				onResponse(r, ResponseData{
					StatusCode: m.Code,
					Duration:   m.Duration,
					Written:    m.Written,
				})
			}
		})
	}
}

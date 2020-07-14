package logging

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	// Prefix is prepended to all log statements
	Prefix string

	// Log100s will determine whether an HTTP 1xx response is logged
	Log100s bool

	// Log200s will determine whether an HTTP 2xx response is logged
	Log200s bool

	// Log300s will determine whether an HTTP 3xx response is logged
	Log300s bool

	// Log400s will determine whether an HTTP 4xx response is logged
	Log400s bool

	// Log500s will determine whether an HTTP 5xx response is logged
	Log500s bool
}

func (c Config) ShouldLog(statusCode int) bool {
	if statusCode >= 100 && statusCode < 200 {
		return c.Log100s
	}
	if statusCode >= 200 && statusCode < 300 {
		return c.Log200s
	}
	if statusCode >= 300 && statusCode < 400 {
		return c.Log300s
	}
	if statusCode >= 400 && statusCode < 500 {
		return c.Log400s
	}
	if statusCode >= 500 {
		return c.Log500s
	}
	return false
}

func EndpointLoggerMiddleware(cfg Config) mux.MiddlewareFunc {
	stdNoTime := log.New(os.Stderr, cfg.Prefix, 0)

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := &endpointLoggerWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
			handler.ServeHTTP(wrapped, r)
			if cfg.ShouldLog(wrapped.statusCode) {
				stdNoTime.Printf("%s %d %s %s%s", time.Since(start), wrapped.statusCode, r.Method, r.RequestURI, wrapped.ctx)
			}
		})
	}
}

type endpointLoggerWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	ctx        string
}

func (w *endpointLoggerWriterWrapper) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *endpointLoggerWriterWrapper) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

func (w *endpointLoggerWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

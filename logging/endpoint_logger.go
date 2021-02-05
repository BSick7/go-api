package logging

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func EndpointLoggerMiddleware(cfg Config) mux.MiddlewareFunc {
	stdNoTime := log.New(os.Stderr, cfg.Prefix, 0)

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := &endpointLoggerWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
			handler.ServeHTTP(wrapped, r)
			dur := time.Since(start)
			if cfg.OnRequest != nil {
				cfg.OnRequest(r, wrapped, dur)
			}
			if cfg.ShouldLog(wrapped.statusCode) {
				stdNoTime.Printf("%s %d %s %s %s", dur, wrapped.StatusCode(), r.Method, r.RequestURI, wrapped.Body())
			}
		})
	}
}

var _ ResponseData = &endpointLoggerWriterWrapper{}

type endpointLoggerWriterWrapper struct {
	http.ResponseWriter
	statusCode   int
	capturedData []string
}

func (w *endpointLoggerWriterWrapper) StatusCode() int {
	return w.statusCode
}

func (w *endpointLoggerWriterWrapper) Body() string {
	return strings.Join(w.capturedData, "")
}

func (w *endpointLoggerWriterWrapper) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *endpointLoggerWriterWrapper) Write(data []byte) (int, error) {
	if w.statusCode >= 400 {
		w.capturedData = append(w.capturedData, string(data))
	}
	return w.ResponseWriter.Write(data)
}

func (w *endpointLoggerWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

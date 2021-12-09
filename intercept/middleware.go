package intercept

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type ResponseData interface {
	StatusCode() int
	Body() string
}

type OnResponseFunc func(r *http.Request, data ResponseData, duration time.Duration)

func Middleware(onResponses ...OnResponseFunc) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := &responseWriterInterceptor{ResponseWriter: w, statusCode: http.StatusOK}
			handler.ServeHTTP(wrapped, r)
			for _, onResponse := range onResponses {
				onResponse(r, wrapped, time.Since(start))
			}
		})
	}
}

var _ ResponseData = &responseWriterInterceptor{}
var _ http.Hijacker = &responseWriterInterceptor{}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode   int
	capturedData []string
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, fmt.Errorf("can't switch protocols using non-Hijacker ResponseWriter type %T", w.ResponseWriter)
}

func (w *responseWriterInterceptor) StatusCode() int {
	return w.statusCode
}

func (w *responseWriterInterceptor) Body() string {
	return strings.Join(w.capturedData, "")
}

func (w *responseWriterInterceptor) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *responseWriterInterceptor) Write(data []byte) (int, error) {
	if w.statusCode >= 400 {
		w.capturedData = append(w.capturedData, string(data))
	}
	return w.ResponseWriter.Write(data)
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

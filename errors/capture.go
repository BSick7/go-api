package errors

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

// CaptureFunc is injected in the middleware
// This function is used to react to API erros
type CaptureFunc func(r *http.Request, statusCode int, err error)

// OnCaptureFunc is used by producers of errors to signal an error
// This is used as a wrapper in the capture middleware to forward the http.Request into the CaptureFunc
// This is necessary because the producer (i.e. ResponseWriter) doesn't have access to the http.Request
type OnCaptureFunc func(statusCode int, err error)

type captureContextKey struct{}

func CapturerFromContext(ctx context.Context) OnCaptureFunc {
	if fn, ok := ctx.Value(captureContextKey{}).(OnCaptureFunc); ok {
		return fn
	}
	return func(statusCode int, err error) {}
}

func ContextWithCapturer(ctx context.Context, capturer OnCaptureFunc) context.Context {
	if capturer == nil {
		return ctx
	}
	return context.WithValue(ctx, captureContextKey{}, capturer)
}

// CaptureMiddleware enables the API server to capture the body of error responses
// These errors can be used to perform any action like reporting an error to a telemetry provider.
func CaptureMiddleware(capturer CaptureFunc) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := ContextWithCapturer(r.Context(), func(statusCode int, err error) {
				capturer(r, statusCode, err)
			})
			handler.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

package errors

import (
	"context"
	"github.com/BSick7/go-api/logging"
	"github.com/gorilla/mux"
	"net/http"
)

type obscurerContextKey struct{}

type Obscurer struct {
	ObscureInternal    bool
	LogOriginalErrorFn LogOriginalErrorFunc
}

func (c Obscurer) Obscure(err error) ApiError {
	if !c.ObscureInternal {
		return ApiError{Err: err}
	}
	if c.LogOriginalErrorFn != nil {
		c.LogOriginalErrorFn(err)
	}
	return ApiError{Err: nil}
}

type LogOriginalErrorFunc func(err error)

func ObscurerFromContext(ctx context.Context) Obscurer {
	if val, ok := ctx.Value(obscurerContextKey{}).(Obscurer); ok {
		return val
	}
	return Obscurer{}
}

func ContextWithObscurer(ctx context.Context, obscurer Obscurer) context.Context {
	return context.WithValue(ctx, obscurerContextKey{}, obscurer)
}

func ObscureInternalErrorsMiddleware(logOriginal bool) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			obscurer := Obscurer{
				ObscureInternal:    true,
				LogOriginalErrorFn: createErrorLogger(r, logOriginal),
			}
			handler.ServeHTTP(w, r.WithContext(ContextWithObscurer(r.Context(), obscurer)))
		})
	}
}

func createErrorLogger(r *http.Request, logOriginal bool) LogOriginalErrorFunc {
	if !logOriginal {
		return nil
	}
	logger := logging.LoggerFromContext(r.Context())
	return func(err error) {
		logger.Error(err.Error())
	}
}

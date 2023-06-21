package errors

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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
			obscurer := Obscurer{ObscureInternal: true}
			if logOriginal {
				logger := log.New(os.Stderr, "", 0)
				if requestId := r.Header.Get("X-Request-ID"); requestId != "" {
					logger.SetPrefix(fmt.Sprintf("[%s] ", requestId))
				}
				obscurer.LogOriginalErrorFn = func(err error) {
					logger.Println(err.Error())
				}
			}
			handler.ServeHTTP(w, r.WithContext(ContextWithObscurer(r.Context(), obscurer)))
		})
	}
}

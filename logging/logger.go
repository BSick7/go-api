package logging

import (
	"context"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
)

type contextKey struct{}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	val, _ := ctx.Value(contextKey{}).(*slog.Logger)
	if val == nil {
		return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
	}
	return val
}

func ContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func AddLogger() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{ReplaceAttr: removeKeys(slog.TimeKey)}))
			req := r.WithContext(ContextWithLogger(r.Context(), logger))
			handler.ServeHTTP(w, req)
		})
	}
}

func removeKeys(keys ...string) func([]string, slog.Attr) slog.Attr {
	return func(_ []string, a slog.Attr) slog.Attr {
		for _, k := range keys {
			if a.Key == k {
				return slog.Attr{}
			}
		}
		return a
	}
}

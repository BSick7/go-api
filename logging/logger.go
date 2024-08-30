package logging

import (
	"context"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
)

const (
	IdHeader = "X-Request-Id"
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
			logger = logger.With(
				slog.String("scheme", detectScheme(r)),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
			)
			if requestId := r.Header.Get(IdHeader); requestId != "" {
				logger = logger.With(slog.String("x-request-id", requestId))
			}
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

func detectScheme(r *http.Request) string {
	isTls := r.TLS != nil

	// Check for WebSocket upgrade request
	if r.Header.Get("Connection") == "Upgrade" && r.Header.Get("Upgrade") == "websocket" {
		if isTls {
			return "wss"
		}
		return "ws"
	}

	// If behind a proxy, check for X-Forwarded-Proto header
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		return proto
	}

	if isTls {
		return "https "
	}
	return "http"
}

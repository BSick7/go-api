package logging

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"
)

func FallbackBehavior(router *mux.Router, logger *slog.Logger) {
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(r.Context(), "not found",
			slog.Duration("duration", time.Duration(0)),
			slog.Int("status_code", http.StatusNotFound),
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI))

		http.NotFound(w, r)
	})
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(r.Context(), "method not allowed",
			slog.Duration("duration", time.Duration(0)),
			slog.Int("status_code", http.StatusMethodNotAllowed),
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI))

		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

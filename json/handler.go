package json

import (
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/logging"
	"github.com/BSick7/go-api/request"
	"log/slog"
	"net/http"
	"time"
)

type HandlerFunc func(res *ResponseWriter, req *Request)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			Obscurer:       errors.ObscurerFromContext(r.Context()),
			ErrorCapturer:  errors.CapturerFromContext(r.Context()),
		}
		logger := logging.LoggerFromContext(r.Context())
		logger = logger.With(
			slog.String("scheme", r.URL.Scheme),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)
		if requestId := r.Header.Get(request.IdHeader); requestId != "" {
			logger = logger.With(slog.String("x-request-id", requestId))
		}
		req := &Request{
			Request: r.WithContext(logging.ContextWithLogger(r.Context(), logger)),
			Logger:  logger,
		}
		handler(res, req)
	})
}

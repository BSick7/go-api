package logging

import (
	"github.com/BSick7/go-api/intercept"
	"log/slog"
	"net/http"
	"time"
)

func LogAllRequests(logger *slog.Logger) intercept.OnResponseFunc {
	return func(r *http.Request, data intercept.ResponseData, duration time.Duration) {
		logger.InfoContext(r.Context(), data.Body(),
			slog.Duration("duration", duration),
			slog.Int("status_code", data.StatusCode()),
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI))
	}
}

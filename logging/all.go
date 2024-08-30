package logging

import (
	"github.com/BSick7/go-api/intercept"
	"log/slog"
	"net/http"
)

func LogAllRequests() intercept.OnResponseFunc {
	return func(r *http.Request, data intercept.ResponseData) {
		logger := LoggerFromContext(r.Context())
		logger = logger.With(
			slog.Int("status-code", data.StatusCode),
			slog.Duration("duration", data.Duration),
		)
		logger.Info("completed")
	}
}

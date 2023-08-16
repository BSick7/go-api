package recovery

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

type PanicRecoveryFunc func(req *http.Request, err interface{})

func PanicMiddleware(logger *slog.Logger, fns ...PanicRecoveryFunc) mux.MiddlewareFunc {
	panicLogger := logger.With("source", "PANIC")
	return func(next http.Handler) http.Handler {
		return &panicRecoveryHandler{
			next: next,
			fn: func(req *http.Request, err interface{}) {
				panicLogger.Error(fmt.Sprintf("%s", err))
				debug.PrintStack()
				for _, fn := range fns {
					fn(req, err)
				}
			},
		}
	}
}

type panicRecoveryHandler struct {
	next http.Handler
	fn   PanicRecoveryFunc
}

func (h *panicRecoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.fn(req, err)
		}
	}()
	h.next.ServeHTTP(w, req)
}

package recovery

import (
	"fmt"
	"github.com/BSick7/go-api/logging"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

type PanicRecoveryFunc func(req *http.Request, err PanicError)

var _ error = PanicError{}

type PanicError struct {
	stack    []byte
	rawError interface{}
}

func (e PanicError) StackTrace() []byte {
	return e.stack
}

func (e PanicError) Error() string {
	if v, ok := e.rawError.(error); ok {
		return v.Error()
	}
	return fmt.Sprintf("%s", e.rawError)
}

func PanicMiddleware(fns ...PanicRecoveryFunc) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return &panicRecoveryHandler{
			next: next,
			fn: func(req *http.Request, err PanicError) {
				logger := logging.LoggerFromContext(req.Context())
				logger = logger.With(
					slog.String("panic", "PANIC"),
					slog.String("stack-trace", string(err.StackTrace())),
				)
				logger.Error(err.Error())
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
			h.fn(req, PanicError{rawError: err, stack: debug.Stack()})
		}
	}()
	h.next.ServeHTTP(w, req)
}

package recovery

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	panicLogger := log.New(os.Stderr, "[PANIC] ", 0)
	return func(next http.Handler) http.Handler {
		return &panicRecoveryHandler{
			next: next,
			fn: func(req *http.Request, err PanicError) {
				panicLogger.Println(err, string(err.StackTrace()))
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
			h.fn(req, PanicError{rawError: err, stack: debug.Stack()})
		}
	}()
	h.next.ServeHTTP(w, req)
}

package recovery

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/BSick7/go-api/errors"
	"github.com/gorilla/mux"
)

type PanicError struct {
	Err   interface{}
	Stack []byte
}

func (e PanicError) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, e.Stack)
}

func PanicMiddleware() mux.MiddlewareFunc {
	panicLogger := log.New(os.Stderr, "[PANIC] ", 0)
	report := func(err interface{}, r *http.Request) {
		perr := PanicError{
			Err:   err,
			Stack: debug.Stack(),
		}
		container := errors.ContainerFromContext(r.Context())
		if container != nil {
			container.AddError(perr)
		}
		panicLogger.Println(perr)
	}

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					report(err, r)
				}
			}()

			handler.ServeHTTP(w, r)
		})
	}
}

package recovery

import (
	"log"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func PanicMiddleware() mux.MiddlewareFunc {
	panicLogger := log.New(os.Stderr, "[PANIC] ", 0)
	return handlers.RecoveryHandler(handlers.RecoveryLogger(panicLogger),
		handlers.PrintRecoveryStack(true))
}

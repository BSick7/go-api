package logging

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func FallbackBehavior(router *mux.Router) {
	stdNoTime := log.New(os.Stderr, "", 0)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stdNoTime.Printf("%s %d %s %s", time.Duration(0), http.StatusNotFound, r.Method, r.RequestURI)
		http.NotFound(w, r)
	})
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stdNoTime.Printf("%s %d %s %s", time.Duration(0), http.StatusMethodNotAllowed, r.Method, r.RequestURI)
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}

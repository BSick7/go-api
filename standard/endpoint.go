package standard

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
)

func NewEndpoint(method string, path string, handler http.Handler) Endpoint {
	return Endpoint{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

type Endpoint struct {
	Method      string
	Path        string
	Handler     http.Handler
	handlerName string
}

func (e Endpoint) Identifier() string {
	if e.handlerName == "" {
		e.handlerName = runtime.FuncForPC(reflect.ValueOf(e.Handler).Pointer()).Name()
	}
	return fmt.Sprintf("%s\t%s\t%s", e.Method, e.Path, e.handlerName)
}

func (e Endpoint) Register(router *mux.Router) {
	cleanPath := strings.TrimSuffix(e.Path, "/")

	// First registration matches without a trailing slash
	router.Methods(e.Method).
		Path(cleanPath).
		Handler(e.Handler)

	// Second registration matches with a trailing slash
	router.Methods(e.Method).
		Path(cleanPath + "/").
		Handler(e.Handler)
}

package raw

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// NewEndpoint creates a raw endpoint, accepting a plain http.Handler
func NewEndpoint(method string, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		method:      method,
		path:        path,
		handler:     handler,
		handlerName: runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
	}
}

type Endpoint struct {
	method      string
	path        string
	handlerName string
	handler     http.HandlerFunc
}

func (e Endpoint) Method() string      { return e.method }
func (e Endpoint) Path() string        { return e.path }
func (e Endpoint) HandlerName() string { return e.handlerName }
func (e Endpoint) String() string {
	return fmt.Sprintf("%s %s %s", e.method, e.path, e.handlerName)
}

func (e Endpoint) Handler() http.HandlerFunc {
	return e.handler
}

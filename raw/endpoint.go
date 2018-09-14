package raw

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"
)

var stdNoTime = log.New(os.Stderr, "", 0)

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
	unlogged    bool
}

func (e Endpoint) Method() string      { return e.method }
func (e Endpoint) Path() string        { return e.path }
func (e Endpoint) HandlerName() string { return e.handlerName }
func (e Endpoint) String() string {
	return fmt.Sprintf("%s %s %s", e.method, e.path, e.handlerName)
}

func (e *Endpoint) Unlogged() *Endpoint {
	e.unlogged = true
	return e
}

func (e Endpoint) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		e.handler(w, r)
		if !e.unlogged {
			stdNoTime.Printf("%s %s %s", time.Since(start), r.RequestURI, e)
		}
	}
}

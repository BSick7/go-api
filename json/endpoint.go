package json

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/BSick7/go-api"
)

var stdNoTime = log.New(os.Stderr, "", 0)

func NewEndpoint(method string, path string, handler api.EndpointHandler) *Endpoint {
	return &Endpoint{
		method:      method,
		path:        path,
		handlerName: runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
		handler:     handler,
	}
}

type Endpoint struct {
	method      string
	path        string
	handlerName string
	handler     api.EndpointHandler
}

func (e Endpoint) Method() string      { return e.method }
func (e Endpoint) Path() string        { return e.path }
func (e Endpoint) HandlerName() string { return e.handlerName }
func (e Endpoint) String() string {
	return fmt.Sprintf("%s %s %s", e.method, e.path, e.handlerName)
}
func (e Endpoint) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		res := NewResponder(e, w, r)
		req := NewRequest(r)

		e.handler(res, req)

		statusCode, statusCtx := res.Status()
		stdNoTime.Printf("%s %d %s %s%s", time.Since(start), statusCode, r.RequestURI, e, statusCtx)
	}
}

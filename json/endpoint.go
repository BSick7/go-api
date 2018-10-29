package json

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/BishopFox/go-api"
	"github.com/BishopFox/go-api/logging"
)

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
		w.Header().Set("Content-Type", "application/json")
		res := NewResponder(e, w, r)
		req := NewRequest(r)
		e.handler(res, req)
		if rcw, ok := w.(logging.ResponseContextWriter); ok {
			_, ctx := res.Status()
			rcw.WriteContext(ctx)
		}
	}
}

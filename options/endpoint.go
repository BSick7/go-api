package options

import (
	"net/http"

	"github.com/gorilla/mux"
)


const DefaultAllowOrigin = "*"
const DefaultAllowMethods = "DELETE, GET, OPTIONS, POST, PUT"
const DefaultAllowHeaders = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"

func NewEndpoint(allowOrigin, allowMethods, allowHeaders string) Endpoint {
	if allowOrigin == "" {
		allowOrigin = DefaultAllowOrigin
	}
	if allowMethods == "" {
		allowMethods = DefaultAllowMethods
	}
	if allowHeaders == "" {
		allowHeaders = DefaultAllowHeaders
	}
	return Endpoint{
		allowOrigin:  allowOrigin,
		allowMethods: allowMethods,
		allowHeaders: allowHeaders,
	}
}

type Endpoint struct {
	allowOrigin  string
	allowMethods string
	allowHeaders string
}

func (e Endpoint) Description() string {
	return "/*\tOPTIONS%s\tOptions"
}

func (e Endpoint) Register(router *mux.Router) {
	router.Methods("OPTIONS").HandlerFunc(e.handler)
}

func (e Endpoint) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", e.allowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", e.allowHeaders)
	w.Header().Set("Access-Control-Allow-Methods", e.allowMethods)
	w.WriteHeader(http.StatusNoContent)
}

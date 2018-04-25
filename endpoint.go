package api

import (
	"net/http"
)

type Endpoint interface {
	Method() string
	Path() string
	HandlerName() string
	String() string
	Handler() http.HandlerFunc
}

type EndpointHandler func(res Responder, req Request)

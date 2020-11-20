package api

import (
	"fmt"

	"github.com/gorilla/mux"
)

var _ Endpoint = &EndpointGroup{}

type EndpointGroup struct {
	Prefix    string
	Endpoints []Endpoint
}

func (e EndpointGroup) Identifier() string {
	return fmt.Sprintf("\t%s\t", e.Prefix)
}

func (e EndpointGroup) Register(router *mux.Router) {
	subrouter := router.Path(e.Prefix).Subrouter()
	for _, endpoint := range e.Endpoints {
		endpoint.Register(subrouter)
	}
}

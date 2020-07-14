package api

import (
	"github.com/gorilla/mux"
)

type Endpoint interface {
	// Identifier provides identification in logs
	Identifier() string

	// Register is responsible for registering this endpoint with a gorilla router
	Register(router *mux.Router)
}

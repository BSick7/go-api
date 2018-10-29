package api_test

import (
	"testing"

	"github.com/BishopFox/go-api"
	"github.com/BishopFox/go-api/json"
	"github.com/BishopFox/go-api/logging"
	"github.com/gorilla/mux"
)

func TestServer(t *testing.T) {
	apiServer := api.NewServer(mux.NewRouter().StrictSlash(false))
	apiServer.Register(json.NewEndpoint("GET", "/", func(res api.Responder, req api.Request) {
		res.Send(nil)
	}))
	apiServer.AttachMiddleware(logging.EndpointLoggerMiddleware(logging.EndpointLoggerConfig{
		Prefix:     "[example-api] ",
		LogInitial: true,
	}))
}

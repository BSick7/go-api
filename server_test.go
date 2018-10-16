package api_test

import (
	"testing"

	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/json"
	"github.com/BSick7/go-api/logging"
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

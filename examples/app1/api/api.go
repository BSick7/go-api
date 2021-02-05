package app1

import (
	"net/http"
	
	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/cors"
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/gzip"
	"github.com/BSick7/go-api/json"
	"github.com/BSick7/go-api/logging"
	"github.com/BSick7/go-api/recovery"
	"github.com/BSick7/go-api/standard"
	"github.com/gorilla/mux"
)

func Server() *api.Server {
	loggingCfg := logging.Config{
		Prefix:  "[app1]",
		Log100s: false,
		Log200s: false,
		Log300s: false,
		Log400s: false,
		Log500s: false,
	}

	apiServer := &api.Server{
		Router: mux.NewRouter().
			StrictSlash(false).
			SkipClean(true).
			UseEncodedPath(),
	}
	api.DefaultFallbackBehavior(apiServer)
	apiServer.Use(logging.EndpointLoggerMiddleware(loggingCfg))
	apiServer.Use(gzip.Middleware())
	apiServer.Use(errors.ReporterMiddleware(func(container *errors.Container, r *http.Request) {
		// NOTE: An api would typically report errors here
	}))
	apiServer.Use(recovery.PanicMiddleware())
	apiServer.Use(cors.Middleware(cors.DefaultSettings))
	apiServer.Register(endpoints...)
	apiServer.Register(cors.Preflight())
	return apiServer
}

var endpoints = []api.Endpoint{
	standard.Endpoint{
		Method:  "GET",
		Path:    "/simple",
		Handler: json.Handler(Simple),
	},
	standard.Endpoint{
		Method:  "GET",
		Path:    "/paths/{path}/detail",
		Handler: json.Handler(UrlEncodedPath),
	},
}

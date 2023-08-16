package app1

import (
	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/cors"
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/gzip"
	"github.com/BSick7/go-api/intercept"
	"github.com/BSick7/go-api/json"
	"github.com/BSick7/go-api/jwt"
	"github.com/BSick7/go-api/logging"
	"github.com/BSick7/go-api/recovery"
	"github.com/BSick7/go-api/standard"
	"github.com/gorilla/mux"
	"log/slog"
)

func Server(logger *slog.Logger) *api.Server {
	if logger == nil {
		logger = slog.Default()
	}

	apiServer := &api.Server{
		Router: mux.NewRouter().
			StrictSlash(false).
			SkipClean(true).
			UseEncodedPath(),
	}
	api.DefaultFallbackBehavior(apiServer)
	apiServer.Use(gzip.Middleware())
	apiServer.Use(recovery.PanicMiddleware(logger))
	apiServer.Use(cors.Middleware(cors.DefaultSettings))
	apiServer.Use(jwt.Middleware())
	apiServer.Use(intercept.Middleware(false, logging.LogAllRequests(logger)))
	apiServer.Use(errors.ObscureInternalErrorsMiddleware(logger, true))
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

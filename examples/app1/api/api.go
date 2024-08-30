package app1

import (
	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/cors"
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/gzip"
	"github.com/BSick7/go-api/intercept"
	"github.com/BSick7/go-api/json"
	ga_jwt "github.com/BSick7/go-api/jwt"
	"github.com/BSick7/go-api/logging"
	"github.com/BSick7/go-api/recovery"
	"github.com/BSick7/go-api/standard"
	"github.com/cristalhq/jwt/v3"
	"github.com/gorilla/mux"
	"log"
	"log/slog"
	"net/http"
)

func Server() *api.Server {
	apiServer := &api.Server{
		Router: mux.NewRouter().
			StrictSlash(false).
			SkipClean(true).
			UseEncodedPath(),
	}
	api.DefaultFallbackBehavior(apiServer)
	apiServer.Use(logging.AddLogger(nil))
	apiServer.Use(gzip.Middleware())
	apiServer.Use(recovery.PanicMiddleware())
	apiServer.Use(cors.Middleware(cors.DefaultSettings))
	apiServer.Use(errors.CaptureMiddleware(func(r *http.Request, statusCode int, err error) {
		logger := logging.LoggerFromContext(r.Context())
		logger.Info("captured api error", slog.String("error", err.Error()))
	}))
	apiServer.Use(ga_jwt.ClaimsMiddleware[jwt.StandardClaims](handleJwtError))
	apiServer.Use(intercept.Middleware(logging.LogAllRequests()))
	apiServer.Use(errors.ObscureInternalErrorsMiddleware(true))
	apiServer.Register(endpoints...)
	apiServer.Register(cors.Preflight())
	return apiServer
}

func handleJwtError(w http.ResponseWriter, r *http.Request, err error) bool {
	log.Println("handleJwtError", err.Error())
	return true
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

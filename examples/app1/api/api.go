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
	apiServer.Use(gzip.Middleware())
	apiServer.Use(recovery.PanicMiddleware())
	apiServer.Use(cors.Middleware(cors.DefaultSettings))
	apiServer.Use(errors.CaptureMiddleware(func(r *http.Request, statusCode int, err error) {
		log.Printf("captured api error [code = %d, err = %s]\n", statusCode, err)
	}))
	apiServer.Use(ga_jwt.ClaimsMiddleware[jwt.StandardClaims](handleJwtError))
	apiServer.Use(intercept.Middleware(logging.LogAllRequests("[app1] ")))
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

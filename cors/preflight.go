package cors

import (
	"log"
	"net/http"

	"github.com/BSick7/go-api"
	"github.com/gorilla/mux"
)

func Preflight() api.Endpoint {
	return preflight{}
}

type preflight struct {
}

func (e preflight) Identifier() string {
	return "/*\tOPTIONS\tOptions"
}
func (e preflight) Register(router *mux.Router) {
	// This CORS preflight is meant to ensure that the OPTIONS middleware is executed on routes
	// We use a specific matcher function instead of .Methods(OPTIONS) to prevent the router from responding with 405 for unregistered routes
	matcherFunc := func(r *http.Request, match *mux.RouteMatch) bool {
		return r.Method == http.MethodOptions
	}

	router.MatcherFunc(matcherFunc).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("This OPTIONS handler is only used to match routes and should never hit. The CORS middleware intercepts OPTIONS requests and handles requests.")
		})
}

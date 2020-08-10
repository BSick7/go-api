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
	router.Methods("OPTIONS").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("This OPTIONS handler is only used to match routes and should never hit. The CORS middleware intercepts OPTIONS requests and handles requests.")
		})
}

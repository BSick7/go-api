package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Authorizer func(r *http.Request) (*http.Request, error)

// AuthorizeMiddleware creates an authorization middleware to respond based on the result of an authorization sequence
// If using gorilla mux router, this middleware will only hit when a route is matched
//   This is because gorilla matches a route, *then* runs the full chain of middlewares + handler
// Because of this, this makes mux.Vars(r) available to use to grab path parameters that gorilla parsed
func AuthorizeMiddleware(authorize Authorizer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newRequest, err := authorize(r)
			if err != nil {
				log.Println(fmt.Sprintf("forbidden: %s", err))
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`Forbidden`))
				return
			}
			next.ServeHTTP(w, newRequest)
		})
	}
}

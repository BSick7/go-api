package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Authenticator func(r *http.Request) (*http.Request, error)

func AuthenticateMiddleware(authenticate Authenticator) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newRequest, err := authenticate(r)
			if err != nil {
				// Don't hate me for using Unauthorized for Unauthenticated...
				// See https://stackoverflow.com/questions/3297048/403-forbidden-vs-401-unauthorized-http-responses
				log.Println(fmt.Sprintf("unauthenticated: %s", err))
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`Unauthenticated`))
				return
			}

			next.ServeHTTP(w, newRequest)
		})
	}
}

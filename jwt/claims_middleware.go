package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type claimsContextKey struct{}

func ContextWithClaims[T any](ctx context.Context, claims *T) context.Context {
	return context.WithValue(ctx, claimsContextKey{}, claims)
}

func ClaimsFromContext[T any](ctx context.Context) *T {
	if val, ok := ctx.Value(claimsContextKey{}).(*T); ok {
		return val
	}
	return nil
}

// ErrorHandlerFunc intercepts error handling for ClaimsMiddleware
// Return false if you want to terminate the request and respond to the user
// Return true to continue the request
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error) bool

func ClaimsMiddleware[T any](errorHandler ErrorHandlerFunc) mux.MiddlewareFunc {
	if errorHandler == nil {
		errorHandler = func(w http.ResponseWriter, r *http.Request, err error) bool {
			log.Println(err.Error())
			return true
		}
	}

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := ExtractBearerTokenFromRequest(r)
			if err != nil {
				innerErr := fmt.Errorf("error reading jwt token from Authorization Bearer token: %s", err)
				if !errorHandler(w, r, innerErr) {
					return
				}
			}

			if token == nil {
				handler.ServeHTTP(w, r)
				return
			}

			ctx := r.Context()
			ctx = ContextWithToken(ctx, token)

			claims, err := ParseClaims[T](token)
			if err != nil {
				innerErr := fmt.Errorf("error parsing claims from jwt token: %s", err)
				if !errorHandler(w, r, innerErr) {
					return
				}
			}
			if claims != nil {
				ctx = ContextWithClaims(ctx, claims)
			}

			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ParseClaims[T any](token *jwt.Token) (*T, error) {
	var claims T
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}

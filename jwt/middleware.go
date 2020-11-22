package jwt

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/cristalhq/jwt/v3"
	"github.com/gorilla/mux"
)

type tokenContextKey struct{}

func ContextWithToken(ctx context.Context, claims *jwt.Token) context.Context {
	return context.WithValue(ctx, tokenContextKey{}, claims)
}

func TokenFromContext(ctx context.Context) *jwt.Token {
	if val, ok := ctx.Value(tokenContextKey{}).(*jwt.Token); ok {
		return val
	}
	return nil
}

// Middleware parses JWT token from Authorization Bearer token
// This middleware does *not* run JWT verification
func Middleware() mux.MiddlewareFunc {
	extractToken := func(r *http.Request) (*jwt.Token, error) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			return nil, nil
		}
		token, err := jwt.ParseString(strings.TrimPrefix(authorization, "Bearer "))
		if err != nil {
			return nil, err
		}
		return token, nil
	}

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := extractToken(r)
			if err != nil {
				log.Printf("error reading jwt token from Authorization Bearer token: %s\n", err)
			}
			if token != nil {
				wrappedRequest := r.WithContext(ContextWithToken(r.Context(), token))
				handler.ServeHTTP(w, wrappedRequest)
			} else {
				handler.ServeHTTP(w, r)
			}
		})
	}
}

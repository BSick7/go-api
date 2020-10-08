package cors

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Middleware(settings Settings) mux.MiddlewareFunc {
	opts := []handlers.CORSOption{
		handlers.AllowedOrigins(settings.AllowedOrigins),
		handlers.AllowedHeaders(settings.AllowedHeaders),
		handlers.AllowedMethods(settings.AllowedMethods),
		handlers.ExposedHeaders(settings.ExposedHeaders),
	}
	if settings.AllowCredentials {
		opts = append(opts, handlers.AllowCredentials())
	}
	return handlers.CORS(opts...)
}

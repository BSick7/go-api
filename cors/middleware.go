package cors

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Middleware(settings Settings) mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins(settings.AllowedOrigins),
		handlers.AllowedHeaders(settings.AllowedHeaders),
		handlers.AllowedMethods(settings.AllowedMethods),
		handlers.ExposedHeaders(settings.ExposedHeaders),
	)
}

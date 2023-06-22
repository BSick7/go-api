package json

import (
	"github.com/BSick7/go-api/errors"
	"net/http"
	"time"
)

type HandlerFunc func(res *ResponseWriter, req *http.Request)

type ReturnHandlerFunc[T any] func(res http.ResponseWriter, req *http.Request) (T, error)

func Handler(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			Obscurer:       errors.ObscurerFromContext(r.Context()),
		}
		handler(res, r)
	})
}

func ReturnHandler[T any](handler ReturnHandlerFunc[T]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
			Obscurer:       errors.ObscurerFromContext(r.Context()),
		}
		if result, err := handler(res, r); err != nil {
			res.SendError(err)
		} else {
			res.Send(result)
		}
	})
}

package json

import (
	"net/http"
	"time"
)

type ControllerFunc[T any] func(req *Request) (T, error)

func Controller[T any](controller ControllerFunc[T]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := &ResponseWriter{ResponseWriter: w, start: time.Now()}
		req := &Request{Request: r}
		if out, err := controller(req); err != nil {
			res.SendError(err)
		} else {
			res.Send(out)
		}
	})
}

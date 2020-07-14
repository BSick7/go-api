package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	start      time.Time
	statusCode int
}

func (r *ResponseWriter) SendError(statusCode int, err error) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
	encoder := json.NewEncoder(r.ResponseWriter)
	encoder.Encode(map[string]interface{}{"error": err.Error()})
}

func (r *ResponseWriter) SendNotFound(msg string) {
	r.SendError(http.StatusNotFound, fmt.Errorf(msg))
}

func (r *ResponseWriter) Send(data interface{}) {
	if data == nil {
		r.statusCode = http.StatusNoContent
		r.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		r.statusCode = http.StatusOK
		encoder := json.NewEncoder(r.ResponseWriter)
		encoder.Encode(data)
	}
}

func (r *ResponseWriter) Status() int {
	return r.statusCode
}

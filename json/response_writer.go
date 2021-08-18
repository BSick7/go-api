package json

import (
	"encoding/json"
	"github.com/BSick7/go-api/errors"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	start      time.Time
	statusCode int
}

func (r *ResponseWriter) SendError(err error) {
	if isc, ok := err.(StatusCoder); ok {
		r.statusCode = isc.StatusCode()
		r.ResponseWriter.WriteHeader(r.statusCode)
	} else {
		r.statusCode = http.StatusInternalServerError
		r.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(r.ResponseWriter)
	payloader, ok := err.(ResponsePayloader)
	if !ok {
		payloader = errors.ApiError{Err: err}
	}
	encoder.Encode(payloader.Payload())
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

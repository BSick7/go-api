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
	requestId  string
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
	if payloader, ok := err.(ResponsePayloader); ok {
		payload := payloader.Payload()
		payload["request_id"] = r.requestId
		encoder.Encode(payload)
	} else {
		encoder.Encode(err)
	}
}

func (r *ResponseWriter) SendRawError(statusCode int, err error) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
	encoder := json.NewEncoder(r.ResponseWriter)
	encoder.Encode(map[string]interface{}{"error": err.Error()})
}

func (r *ResponseWriter) SendRawNotFound(msg string) {
	r.SendRawError(http.StatusNotFound, fmt.Errorf(msg))
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

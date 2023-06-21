package json

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	api_errors "github.com/BSick7/go-api/errors"
	"net"
	"net/http"
	"time"
)

type StatusCoder interface {
	StatusCode() int
}

type ResponsePayloader interface {
	Payload() map[string]interface{}
}

var _ http.Hijacker = &ResponseWriter{}

type ResponseWriter struct {
	http.ResponseWriter
	start      time.Time
	statusCode int
}

func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, fmt.Errorf("can't switch protocols using non-Hijacker ResponseWriter type %T", w.ResponseWriter)
}

func (w *ResponseWriter) SendError(err error) {
	if isc, ok := err.(StatusCoder); ok {
		w.statusCode = isc.StatusCode()
		w.ResponseWriter.WriteHeader(w.statusCode)
	} else {
		w.statusCode = http.StatusInternalServerError
		w.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}

	var rerr api_errors.ResponseErrorer
	if errors.As(err, &rerr) {
		err = rerr.ResponseError()
	}

	encoder := json.NewEncoder(w.ResponseWriter)
	payloader, ok := err.(ResponsePayloader)
	if !ok {
		payloader = api_errors.ApiError{Err: err}
	}
	if err := encoder.Encode(payloader.Payload()); err != nil {
		fmt.Printf("[go-api/json/response_writer] Error encoding error payload: %s\n", err)
	}
}

func (w *ResponseWriter) Send(data interface{}) {
	if data == nil {
		w.statusCode = http.StatusNoContent
		w.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		w.statusCode = http.StatusOK
		encoder := json.NewEncoder(w.ResponseWriter)
		if err := encoder.Encode(data); err != nil {
			fmt.Printf("[go-api/json/response_writer] Error encoding payload: %s\n", err)
		}
	}
}

func (w *ResponseWriter) Status() int {
	return w.statusCode
}

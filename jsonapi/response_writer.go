package jsonapi

import (
	"bufio"
	"errors"
	"fmt"
	api_errors "github.com/BSick7/go-api/errors"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/svanharmelen/jsonapi"
)

type HttpError struct {
	StatusCode int
	Errs       []*jsonapi.ErrorObject
}

func (e HttpError) Error() string {
	return fmt.Sprintf("http error (%d): %+v", e.StatusCode, e.Errs)
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

func (w *ResponseWriter) SendJsonApiErrors(errs []*jsonapi.ErrorObject) {
	if len(errs) < 1 {
		return
	}

	statusCode, _ := strconv.Atoi(errs[0].Code)
	w.statusCode = statusCode
	w.WriteHeader(statusCode)

	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Printf("error marshaling jsonapi errors to response: %s (%+v)", err, errs)
		w.statusCode = http.StatusInternalServerError
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (w *ResponseWriter) SendJsonApiError(err *jsonapi.ErrorObject) {
	w.SendJsonApiErrors([]*jsonapi.ErrorObject{err})
}

type StatusCoder interface {
	StatusCode() int
}

type ResponsePayloader interface {
	Payload() map[string]interface{}
}

func (w *ResponseWriter) SendError(err error) {
	var jaerr *jsonapi.ErrorObject
	if !errors.As(err, &jaerr) {
		jaerr = &jsonapi.ErrorObject{
			ID:     "",
			Title:  "General Error",
			Detail: err.Error(),
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Status: http.StatusText(http.StatusInternalServerError),
		}
		if isc, ok := err.(StatusCoder); ok {
			jaerr.Code = strconv.Itoa(isc.StatusCode())
			jaerr.Status = http.StatusText(isc.StatusCode())
		}
		var rerr api_errors.ResponseErrorer
		if errors.As(err, &rerr) {
			err = rerr.ResponseError()
		}
		if payloader, ok := err.(ResponsePayloader); ok {
			payload := payloader.Payload()
			jaerr.Title = payload["title"].(string)
			jaerr.Detail = payload["message"].(string)
		}
	}
	w.SendJsonApiError(jaerr)
}

func (w *ResponseWriter) Send(data interface{}) {
	if data == nil {
		w.statusCode = http.StatusNoContent
		w.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		if err := jsonapi.MarshalPayload(w.ResponseWriter, data); err != nil {
			w.statusCode = http.StatusInternalServerError
			http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		} else {
			w.statusCode = http.StatusOK
		}
	}
}

func (w *ResponseWriter) Status() int {
	return w.statusCode
}

package jsonapi

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/svanharmelen/jsonapi"
)

type ResponseWriter struct {
	http.ResponseWriter
	start      time.Time
	statusCode int
}

// SendErrorIfJsonApiError will respond to the request if the input err is a jsonapi.ErrorObject
func (r *ResponseWriter) SendErrorIfJsonApiError(err error) bool {
	jaerr, ok := err.(*jsonapi.ErrorObject)
	if !ok {
		return false
	}
	statusCode, _ := strconv.Atoi(jaerr.Code)
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)

	if err := jsonapi.MarshalPayload(r.ResponseWriter, jaerr); err != nil {
		r.statusCode = http.StatusInternalServerError
		http.Error(r.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
	return true
}

func (r *ResponseWriter) SendError(id string, statusCode int, title string, err error) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)

	errObject := &jsonapi.ErrorObject{
		ID:     id,
		Title:  title,
		Detail: err.Error(),
		Status: http.StatusText(statusCode),
		Code:   strconv.Itoa(statusCode),
		Meta:   nil,
	}

	if err := jsonapi.MarshalPayload(r.ResponseWriter, errObject); err != nil {
		r.statusCode = http.StatusInternalServerError
		http.Error(r.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (r *ResponseWriter) SendNotFound(id string, msg string) {
	r.SendError(id, http.StatusNotFound, "not found", fmt.Errorf(msg))
}

func (r *ResponseWriter) Send(data interface{}) {
	if data == nil {
		r.statusCode = http.StatusNoContent
		r.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		if err := jsonapi.MarshalPayload(r.ResponseWriter, data); err != nil {
			r.statusCode = http.StatusInternalServerError
			http.Error(r.ResponseWriter, err.Error(), http.StatusInternalServerError)
		} else {
			r.statusCode = http.StatusOK
		}
	}
}

func (r *ResponseWriter) Status() int {
	return r.statusCode
}
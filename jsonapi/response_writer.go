package jsonapi

import (
	"fmt"
	"log"
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

func (r *ResponseWriter) SendJsonApiErrors(errs []*jsonapi.ErrorObject) {
	if len(errs) < 1 {
		return
	}

	statusCode, _ := strconv.Atoi(errs[0].Code)
	r.statusCode = statusCode
	r.WriteHeader(statusCode)

	if err := jsonapi.MarshalErrors(r, errs); err != nil {
		log.Printf("error marshaling jsonapi errors to response: %s (%+v)", err, errs)
		r.statusCode = http.StatusInternalServerError
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
}

func (r *ResponseWriter) SendJsonApiError(err *jsonapi.ErrorObject) {
	r.SendJsonApiErrors([]*jsonapi.ErrorObject{err})
}

func (r *ResponseWriter) SendError(id string, statusCode int, title string, err error) {
	r.SendJsonApiError(&jsonapi.ErrorObject{
		ID:     id,
		Title:  title,
		Detail: err.Error(),
		Status: http.StatusText(statusCode),
		Code:   strconv.Itoa(statusCode),
	})
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

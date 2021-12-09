package jsonapi

import (
	"bufio"
	"fmt"
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

func (w *ResponseWriter) SendError(id string, statusCode int, title string, err error) {
	w.SendJsonApiError(&jsonapi.ErrorObject{
		ID:     id,
		Title:  title,
		Detail: err.Error(),
		Status: http.StatusText(statusCode),
		Code:   strconv.Itoa(statusCode),
	})
}

func (w *ResponseWriter) SendNotFound(id string, msg string) {
	w.SendError(id, http.StatusNotFound, "not found", fmt.Errorf(msg))
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

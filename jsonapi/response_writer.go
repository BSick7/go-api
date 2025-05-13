package jsonapi

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	api_errors "github.com/BSick7/go-api/errors"
	"github.com/hashicorp/jsonapi"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
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
	Obscurer      api_errors.Obscurer
	ErrorCapturer api_errors.OnCaptureFunc
	start         time.Time
	statusCode    int
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
		w.ErrorCapturer(w.statusCode, err)
	}
}

func (w *ResponseWriter) SendJsonApiError(err *jsonapi.ErrorObject) {
	w.SendJsonApiErrors([]*jsonapi.ErrorObject{err})
	w.ErrorCapturer(w.statusCode, err)
}

type StatusCoder interface {
	StatusCode() int
}

type ResponsePayloader interface {
	Payload() map[string]any
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

		// The interfaces defined in the errors package are not specific to jsonapi serialization
		// This ResponsePayloader allows us to structure content before emitting in the jsonapi format
		// Additionally, errors that aren't a ResponsePayloader are obscured (configurable through middleware)
		payloader, ok := err.(ResponsePayloader)
		if !ok {
			payloader = w.Obscurer.Obscure(err)
		}

		payload := payloader.Payload()
		jaerr.Title = payload["title"].(string)
		jaerr.Detail = payload["message"].(string)
	}
	w.SendJsonApiError(jaerr)
}

// DataDecorator allows you to emit `links` and `meta` along with your payload
// This is useful for things like adding pagination metadata to a list response
// To do things like this, implement jsonapi.Metable (for `meta`) and jsonapi.Linkable (for `links`)
type DataDecorator interface {
	Data() any
}

func (w *ResponseWriter) Send(data interface{}) {
	if data == nil {
		w.statusCode = http.StatusNoContent
		w.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		if decorator, ok := data.(DataDecorator); ok {
			w.sendDecorated(decorator)
		} else {
			w.sendNormal(data)
		}
	}
}

func (w *ResponseWriter) sendNormal(data any) {
	if err := jsonapi.MarshalPayload(w.ResponseWriter, data); err != nil {
		w.statusCode = http.StatusInternalServerError
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
	} else {
		w.statusCode = http.StatusOK
	}
}

func (w *ResponseWriter) sendDecorated(decorator DataDecorator) {
	payload, err := jsonapi.Marshal(decorator.Data())
	if err != nil {
		w.statusCode = http.StatusInternalServerError
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	linkable, isLinkable := decorator.(jsonapi.Linkable)
	metable, isMetable := decorator.(jsonapi.Metable)
	if manyPayload, ok := payload.(*jsonapi.ManyPayload); ok {
		if isLinkable {
			manyPayload.Links = linkable.JSONAPILinks()
		}
		if isMetable {
			manyPayload.Meta = metable.JSONAPIMeta()
		}
		payload = manyPayload
	}
	if onePayload, ok := payload.(*jsonapi.OnePayload); ok {
		if isLinkable {
			onePayload.Links = linkable.JSONAPILinks()
		}
		if isMetable {
			onePayload.Meta = metable.JSONAPIMeta()
		}
		payload = onePayload
	}

	if err := json.NewEncoder(w.ResponseWriter).Encode(payload); err != nil {
		w.statusCode = http.StatusInternalServerError
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
	} else {
		w.statusCode = http.StatusOK
	}
}

func (w *ResponseWriter) Status() int {
	return w.statusCode
}

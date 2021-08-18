package json

import (
	"bytes"
	"encoding/json"
	"github.com/BSick7/go-api/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Tests []Test

func (s Tests) Run(t *testing.T, server http.Handler) {
	for _, test := range s {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			server.ServeHTTP(w, test.Request)
			test.Want.Verify(t, w)
		})
	}
}

type Test struct {
	Name    string
	Request *http.Request
	Want    TestExpectation
}

type TestExpectation interface {
	Verify(t *testing.T, recorder *httptest.ResponseRecorder)
}

type ExpectResponse struct {
	Code    int
	Payload interface{}
	Headers map[string]string
}

func (e ExpectResponse) WithHeader(key, value string) ExpectResponse {
	if e.Headers == nil {
		e.Headers = make(map[string]string)
	}
	e.Headers[key] = value
	return e
}

func (e ExpectResponse) Verify(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, e.Code, w.Code, "mismatched status code")
	if e.Payload == nil {
		// We expect nil, let's confirm there is no body
		if assert.NotNil(t, w.Body) {
			assert.LessOrEqualf(t, w.Body.Len(), 0, "got response body, want empty body")
		}
	} else {
		rawWant := bytes.NewBufferString("")
		encoder := json.NewEncoder(rawWant)
		if payloader, ok := e.Payload.(ResponsePayloader); ok {
			require.NoError(t, encoder.Encode(payloader.Payload()))
		} else {
			require.NoError(t, encoder.Encode(e.Payload))
		}
		rawGot, err := ioutil.ReadAll(w.Body)
		require.NoError(t, err)
		assert.Equal(t, rawWant.String(), string(rawGot), "mismatched response")
	}
	if e.Headers != nil {
		for k, want := range e.Headers {
			got := w.Header().Get(k)
			assert.Equal(t, want, got, "unexpected value in header")
		}
	}
}

func ExpectOK(payload interface{}) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusOK,
		Payload: payload,
	}
}

func ExpectBadRequest(details map[string]string) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusBadRequest,
		Payload: errors.BadRequestError{Details: details},
	}
}

func ExpectInvalidRequest(validationErrors errors.ValidationErrors) ExpectResponse {
	return ExpectResponse{
		Code: http.StatusUnprocessableEntity,
		Payload: errors.InvalidRequestError{
			Errors: validationErrors,
		},
	}
}

func ExpectForbidden() ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusForbidden,
		Payload: errors.AuthorizationError{},
	}
}

func ExpectInternalError(err error) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusInternalServerError,
		Payload: errors.ApiError{Err: err},
	}
}

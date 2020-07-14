package json

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	if e.Code != w.Code {
		t.Errorf("mismatched status code, got %d, want %d", w.Code, e.Code)
	}
	if e.Payload == nil {
		// We expect nil, let's confirm there is no body
		if w.Body != nil || w.Body.Len() > 0 {
			t.Errorf("got response body, want empty payload")
		}
	} else {
		rawWant := bytes.NewBufferString("")
		encoder := json.NewEncoder(rawWant)
		if err := encoder.Encode(e.Payload); err != nil {
			t.Fatalf("unexpected error marshaling expected response body: %s", err)
		}
		if rawGot, err := ioutil.ReadAll(w.Body); err != nil {
			t.Fatalf("unexpected error attempting to read response body: %s", err)
		} else if diff := cmp.Diff(rawWant.String(), string(rawGot)); diff != "" {
			t.Errorf("unexpected response body, (-want, +got):\n%s", diff)
		}
	}
	if e.Headers != nil {
		for k, want := range e.Headers {
			if got := w.Header().Get(k); got != want {
				t.Errorf("unexpected value for header %q, wanted %q got %q", k, want, got)
			}
		}
	}
}

func ExpectOK(payload interface{}) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusOK,
		Payload: payload,
	}
}

func ExpectBadRequest(message string) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusBadRequest,
		Payload: Error{ErrorMessage: message},
	}
}

func ExpectInternalError(message string) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusInternalServerError,
		Payload: Error{ErrorMessage: message},
	}
}

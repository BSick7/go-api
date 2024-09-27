package json

import (
	"encoding/json"
	"github.com/BSick7/go-api/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	Verify(t *testing.T, recorder *httptest.ResponseRecorder, opts ...cmp.Option)
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

func (e ExpectResponse) Verify(t *testing.T, w *httptest.ResponseRecorder, opts ...cmp.Option) {
	assert.Equal(t, e.Code, w.Code, "mismatched status code")
	if e.Payload == nil {
		// We expect nil, let's confirm there is no response body
		if assert.NotNil(t, w.Body) {
			assert.LessOrEqualf(t, w.Body.Len(), 0, "got response body, want empty body")
		}
	} else {
		var want string
		if payloader, ok := e.Payload.(ResponsePayloader); ok {
			raw, err := json.Marshal(payloader.Payload())
			require.NoError(t, err, "error serializing want: %w", err)
			want = string(raw)
		} else {
			raw, err := json.Marshal(e.Payload)
			require.NoError(t, err, "error serializing want: %w", err)
			want = string(raw)
		}

		var temp any
		decoder := json.NewDecoder(w.Body)
		require.NoError(t, decoder.Decode(&temp), "error reading response body")
		raw, err := json.Marshal(temp)
		require.NoError(t, err, "error normalizing response body")
		got := string(raw)
		if diff := cmp.Diff(want, got, opts...); diff != "" {
			t.Errorf("mismatched config (-want +got):\n%s", diff)
		}
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

func ExpectBadRequest(errorCode int, details map[string]string) ExpectResponse {
	return ExpectResponse{
		Code:    http.StatusBadRequest,
		Payload: errors.BadRequestError{ApiError: errors.ApiError{ErrorCode: errorCode}, Details: details},
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

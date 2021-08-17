package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	request *http.Request
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[%s] http error (%d)", e.RequestId(), e.StatusCode())
}

func (e ApiError) StatusCode() int {
	// By default, API Error returns 500
	return http.StatusInternalServerError
}

func (e ApiError) RequestId() string {
	if e.request != nil {
		return e.request.Header.Get("X-Request-Id")
	}
	return ""
}

var _ json.Marshaler = ApiError{}

func (e ApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"request_id": e.RequestId(),
		"title":      "General Error",
		"type":       "problems/general-error",
		"code":       e.StatusCode(),
		"message":    "We have encountered an unexpected error.",
	})
}

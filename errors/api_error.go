package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Err error
}

func NewApiError(err error) ApiError {
	return ApiError{Err: err}
}

func (e ApiError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("http error (%d)", e.StatusCode())
	}
	return e.Err.Error()
}

func (e ApiError) StatusCode() int {
	// By default, API Error returns 500
	return http.StatusInternalServerError
}

func (e ApiError) Payload() map[string]interface{} {
	message := "We have encountered an unexpected error."
	if e.Err != nil {
		message = e.Err.Error()
	}
	return map[string]interface{}{
		"title":   "General Error",
		"type":    "problems/general-error",
		"code":    e.StatusCode(),
		"message": message,
	}
}

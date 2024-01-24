package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	// ErrorCode is different from the HTTP Status Code as it represents a Nullstone-specific error code
	ErrorCode int

	Err error
}

func NewApiError(errorCode int, err error) ApiError {
	return ApiError{Err: err, ErrorCode: errorCode}
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

func (e ApiError) Payload() map[string]any {
	message := "We have encountered an unexpected error."
	if e.Err != nil {
		message = e.Err.Error()
	}
	return map[string]any{
		"title":      "General Error",
		"type":       "problems/general-error",
		"code":       e.StatusCode(),
		"message":    message,
		"error_code": e.ErrorCode,
	}
}

package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct{}

func (e ApiError) Error() string {
	return fmt.Sprintf("http error (%d)", e.StatusCode())
}

func (e ApiError) StatusCode() int {
	// By default, API Error returns 500
	return http.StatusInternalServerError
}

func (e ApiError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":   "General Error",
		"type":    "problems/general-error",
		"code":    e.StatusCode(),
		"message": "We have encountered an unexpected error.",
	}
}

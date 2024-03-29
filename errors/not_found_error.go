package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ApiError
}

func NewNotFoundError(msg string) NotFoundError {
	if msg == "" {
		msg = "We could not find this resource."
	}
	return NotFoundError{
		ApiError: ApiError{
			Err: fmt.Errorf(msg),
		},
	}
}

func (e NotFoundError) Error() string {
	if e.ApiError.Err == nil {
		return "We could not find this resource."
	}
	return e.ApiError.Err.Error()
}

func (e NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

func (e NotFoundError) Payload() map[string]any {
	return map[string]any{
		"title":      "Resource Not Found",
		"type":       "problems/not-found-error",
		"code":       e.StatusCode(),
		"message":    e.Error(),
		"error_code": e.ErrorCode,
	}
}

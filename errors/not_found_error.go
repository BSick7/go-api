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
	return e.ApiError.Error()
}

func (e NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

func (e NotFoundError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":   "Resource Not Found",
		"type":    "problems/not-found-error",
		"code":    e.StatusCode(),
		"message": e.Error(),
	}
}

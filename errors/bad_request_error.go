package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

type BadRequestError struct {
	ApiError

	// Details provides specific messages and information to detail the application error
	Details map[string]string
}

func NewBadRequestError(errorCode int, details map[string]string) BadRequestError {
	return BadRequestError{
		ApiError: ApiError{ErrorCode: errorCode},
		Details:  details,
	}
}

func (e BadRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprint(buf, "bad request:")
	for key, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s = %s", key, value)
	}
	return buf.String()
}

func (e BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

func (e BadRequestError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":      "Bad Request",
		"type":       "problems/bad-request",
		"code":       e.StatusCode(),
		"message":    "Your request could not be processed.",
		"details":    e.Details,
		"error_code": e.ErrorCode,
	}
}

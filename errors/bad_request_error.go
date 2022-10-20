package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

type BadRequestError struct {
	ApiError
	Details []string
}

func (e BadRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprint(buf, "bad request:")
	for _, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s", value)
	}
	return buf.String()
}

func (e BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

func (e BadRequestError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":   "Bad Request",
		"type":    "problems/bad-request",
		"code":    e.StatusCode(),
		"message": "Your request could not be processed.",
		"details": e.Details,
	}
}

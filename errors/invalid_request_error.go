package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

type InvalidRequestError struct {
	ApiError
	Errors ValidationErrors
}

func (e InvalidRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprint(buf, "invalid request:")
	for field, errs := range e.Errors.ToJson() {
		fmt.Fprintf(buf, "\n  %s: %s", field, errs)
	}
	return buf.String()
}

func (e InvalidRequestError) StatusCode() int {
	return http.StatusUnprocessableEntity
}

func (e InvalidRequestError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":             "Invalid Request",
		"type":              "problems/invalid-request",
		"code":              e.StatusCode(),
		"message":           "Your request is invalid and could not be processed.",
		"validation_errors": e.Errors.ToJson(),
	}
}

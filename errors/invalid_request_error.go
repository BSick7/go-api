package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvalidRequestError struct {
	ApiError
	Errors ValidationErrors
}

func (e InvalidRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s] invalid request:", e.RequestId())
	for field, errs := range e.Errors {
		for _, err := range errs {
			fmt.Fprintf(buf, "\n  %s: %s", field, err)
		}
	}
	return buf.String()
}

func (e InvalidRequestError) StatusCode() int {
	return http.StatusUnprocessableEntity
}

var _ json.Marshaler = InvalidRequestError{}

func (e InvalidRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"request_id":        e.RequestId(),
		"title":             "Invalid Request",
		"type":              "problems/invalid-request",
		"code":              e.StatusCode(),
		"message":           "Your request is invalid and could not be processed.",
		"validation_errors": e.Errors,
	})
}

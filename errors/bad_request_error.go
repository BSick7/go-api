package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BadRequestError struct {
	ApiError
	Details map[string]string
}

func (e BadRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s] bad request:", e.RequestId())
	for key, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s: %s", key, value)
	}
	return buf.String()
}

func (e BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

var _ json.Marshaler = BadRequestError{}

func (e BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"request_id": e.RequestId(),
		"title":      "Bad Request",
		"type":       "problems/bad-request",
		"code":       e.StatusCode(),
		"message":    "Your request could not be processed.",
		"details":    e.Details,
	})
}

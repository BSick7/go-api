package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ApiError
}

func (e NotFoundError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s] not found", e.RequestId())
	return buf.String()
}

func (e NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

var _ json.Marshaler = NotFoundError{}

func (e NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"request_id": e.RequestId(),
		"title":      "Resource Not Found",
		"type":       "problems/not-found-error",
		"code":       e.StatusCode(),
		"message":    "We could not find this resource.",
	})
}

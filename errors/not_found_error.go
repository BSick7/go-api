package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ApiError
}

func (e NotFoundError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "not found")
	return buf.String()
}

func (e NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

func (e NotFoundError) Payload() map[string]interface{} {
	return map[string]interface{}{
		"title":   "Resource Not Found",
		"type":    "problems/not-found-error",
		"code":    e.StatusCode(),
		"message": "We could not find this resource.",
	}
}

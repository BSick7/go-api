package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthorizationError struct {
	ApiError
}

func (e AuthorizationError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s] not found", e.RequestId())
	return buf.String()
}

func (e AuthorizationError) StatusCode() int {
	return http.StatusForbidden
}

var _ json.Marshaler = AuthorizationError{}

func (e AuthorizationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"request_id": e.RequestId(),
		"title":      "Access Denied",
		"type":       "problems/authorization-error",
		"code":       e.StatusCode(),
		"message":    "You do not have the proper authorization to access this resource.",
	})
}

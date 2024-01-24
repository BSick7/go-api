package errors

import (
	"bytes"
	"fmt"
	"net/http"
)

type AuthorizationError struct {
	ApiError
}

func (e AuthorizationError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprint(buf, "forbidden")
	return buf.String()
}

func (e AuthorizationError) StatusCode() int {
	return http.StatusForbidden
}

func (e AuthorizationError) Payload() map[string]any {
	return map[string]any{
		"title":      "Access Denied",
		"type":       "problems/authorization-error",
		"code":       e.StatusCode(),
		"message":    "You do not have the proper authorization to access this resource.",
		"error_code": e.ErrorCode,
	}
}

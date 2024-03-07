package errors

import (
	"fmt"
)

const (
	ErrorCodeInvalidPathParameter  = 10000
	ErrorCodeInvalidQueryParameter = 10001
	ErrorCodeInvalidPayload        = 10002
)

func InvalidPathParameter(part string, message string) BadRequestError {
	details := map[string]string{
		"parameter": part,
		"message":   message,
	}
	return NewBadRequestError(ErrorCodeInvalidPathParameter, details)
}

func InvalidQueryParameter(field string, message string) BadRequestError {
	details := map[string]string{
		"parameter": field,
		"message":   message,
	}
	return NewBadRequestError(ErrorCodeInvalidQueryParameter, details)
}

func InvalidPayload(err error) BadRequestError {
	details := map[string]string{
		"message": fmt.Sprintf("invalid payload: %s", err),
	}
	return NewBadRequestError(ErrorCodeInvalidPayload, details)
}

func InvalidPayloadAttribute(attribute string, message string) BadRequestError {
	details := map[string]string{
		"attribute": attribute,
		"message":   message,
	}
	return NewBadRequestError(ErrorCodeInvalidPayload, details)
}

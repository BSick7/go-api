package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DecodeStructBody[T any](r *http.Request) (*T, error) {
	var t T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err == io.EOF {
		return nil, fmt.Errorf("missing body content")
	}
	return &t, err
}

func DecodeSliceBody[T any](r *http.Request) ([]T, error) {
	var t []T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err == io.EOF {
		return nil, fmt.Errorf("missing body content")
	}
	return t, err
}

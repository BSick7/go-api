package jsonapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svanharmelen/jsonapi"
)

type Request struct {
	*http.Request
}

func (r *Request) Var(key string) string {
	if v := mux.Vars(r.Request); v != nil {
		return v[key]
	}
	return ""
}

func (r *Request) DecodeBody(v interface{}) error {
	err := jsonapi.UnmarshalPayload(r.Request.Body, v)
	if err == io.EOF {
		return fmt.Errorf("missing body content")
	}
	return err
}

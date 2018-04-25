package json

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRequest(req *http.Request) *Request {
	return &Request{
		req:  req,
		vars: mux.Vars(req),
	}
}

type Request struct {
	req  *http.Request
	vars map[string]string
}

func (r *Request) Var(key string) string {
	return r.vars[key]
}

func (r *Request) DecodeBody(v interface{}) error {
	decoder := json.NewDecoder(r.req.Body)
	err := decoder.Decode(v)
	if err == io.EOF {
		return fmt.Errorf("missing body content")
	}
	return err
}

func (r *Request) Context() context.Context {
	return r.req.Context()
}

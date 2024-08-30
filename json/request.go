package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	*http.Request
	Logger *slog.Logger
}

func (r *Request) Var(key string) string {
	if v := mux.Vars(r.Request); v != nil {
		return v[key]
	}
	return ""
}

func (r *Request) DecodeBody(v interface{}) error {
	decoder := json.NewDecoder(r.Request.Body)
	err := decoder.Decode(v)
	if err == io.EOF {
		return fmt.Errorf("missing body content")
	}
	return err
}

package jsonapi

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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

func (r *Request) Pagination() (*Pagination, error) {
	p := &Pagination{}

	var err error
	if p.Limit, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageLimit); err != nil {
		return nil, err
	}
	if p.Offset, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageOffset); err != nil {
		return nil, err
	}
	if p.Size, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageSize); err != nil {
		return nil, err
	}
	if p.Number, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageNumber); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Request) parseQueryParamPositiveInt(param string) (*int, error) {
	num, err := r.parseQueryParamInt(param)
	if err != nil {
		return nil, err
	}
	if num != nil && *num < 0 {
		return nil, &jsonapi.ErrorObject{
			Title:  "Invalid Parameter.",
			Detail: fmt.Sprintf("%s must be a positive integer", param),
			Status: http.StatusText(http.StatusBadRequest),
			Code:   strconv.Itoa(http.StatusBadRequest),
		}
	}
	return num, nil
}

func (r *Request) parseQueryParamInt(param string) (*int, error) {
	val := r.URL.Query().Get(param)
	if val == "" {
		return nil, nil
	}
	limit, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

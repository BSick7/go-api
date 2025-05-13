package jsonapi

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/jsonapi"
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

func (r *Request) Pagination() (*Pagination, []*jsonapi.ErrorObject) {
	p := &Pagination{}

	errs := make([]*jsonapi.ErrorObject, 0)
	var err *jsonapi.ErrorObject
	if p.Limit, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageLimit); err != nil {
		errs = append(errs, err)
	}
	if p.Offset, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageOffset); err != nil {
		errs = append(errs, err)
	}
	if p.Size, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageSize); err != nil {
		errs = append(errs, err)
	}
	if p.Number, err = r.parseQueryParamPositiveInt(jsonapi.QueryParamPageNumber); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return nil, errs
	}
	return p, nil
}

func (r *Request) parseQueryParamPositiveInt(param string) (*int, *jsonapi.ErrorObject) {
	val := r.URL.Query().Get(param)
	if val == "" {
		return nil, nil
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return nil, &jsonapi.ErrorObject{
			Title:  "Invalid Parameter.",
			Detail: fmt.Sprintf("%s must be a positive integer", param),
			Status: http.StatusText(http.StatusBadRequest),
			Code:   strconv.Itoa(http.StatusBadRequest),
		}
	}
	if num < 0 {
		return nil, &jsonapi.ErrorObject{
			Title:  "Invalid Parameter.",
			Detail: fmt.Sprintf("%s must be a positive integer", param),
			Status: http.StatusText(http.StatusBadRequest),
			Code:   strconv.Itoa(http.StatusBadRequest),
		}
	}
	return &num, nil
}

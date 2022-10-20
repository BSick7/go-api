package app1

import (
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/json"
	"net/url"
)

func UrlEncodedPath(res *json.ResponseWriter, req *json.Request) {
	path := req.Var("path")
	if path == "" {
		res.SendError(errors.BadRequestError{Details: []string{
			"missing path",
		}})
		return
	}
	escaped, err := url.PathUnescape(path)
	if err != nil {
		res.SendError(errors.BadRequestError{Details: []string{
			"invalid path",
		}})
		return
	}
	res.Send(escaped)
}

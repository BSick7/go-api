package app1

import (
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/json"
	"net/url"
)

func UrlEncodedPath(res *json.ResponseWriter, req *json.Request) {
	path := req.Var("path")
	if path == "" {
		res.SendError(errors.BadRequestError{Details: map[string]string{
			"path": "missing path",
		}})
		return
	}
	escaped, err := url.PathUnescape(path)
	if err != nil {
		res.SendError(errors.BadRequestError{Details: map[string]string{
			"path": "invalid path",
		}})
		return
	}
	res.Send(escaped)
}

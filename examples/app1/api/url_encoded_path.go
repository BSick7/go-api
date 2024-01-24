package app1

import (
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/json"
	"net/url"
)

func UrlEncodedPath(res *json.ResponseWriter, req *json.Request) {
	path := req.Var("path")
	if path == "" {
		res.SendError(errors.NewBadRequestError(1, map[string]string{"message": "missing path"}))
		return
	}
	escaped, err := url.PathUnescape(path)
	if err != nil {
		res.SendError(errors.NewBadRequestError(2, map[string]string{"message": "invalid path"}))
		return
	}
	res.Send(escaped)
}

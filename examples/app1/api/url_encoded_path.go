package app1

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/BSick7/go-api/json"
	"github.com/pkg/errors"
)

func UrlEncodedPath(res *json.ResponseWriter, req *json.Request) {
	path := req.Var("path")
	if path == "" {
		res.SendRawError(http.StatusBadRequest, fmt.Errorf("missing path"))
		return
	}
	escaped, err := url.PathUnescape(path)
	if err != nil {
		res.SendRawError(http.StatusBadRequest, errors.Wrap(err, "invalid 'path'"))
		return
	}
	res.Send(escaped)
}

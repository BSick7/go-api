package app1

import (
	"github.com/BSick7/go-api/errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func UrlEncodedPath(_ http.ResponseWriter, req *http.Request) (string, error) {
	path := mux.Vars(req)["path"]
	if path == "" {
		return "", errors.NewBadRequestError("missing path")
	}
	escaped, err := url.PathUnescape(path)
	if err != nil {
		return "", errors.NewBadRequestError("invalid path")
	}
	return escaped, nil
}

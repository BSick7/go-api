package request

import (
	"fmt"
	"github.com/BSick7/go-api/errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Int64PathParameter(r *http.Request, name string) (int64, error) {
	vars := mux.Vars(r)
	val, err := strconv.ParseInt(vars[name], 10, 64)
	if err != nil {
		log.Printf("%s is not a valid int64: %s\n", name, err)
		return 0, errors.InvalidPathParameter(name, fmt.Sprintf("%s must be an integer", name))
	}
	return val, nil
}

func Int64QueryParam(r *http.Request, name string) (int64, error) {
	q := r.URL.Query()
	val, err := strconv.ParseInt(q.Get(name), 10, 64)
	if err != nil {
		log.Printf("%s is not a valid int64: %s\n", name, err)
		return 0, errors.InvalidQueryParameter(name, fmt.Sprintf("%s must be an integer", name))
	}
	return val, nil
}

func StringSliceQueryParam(r *http.Request, name string) []string {
	q := r.URL.Query()
	if val := q.Get(name); val != "" {
		return strings.Split(val, ",")
	}
	return nil
}

func OptionalTimeQueryParam(r *http.Request, name string) (*time.Time, error) {
	q := r.URL.Query()
	if val := q.Get(name); val != "" {
		if rawEnd, err := time.Parse(time.RFC3339, val); err != nil {
			return nil, errors.InvalidQueryParameter(name, err.Error())
		} else {
			return &rawEnd, nil
		}
	}
	return nil, nil
}

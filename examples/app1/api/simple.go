package app1

import (
	"fmt"
	"github.com/BSick7/go-api/errors"
	"net/http"
	"strconv"
)

func Simple(res http.ResponseWriter, req *http.Request) (int, error) {
	data := req.URL.Query().Get("data")
	if data == "" {
		return 0, errors.NewBadRequestError("missing data")
	}

	if i, err := strconv.Atoi(data); err != nil {
		return 0, fmt.Errorf("invalid syntax")
	} else {
		res.Header().Set("Have", "data")
		return i, nil
	}
}

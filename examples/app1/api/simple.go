package app1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BSick7/go-api/json"
)

func Simple(res *json.ResponseWriter, req *json.Request) {
	data := req.Request.URL.Query().Get("data")
	if data == "" {
		res.SendError(http.StatusBadRequest, fmt.Errorf("missing data"))
		return
	}

	if i, err := strconv.Atoi(data); err != nil {
		res.SendError(http.StatusInternalServerError, errors.New("invalid syntax"))
	} else {
		res.Header().Set("Have", "data")
		res.Send(i)
	}
}

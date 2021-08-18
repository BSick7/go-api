package app1

import (
	"fmt"
	"github.com/BSick7/go-api/errors"
	"github.com/BSick7/go-api/json"
	"strconv"
)

func Simple(res *json.ResponseWriter, req *json.Request) {
	data := req.Request.URL.Query().Get("data")
	if data == "" {
		res.SendError(errors.BadRequestError{Details: map[string]string{"data": "missing data"}})
		return
	}

	if i, err := strconv.Atoi(data); err != nil {
		res.SendError(fmt.Errorf("invalid syntax"))
	} else {
		res.Header().Set("Have", "data")
		res.Send(i)
	}
}

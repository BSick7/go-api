package examples

import (
	"fmt"
	"github.com/BSick7/go-api"
	"net/http"
	"strconv"
)

func Simple(res api.Responder, req api.Request) {
	data := req.Var("data")
	if data == "" {
		res.SendError(http.StatusBadRequest, fmt.Errorf("missing data"))
		return
	}

	if i, err := strconv.Atoi(data); err != nil {
		res.SendError(http.StatusInternalServerError, err)
	} else {
		res.Send(i)
	}
}

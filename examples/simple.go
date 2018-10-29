package examples

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BishopFox/go-api"
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

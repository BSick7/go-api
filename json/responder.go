package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Responder struct {
	encoder    *json.Encoder
	w          http.ResponseWriter
	start      time.Time
	statusCode int
	statusCtx  []string
}

func NewResponder(e Endpoint, w http.ResponseWriter, r *http.Request) *Responder {
	encoder := json.NewEncoder(w)
	prettyJson(r, encoder)
	return &Responder{
		encoder: encoder,
		w:       w,
		start:   time.Now(),
	}
}

func (r *Responder) SendError(statusCode int, err error, context ...string) {
	r.statusCode = statusCode
	r.statusCtx = append(context, err.Error())
	r.w.WriteHeader(statusCode)
	r.encoder.Encode(map[string]interface{}{"error": err.Error()})
}

func (r *Responder) SendNotFound(msg string, context ...string) {
	r.SendError(http.StatusNotFound, fmt.Errorf(msg), context...)
}

func (r *Responder) Send(data interface{}, context ...string) {
	if data == nil {
		r.statusCode = http.StatusNoContent
		r.statusCtx = context
		r.w.WriteHeader(http.StatusNoContent)
	} else {
		r.statusCode = http.StatusOK
		r.statusCtx = context
		r.encoder.Encode(data)
	}
}

func (r *Responder) Status() (code int, ctx string) {
	if len(r.statusCtx) > 0 {
		ctx = fmt.Sprintf(" %+v", r.statusCtx)
	}
	return r.statusCode, ctx
}

func prettyJson(r *http.Request, encoder *json.Encoder) {
	indent := r.URL.Query().Get("indent")
	if indent == "" {
		return
	}
	i, err := strconv.Atoi(indent)
	if err != nil {
		i = 2
	}
	encoder.SetIndent("", strings.Repeat(" ", i))
}

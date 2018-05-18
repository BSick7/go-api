package json

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Responder struct {
	endpoint Endpoint
	encoder  *json.Encoder
	w        http.ResponseWriter
	req      *http.Request
	start    time.Time
}

func NewResponder(e Endpoint, w http.ResponseWriter, r *http.Request) *Responder {
	encoder := json.NewEncoder(w)
	prettyJson(r, encoder)
	return &Responder{
		endpoint: e,
		encoder:  encoder,
		w:        w,
		req:      r,
		start:    time.Now(),
	}
}

func (r *Responder) SendError(statusCode int, err error, context ...string) {
	r.logWithContext(statusCode, append(context, err.Error())...)
	r.w.WriteHeader(statusCode)
	r.encoder.Encode(map[string]interface{}{"error": err.Error()})
}

func (r *Responder) SendNotFound(msg string, context ...string) {
	r.SendError(http.StatusNotFound, fmt.Errorf(msg), context...)
}

func (r *Responder) Send(data interface{}, context ...string) {
	if data == nil {
		r.logWithContext(http.StatusNoContent, context...)
		r.w.WriteHeader(http.StatusNoContent)
	} else {
		r.logWithContext(http.StatusOK, context...)
		r.encoder.Encode(data)
	}
}

func (r *Responder) logWithContext(code int, context ...string) {
	ctx := ""
	if len(context) > 0 {
		ctx = fmt.Sprintf(" %+v", context)
	}
	log.Printf("%s %d %s %s%s", time.Since(r.start), code, r.req.RequestURI, r.endpoint, ctx)
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

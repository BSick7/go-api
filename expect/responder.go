package expect

import (
	"reflect"
	"strings"
	"testing"
)

func NewResponder(t *testing.T, prefix string, want Response) *Responder {
	return &Responder{
		t:           t,
		prefix:      prefix,
		hitNotFound: false,
		want:        want,
	}
}

type Responder struct {
	t           *testing.T
	prefix      string
	hitNotFound bool
	want        Response
}

func (r *Responder) Send(data interface{}, context ...string) {
	if r.want.Data == nil {
		return
	}
	if !reflect.DeepEqual(data, r.want.Data) {
		r.t.Errorf("%smismatched data, got %+v, want %+v", r.prefix, data, r.want.Data)
	}
}

func (r *Responder) SendNotFound(msg string, context ...string) {
	r.hitNotFound = true
	if r.want.StatusCode != 404 {
		r.t.Errorf("%sunexpected not found", r.prefix)
	}
}

func (r *Responder) SendError(statusCode int, err error, context ...string) {
	if r.want.Err == "" && err.Error() == "" {
		return
	}
	if r.want.Err == "" && err.Error() != "" {
		r.t.Errorf("%sgot error %s, expected no message", r.prefix, err.Error())
	}
	if statusCode != r.want.StatusCode {
		r.t.Errorf("%smismatched status code, got %d, want %d", r.prefix, statusCode, r.want.StatusCode)
	}
	if !strings.Contains(err.Error(), r.want.Err) {
		r.t.Errorf("%smismatched error message, got %s, want %s", r.prefix, err.Error(), r.want.Err)
	}
}

func (r *Responder) VerifyNotFound() {
	if r.want.StatusCode == 404 && !r.hitNotFound {
		r.t.Errorf("%sexpected not found", r.prefix)
	}
	r.hitNotFound = false
}

func (r *Responder) Status() (code int, ctx string) {
	return r.want.StatusCode, ""
}

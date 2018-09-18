package expect

import (
	"context"
	"fmt"
)

func NewRequest(data RequestData, ctx context.Context) *Request {
	return &Request{
		Data: data,
		Ctx:  ctx,
	}
}

type Request struct {
	Data        RequestData
	CopyPointer func(a, b interface{})
	Ctx         context.Context
}

func (r *Request) Var(key string) string {
	return r.Data.Vars[key]
}

func (r *Request) Query(key string) string {
	return r.Data.Query[key]
}

func (r *Request) DecodeBody(v interface{}) error {
	r.CopyPointer(v, r.Data.Body)
	if r.Data.DecodeErr != "" {
		return fmt.Errorf(r.Data.DecodeErr)
	}
	return nil
}

func (r *Request) Context() context.Context {
	return r.Ctx
}

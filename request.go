package api

import "context"

type Request interface {
	Var(key string) string
	DecodeBody(v interface{}) error
	Context() context.Context
}

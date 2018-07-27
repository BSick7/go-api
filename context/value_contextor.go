package context

import (
	"context"
)

type ValueContextor struct {
	context.Context
	Key interface{}
}

func (c ValueContextor) With(parentCtx context.Context, value interface{}) context.Context {
	return context.WithValue(parentCtx, c.Key, value)
}

func (c ValueContextor) Wrapper(value interface{}) Wrapper {
	return func(parentCtx context.Context) context.Context {
		return context.WithValue(parentCtx, c.Key, value)
	}
}

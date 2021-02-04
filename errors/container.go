package errors

import (
	"context"
)

type Container struct {
	Errors []error
}

func (c *Container) AddError(err error) {
	if c.Errors == nil {
		c.Errors = []error{err}
	} else {
		c.Errors = append(c.Errors, err)
	}
}

type containerContextKey struct{}

func ContextWithContainer(ctx context.Context, container *Container) context.Context {
	return context.WithValue(ctx, containerContextKey{}, container)
}

func ContainerFromContext(ctx context.Context) *Container {
	if val, ok := ctx.Value(containerContextKey{}).(*Container); ok {
		return val
	}
	return nil
}

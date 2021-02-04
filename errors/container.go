package errors

import (
	"context"
)

type Container struct {
	Error error
}

func (c *Container) AddError(err error) {

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
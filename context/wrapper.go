package context

import (
	"context"
)

type Wrapper func(ctx context.Context) context.Context

func WrapMany(wrappers ...Wrapper) Wrapper {
	return func(parentCtx context.Context) context.Context {
		newCtx := parentCtx
		for _, wrapper := range wrappers {
			newCtx = wrapper(newCtx)
		}
		return newCtx
	}
}

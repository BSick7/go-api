package context

import "context"

type Values map[interface{}]interface{}

func WithValues(parentCtx context.Context, values Values) context.Context {
	return &valuesContext{parentCtx, values}
}

type valuesContext struct {
	context.Context
	values Values
}

func (c *valuesContext) Value(key interface{}) interface{} {
	if val, ok := c.values[key]; ok {
		return val
	}
	return c.Context.Value(key)
}

package util

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"time"
)

var SystemContext = WithSystemContext(context.TODO())

func WithSystemContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, abs.SystemContextKey, true)
}

func IsSystemContext(ctx context.Context) bool {
	return ctx.Value(abs.SystemContextKey) != nil && ctx.Value(abs.SystemContextKey).(bool)
}

type contextWithValues struct {
	parent     context.Context
	valuesFrom context.Context
}

func (c contextWithValues) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c contextWithValues) Done() <-chan struct{} {
	return c.parent.Done()
}

func (c contextWithValues) Err() error {
	return c.parent.Err()
}

func (c contextWithValues) Value(key any) any {
	return c.valuesFrom.Value(key)
}

func NewContextWithValues(parent context.Context, valuesFrom context.Context) context.Context {
	return &contextWithValues{
		parent:     parent,
		valuesFrom: valuesFrom,
	}
}

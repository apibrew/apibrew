package util

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
)

var SystemContext = WithSystemContext(context.TODO())

func WithSystemContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, abs.SystemContextKey, true)
}

func IsSystemContext(ctx context.Context) bool {
	return ctx.Value(abs.SystemContextKey) != nil && ctx.Value(abs.SystemContextKey).(bool)
}

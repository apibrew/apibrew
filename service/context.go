package service

import (
	"context"
	"data-handler/stub/model"
	"errors"
)

var systemResourceAccessError = errors.New("system resource is accessed outside of system context")

const systemContextKey = "SYSTEM_CTX"

func withSystemContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, systemContextKey, true)
}

func isSystemContext(ctx context.Context) bool {
	return ctx.Value(systemContextKey) != nil && ctx.Value(systemContextKey).(bool)
}

func checkSystemResourceAccess(ctx context.Context, resource *model.Resource) error {
	if resource.Type == model.DataType_SYSTEM {
		if !isSystemContext(ctx) {
			return systemResourceAccessError
		}
	}

	return nil
}

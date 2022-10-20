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

type HasDataType interface {
	GetType() model.DataType
}

func checkSystemResourceAccess(ctx context.Context, objs ...HasDataType) error {
	for _, obj := range objs {
		if obj.GetType() == model.DataType_SYSTEM {
			if !isSystemContext(ctx) {
				return systemResourceAccessError
			}
		}
	}

	return nil
}

package logging

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	log "github.com/sirupsen/logrus"
)

func WithLogFields(ctx context.Context, fields log.Fields) context.Context {
	currentFields := CtxFields(ctx)

	data := make(log.Fields, len(currentFields)+len(fields))
	for k, v := range currentFields {
		data[k] = v
	}
	for k, v := range fields {
		data[k] = v
	}

	return context.WithValue(ctx, abs.LogFieldsContextKey, data)
}

func WithLogField(ctx context.Context, key string, value interface{}) context.Context {
	return WithLogFields(ctx, map[string]interface{}{
		key: value,
	})
}

func CtxFields(ctx context.Context) log.Fields {
	if ctx.Value(abs.LogFieldsContextKey) != nil {
		return ctx.Value(abs.LogFieldsContextKey).(log.Fields)
	} else {
		return make(map[string]interface{})
	}
}

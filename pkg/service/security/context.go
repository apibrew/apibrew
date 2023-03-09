package security

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"
)

var SystemContext = WithSystemContext(context.TODO())

func WithSystemContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, abs.SystemContextKey, true)
}

func WithUserDetails(ctx context.Context, userDetails abs.UserDetails) context.Context {
	return context.WithValue(ctx, abs.UserContextKey, userDetails)
}

func GetUserDetailsFromContext(ctx context.Context) *abs.UserDetails {
	if ctx.Value(abs.UserContextKey) == nil {
		return nil
	}

	var res = new(abs.UserDetails)

	*res = ctx.Value(abs.UserContextKey).(abs.UserDetails)

	return res
}

func IsSystemContext(ctx context.Context) bool {
	return ctx.Value(abs.SystemContextKey) != nil && ctx.Value(abs.SystemContextKey).(bool)
}

type HasDataType interface {
	GetDataType() model.DataType
}

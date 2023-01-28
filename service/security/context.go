package security

import (
	"context"
	"github.com/tislib/data-handler/model"
)

const systemContextKey = "SYSTEM_CTX"
const userContextKey = "USER"

var SystemContext = WithSystemContext(context.TODO())

func WithSystemContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, systemContextKey, true)
}

func WithUserDetails(ctx context.Context, userDetails UserDetails) context.Context {
	return context.WithValue(ctx, userContextKey, userDetails)
}

func GetUserDetailsFromContext(ctx context.Context) *UserDetails {
	if ctx.Value(userContextKey) == nil {
		return nil
	}

	var res = new(UserDetails)

	*res = ctx.Value(userContextKey).(UserDetails)

	return res
}

func IsSystemContext(ctx context.Context) bool {
	return ctx.Value(systemContextKey) != nil && ctx.Value(systemContextKey).(bool)
}

type HasDataType interface {
	GetDataType() model.DataType
}

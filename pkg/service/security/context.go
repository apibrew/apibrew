package security

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
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

func GetUserPrincipalFromContext(ctx context.Context) string {
	userDetails := GetUserDetailsFromContext(ctx)

	if userDetails == nil {
		if IsSystemContext(ctx) {
			return "system"
		} else {
			return "guest"
		}
	}

	return userDetails.Username
}

func IsSystemContext(ctx context.Context) bool {
	return ctx.Value(abs.SystemContextKey) != nil && ctx.Value(abs.SystemContextKey).(bool)
}

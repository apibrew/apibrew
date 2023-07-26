package jwt_model

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/util"
)

func WithUserDetails(ctx context.Context, userDetails UserDetails) context.Context {
	return context.WithValue(ctx, abs.UserContextKey, userDetails)
}

func GetUserDetailsFromContext(ctx context.Context) *UserDetails {
	if ctx.Value(abs.UserContextKey) == nil {
		return nil
	}

	var res = new(UserDetails)

	*res = ctx.Value(abs.UserContextKey).(UserDetails)

	return res
}

func GetUserPrincipalFromContext(ctx context.Context) string {
	userDetails := GetUserDetailsFromContext(ctx)

	if userDetails == nil {
		if util.IsSystemContext(ctx) {
			return "system"
		} else {
			return "guest"
		}
	}

	return userDetails.Username
}

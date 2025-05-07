package pkg

import "context"

func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(ctx, "tenant", tenant)
}

func GetTenant(ctx context.Context) string {
	if ctx.Value("tenant") == nil {
		return ""
	}
	return ctx.Value("tenant").(string)
}

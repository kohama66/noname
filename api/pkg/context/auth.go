package context

import "context"

var (
	authIDKey = "AUTH_ID"
)

func AuthID(ctx context.Context) string {
	if v := ctx.Value(&authIDKey); v != nil {
		return v.(string)
	}
	return ""
}

package middleware

import (
	"net/http"

	"github.com/myapp/noname/api/pkg/context"
)

func AuthAPI(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		debugID := r.Header.Get("X-DEBUG-ID")
		r = r.WithContext(context.WithAuthID(ctx, debugID))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

package middleware

import (
	"net/http"

	"github.com/myapp/noname/api/pkg/context"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

func AuthAPI(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		debugID := r.Header.Get("X-DEBUG-ID")
		if debugID == "" {
			log.Errorf(ctx, "Auth.VerifyIDToken empty")
			factory.ErrorJSON(w, "Auth.VerifyIDToken empty", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithAuthID(ctx, debugID))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

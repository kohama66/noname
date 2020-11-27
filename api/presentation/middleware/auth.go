package middleware

import (
	"net/http"
	"strings"

	"github.com/myapp/noname/api/infrastructure/firebase/auth"
	"github.com/myapp/noname/api/pkg/context"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// AuthAPI 認証ユーザー判断ミドルウェア
func AuthAPI(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Errorf(ctx, "Auth.VerifyIDToken %v", err)
			factory.ErrorJSON(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithAuthID(ctx, token.UID))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// func AuthAPI(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		debugID := r.Header.Get("X-DEBUG-ID")
// 		if debugID == "" {
// 			log.Errorf(ctx, "Auth.VerifyIDToken empty")
// 			factory.ErrorJSON(w, "Auth.VerifyIDToken empty", http.StatusUnauthorized)
// 			return
// 		}
// 		r = r.WithContext(context.WithAuthID(ctx, debugID))
// 		next.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(fn)
// }

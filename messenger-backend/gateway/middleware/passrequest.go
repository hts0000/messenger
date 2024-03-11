package middleware

import (
	"context"
	"net/http"
)

type PassRequestKey struct{}

// 传递request到context中
func PassRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/auth/login" {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = context.WithValue(ctx, PassRequestKey{}, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

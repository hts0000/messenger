package middleware

import (
	"net/http"
)

// 为mux设置一系列中间件
func NewHandler(mux http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		mux = middleware(mux)
	}
	return mux
}

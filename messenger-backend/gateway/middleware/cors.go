package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// 设置跨域
func Cors(next http.Handler) http.Handler {
	return cors.Default().Handler(next)
}

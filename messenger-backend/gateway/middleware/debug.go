package middleware

import (
	"log"
	"messenger-backend/share/server"
	"net/http"

	"go.uber.org/zap"
)

// 打印请求路径和方法
func Debug(next http.Handler) http.Handler {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		logger.Debug("handle", zap.String("method", method), zap.String("path", path))
		next.ServeHTTP(w, r)
	})
}

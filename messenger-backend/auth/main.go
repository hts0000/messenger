package main

import (
	"io"
	"log"
	authpb "messenger-backend/auth/api/gen/v1"
	"messenger-backend/auth/auth"
	"messenger-backend/auth/dao"
	"messenger-backend/auth/password"
	"messenger-backend/auth/token"
	"messenger-backend/share/server"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	dsn := "root:123456@tcp(127.0.0.1:13306)/messenger?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("cannot connect mysql", zap.Error(err))
	}

	pkFp, err := os.Open("auth/auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key file", zap.Error(err))
	}
	defer pkFp.Close()

	pkByte, err := io.ReadAll(pkFp)
	if err != nil {
		logger.Fatal("cannot read private key file", zap.Error(err))
	}

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(pkByte)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	logger.Fatal("run grpc server failed", zap.Error(server.RunGrpcServer(&server.GrpcConfig{
		Name:   "auth",
		Addr:   ":18081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				Logger:            logger,
				MySQL:             dao.NewMySQL(mysqlDB),
				PasswordEncryptor: password.NewArgon2Gen(password.WithDefaultSalt()),
				TokenGenerator:    token.NewJwtTokenGen("messenger/auth", pk),
				TokenExpire:       2 * time.Hour,
			})
		},
	})))
}

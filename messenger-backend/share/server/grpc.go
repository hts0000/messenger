package server

import (
	"messenger-backend/share/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGrpcServer(cfg *GrpcConfig) error {
	nameField := zap.String("name", cfg.Name)
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		cfg.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption
	// 如果需要解析token，则添加auth拦截器，用于解析token里的uid
	if cfg.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(cfg.AuthPublicKeyFile)
		if err != nil {
			cfg.Logger.Fatal("cannot create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	cfg.RegisterFunc(s)

	cfg.Logger.Info("server started", nameField, zap.String("addr", cfg.Addr))
	return s.Serve(lis)
}

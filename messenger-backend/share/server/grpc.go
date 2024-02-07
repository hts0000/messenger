package server

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(cfg *GRPCConfig) error {
	nameFiled := zap.String("name", cfg.Name)
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		cfg.Logger.Fatal("cannot listen", nameFiled, zap.Error(err))
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	cfg.RegisterFunc(s)

	cfg.Logger.Info("server started", nameFiled, zap.String("addr", cfg.Addr))
	return s.Serve(lis)
}

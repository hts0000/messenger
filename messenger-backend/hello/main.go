package main

import (
	"context"
	"log"
	hellopb "messenger-backend/hello/api/gen/v1"
	"messenger-backend/share/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Hello struct {
	Logger *zap.Logger
	hellopb.UnimplementedGreeterServer
}

func (h *Hello) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	h.Logger.Info("received message", zap.String("name", req.GetName()))
	return &hellopb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	logger.Fatal("run grpc server failed", zap.Error(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "hello",
		Addr:   ":18081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			hellopb.RegisterGreeterServer(s, &Hello{
				Logger: logger,
			})
		},
	})))
}

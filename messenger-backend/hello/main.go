package main

import (
	"context"
	"fmt"
	"log"
	hellopb "messenger-backend/hello/api/gen/v1"
	"messenger-backend/share/server"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Hello struct {
	Logger *zap.Logger
	hellopb.UnimplementedGreeterServer
}

func (h *Hello) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	st := status.New(codes.InvalidArgument, "need name argument")
	// 附加额外的错误信息
	detial, err := st.WithDetails(&errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			{
				Field:       "service_name",
				Description: "hello",
			},
			{
				Field:       "error_msg",
				Description: "invalid argument, need name argument",
			},
		},
	})
	// 附加额外错误信息成功，返回附加了新信息的错误
	if err == nil {
		fmt.Println("######", detial.Err())
		return nil, detial.Err()
	}
	fmt.Println("*****", st.Err())
	// 如果附加额外信息失败，返回常规错误
	return nil, st.Err()
}

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	logger.Fatal("run grpc server failed", zap.Error(server.RunGrpcServer(&server.GrpcConfig{
		Name:   "hello",
		Addr:   ":28081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			hellopb.RegisterGreeterServer(s, &Hello{
				Logger: logger,
			})
		},
	})))
}

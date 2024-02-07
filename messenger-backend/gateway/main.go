package main

import (
	"context"
	"log"
	hellopb "messenger-backend/hello/api/gen/v1"
	"messenger-backend/share/server"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseEnumNumbers: true, // 返回常量时，返回对应的number而不是string
					UseProtoNames:  true, // 返回JSON时，字段名格式为xxx_xxx格式
				},
			},
		),
	)

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "hello",
			addr:         "localhost:18081",
			registerFunc: hellopb.RegisterGreeterHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			logger.Fatal("cannot register service", zap.String("service", s.name), zap.Error(err))
		}
	}
	// 设置跨域
	handler := cors.Default().Handler(mux)
	addr := ":18080"
	logger.Info("grpc gateway started", zap.String("addr", addr))
	logger.Fatal("cannot listen and server", zap.Error(http.ListenAndServe(addr, handler)))
}

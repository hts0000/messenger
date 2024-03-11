package auth

import (
	"context"
	"fmt"
	"io"
	"messenger-backend/share/auth/token"
	"messenger-backend/share/id"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

// 返回一个grpc interceptor
// 将token中的uid设置到ctx中
// 后续grpc服务可以从ctx中拿到uid
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file, err: %v", err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file, err: %v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key file, err: %v", err)
	}

	i := &authInterceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}

	return i.HandleReq, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type authInterceptor struct {
	verifier tokenVerifier
}

// grpc interceptor 解析http请求header中的authorizationHeader
func (i *authInterceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	idStr, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	uid, err := id.String2UID(idStr)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot parse uid, err: %v", err)
	}

	resp, err = handler(ContextWithUID(ctx, uid), req)
	return resp, err
}

// grcp-gateway获取到的http header会转换成metadata存在ctx中
func tokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}

type uidKey struct{}

func ContextWithUID(ctx context.Context, uid id.UID) context.Context {
	return context.WithValue(ctx, uidKey{}, uid)
}

func UIDFromContext(ctx context.Context) (id.UID, error) {
	v := ctx.Value(uidKey{})
	uid, ok := v.(id.UID)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "")
	}
	return uid, nil
}

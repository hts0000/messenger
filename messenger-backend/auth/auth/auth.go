package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	authpb "messenger-backend/auth/api/gen/v1"
	"messenger-backend/auth/dao"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service struct {
	Logger      *zap.Logger
	MySQL       *dao.MySQL
	TokenExpire time.Duration

	PasswordEncryptor PasswordEncryptor
	TokenGenerator    TokenGenerator

	authpb.UnimplementedAuthServiceServer
}

type PasswordEncryptor interface {
	EncryptPassword(password string, salt []byte) (encryptedPassword []byte, err error)
}

type TokenGenerator interface {
	GenerateToken(uid string, expire time.Duration) (token string, err error)
}

func (s *Service) Login(ctx context.Context, req *authpb.AuthRequest) (*authpb.AuthResponse, error) {
	s.Logger.Info("user login", zap.String("user_email", req.Email))

	// check user exists
	user, err := s.MySQL.GetUser(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Error("user not existed", zap.String("email", req.Email))
			return nil, status.Error(codes.InvalidArgument, "user not existed")
		}
		s.Logger.Error("get user failed", zap.String("email", req.Email), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// validate password
	encryptedPassword, err := s.PasswordEncryptor.EncryptPassword(req.Password, user.Salt)
	if err != nil {
		s.Logger.Error("encrypt password failed", zap.Any("request", req), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if !bytes.Equal(encryptedPassword, user.Password) {
		return nil, status.Error(codes.InvalidArgument, "password incorrect")
	}

	// gen jwt token
	uid := strconv.FormatUint(user.ID, 10)
	tkn, err := s.TokenGenerator.GenerateToken(uid, s.TokenExpire)
	if err != nil {
		s.Logger.Error("generate token failed", zap.Uint64("uid", user.ID), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// set token to metadata
	md := metadata.MD{}
	// Grpc-Metadata-Token
	md.Set("token", tkn)
	if err := grpc.SetHeader(ctx, md); err != nil {
		s.Logger.Error("set token to metadata failed", zap.Any("metadata", md))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &authpb.AuthResponse{}, nil
}

func (s *Service) Register(ctx context.Context, req *authpb.AuthRequest) (*authpb.AuthResponse, error) {
	s.Logger.Info("user login", zap.String("user_email", req.Email))

	// check username
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is invalid")
	}

	// check email
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is invalid")
	}

	// check password
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is invalid")
	}

	// check user exists
	user, err := s.MySQL.GetUser(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Logger.Error("get user failed", zap.String("email", req.Email), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if user != nil {
		s.Logger.Error("user already existed", zap.String("email", req.Email), zap.Any("user", user))
		return nil, status.Error(codes.InvalidArgument, "user already existed")
	}

	// encrypt password
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		s.Logger.Error("gen rand salt failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	encryptedPassword, err := s.PasswordEncryptor.EncryptPassword(req.Password, salt)
	if err != nil {
		s.Logger.Error("encrypt password failed", zap.Any("request", req), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// create user
	_, err = s.MySQL.CreateUser(ctx, req.Name, req.Email, encryptedPassword, salt)
	if err != nil {
		s.Logger.Error("create user failed", zap.Any("request", req), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// return user
	return &authpb.AuthResponse{}, nil
}

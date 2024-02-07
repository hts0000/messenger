package server

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.EncoderConfig.FunctionKey = "F"
	return cfg.Build()
}

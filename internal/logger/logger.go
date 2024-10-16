package logger

import (
	"github.com/mechatron-x/atehere/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(conf *config.App) *zap.Logger {
	loggerConf := conf.Logger
	return zap.Must(loggerConf.Build())
}

func ErrorReason(err error) zapcore.Field {
	return zap.String("reason", err.Error())
}

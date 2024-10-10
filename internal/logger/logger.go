package logger

import (
	"github.com/mechatron-x/8here/internal/config"
	"go.uber.org/zap"
)

func New(conf *config.App) *zap.Logger {
	loggerConf := conf.Logger
	return zap.Must(loggerConf.Build())
}

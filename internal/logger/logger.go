package logger

import (
	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

func Config(conf zap.Config) {
	log = zap.Must(conf.Build())
}

func Instance() *zap.Logger {
	if log != nil {
		return log
	}

	panic("logger is not initialized")
}

func Error(msg string, reason error) {
	if log == nil {
		return
	}

	log.Error(msg, zap.String("reason", reason.Error()))
}

func Fatal(msg string, reason error) {
	if log == nil {
		return
	}

	log.Fatal(msg, zap.String("reason", reason.Error()))
}

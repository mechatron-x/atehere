package middleware

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/header"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func Recover(next http.Handler) http.Handler {
	log := logger.Instance()
	h := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Recovered from panic", zap.Any("reason", err))
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set(header.ConnectionKey, header.ConnectionClose)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(h)
}

package middleware

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/header"
)

func Header(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		w.Header().Add(header.AccessControlAllowOriginKey, "*")
	}

	return http.HandlerFunc(f)
}

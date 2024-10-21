package httpserver

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/httpserver/middleware"
	"go.uber.org/zap"
)

type (
	Params struct {
		Conf *config.App
		Mux  *http.ServeMux
	}
)

func NewHTTP(apiConf config.Api, mux *http.ServeMux, log *zap.Logger) error {
	url := fmt.Sprintf("%s:%s", apiConf.Host, apiConf.Port)
	srv := &http.Server{
		Addr:              url,
		Handler:           mux,
		ReadHeaderTimeout: time.Millisecond,
	}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	log.Info("Starting HTTP server at", zap.String("address", srv.Addr))
	return srv.Serve(ln)
}

func NewServeMux(routes []handler.Router, log *zap.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		mux.Handle(route.Pattern(),
			middleware.Logger(
				middleware.Header(
					route,
				),
				log,
			),
		)
	}

	return mux
}

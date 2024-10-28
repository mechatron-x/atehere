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

func NewServeMux(
	conf config.Api,
	log *zap.Logger,
	hh handler.Health,
	ch handler.Customer,
	mh handler.Manager,
) *http.ServeMux {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()
	versionMux := http.NewServeMux()

	// Health endpoints
	apiMux.HandleFunc("GET /api/health", hh.GetHealth)

	// Customer endpoints
	versionMux.HandleFunc("GET /customer/profile", ch.GetProfile)
	versionMux.HandleFunc("PATCH /customer/profile", ch.UpdateProfile)
	versionMux.HandleFunc("POST /customer/auth/signup", ch.SignUp)

	// Manager endpoints
	versionMux.HandleFunc("GET /manager/profile", mh.GetProfile)
	versionMux.HandleFunc("PATCH /manager/profile", mh.UpdateProfile)
	versionMux.HandleFunc("POST /manager/auth/signup", mh.SignUp)

	// Routers
	mux.Handle("/", middleware.Header(middleware.Logger(apiMux, log)))
	apiMux.Handle("/api/", http.StripPrefix("/api", versionMux))
	versionMux.Handle(fmt.Sprintf("/%s/", conf.Version), http.StripPrefix(fmt.Sprintf("/%s", conf.Version), versionMux))

	return mux
}

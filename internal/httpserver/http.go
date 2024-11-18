package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/httpserver/middleware"
	"github.com/mechatron-x/atehere/internal/logger"
)

type (
	Params struct {
		Conf *config.App
		Mux  *http.ServeMux
	}
)

func New(apiConf config.Api, mux *http.ServeMux) (*http.Server, error) {
	url := fmt.Sprintf("%s:%s", apiConf.Host, apiConf.Port)
	srv := &http.Server{
		Addr:              url,
		Handler:           mux,
		ReadHeaderTimeout: time.Millisecond,
	}

	return srv, nil
}

func NewServeMux(
	conf config.Api,
	dh handler.Default,
	hh handler.Health,
	ch handler.Customer,
	mh handler.Manager,
	rh handler.Restaurant,
	rmh handler.Menu,
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

	// Manager restaurant endpoints
	versionMux.HandleFunc("GET /manager/restaurants", rh.ListForManager)
	versionMux.HandleFunc("POST /manager/restaurants", rh.Create)
	versionMux.HandleFunc("DELETE /manager/restaurants/{restaurant_id}", rh.Delete)

	// Restaurant endpoints
	versionMux.HandleFunc("POST /restaurants", rh.ListForCustomer)
	versionMux.HandleFunc("GET /restaurants/{restaurant_id}", rh.GetOneForCustomer)

	// Menu endpoints
	versionMux.HandleFunc("POST /menu", rmh.Create)
	versionMux.HandleFunc("PATCH /menu/item", rmh.AddMenuItem)
	versionMux.HandleFunc("GET /restaurants/{restaurant_id}/menus", rmh.ListForCustomer)

	// Default handler
	apiMux.HandleFunc("/", dh.NoHandler)
	versionMux.HandleFunc("/", dh.NoHandler)

	// Routers
	mux.Handle("/", middleware.Header(middleware.Logger(apiMux, logger.Instance())))
	apiMux.Handle("/api/", http.StripPrefix("/api", versionMux))
	apiMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(conf.StaticRoot))))
	versionMux.Handle(fmt.Sprintf("/%s/", conf.Version), http.StripPrefix(fmt.Sprintf("/%s", conf.Version), versionMux))

	return mux
}

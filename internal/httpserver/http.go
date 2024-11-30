package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/httpserver/middleware"
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
	sh handler.Session,
) *http.ServeMux {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()
	versionMux := http.NewServeMux()

	// Health endpoints
	apiMux.HandleFunc("GET /api/health", hh.GetHealth)

	// Customer endpoints
	versionMux.HandleFunc("GET /customers", ch.GetProfile)
	versionMux.HandleFunc("PATCH /customers", ch.UpdateProfile)
	versionMux.HandleFunc("POST /customers/auth/signup", ch.SignUp)

	// Manager endpoints
	versionMux.HandleFunc("GET /managers", mh.GetProfile)
	versionMux.HandleFunc("PATCH /managers", mh.UpdateProfile)
	versionMux.HandleFunc("POST /managers/auth/signup", mh.SignUp)

	// Manager restaurant endpoints
	versionMux.HandleFunc("GET /managers/restaurants", rh.ListForManager)
	versionMux.HandleFunc("POST /managers/restaurants", rh.Create)
	versionMux.HandleFunc("DELETE /managers/restaurants/{restaurant_id}", rh.Delete)

	// Restaurant endpoints
	versionMux.HandleFunc("POST /restaurants", rh.ListForCustomer)
	versionMux.HandleFunc("GET /restaurants/{restaurant_id}", rh.GetOneForCustomer)

	// Menu endpoints
	versionMux.HandleFunc("POST /menus", rmh.Create)
	versionMux.HandleFunc("PUT /menus/{menu_id}/items", rmh.AddMenuItem)
	versionMux.HandleFunc("GET /restaurants/{restaurant_id}/menus", rmh.ListForCustomer)

	// Session endpoints
	versionMux.HandleFunc("POST /tables/{table_id}/checkout", sh.Checkout)
	versionMux.HandleFunc("PUT /tables/{table_id}/orders", sh.PlaceOrder)
	versionMux.HandleFunc("GET /tables/{table_id}/customer/orders", sh.ListForCustomer)
	versionMux.HandleFunc("GET /tables/{table_id}/orders", sh.ListForTable)

	// Default handler
	apiMux.HandleFunc("/", dh.NoHandler)
	versionMux.HandleFunc("/", dh.NoHandler)

	// Routers
	mux.Handle("/",
		middleware.Header(
			middleware.Logger(
				middleware.Recover(apiMux),
			),
		),
	)
	apiMux.Handle("/api/", http.StripPrefix("/api", versionMux))
	apiMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(conf.StaticRoot))))
	versionMux.Handle(fmt.Sprintf("/%s/", conf.Version), http.StripPrefix(fmt.Sprintf("/%s", conf.Version), versionMux))

	return mux
}

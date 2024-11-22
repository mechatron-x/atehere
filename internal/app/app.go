package app

import (
	"fmt"
	"net/http"

	"github.com/mechatron-x/atehere/internal/app/ctx"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/sqldb"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type App struct {
	conf       *config.App
	httpServer *http.Server
}

func New(conf *config.App) (*App, error) {
	logger.Config(conf.Logger)

	db, err := sqldb.Connect(conf.DB)
	if err != nil {
		return nil, err
	}

	err = sqldb.Migrate(
		db,
		&model.Customer{},
		&model.Manager{},
		&model.Restaurant{},
		&model.RestaurantTable{},
		&model.Menu{},
		&model.MenuItem{},
		&model.Session{},
		&model.SessionOrder{},
	)
	if err != nil {
		return nil, err
	}

	diskFileManager := infrastructure.NewDiskFileManager()
	imageStorage, err := infrastructure.NewImageStorage(diskFileManager, conf.Api.StaticRoot)
	if err != nil {
		return nil, err
	}

	var authenticator port.Authenticator

	if conf.Environment == config.PROD {
		authenticator, err = infrastructure.NewFirebaseAuthenticator(conf.Firebase)
		if err != nil {
			return nil, err
		}
	} else {
		authenticator, err = infrastructure.NewMockAuthenticator(conf.Api, diskFileManager)
		if err != nil {
			return nil, err
		}
	}

	eventNotifier, err := infrastructure.NewFirebaseEventNotifier()
	if err != nil {
		return nil, err
	}

	customerCtx := ctx.NewCustomer(db, authenticator)
	managerCtx := ctx.NewManager(db, authenticator)
	restaurantCtx := ctx.NewRestaurant(db, authenticator, imageStorage, conf.Api)
	menuCtx := ctx.NewMenu(db, authenticator, imageStorage, conf.Api)
	sessionCtx := ctx.NewSession(db, authenticator, eventNotifier)

	mux := httpserver.NewServeMux(
		conf.Api,
		handler.NewDefault(),
		handler.NewHealth(),
		customerCtx.Handler(),
		managerCtx.Handler(),
		restaurantCtx.Handler(),
		menuCtx.Handler(),
		sessionCtx.Handler(),
	)

	httpServer, err := httpserver.New(conf.Api, mux)
	if err != nil {
		return nil, err
	}

	return &App{
		conf:       conf,
		httpServer: httpServer,
	}, nil
}

func (a *App) Start() error {
	log := logger.Instance()
	log.Info(fmt.Sprintf("Starting HTTP server at: %s", a.httpServer.Addr))
	return a.httpServer.ListenAndServe()
}

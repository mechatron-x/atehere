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
)

type App struct {
	conf       *config.App
	dbManager  *sqldb.DbManager
	httpServer *http.Server
}

func New(conf *config.App) (*App, error) {
	logger.Config(conf.Logger)

	dbManager, err := newDBManager(conf.DB)
	if err != nil {
		return nil, err
	}

	firebaseAuth, err := infrastructure.NewFirebaseAuthenticator(conf.Firebase)
	if err != nil {
		return nil, err
	}

	customerCtx := ctx.NewCustomer(dbManager.DB(), firebaseAuth)
	managerCtx := ctx.NewManager(dbManager.DB(), firebaseAuth)
	restaurantCtx := ctx.NewRestaurant(dbManager.DB(), firebaseAuth)

	mux := httpserver.NewServeMux(
		conf.Api,
		handler.NewDefault(),
		handler.NewHealth(),
		customerCtx.Handler(),
		managerCtx.Handler(),
		restaurantCtx.Handler(),
	)

	httpServer, err := httpserver.New(conf.Api, mux)
	if err != nil {
		return nil, err
	}

	return &App{
		conf:       conf,
		dbManager:  dbManager,
		httpServer: httpServer,
	}, nil
}

func (a *App) Start() error {
	log := logger.Instance()
	log.Info(fmt.Sprintf("Starting HTTP server at: %s", a.httpServer.Addr))
	return a.httpServer.ListenAndServe()
}

func (a *App) Shutdown() error {
	return a.dbManager.MigrateDown()
}

func newDBManager(dbConf config.DB) (*sqldb.DbManager, error) {
	dm := sqldb.New(dbConf)
	if err := dm.Connect(); err != nil {
		return nil, err
	}
	if err := dm.MigrateUp(); err != nil {
		return nil, err
	}

	return dm, nil
}

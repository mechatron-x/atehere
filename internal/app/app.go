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
		&model.RestaurantWorkingDay{},
		&model.Menu{},
		&model.MenuItem{},
		&model.MenuItemIngredient{},
	)
	if err != nil {
		return nil, err
	}

	firebaseAuth, err := infrastructure.NewFirebaseAuthenticator(conf.Firebase)
	if err != nil {
		return nil, err
	}

	diskFileSaver := infrastructure.NewDiskFileSaver()
	imageStorage, err := infrastructure.NewImageStorage(diskFileSaver, conf.Api.StaticRoot)
	if err != nil {
		return nil, err
	}

	customerCtx := ctx.NewCustomer(db, firebaseAuth)
	managerCtx := ctx.NewManager(db, firebaseAuth)
	restaurantCtx := ctx.NewRestaurant(db, firebaseAuth, imageStorage, conf.Api)
	menuCtx := ctx.NewMenu(db, firebaseAuth, imageStorage, conf.Api)

	mux := httpserver.NewServeMux(
		conf.Api,
		handler.NewDefault(),
		handler.NewHealth(),
		customerCtx.Handler(),
		managerCtx.Handler(),
		restaurantCtx.Handler(),
		menuCtx.Handler(),
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

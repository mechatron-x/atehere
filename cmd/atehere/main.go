package main

import (
	"github.com/mechatron-x/atehere/internal/cmdarg"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure"
	"github.com/mechatron-x/atehere/internal/logger"
	restaurantsrv "github.com/mechatron-x/atehere/internal/restaurant/service"
	"github.com/mechatron-x/atehere/internal/sqldb"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	mngntsrv "github.com/mechatron-x/atehere/internal/usermanagement/service"
)

var (
	conf *config.App
)

func init() {
	flags := cmdarg.Setup()
	c, err := config.Load(flags.ConfPath)
	if err != nil {
		panic(err)
	}

	conf = c
}

func main() {
	// Logger config
	logger.Config(conf.Logger)

	// DB Connection and Migrations
	dm := sqldb.New(conf.DB)
	if err := dm.Connect(); err != nil {
		logger.Fatal("Unable to connect to the db", err)
	}
	if err := dm.MigrateUp(); err != nil {
		logger.Fatal("Unable to migrate the db", err)
	}

	// Repositories
	customerRepository := repository.NewCustomer(dm.DB())
	managerRepository := repository.NewManager(dm.DB())

	// Infrastructure services
	firebaseAuthenticator, err := infrastructure.NewFirebaseAuthenticator(conf.Firebase)
	if err != nil {
		logger.Fatal("Firebase initialization error", err)
	}

	// Services
	customerService := mngntsrv.NewCustomer(customerRepository, firebaseAuthenticator)
	managerService := mngntsrv.NewManager(managerRepository, firebaseAuthenticator)
	restaurantService := restaurantsrv.NewRestaurant(nil, firebaseAuthenticator)

	// HTTP handlers
	healthHandler := handler.NewHealth()
	customerHandler := handler.NewCustomerHandler(*customerService)
	managerHandler := handler.NewManagerHandler(*managerService)
	restaurantHandler := handler.NewRestaurantHandler(*restaurantService)

	// Start HTTP server
	mux := httpserver.NewServeMux(
		conf.Api,
		healthHandler,
		customerHandler,
		managerHandler,
		restaurantHandler,
	)
	err = httpserver.NewHTTP(conf.Api, mux)
	if err != nil {
		logger.Fatal("Cannot start HTTP server", err)
	}
}

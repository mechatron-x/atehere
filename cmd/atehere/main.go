package main

import (
	"log"

	"github.com/mechatron-x/atehere/internal/cmdarg"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/sqldb"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
	"go.uber.org/zap"
)

func main() {
	flags := cmdarg.Setup()
	conf, err := config.Load(flags.ConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log := logger.New(conf)

	// DB Connection and Migrations
	dm := sqldb.New(conf.DB, log)
	if err = dm.Connect(); err != nil {
		log.Fatal("Unable to connect to the db", logger.ErrorReason(err))
	}
	if err = dm.MigrateUp(); err != nil {
		log.Fatal("Unable to migrate the db", logger.ErrorReason(err))
	}

	// Repositories
	userRepo := repository.NewCustomer(dm.DB())

	// Infrastructure services
	firebaseAuthenticator, err := infrastructure.NewFirebaseAuthenticator(conf.Firebase, log)
	if err != nil {
		log.Fatal("Firebase initialization error", logger.ErrorReason(err))
	}

	// Services
	userService := service.NewCustomer(userRepo, firebaseAuthenticator)

	// HTTP handlers
	handlers := make([]handler.Router, 0)
	handlers = append(
		handlers,
		handler.NewHealth(),
		handler.NewCustomerSignUp(userService),
		handler.NewCustomerProfile(userService),
	)

	// Start HTTP server
	mux := httpserver.NewServeMux(handlers, log)
	err = httpserver.NewHTTP(conf.Api, mux, log)
	if err != nil {
		log.Fatal("Cannot start HTTP server", zap.String("reason", err.Error()))
	}

}

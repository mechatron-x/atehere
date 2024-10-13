package main

import (
	"log"

	"github.com/mechatron-x/atehere/internal/cmdarg"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/sqldb"
	"go.uber.org/zap"
)

func main() {
	flags := cmdarg.Setup()
	conf, err := config.Load(flags.ConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log := logger.New(conf)

	db, err := sqldb.Connect(conf.DB, log)
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = db.Ping()

	handlers := make([]handler.Route, 0)
	handlers = append(
		handlers,
		handler.NewHealth(),
	)

	mux := httpserver.NewServeMux(handlers, log)
	err = httpserver.NewHTTP(conf.Api, mux, log)
	if err != nil {
		log.Fatal("Cannot start HTTP server", zap.String("reason", err.Error()))
	}
}

package main

import (
	"github.com/mechatron-x/8here/internal/cmdarg"
	"github.com/mechatron-x/8here/internal/config"
	"github.com/mechatron-x/8here/internal/httpserver"
	"github.com/mechatron-x/8here/internal/httpserver/handler"
	"github.com/mechatron-x/8here/internal/logger"
	"github.com/mechatron-x/8here/internal/sqldb"
	"go.uber.org/zap"
)

func main() {
	flags, err := cmdarg.Setup()
	if err != nil {
		panic(err)
	}

	conf := config.Load(flags.ConfPath)
	log := logger.New(conf)

	db, err := sqldb.Connect(conf.DB, log)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Ping()

	mux := httpserver.NewServeMux([]handler.Route{}, log)
	err = httpserver.NewHTTP(conf.Api, mux, log)
	if err != nil {
		log.Fatal("Cannot start HTTP server", zap.String("reason", err.Error()))
	}
}

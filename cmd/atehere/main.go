package main

import (
	"log"

	"github.com/mechatron-x/atehere/internal/app"
	"github.com/mechatron-x/atehere/internal/cmdarg"
	"github.com/mechatron-x/atehere/internal/config"
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
	app, err := app.New(conf)
	if err != nil {
		log.Fatal("Error starting application:", err)
	}

	if err := app.Start(); err != nil {
		log.Fatal("Failed to start application:", err)
	}

	defer app.Shutdown()
}

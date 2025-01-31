package app

import (
	"fmt"
	"net/http"

	"github.com/mechatron-x/atehere/internal/app/ctx"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/authenticator"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/httpserver"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"github.com/mechatron-x/atehere/internal/infrastructure/notifier"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/infrastructure/storage"
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
		&model.RestaurantLocation{},
		&model.Menu{},
		&model.MenuItem{},
		&model.Session{},
		&model.SessionOrder{},
		&model.Bill{},
		&model.BillItem{},
		&model.BillItemPayments{},
	)
	if err != nil {
		return nil, err
	}

	diskFileManager := storage.NewFile()
	imageStorage, err := storage.NewImage(diskFileManager, conf.Api.StaticRoot)
	if err != nil {
		return nil, err
	}

	var auth port.Authenticator

	if conf.Environment == config.PROD {
		auth, err = authenticator.NewFirebase(conf.Firebase)
		if err != nil {
			return nil, err
		}
	} else {
		auth, err = authenticator.NewMock(conf.Api, diskFileManager)
		if err != nil {
			return nil, err
		}
	}

	eventNotifier, err := notifier.NewFirestore(conf.Firebase)
	if err != nil {
		return nil, err
	}

	// Publishers
	newOrderEventPublisher := broker.NewPublisher[core.NewOrderEvent]()
	checkoutEventPublisher := broker.NewPublisher[core.CheckoutEvent]()
	allPaymentsDonePublisher := broker.NewPublisher[core.AllPaymentsDoneEvent]()

	// Contexts
	customerCtx := ctx.NewCustomer(db, auth)
	managerCtx := ctx.NewManager(db, auth)
	restaurantCtx := ctx.NewRestaurant(db, auth, imageStorage, conf.Api)
	menuCtx := ctx.NewMenu(db, auth, imageStorage, conf.Api)
	sessionCtx := ctx.NewSession(db, auth, newOrderEventPublisher, checkoutEventPublisher)
	billingCtx := ctx.NewBilling(db, auth, allPaymentsDonePublisher)

	// Consumers
	newOrderEventPublisher.AddConsumer(
		ctx.NewNotifyOrderConsumer(db, eventNotifier),
	)
	checkoutEventPublisher.AddConsumer(
		ctx.NewCheckoutConsumer(db, eventNotifier),
		ctx.NewCreateBillConsumer(db),
	)
	allPaymentsDonePublisher.AddConsumer(
		ctx.NewSessionClosedConsumer(db),
	)

	mux := httpserver.NewServeMux(
		conf.Api,
		handler.NewDefault(),
		handler.NewHealth(),
		customerCtx.Handler(),
		managerCtx.Handler(),
		restaurantCtx.Handler(),
		menuCtx.Handler(),
		sessionCtx.Handler(),
		billingCtx.Handler(),
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

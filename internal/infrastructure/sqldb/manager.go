package sqldb

import (
	"errors"
	"fmt"
	"time"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/view"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

const (
	defaultConnTimeout = time.Second * 2
)

func Connect(config config.DB) (*gorm.DB, error) {
	log := logger.Instance()
	dsn := createDsn(config)

	db, err := setupConnection(dsn, log, config)
	if err != nil {
		return nil, err
	}

	log.Info("DB connection established")
	return db, nil
}

func Migrate(db *gorm.DB, model ...any) error {
	if err := db.AutoMigrate(); err != nil {
		return err
	}

	if err := db.Migrator().CreateView("table_orders", gorm.ViewOption{Query: view.TableOrdersView(db), Replace: true}); err != nil {
		return err
	}

	if err := db.Migrator().CreateView("manager_orders", gorm.ViewOption{Query: view.ManagerOrdersView(db), Replace: true}); err != nil {
		return err
	}

	return nil
}

func createDsn(config config.DB) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Driver,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}

func setupConnection(dsn string, log *zap.Logger, config config.DB) (*gorm.DB, error) {
	timeout, err := time.ParseDuration(config.Timeout)
	if err != nil {
		log.Warn(fmt.Sprintf("Using default timeout: %s", defaultConnTimeout), zap.String("reason", err.Error()))
		timeout = defaultConnTimeout
	}

	db, err := connect(dsn, config.TryCount, timeout)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect(dsn string, tyrCount uint, timeout time.Duration) (*gorm.DB, error) {
	if tyrCount <= 0 {
		return nil, errors.New("connection failed, all retries exhausted")
	}

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		time.Sleep(timeout)
		return connect(dsn, tyrCount-1, timeout)
	}

	db.Logger.LogMode(glogger.Silent)
	return db, nil
}

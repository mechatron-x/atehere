package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/logger"
	"go.uber.org/zap"
)

const (
	defaultConnTimeout = time.Second * 2
)

type DbManager struct {
	conf    config.DB
	log     *zap.Logger
	db      *sql.DB
	migrate *migrate.Migrate
}

func New(conf config.DB) *DbManager {
	return &DbManager{
		conf: conf,
		log:  logger.Instance(),
	}
}

func (dm *DbManager) Connect() error {
	log := dm.log
	dbConf := dm.conf

	if err := dm.setupConnection(); err != nil {
		return err
	}

	if err := dm.setupMigration(); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("DB connection established at: %s", dbConf.DSN))
	return nil
}

func (dm *DbManager) DB() *sql.DB {
	return dm.db
}

func (dm *DbManager) MigrateUp() error {
	err := dm.migrate.Up()
	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (dm *DbManager) MigrateDown() error {
	return dm.migrate.Down()
}

func (dm *DbManager) setupConnection() error {
	dbConf := dm.conf
	log := dm.log

	timeout, err := time.ParseDuration(dbConf.Timeout)
	if err != nil {
		log.Warn(fmt.Sprintf("Using default timeout: %s", defaultConnTimeout), zap.String("reason", err.Error()))
		timeout = defaultConnTimeout
	}

	db, err := connect(dbConf.DSN, dbConf.TryCount, timeout)
	if err != nil {
		return err
	}

	dm.db = db
	return nil
}

func (dm *DbManager) setupMigration() error {
	dbConf := dm.conf

	db, err := sql.Open(dbConf.Driver, dbConf.DSN)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		dbConf.Migrations,
		dbConf.Driver,
		driver)

	if err != nil {
		return err
	}

	dm.migrate = m
	return nil
}

func connect(dsn string, tyrCount uint, timeout time.Duration) (*sql.DB, error) {
	if tyrCount <= 0 {
		return nil, errors.New("connection failed, all retries exhausted")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		time.Sleep(timeout)
		return connect(dsn, tyrCount-1, timeout)
	}

	return db, nil
}

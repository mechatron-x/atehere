package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mechatron-x/atehere/internal/config"
	"go.uber.org/zap"
)

const (
	defaultConnTimeout = time.Second * 2
)

type DbManager struct {
	conf    config.DB
	log     *zap.Logger
	pool    *pgxpool.Pool
	migrate *migrate.Migrate
}

func New(conf config.DB, log *zap.Logger) *DbManager {
	return &DbManager{
		conf: conf,
		log:  log,
	}
}

func (dm *DbManager) Connect() error {
	log := dm.log
	dbConf := dm.conf

	if err := dm.setupConnection(); err != nil {
		return err
	}

	if err := dm.setupMigrate(); err != nil {
		return err
	}

	log.Info("DB connection established", zap.String("address", dbConf.DSN))
	return nil
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

	connPool, err := connect(dbConf.DSN, dbConf.TryCount, timeout)
	if err != nil {
		return err
	}

	dm.pool = connPool
	return nil
}

func (dm *DbManager) setupMigrate() error {
	dbConf := dm.conf

	db, err := sql.Open(dbConf.Driver, dbConf.DSN)
	if err != nil {
		return err
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})
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

func Connect(dbConf config.DB, log *zap.Logger) (*pgxpool.Pool, error) {
	timeout, err := time.ParseDuration(dbConf.Timeout)
	if err != nil {
		log.Warn(fmt.Sprintf("Using default timeout: %s", defaultConnTimeout), zap.String("reason", err.Error()))
		timeout = defaultConnTimeout
	}

	connPool, err := connect(dbConf.DSN, dbConf.TryCount, timeout)
	if err != nil {
		return nil, err
	}

	log.Info("DB connection established", zap.String("address", dbConf.DSN))

	return connPool, nil
}

func connect(dsn string, tyrCount uint, timeout time.Duration) (*pgxpool.Pool, error) {
	if tyrCount <= 0 {
		return nil, errors.New("connection failed, all retries exhausted")
	}

	connPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v\n", err)
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		time.Sleep(timeout)
		return connect(dsn, tyrCount-1, timeout)
	}

	return connPool, nil
}

package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/mechatron-x/8here/internal/config"
	"go.uber.org/zap"
)

const (
	defaultConnTimeout = time.Second * 2
)

func Connect(dbConf config.DB, log *zap.Logger) (*sql.DB, error) {
	timeout, err := time.ParseDuration(dbConf.Timeout)
	if err != nil {
		timeout = defaultConnTimeout
	}

	db, err := connect(dbConf.Driver, dbConf.DSN, dbConf.TryCount, timeout)
	if err != nil {
		return nil, err
	}

	log.Info("DB connection established", zap.String("address", dbConf.DSN))

	return db, nil
}

func connect(driver, dsn string, tyrCount uint, timeout time.Duration) (*sql.DB, error) {
	if tyrCount <= 0 {
		return nil, errors.New("connection failed, all retries exhausted")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("invalid dsn configuration for driver: %s", driver)
	}

	err = db.Ping()
	if err != nil {
		time.Sleep(timeout)
		return connect(driver, dsn, tyrCount-1, timeout)
	}

	return db, nil
}

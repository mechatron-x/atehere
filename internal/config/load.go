package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mechatron-x/atehere/internal/core"
)

type EnvKey string

func (ek EnvKey) String() string {
	return string(ek)
}

const (
	FIREBASE_PRIVATE_KEY EnvKey = "FIREBASE_PRIVATE_KEY"
	DB_USER              EnvKey = "POSTGRES_USER"
	DB_PASSWORD          EnvKey = "POSTGRES_PASSWORD"
	DB_NAME              EnvKey = "POSTGRES_DB"
)

func Load(confPath string) (*App, error) {
	conf := new(App)

	confPath = filepath.Clean(confPath)
	bytes, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %v", err)
	}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %v", err)
	}

	envMap, err := loadEnv(FIREBASE_PRIVATE_KEY, DB_USER, DB_PASSWORD, DB_NAME)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %v", err)
	}

	conf.Firebase.PrivateKey = envMap[FIREBASE_PRIVATE_KEY]
	conf.DB.User = envMap[DB_USER]
	conf.DB.Password = envMap[DB_PASSWORD]
	conf.DB.Name = envMap[DB_NAME]

	return conf, nil
}

func loadEnv(keys ...EnvKey) (map[EnvKey]string, error) {
	envMap := make(map[EnvKey]string)

	for _, key := range keys {
		val, ok := os.LookupEnv(key.String())
		if !ok {
			return nil, fmt.Errorf("environment variable %s is missing; please ensure it is set", key.String())
		}
		if core.IsEmptyString(val) {
			return nil, fmt.Errorf("environment variable %s is set but empty; please provide a valid value", key.String())
		}

		envMap[key] = val
	}

	return envMap, nil
}

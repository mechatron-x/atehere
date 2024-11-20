package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type EnvKey string

func (ek EnvKey) String() string {
	return string(ek)
}

const (
	FIREBASE_PRIVATE_KEY EnvKey = "FIREBASE_PRIVATE_KEY"
	DB_USER              EnvKey = "POSTGRES_USER"
	DB_KEY               EnvKey = "POSTGRES_PASSWORD"
	DB_NAME              EnvKey = "POSTGRES_DB"
	APP_ENV              EnvKey = "APP_ENV"
)

var envKeys []EnvKey = []EnvKey{
	FIREBASE_PRIVATE_KEY,
	DB_USER,
	DB_KEY,
	DB_NAME,
	APP_ENV,
}

func Load(confPath string) (*App, error) {
	conf := new(App)

	confPath = filepath.Clean(confPath)
	bytes, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	envMap, err := loadEnv(envKeys...)
	if err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	conf.Firebase.PrivateKey = envMap[FIREBASE_PRIVATE_KEY]

	conf.DB.User = envMap[DB_USER]
	conf.DB.Password = envMap[DB_KEY]
	conf.DB.Name = envMap[DB_NAME]

	conf.Environment = parseEnvironment(envMap[APP_ENV])

	return conf, nil
}

func loadEnv(keys ...EnvKey) (map[EnvKey]string, error) {
	envMap := make(map[EnvKey]string)

	for _, key := range keys {
		val, ok := os.LookupEnv(key.String())
		if !ok {
			return nil, fmt.Errorf("environment variable %s is missing; please ensure it is set", key.String())
		}

		envMap[key] = val
	}

	return envMap, nil
}

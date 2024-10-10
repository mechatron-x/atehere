package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Load(confPath string) (*App, error) {
	conf := new(App)

	if !filepath.IsLocal(confPath) {
		return nil, errors.New("config.Load: config filepath is not local")
	}

	confPath = filepath.Clean(confPath)
	bytes, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %v", err)
	}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %v", err)
	}

	return conf, nil
}

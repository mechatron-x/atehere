package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

	return conf, nil
}

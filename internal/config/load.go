package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(confPath string) *App {
	conf := new(App)

	bytes, err := os.ReadFile(confPath)
	if err != nil {
		panic(fmt.Sprintf("config.Load: %v", err))
	}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		panic(fmt.Sprintf("config.Load: %v", err))
	}

	return conf
}

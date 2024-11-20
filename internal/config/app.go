package config

import "go.uber.org/zap"

type App struct {
	Environment Environment `json:"environment"`
	Logger      zap.Config  `json:"logger"`
	Api         Api         `json:"api"`
	DB          DB          `json:"db"`
	Firebase    Firebase    `json:"firebase"`
}

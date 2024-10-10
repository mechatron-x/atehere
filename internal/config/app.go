package config

import "go.uber.org/zap"

type Environment string

type App struct {
	Environment Environment `json:"environment"`
	Logger      zap.Config  `json:"logger"`
	Api         Api         `json:"api"`
	DB          DB          `json:"db"`
}

const (
	Dev  Environment = "dev"
	Qa   Environment = "qa"
	Prod Environment = "prod"
)

func (rcv Environment) String() string {
	return string(rcv)
}

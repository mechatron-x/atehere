package config

import "strings"

type Environment int

const (
	DEV Environment = iota
	PROD
)

const (
	defaultEnv Environment = PROD
)

func parseEnvironment(env string) Environment {
	env = strings.TrimSpace(env)

	switch env {
	case "dev", "development", "local":
		return DEV
	case "prod", "live":
		return PROD
	default:
		return defaultEnv
	}
}

func (e Environment) String() string {
	switch e {
	case DEV:
		return "dev"
	case PROD:
		return "prod"
	default:
		return defaultEnv.String()
	}
}

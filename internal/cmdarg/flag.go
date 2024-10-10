package cmdarg

import (
	"flag"
	"os"
)

type Flags struct {
	ConfPath string
}

func Setup() Flags {
	var confPath string

	flag.StringVar(&confPath, "c", os.Getenv("8HERE_CONF_PATH"), "configuration file path")
	flag.Parse()

	return Flags{
		ConfPath: confPath,
	}
}

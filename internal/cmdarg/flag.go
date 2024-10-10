package cmdarg

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
)

type Flags struct {
	ConfPath string
}

func Setup() (Flags, error) {
	var confPath string

	flag.StringVar(&confPath, "c", os.Getenv("8HERE_CONF_PATH"), "path of the configuration location")
	flag.Parse()

	confPath = filepath.Clean(confPath)
	if !filepath.IsLocal(confPath) {
		return Flags{}, errors.New("cmdarg.Setup: config filepath is not local")
	}

	return Flags{
		ConfPath: confPath,
	}, nil
}

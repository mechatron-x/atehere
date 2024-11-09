package infrastructure

import (
	"os"
	"path/filepath"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
)

type DiskFileSaver struct{}

func (dfs DiskFileSaver) Save(filename string, data []byte, conf config.Api) error {
	savePath := filepath.Join(conf.StaticRoot, filepath.Clean(filename))

	imageFile, err := os.Create(savePath)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}
	defer imageFile.Close()

	err = os.WriteFile(savePath, data, 0600)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}

	return nil
}

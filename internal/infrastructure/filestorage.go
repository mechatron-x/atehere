package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mechatron-x/atehere/internal/config"
)

type DiskFileSaver struct{}

func (dfs DiskFileSaver) Save(filename string, data []byte, conf config.Api) error {
	savePath := filepath.Join(conf.StaticRoot, filepath.Clean(filename))

	savePath = filepath.Clean(savePath)
	if !strings.HasPrefix(savePath, conf.StaticRoot) {
		return fmt.Errorf("invalid static root")
	}

	imageFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	err = os.WriteFile(savePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DiskFileSaver struct{}

func NewDiskFileSaver() DiskFileSaver {
	return DiskFileSaver{}
}

func (dfs DiskFileSaver) Save(filename string, data []byte, location string) error {
	savePath := filepath.Join(location, filepath.Clean(filename))

	savePath = filepath.Clean(savePath)
	if !strings.HasPrefix(savePath, location) {
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

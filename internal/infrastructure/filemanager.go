package infrastructure

import (
	"os"
)

type DiskFileManager struct{}

func NewDiskFileManager() DiskFileManager {
	return DiskFileManager{}
}

func (dfs DiskFileManager) Save(savePath string, data []byte) error {
	imageFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	err = os.WriteFile(savePath, data, os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

func (fds DiskFileManager) Delete(deletePath string) error {
	return os.Remove(deletePath)
}

func (fds DiskFileManager) Read(readPath string) ([]byte, error) {
	return os.ReadFile(readPath)
}

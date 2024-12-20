package storage

import (
	"errors"
	"os"
	"path/filepath"
)

type FileStorage struct{}

func NewFile() *FileStorage {
	return &FileStorage{}
}

func (rcv *FileStorage) Save(savePath string, data []byte) error {
	savePath = filepath.Clean(savePath)

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

func (rcv *FileStorage) Delete(deletePath string) error {
	deletePath = filepath.Clean(deletePath)

	return os.Remove(deletePath)
}

func (rcv *FileStorage) Read(readPath string) ([]byte, error) {
	readPath = filepath.Clean(readPath)

	fileInfo, err := os.Stat(readPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []byte{}, nil
		}

		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, errors.New("cannot read from directory")
	}

	return os.ReadFile(readPath)
}

package infrastructure

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type (
	ImageStorage struct {
		location string
		fs       FileSaver
	}

	FileSaver interface {
		Save(savePath string, data []byte) error
	}
)

func NewImageStorage(fileSaver FileSaver, location string) (*ImageStorage, error) {
	_, err := os.ReadDir(location)
	if err != nil {
		return nil, fmt.Errorf("cannot read image storage location, reason: %v", err)
	}

	return &ImageStorage{
		location: location,
		fs:       fileSaver,
	}, nil
}

func (is *ImageStorage) Save(filename string, data string) (string, error) {
	imageDecoded, imageType, err := decodeBase64Image(data)
	if err != nil {
		return "", err
	}

	imageName := fmt.Sprintf("%s.%s", filename, imageType)
	savePath, err := createAbsPath(is.location, imageName)
	if err != nil {
		return "", err
	}

	err = is.fs.Save(savePath, imageDecoded)
	if err != nil {
		return "", err
	}

	return imageName, nil
}

func decodeBase64Image(base64Str string) ([]byte, string, error) {
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %v", err)
	}

	fileExtension, err := fileType(imageData)
	if err != nil {
		return nil, "", fmt.Errorf("failed to detect image type: %v", err)
	}

	return imageData, fileExtension, nil
}

func fileType(data []byte) (string, error) {
	mimeType := http.DetectContentType(data)

	extension := ""
	switch mimeType {
	case "image/jpeg":
		extension = "jpg"
	case "image/png":
		extension = "png"
	case "image/gif":
		extension = "gif"
	case "image/webp":
		extension = "webp"
	default:
		return "", fmt.Errorf("unsupported file type: %s", mimeType)
	}

	return extension, nil
}

func createAbsPath(location, filename string) (string, error) {
	savePath := filepath.Join(location, filepath.Clean(filename))
	savePath = filepath.Clean(savePath)
	if !strings.HasPrefix(savePath, location) {
		return "", fmt.Errorf("invalid static root")
	}

	return savePath, nil
}

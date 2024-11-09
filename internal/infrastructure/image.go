package infrastructure

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/mechatron-x/atehere/internal/config"
)

type (
	ImageStorage struct {
		conf config.Api
		fs   FileSaver
	}

	FileSaver interface {
		Save(filename string, data []byte, conf config.Api) error
	}
)

func NewImageStorage(conf config.Api) *ImageStorage {
	return &ImageStorage{
		conf: conf,
		fs:   DiskFileSaver{},
	}
}

func (is *ImageStorage) Save(fileName string, data string) (string, error) {
	imageDecoded, imageType, err := decodeBase64Image(data)
	if err != nil {
		return "", err
	}

	imageName := fmt.Sprintf("%s.%s", fileName, imageType)

	err = is.fs.Save(imageName, imageDecoded, is.conf)
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

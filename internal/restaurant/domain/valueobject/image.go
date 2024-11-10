package valueobject

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mechatron-x/atehere/internal/core"
)

type Image struct {
	name      string
	extension string
}

const (
	defaultImageName      string = "default_restaurant_image"
	defaultImageExtension string = "png"
	imageSeparator               = "."
)

func NewImage(image string) (Image, error) {
	if core.IsEmptyString(image) {
		return Image{
			name:      defaultImageName,
			extension: defaultImageExtension,
		}, nil
	}

	parts := strings.Split(image, imageSeparator)
	if len(parts) < 2 {
		return Image{}, errors.New("invalid image name format")
	}

	return Image{
		name:      parts[0],
		extension: parts[1],
	}, nil
}

func (in Image) Name() string {
	return in.name
}

func (in Image) Extension() string {
	return in.extension
}

func (i Image) String() string {
	return fmt.Sprintf("%s%s%s", i.name, imageSeparator, i.extension)
}

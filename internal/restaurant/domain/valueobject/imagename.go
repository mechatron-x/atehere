package valueobject

type ImageName struct {
	name string
}

func NewImageName(name string) (ImageName, error) {
	return ImageName{name}, nil
}

func (i ImageName) Name() string {
	return i.name
}

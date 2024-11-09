package valueobject

type ImageName struct {
	name string
}

func NewImageName(name string) ImageName {
	return ImageName{name}
}

func (i ImageName) String() string {
	return i.name
}

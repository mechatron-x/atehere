package valueobject

import (
	"errors"
	"time"
)

type BirthDate struct {
	date time.Time
}

const (
	BirthDateLayout = "02-01-2006" // dd-mm-yyyy
)

func NewBirthDate(date string) (BirthDate, error) {
	birthDate, err := time.Parse(BirthDateLayout, date)
	if err != nil {
		return BirthDate{}, err
	}

	if birthDate.Compare(time.Now()) > 0 {
		return BirthDate{time.Time{}}, errors.New("birth date cannot be in future")
	}

	return BirthDate{birthDate}, nil
}

func (b BirthDate) CalculateAge() int {
	birthDate := b.date
	return time.Now().Year() - birthDate.Year()
}

func (b BirthDate) Date() time.Time {
	return b.date
}

func (b BirthDate) String() string {
	return b.date.Format(BirthDateLayout)
}

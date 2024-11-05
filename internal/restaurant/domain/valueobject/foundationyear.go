package valueobject

import (
	"errors"
	"time"
)

type FoundationYear struct {
	date time.Time
}

const (
	FoundationYearLayout = "2006"
	FoundationYearFormat = "yyyy"
)

func NewFoundationYear(year string) (FoundationYear, error) {
	foundationYear, err := time.Parse(FoundationYearLayout, year)
	if err != nil {
		return FoundationYear{}, err
	}

	if foundationYear.Compare(time.Now()) > 0 {
		return FoundationYear{time.Time{}}, errors.New("foundation year cannot be in future")
	}

	return FoundationYear{foundationYear}, nil
}

func (f FoundationYear) YearsSinceFoundation() int {
	foundationDate := f.date
	return time.Now().Year() - foundationDate.Year()
}

func (f FoundationYear) Date() time.Time {
	return f.date
}

func (f FoundationYear) String() string {
	return f.date.Format(FoundationYearLayout)
}

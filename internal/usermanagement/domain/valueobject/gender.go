package valueobject

import "strings"

type Gender int64

const (
	Undefined Gender = iota
	Male
	Female
	Other
)

func ParseGender(gender string) Gender {
	gender = strings.TrimSpace(gender)
	gender = strings.ToLower(gender)

	switch gender {
	case "male":
		return Male
	case "female":
		return Female
	case "other":
		return Other
	default:
		return Undefined
	}
}

func GetGenders() []string {
	return []string{
		Undefined.String(),
		Male.String(),
		Female.String(),
		Other.String(),
	}
}

func (g Gender) String() string {
	switch g {
	case 1:
		return "MALE"
	case 2:
		return "FEMALE"
	case 3:
		return "OTHER"
	default:
		return "UNDEFINED"
	}
}

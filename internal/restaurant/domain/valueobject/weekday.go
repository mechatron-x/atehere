package valueobject

import (
	"fmt"
	"strings"
	"time"
)

func ParseWeekday(weekday string) (time.Weekday, error) {
	weekday = strings.TrimSpace(weekday)
	weekday = strings.ToLower(weekday)

	switch weekday {
	case "sunday":
		return time.Sunday, nil
	case "monday":
		return time.Monday, nil
	case "tuesday":
		return time.Tuesday, nil
	case "wednesday":
		return time.Wednesday, nil
	case "thursday":
		return time.Thursday, nil
	case "friday":
		return time.Friday, nil
	case "saturday":
		return time.Saturday, nil
	}

	return -1, fmt.Errorf("invalid weekday %s", weekday)
}

func AvailableWeekdays() []string {
	return []string{
		time.Sunday.String(),
		time.Monday.String(),
		time.Tuesday.String(),
		time.Wednesday.String(),
		time.Thursday.String(),
		time.Friday.String(),
		time.Saturday.String(),
	}
}

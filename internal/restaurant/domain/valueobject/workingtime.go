package valueobject

import (
	"fmt"
	"time"
)

type WorkTime struct {
	time time.Time
}

const (
	WorkTimeLayout    = "15:04"
	WorkingTimeFormat = "hh-mm"
)

func NewWorkTime(workTime string) (WorkTime, error) {
	wt, err := time.Parse(WorkTimeLayout, workTime)
	if err != nil {
		return WorkTime{}, fmt.Errorf("invalid work time %v", err)
	}

	return WorkTime{wt}, nil
}

func (b WorkTime) CalculateWorkingHours(workTime WorkTime) int {
	wt1 := b.time
	wt2 := workTime.time

	duration := wt2.Sub(wt1)
	hours := int(duration.Hours())

	return hours
}

func (b WorkTime) Time() time.Time {
	return b.time
}

func (b WorkTime) String() string {
	return b.time.Format(WorkTimeLayout)
}

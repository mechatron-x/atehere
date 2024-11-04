package aggregate

import (
	"slices"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
)

type Restaurant struct {
	core.Aggregate
	owner          *entity.Owner
	name           valueobject.RestaurantName
	foundationDate valueobject.FoundationYear
	phoneNumber    valueobject.PhoneNumber
	openingTime    valueobject.WorkTime
	closingTime    valueobject.WorkTime
	workingDays    []time.Weekday
}

func NewRestaurant() *Restaurant {
	return &Restaurant{
		Aggregate: core.NewAggregate(),
	}
}

func (r *Restaurant) Owner() *entity.Owner {
	return r.owner
}

func (r *Restaurant) Name() valueobject.RestaurantName {
	return r.name
}

func (r *Restaurant) FoundationDate() valueobject.FoundationYear {
	return r.foundationDate
}

func (r *Restaurant) PhoneNumber() valueobject.PhoneNumber {
	return r.phoneNumber
}

func (r *Restaurant) OpeningTime() valueobject.WorkTime {
	return r.openingTime
}

func (r *Restaurant) ClosingTime() valueobject.WorkTime {
	return r.closingTime
}

func (r *Restaurant) WorkingDays() []time.Weekday {
	return r.workingDays
}

func (r *Restaurant) SetOwner(owner *entity.Owner) {
	r.owner = owner
}

func (r *Restaurant) SetName(name valueobject.RestaurantName) {
	r.name = name
}

func (r *Restaurant) SetFoundationDate(foundationDate valueobject.FoundationYear) {
	r.foundationDate = foundationDate
}

func (r *Restaurant) SetPhoneNumber(phoneNumber valueobject.PhoneNumber) {
	r.phoneNumber = phoneNumber
}

func (r *Restaurant) SetOpeningTime(openingTime valueobject.WorkTime) {
	r.openingTime = openingTime
}

func (r *Restaurant) SetClosingTime(closingTime valueobject.WorkTime) {
	r.closingTime = closingTime
}

func (r *Restaurant) AddWorkingDays(workingDays ...time.Weekday) {
	for _, workingDay := range workingDays {
		r.addWorkingDay(workingDay)
	}
}

func (r *Restaurant) addWorkingDay(workingDay time.Weekday) {
	if slices.Contains(r.workingDays, workingDay) {
		return
	}

	r.workingDays = append(r.workingDays, workingDay)
}

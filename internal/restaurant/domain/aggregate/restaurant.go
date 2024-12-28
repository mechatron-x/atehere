package aggregate

import (
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
)

type Restaurant struct {
	core.Aggregate
	ownerID        uuid.UUID
	name           valueobject.RestaurantName
	foundationYear valueobject.FoundationYear
	phoneNumber    valueobject.PhoneNumber
	openingTime    valueobject.WorkTime
	closingTime    valueobject.WorkTime
	workingDays    []time.Weekday
	imageName      valueobject.Image
	tables         []entity.Table
	locations      valueobject.Locations
}

func NewRestaurant() *Restaurant {
	return &Restaurant{
		Aggregate:   core.NewAggregate(),
		workingDays: make([]time.Weekday, 0),
		tables:      make([]entity.Table, 0),
		locations:   make(valueobject.Locations, 0),
	}
}

func (r *Restaurant) OwnerID() uuid.UUID {
	return r.ownerID
}

func (r *Restaurant) Name() valueobject.RestaurantName {
	return r.name
}

func (r *Restaurant) FoundationYear() valueobject.FoundationYear {
	return r.foundationYear
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

func (r *Restaurant) ImageName() valueobject.Image {
	return r.imageName
}

func (r *Restaurant) Tables() []entity.Table {
	return r.tables
}

func (r *Restaurant) Locations() valueobject.Locations {
	return r.locations
}

func (r *Restaurant) SetOwner(ownerID uuid.UUID) {
	r.ownerID = ownerID
}

func (r *Restaurant) SetName(name valueobject.RestaurantName) {
	r.name = name
}

func (r *Restaurant) SetFoundationYear(foundationDate valueobject.FoundationYear) {
	r.foundationYear = foundationDate
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

func (r *Restaurant) AddTables(tables ...entity.Table) {
	for _, t := range tables {
		r.addTable(t)
	}
}

func (r *Restaurant) AddLocations(locations ...valueobject.Location) {
	r.locations = append(r.locations, locations...)
}

func (r *Restaurant) SetImageName(imageName valueobject.Image) {
	r.imageName = imageName
}

func (r *Restaurant) IsOwner(ownerID uuid.UUID) bool {
	return r.ownerID == ownerID
}

func (r *Restaurant) IsInRadius(target valueobject.Location, radius float64) bool {
	return r.locations.IsInRadius(target, radius)
}

func (r *Restaurant) DeleteNow() {
	r.tables = []entity.Table{}
	r.SetDeletedAt(time.Now())
}

func (r *Restaurant) addWorkingDay(workingDay time.Weekday) {
	if slices.Contains(r.workingDays, workingDay) {
		return
	}

	r.workingDays = append(r.workingDays, workingDay)
}

func (r *Restaurant) addTable(table entity.Table) {
	for _, t := range r.tables {
		if t.ID() == table.ID() {
			return
		}
	}

	r.tables = append(r.tables, table)
}

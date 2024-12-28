package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"gorm.io/gorm"
)

type Restaurant struct{}

func (r Restaurant) FromModel(model *model.Restaurant) (*aggregate.Restaurant, error) {
	restaurant := aggregate.NewRestaurant()

	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	ownerID, err := uuid.Parse(model.OwnerID)
	if err != nil {
		return nil, err
	}

	name, err := valueobject.NewRestaurantName(model.Name)
	if err != nil {
		return nil, err
	}

	foundationYear, err := valueobject.NewFoundationYear(model.FoundationYear)
	if err != nil {
		return nil, err
	}

	phoneNumber, err := valueobject.NewPhoneNumber(model.PhoneNumber)
	if err != nil {
		return nil, err
	}

	openingTime, err := valueobject.NewWorkTime(model.OpeningTime)
	if err != nil {
		return nil, err
	}

	closingTime, err := valueobject.NewWorkTime(model.ClosingTime)
	if err != nil {
		return nil, err
	}

	workingDays := make([]time.Weekday, 0)
	for _, wd := range model.WorkingDays {
		workingDay, err := valueobject.ParseWeekday(wd)
		if err != nil {
			return nil, err
		}

		workingDays = append(workingDays, workingDay)
	}

	imageName, err := valueobject.NewImage(model.ImageName)
	if err != nil {
		return nil, err
	}

	tables := make([]entity.Table, 0)
	for _, t := range model.Tables {
		id, err := uuid.Parse(t.ID)
		if err != nil {
			return nil, err
		}

		tableName, err := valueobject.NewTableName(t.Name)
		if err != nil {
			return nil, err
		}

		table := entity.NewTable()
		table.SetID(id)
		table.SetName(tableName)
		table.SetCreatedAt(t.CreatedAt)
		table.SetUpdatedAt(t.UpdatedAt)

		tables = append(tables, table)
	}

	locations := make(valueobject.Locations, 0)
	for _, l := range model.Locations {
		location, err := valueobject.NewLocation(l.Latitude, l.Longitude)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	restaurant.SetID(id)
	restaurant.SetOwner(ownerID)
	restaurant.SetName(name)
	restaurant.SetFoundationYear(foundationYear)
	restaurant.SetPhoneNumber(phoneNumber)
	restaurant.SetOpeningTime(openingTime)
	restaurant.SetClosingTime(closingTime)
	restaurant.AddWorkingDays(workingDays...)
	restaurant.SetImageName(imageName)
	restaurant.AddTables(tables...)
	restaurant.AddLocations(locations...)
	restaurant.SetCreatedAt(model.CreatedAt)
	restaurant.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		restaurant.SetDeletedAt(model.DeletedAt.Time)
	}

	return restaurant, nil
}

func (r Restaurant) FromModels(models []model.Restaurant) ([]*aggregate.Restaurant, error) {
	aggregates := make([]*aggregate.Restaurant, 0)

	for _, m := range models {
		aggregate, err := r.FromModel(&m)
		if err != nil {
			return nil, err
		}

		aggregates = append(aggregates, aggregate)
	}

	return aggregates, nil
}

func (r Restaurant) FromAggregate(aggregate *aggregate.Restaurant) *model.Restaurant {
	workingDays := make(pq.StringArray, 0)
	for _, wd := range aggregate.WorkingDays() {
		workingDays = append(workingDays, wd.String())
	}

	tables := make([]model.RestaurantTable, 0)
	for _, t := range aggregate.Tables() {
		table := model.RestaurantTable{
			ID:           t.ID().String(),
			RestaurantID: aggregate.ID().String(),
			Name:         t.Name().String(),
		}

		tables = append(tables, table)
	}

	locations := make([]model.RestaurantLocation, 0)
	for _, l := range aggregate.Locations() {
		location := model.RestaurantLocation{
			RestaurantID: aggregate.ID().String(),
			Latitude:     l.Lat(),
			Longitude:    l.Long(),
		}

		locations = append(locations, location)
	}

	return &model.Restaurant{
		ID:             aggregate.ID().String(),
		OwnerID:        aggregate.OwnerID().String(),
		Name:           aggregate.Name().String(),
		FoundationYear: aggregate.FoundationYear().String(),
		PhoneNumber:    aggregate.PhoneNumber().String(),
		OpeningTime:    aggregate.OpeningTime().String(),
		ClosingTime:    aggregate.ClosingTime().String(),
		WorkingDays:    workingDays,
		ImageName:      aggregate.ImageName().String(),
		Tables:         tables,
		Locations:      locations,
		Model: gorm.Model{
			CreatedAt: aggregate.CreatedAt(),
			UpdatedAt: aggregate.UpdatedAt(),
			DeletedAt: gorm.DeletedAt{
				Time:  aggregate.DeletedAt(),
				Valid: aggregate.IsDeleted(),
			},
		},
	}
}

func (r Restaurant) FromAggregates(aggregates []*aggregate.Restaurant) []*model.Restaurant {
	models := make([]*model.Restaurant, 0)
	for _, a := range aggregates {
		models = append(models, r.FromAggregate(a))
	}

	return models
}

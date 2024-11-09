package mapper

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
)

type (
	Restaurant struct{}
)

func NewRestaurant() Restaurant {
	return Restaurant{}
}

func (rm Restaurant) FromModel(model dal.Restaurant, tables ...dal.RestaurantTable) (*aggregate.Restaurant, error) {
	restaurant := aggregate.NewRestaurant()

	verifiedName, err := valueobject.NewRestaurantName(model.Name)
	if err != nil {
		return nil, err
	}

	verifiedFoundationYear, err := valueobject.NewFoundationYear(model.FoundationYear.String)
	if err != nil {
		return nil, err
	}

	verifiedPhoneNumber, err := valueobject.NewPhoneNumber(model.PhoneNumber.String)
	if err != nil {
		return nil, err
	}

	verifiedOpeningTime, err := valueobject.NewWorkTime(model.OpeningTime)
	if err != nil {
		return nil, err
	}

	verifiedClosingTime, err := valueobject.NewWorkTime(model.ClosingTime)
	if err != nil {
		return nil, err
	}

	verifiedImageName, err := valueobject.NewImage(model.ImageName.String)
	if err != nil {
		return nil, err
	}

	for _, workingDay := range model.WorkingDays {
		verifiedWorkingDay, err := valueobject.ParseWeekday(workingDay)
		if err != nil {
			return nil, err
		}

		restaurant.AddWorkingDays(verifiedWorkingDay)
	}

	for _, t := range tables {
		verifiedTableName, err := valueobject.NewTableName(t.Name)
		if err != nil {
			return nil, err
		}

		table := entity.NewTable()
		table.SetID(model.ID)
		table.SetName(verifiedTableName)
		table.SetCreatedAt(model.CreatedAt)
		table.SetUpdatedAt(model.UpdatedAt)

		restaurant.AddTables(table)
	}

	restaurant.SetID(model.ID)
	restaurant.SetOwner(model.OwnerID)
	restaurant.SetName(verifiedName)
	restaurant.SetFoundationYear(verifiedFoundationYear)
	restaurant.SetPhoneNumber(verifiedPhoneNumber)
	restaurant.SetOpeningTime(verifiedOpeningTime)
	restaurant.SetClosingTime(verifiedClosingTime)
	restaurant.SetImageName(verifiedImageName)
	restaurant.SetCreatedAt(model.CreatedAt)
	restaurant.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		restaurant.SetDeletedAt(model.DeletedAt.Time)
	}

	return restaurant, nil
}

func (rm Restaurant) FromAggregate(restaurant *aggregate.Restaurant) dal.Restaurant {
	workingDays := make([]string, 0)

	for _, wd := range restaurant.WorkingDays() {
		workingDays = append(workingDays, wd.String())
	}

	return dal.Restaurant{
		ID:      restaurant.ID(),
		OwnerID: restaurant.OwnerID(),
		Name:    restaurant.Name().String(),
		FoundationYear: sql.NullString{
			String: restaurant.FoundationYear().String(),
			Valid:  true,
		},
		PhoneNumber: sql.NullString{
			String: restaurant.PhoneNumber().String(),
			Valid:  true,
		},
		OpeningTime: restaurant.OpeningTime().String(),
		ClosingTime: restaurant.ClosingTime().String(),
		WorkingDays: workingDays,
		ImageName: sql.NullString{
			String: restaurant.ImageName().String(),
			Valid:  true,
		},
		CreatedAt: restaurant.CreatedAt(),
		UpdatedAt: restaurant.UpdatedAt(),
		DeletedAt: sql.NullTime{
			Time:  restaurant.DeletedAt(),
			Valid: restaurant.IsDeleted(),
		},
	}
}

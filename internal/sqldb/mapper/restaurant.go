package mapper

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
)

type (
	Restaurant struct{}
)

func NewRestaurant() Restaurant {
	return Restaurant{}
}

func (rm Restaurant) FromModel(model dal.Restaurant) (*aggregate.Restaurant, error) {
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

	for _, workingDay := range model.WorkingDays {
		verifiedWorkingDay, err := valueobject.ParseWeekday(workingDay)
		if err != nil {
			return nil, err
		}

		restaurant.AddWorkingDays(verifiedWorkingDay)
	}

	restaurant.SetID(model.ID)
	restaurant.SetOwner(model.OwnerID)
	restaurant.SetName(verifiedName)
	restaurant.SetFoundationYear(verifiedFoundationYear)
	restaurant.SetPhoneNumber(verifiedPhoneNumber)
	restaurant.SetOpeningTime(verifiedOpeningTime)
	restaurant.SetClosingTime(verifiedClosingTime)
	if model.ImageName.Valid {
		restaurant.SetImageName(model.ImageName.String)
	}
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
			String: restaurant.ImageName(),
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

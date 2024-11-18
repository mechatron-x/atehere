package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Restaurant struct {
		gorm.Model
		ID             string `gorm:"primarykey"`
		OwnerID        string
		Name           string
		FoundationYear string
		PhoneNumber    string
		OpeningTime    string
		ClosingTime    string
		ImageName      string
		WorkingDays    []RestaurantWorkingDay `gorm:"constraint:OnDelete:CASCADE"`
		Tables         []RestaurantTable      `gorm:"constraint:OnDelete:CASCADE"`
		Menus          []Menu                 `gorm:"constraint:OnDelete:CASCADE"`
	}

	RestaurantTable struct {
		ID           string `gorm:"primarykey"`
		RestaurantID string
		Name         string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	RestaurantWorkingDay struct {
		RestaurantID string `gorm:"primaryKey"`
		Day          string `gorm:"primaryKey"`
	}
)

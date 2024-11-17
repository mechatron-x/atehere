package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Restaurant struct {
		gorm.Model
		ID             string `gorm:"primarykey"`
		OwnerID        string `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ManagerID"`
		Name           string
		FoundationYear string
		PhoneNumber    string
		OpeningTime    string
		ClosingTime    string
		WorkingDays    []RestaurantWorkingDay
		ImageName      string
		Tables         []RestaurantTable
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

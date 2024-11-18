package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Restaurant struct {
		ID             string `gorm:"primarykey"`
		OwnerID        string
		Name           string
		FoundationYear string
		PhoneNumber    string
		OpeningTime    string
		ClosingTime    string
		ImageName      string
		WorkingDays    pq.StringArray `gorm:"type:text[]"`
		Tables         []RestaurantTable
		Menus          []Menu
		gorm.Model
	}

	RestaurantTable struct {
		ID           string `gorm:"primarykey"`
		RestaurantID string
		Name         string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Restaurant struct {
		ID             string `gorm:"primarykey;type:uuid"`
		OwnerID        string
		Owner          Manager        `gorm:"foreignKey:OwnerID"`
		Name           string         `gorm:"not null;size:100;unique"`
		FoundationYear string         `gorm:"not null;size:20"`
		PhoneNumber    string         `gorm:"not null;size:255"`
		OpeningTime    string         `gorm:"not null;size:5"`
		ClosingTime    string         `gorm:"not null;size:5"`
		ImageName      string         `gorm:"not null;size:50"`
		WorkingDays    pq.StringArray `gorm:"type:text[];size:10"`
		Tables         []RestaurantTable
		Locations      []RestaurantLocation
		gorm.Model
	}

	RestaurantTable struct {
		ID           string `gorm:"primarykey;type:uuid"`
		RestaurantID string `gorm:"type:uuid"`
		Name         string `gorm:"not null;size:10"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	RestaurantLocation struct {
		RestaurantID string  `gorm:"primarykey;type:uuid"`
		Latitude     float64 `gorm:"primarykey;numeric;precision:10;scale:8"`
		Longitude    float64 `gorm:"primarykey;numeric;precision:11;scale:8"`
	}
)

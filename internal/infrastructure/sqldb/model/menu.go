package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Menu struct {
		ID           string `gorm:"primarykey;type:uuid"`
		RestaurantID string
		Restaurant   Restaurant
		Category     string `gorm:"not null;size:100"`
		MenuItems    []MenuItem
		gorm.Model
	}

	MenuItem struct {
		ID                 string `gorm:"primarykey;type:uuid"`
		MenuID             string
		Name               string `gorm:"not null;size:100"`
		Description        string `gorm:"not null;text"`
		ImageName          string
		PriceAmount        float64        `gorm:"not null;numeric;precision:10;scale:2"`
		PriceCurrency      string         `gorm:"not null;size:5"`
		DiscountPercentage int            `gorm:"not null"`
		Ingredients        pq.StringArray `gorm:"type:text[]"`
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
)

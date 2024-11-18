package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Menu struct {
		ID           string `gorm:"primarykey;"`
		RestaurantID string
		Category     string `gorm:"not null"`
		MenuItems    []MenuItem
		gorm.Model
	}

	MenuItem struct {
		ID                 string `gorm:"primarykey"`
		MenuID             string
		Name               string `gorm:"not null"`
		Description        string `gorm:"not null"`
		ImageName          string
		PriceAmount        float64        `gorm:"not null"`
		PriceCurrency      string         `gorm:"not null"`
		DiscountPercentage int            `gorm:"not null"`
		Ingredients        pq.StringArray `gorm:"type:text[]"`
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
)

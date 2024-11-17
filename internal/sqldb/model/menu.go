package model

import "gorm.io/gorm"

type (
	Menu struct {
		gorm.Model
		ID           string `gorm:"primarykey"`
		RestaurantID string `gorm:"not null"`
		Category     string `gorm:"not null"`
		MenuItems    []MenuItem
	}

	MenuItem struct {
		gorm.Model
		ID                 string `gorm:"primarykey"`
		MenuID             string `gorm:"primarykey"`
		Name               string `gorm:"not null"`
		Description        string `gorm:"not null"`
		ImageName          string
		PriceAmount        float64  `gorm:"not null"`
		PriceCurrency      string   `gorm:"not null"`
		DiscountPercentage int      `gorm:"not null"`
		Ingredients        []string `gorm:"type:text[]"`
	}
)

package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Session struct {
		ID        string `gorm:"primarykey;type:uuid"`
		TableID   string
		Table     RestaurantTable `gorm:"foreignKey:TableID"`
		StartTime time.Time
		EndTime   time.Time
		Orders    []SessionOrder
		gorm.Model
	}

	SessionOrder struct {
		ID         string `gorm:"primarykey;type:uuid"`
		SessionID  string
		MenuItemID string
		MenuItem   MenuItem
		OrderedBy  string
		Customer   Customer `gorm:"foreignKey:OrderedBy"`
		Quantity   int      `gorm:"not null"`
	}
)

package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Bill struct {
		ID        string `gorm:"primarykey;type:uuid"`
		SessionID string `gorm:"unique;type:uuid"`
		BillItems []BillItem
		gorm.Model
	}

	BillItem struct {
		ID        string `gorm:"primarykey;type:uuid"`
		BillID    string
		OwnerID   string `gorm:"type:uuid"`
		ItemName  string
		Quantity  int
		UnitPrice float64 `gorm:"not null;numeric;precision:10;scale:2"`
		PaidPrice float64 `gorm:"not null;numeric;precision:10;scale:2"`
		Currency  string  `gorm:"not null;size:5"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

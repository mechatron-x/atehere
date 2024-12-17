package model

import (
	"time"

	"github.com/lib/pq"
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
		ID           string `gorm:"primarykey;type:uuid"`
		BillID       string `gorm:"type:uuid"`
		OwnerID      string `gorm:"type:uuid"`
		ItemName     string
		UnitPrice    float64        `gorm:"not null;numeric;precision:10;scale:2"`
		Currency     string         `gorm:"not null;size:5"`
		Quantity     int            `gorm:"not null"`
		PaidQuantity int            `gorm:"not null;numeric;"`
		PaidBy       pq.StringArray `gorm:"type:text[]"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

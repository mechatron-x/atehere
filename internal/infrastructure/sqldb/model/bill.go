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
		BillID    string `gorm:"not null;type:uuid"`
		OwnerID   string `gorm:"not null;type:uuid"`
		ItemName  string
		UnitPrice float64 `gorm:"not null;numeric;precision:10;scale:2"`
		Currency  string  `gorm:"not null;size:5"`
		Quantity  int     `gorm:"not null"`
		Payments  []BillItemPayments
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	BillItemPayments struct {
		CustomerID string  `gorm:"primarykey;type:uuid"`
		BillItemID string  `gorm:"primarykey;type:uuid"`
		PaidPrice  float64 `gorm:"not null;numeric;precision:10;scale:2"`
		Currency   string  `gorm:"not null;size:5"`
	}

	PastBillsView struct {
		BillID         string
		OwnerID        string
		RestaurantName string
		ItemName       string
		Quantity       int
		UnitPrice      float64
		OrderPrice     float64
		PaidPrice      float64
		Currency       string
	}
)

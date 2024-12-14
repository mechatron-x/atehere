package model

import (
	"gorm.io/gorm"
)

type PostOrder struct {
	ID         string `gorm:"primarykey;type:uuid"`
	SessionID  string `gorm:"type:uuid"`
	Session    Session
	OrderedBy  string   `gorm:"type:uuid"`
	Customer   Customer `gorm:"foreignKey:OrderedBy"`
	MenuItemID string   `gorm:"type:uuid"`
	MenuItem   MenuItem
	Quantity   int
	gorm.Model
}

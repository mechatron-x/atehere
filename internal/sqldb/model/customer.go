package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	ID        string `gorm:"primarykey;type:uuid"`
	FullName  string `gorm:"not null;size:100"`
	Gender    string `gorm:"not null;size:10"`
	BirthDate string `gorm:"size:20"`
	gorm.Model
}

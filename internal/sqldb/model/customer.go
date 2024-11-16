package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	FullName  string
	Gender    string
	BirthDate string
}

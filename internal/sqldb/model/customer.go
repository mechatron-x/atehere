package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	ID        string `gorm:"primarykey"`
	FullName  string
	Gender    string
	BirthDate string
	gorm.Model
}

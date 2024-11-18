package model

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	ID          string `gorm:"primarykey"`
	FullName    string
	PhoneNumber string
	Restaurants []Restaurant `gorm:"foreignKey:OwnerID"`
}

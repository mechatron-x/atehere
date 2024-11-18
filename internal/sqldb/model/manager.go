package model

import "gorm.io/gorm"

type Manager struct {
	ID          string `gorm:"primarykey;type:uuid"`
	FullName    string `gorm:"not null;size:100"`
	PhoneNumber string
	gorm.Model
}

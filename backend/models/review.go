package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Rating uint `gorm:"not null"`
	Comment string `gorm:"not null"`
}

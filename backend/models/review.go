package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	FullName string `gorm:"not null" json:"full_name"`
	Rating uint `gorm:"not null" json:"rating"`
	Comment string `gorm:"not null" json:"comment"`
}

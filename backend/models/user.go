package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"not null" json:"full_name"`
	Email string `gorm:"not null,unique" json:"email"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
	Password string `gorm:"not null" json:"-"`
	Role string `gorm:"not null" json:"role"`
}

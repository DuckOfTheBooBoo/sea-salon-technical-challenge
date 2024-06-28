package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	FullName    string    `gorm:"not null" json:"full_name"`
	PhoneNumber string    `gorm:"not null" json:"phone_number"`
	Service     string    `gorm:"not null" json:"service"`
	Date        time.Time `gorm:"not null" json:"date"`
	UserID      uint
	BranchID uint
}

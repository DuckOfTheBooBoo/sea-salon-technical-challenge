package models

import (
	"gorm.io/gorm"
	"time"
)

type Branch struct {
	gorm.Model
	BranchName    string        `gorm:"not null" json:"branch_name"`
	BranchAddress string        `gorm:"not null" json:"branch_address"`
	Lat           float64       `gorm:"not null" json:"lat"`
	Lng           float64       `gorm:"not null" json:"lng"`
	OpenAt        time.Time     `gorm:"not null" json:"open_time"`
	ClosedAt      time.Time     `gorm:"not null" json:"close_time"`
	Services      []Service     `gorm:"many2many:branch_services" json:"services"`
	Reservations  []Reservation `gorm:"foreignKey:BranchID" json:"-"`
}

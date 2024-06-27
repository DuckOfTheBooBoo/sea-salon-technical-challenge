package models

import (

	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	BranchName    string     `gorm:"not null" json:"branch_name"`
	BranchAddress string     `gorm:"not null" json:"branch_address"`
	Lat           float64    `gorm:"not null" json:"lat"`
	Lng           float64    `gorm:"not null" json:"lng"`
	OpenAt        string  `gorm:"not null;type:time" json:"open_time"`
	ClosedAt      string  `gorm:"not null;type:time" json:"close_time"`
	Services      []Service `gorm:"many2many:branch_services" json:"services"`
}

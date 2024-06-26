package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	BranchName    string  `gorm:"not null" json:"branch_name"`
	BranchAddress string  `gorm:"not null" json:"branch_address"`
	Lat           float64 `gorm:"not null" json:"lat"`
	Lng           float64 `gorm:"not null" json:"lng"`
}

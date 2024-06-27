package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	ServiceName string `gorm:"not null;unique" json:"service_name"`
	ServiceCode string `gorm:"not null;unique" json:"service_code"`
	Duration uint `gorm:"not null" json:"duration"`
	Branches []Branch `gorm:"many2many:branch_services" json:"branch_services"`
}

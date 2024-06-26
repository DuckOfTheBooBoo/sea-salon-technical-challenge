package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ServiceName string `gorm:"not null" json:"service_name"`
	Branches []*Branch `gorm:"many2many:branch_services" json:"branch_services"`
}

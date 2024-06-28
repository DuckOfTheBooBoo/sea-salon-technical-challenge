package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type model struct {
	gorm.Model
	Time MySQLTime
}

type MySQLTime struct {
	time.Time
}

func (MySQLTime) GormDataType() string {
	return "time"
}

func (MySQLTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

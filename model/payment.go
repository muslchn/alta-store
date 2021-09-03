package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Type    string
	Allowed bool
	Time    time.Time
	Paid    bool
}

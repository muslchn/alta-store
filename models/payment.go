package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID      uint
	Type    string
	Allowed bool
	Time    time.Time
	Paid    bool
}
